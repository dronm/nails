package controllers

import(
	"sync"
	"context"
	"errors"
	"encoding/json"
	"time"
	
	"osbe/logger"
	"osbe/repo/userOperation"
	
	"yclients"	
	"github.com/dronm/ds/pgds"
)

const (
	YCL_STAFF_UPDATE_QUERY = `INSERT INTO ycl_staff (id, data) VALUES
		($1, $2::jsonb)
		ON CONFLICT (id) DO UPDATE SET data = $2::jsonb`
	
	YCL_TRANSACTION_UPDATE_QUERY = `INSERT INTO ycl_transactions (id, data, record_inf_updated) VALUES
		($1, $2::jsonb, TRUE)
		ON CONFLICT (id) DO UPDATE SET data = $2::jsonb`
		
	YCL_TRANSACTION_REC_INF_SELECT_QUERY = `SELECT id, record_id
		FROM ycl_transactions
		WHERE record_inf_updated = FALSE
		LIMIT 100`	
			
	YCL_TRANSACTION_REC_INF_UPDATE_QUERY = `UPDATE ycl_transactions
		SET
			data = data || jsonb_build_object('record', $1::jsonb),
			record_inf_updated = TRUE
		WHERE id = $2`		
		
)

var YClients *YClientsAPI

type YClientsAPI struct {
	CompanyID int	
	m sync.RWMutex
	client *yclients.YClients
	auth *yclients.Auth	
}
func NewYClientsAPI(login, password, partnerToken string, companyID int) *YClientsAPI{	
	ycl := YClientsAPI{CompanyID: companyID}
	ycl.client = yclients.NewYClients(partnerToken)
	ycl.auth = &yclients.Auth{Login: login, Password: password}
	return &ycl
}

func (ycl *YClientsAPI) Auth() error {
	ycl.m.Lock()
	defer ycl.m.Unlock()
	
	return ycl.client.Auth(ycl.auth)	
}

func (ycl *YClientsAPI) UpdateStaff(operationID string, userID int64, dStore *pgds.PgProvider, log logger.Logger) {
	pool_conn, conn_id, err_с := dStore.GetPrimary()
	if err_с != nil {
		log.Errorf("YClientsAPI UpdateStaff() dStore.GetPrimary() failed: %v", err_с)		
		return
	}
	defer dStore.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	userOperation.StartUserOperation(conn, "update_staff", userID, operationID)
	
	if ycl.client.User.UserToken == "" {
		if err := ycl.Auth(); err != nil {
			log.Errorf("YClientsAPI UpdateStaff() yclients.Auth() failed: %v", err)
			userOperation.EndUserOperationWithError(log, conn, userID, operationID, errors.New(ER_INTERNAL))
			return
		}
	}
	
	staff, err := ycl.client.StaffAll(ycl.CompanyID)
	if err == yclients.ErrUnauthorized || err == yclients.ErrForbidden {
		//auth error
		if err := ycl.Auth(); err != nil {
			log.Errorf("YClientsAPI UpdateStaff() yclients.Auth() failed: %v", err)
			userOperation.EndUserOperationWithError(log, conn, userID, operationID, errors.New(ER_INTERNAL))
			return
		}		
	}else if err != nil {
		log.Errorf("YClientsAPI UpdateStaff() yclients.StaffAll() failed: %v", err)
		userOperation.EndUserOperationWithError(log, conn, userID, operationID, errors.New(ER_INTERNAL))
		return		
	}
log.Debugf("YClientsAPI got staff, count: %d", len(staff))		
	for _,s := range staff {
		s_bt, err := json.Marshal(s)
		if err != nil {
			log.Errorf("YClientsAPI UpdateStaff() json.Marshal() failed: %v", err)
			userOperation.EndUserOperationWithError(log, conn, userID, operationID, errors.New(ER_INTERNAL))
			return
		}
		if _, err := conn.Exec(context.Background(),
			YCL_STAFF_UPDATE_QUERY,
			s.ID,
			s_bt,
		); err != nil {
			log.Errorf("YClientsAPI UpdateStaff() conn.Exec() failed: %v", err)
			userOperation.EndUserOperationWithError(log, conn, userID, operationID, errors.New(ER_INTERNAL))
			return
		}
	}
	userOperation.EndUserOperation(log, conn, userID, operationID)
}

type yclTransaction struct {
	ID int `json:"id"`					//number ID транзакции
	DocumentID int `json:"document_id"`			//number ID документа
	Date string `json:"date"`				//Дата создания транзакции
	Amount float32 `json:"amount"`				//Сумма транзакции
	Comment string `json:"comment"`				//Комментарий
	Client *yclients.CatalogRef `json:"client"`			//Клиент	
	Last_change_date string `json:"last_change_date"`
	RecordID int `json:"record_id"`			//Идентификатор записи
	VisitID int `json:"visit_id"`
	Record *yclients.Record `json:"record"`
}

//Вызывается из ГУИ
func (ycl *YClientsAPI) UpdateTransactions(operationID string, userID int64, fromDate time.Time, toDate time.Time,  dStore *pgds.PgProvider, log logger.Logger) {
	pool_conn, conn_id, err_с := dStore.GetPrimary()
	if err_с != nil {
		log.Errorf("YClientsAPI UpdateTransactions() dStore.GetPrimary() failed: %v", err_с)		
		return
	}
	defer dStore.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	userOperation.StartUserOperation(conn, "update_transactions", userID, operationID)
	
	if ycl.client.User.UserToken == "" {
		if err := ycl.Auth(); err != nil {
			log.Errorf("YClientsAPI UpdateTransactions() yclients.Auth() failed: %v", err)
			userOperation.EndUserOperationWithError(log, conn, userID, operationID, errors.New(ER_INTERNAL))
			return
		}
	}

	count := 50
	page := 1
	for {
		flt := 	yclients.FinTransactionFilter{StartDate: yclients.APIDate(fromDate), EndDate: yclients.APIDate(toDate), Page: page, Count: count}
		trans, err := ycl.client.Transactions(ycl.CompanyID, &flt)
		if err == yclients.ErrUnauthorized || err == yclients.ErrForbidden {
			//auth error
			if err := ycl.Auth(); err != nil {
				log.Errorf("YClientsAPI UpdateTransactions() yclients.Auth() failed: %v", err)
				userOperation.EndUserOperationWithError(log, conn, userID, operationID, errors.New(ER_INTERNAL))
				return
			}		
		}else if err != nil {
			log.Errorf("YClientsAPI UpdateTransactions() yclients.StaffAll() failed: %v", err)
			userOperation.EndUserOperationWithError(log, conn, userID, operationID, errors.New(ER_INTERNAL))
			return		
		}
log.Debugf("YClientsAPI UpdateTransactions()len(trans): %d, page: %d", len(trans), page)
		if len(trans) == 0 {
			break //no more data
		}
		
		for _,tr := range trans {
			//add record{staff_id, SeanceLength}
			doc_tr := yclTransaction{ID: tr.ID,
				DocumentID: tr.DocumentID,
				Date: time.Time(tr.Date).Format("2006-01-02T15:04:05-07:00"),
				Amount: tr.Amount,
				Comment: tr.Comment,
				Client: tr.Client,
				VisitID: tr.VisitID,
				Last_change_date:  time.Time(tr.Last_change_date).Format("2006-01-02T15:04:05-07:00"),
				RecordID: tr.RecordID,
			}
			if doc_tr.RecordID > 0 {
				var err error			
				doc_tr.Record, err = ycl.client.Record(ycl.CompanyID, doc_tr.RecordID)
				doc_tr.Record.Staff = nil
				if err != nil {
					log.Errorf("YClientsAPI UpdateTransactions() ycl.client.Record() failed: %v, recordID: %d", err, doc_tr.RecordID)
					userOperation.EndUserOperationWithError(log, conn, userID, operationID, errors.New(ER_INTERNAL))
					return
				}
			}
			tr_bt, err := json.Marshal(doc_tr)
			if err != nil {
				log.Errorf("YClientsAPI UpdateTransactions() json.Marshal() failed: %v", err)
				userOperation.EndUserOperationWithError(log, conn, userID, operationID, errors.New(ER_INTERNAL))
				return
			}

			if _, err := conn.Exec(context.Background(),
				YCL_TRANSACTION_UPDATE_QUERY,
				tr.ID,
				tr_bt,
			); err != nil {
				log.Errorf("YClientsAPI UpdateTransactions() conn.Exec() failed: %v", err)
				userOperation.EndUserOperationWithError(log, conn, userID, operationID, errors.New(ER_INTERNAL))
				return
			}			
		}
		time.Sleep(time.Duration(500) * time.Millisecond)
		page++
	}
	
	userOperation.EndUserOperation(log, conn, userID, operationID)
}

//Вызывается из горутины, обновляет, те, которые получены через коллбак,
//добавляет недостающие данные
func (ycl *YClientsAPI) UpdateTransactionRecordInf(dStore *pgds.PgProvider, log logger.Logger, ctx context.Context) {
	pool_conn, conn_id, err_с := dStore.GetSecondary("")
	if err_с != nil {
		log.Errorf("YClientsAPI UpdateTransactions() dStore.GetPrimary() failed: %v", err_с)		
		return
	}	
	conn := pool_conn.Conn()
	
	if ycl.client.User.UserToken == "" {
		if err := ycl.Auth(); err != nil {
			log.Errorf("YClientsAPI UpdateTransactionRecordInf() yclients.Auth() failed: %v", err)
			dStore.Release(pool_conn, conn_id)
			return
		}
	}

	//make transaction list
	type trans_t struct{
		trans_id int
		record_id int
	}
	trans_list := make([]trans_t, 0)
	rows, err := conn.Query(context.Background(),
		YCL_TRANSACTION_REC_INF_SELECT_QUERY,		
		)
	if err != nil {
		log.Errorf("YClientsAPI UpdateTransactionRecordInf pgx.Conn.Query() failed: %v", err)
		dStore.Release(pool_conn, conn_id)
		return
	}	
	for rows.Next() {
		trans := trans_t{}
		if err := rows.Scan(&trans.trans_id, &trans.record_id); err != nil {
			log.Errorf("YClientsAPI UpdateTransactionRecordInf pgx.Rows.Scan() failed: %v", err)
			dStore.Release(pool_conn, conn_id)
			return
		}
		trans_list = append(trans_list, trans)
	}
	if err := rows.Err(); err != nil {
		log.Errorf("YClientsAPI UpdateTransactionRecordInf() rows.Err() failed: %v", err)
		dStore.Release(pool_conn, conn_id)
		return
	}
	dStore.Release(pool_conn, conn_id)

log.Debugf("***** YClientsAPI UpdateTransactionRecordInf() len(trans_list): %d", len(trans_list))	
	if len(trans_list) == 0 {
		return
	}

	pool_conn, conn_id, err_с = dStore.GetPrimary()
	if err_с != nil {
		log.Errorf("YClientsAPI UpdateTransactions() dStore.GetPrimary() failed: %v", err_с)		
		return
	}
	defer dStore.Release(pool_conn, conn_id)	
	conn = pool_conn.Conn()
	
	for _,trans := range trans_list {
		//retrieve tsansaction inf
		var tr_bt []byte
		if trans.record_id > 0 {
			record, err := ycl.client.Record(ycl.CompanyID, trans.record_id)
			if err != nil {
				log.Errorf("YClientsAPI UpdateTransactionRecordInf() ycl.client.Record() failed: %v, recordID: %d, transactionID: %d", err, trans.record_id, trans.trans_id)
				return
			}
			tr_bt, err = json.Marshal(record)
			if err != nil {
				log.Errorf("YClientsAPI UpdateTransactionRecordInf() json.Marshal() failed: %v", err)
				return
			}
		}
		//update
		if _, err := conn.Exec(ctx,
			YCL_TRANSACTION_REC_INF_UPDATE_QUERY,
			tr_bt,
			trans.trans_id,
		); err != nil {
			log.Errorf("YClientsAPI UpdateTransactionRecordInf() conn.Exec() failed: %v", err)
			return
		}
	}
}

func YClientsServe(dStore *pgds.PgProvider, log logger.Logger, interval int, ctx context.Context) {
	go func(){
		srvLoop:
		for {
			select{
			case <-ctx.Done(): //context cancelled
				break srvLoop
			default:
				YClients.UpdateTransactionRecordInf(dStore, log, ctx)				
				if interval == 0 {
					break srvLoop
				}
				time.Sleep(time.Duration(interval) * time.Millisecond)
			}
		}
	}()
}

//
func YClientsAPIInit(login, password, partnerToken string, companyID int) {
	YClients = NewYClientsAPI(login, password, partnerToken, companyID)
}

