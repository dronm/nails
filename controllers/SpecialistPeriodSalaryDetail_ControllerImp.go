package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 13/12/2023
 *
 * This is SpecialistPeriodSalaryDetail controller implimentation file.
 *
 */

import (
	"reflect"	
	"time"
	"fmt"
	"strings"
	"context"
	
	"nails/models"
	
	"github.com/dronm/ds/pgds"
	"github.com/dronm/clbnkexch"
	
	"osbe"
	"osbe/srv"
	"osbe/socket"
	"osbe/response"	
	"osbe/sql"
	"osbe/srv/httpSrv"
	
	//"github.com/jackc/pgx/v4"
)

const (
	RESP_ER_NO_DOC = 1000
	RESP_ER_NO_DOC_DESCR = "Нет документов для отправки в банк"
)

//Method implemenation insert
func (pm *SpecialistPeriodSalaryDetail_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistPeriodSalaryDetail"], &models.SpecialistPeriodSalaryDetail_keys{}, sock.GetPresetFilter("SpecialistPeriodSalaryDetail"))	
}

//Method implemenation delete
func (pm *SpecialistPeriodSalaryDetail_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistPeriodSalaryDetail"], sock.GetPresetFilter("SpecialistPeriodSalaryDetail"))	
}

//Method implemenation get_object
func (pm *SpecialistPeriodSalaryDetail_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistPeriodSalaryDetailList"], &models.SpecialistPeriodSalaryDetailList{}, sock.GetPresetFilter("SpecialistPeriodSalaryDetailList"))	
}

//Method implemenation get_list
func (pm *SpecialistPeriodSalaryDetail_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistPeriodSalaryDetailList"], &models.SpecialistPeriodSalaryDetailList{}, sock.GetPresetFilter("SpecialistPeriodSalaryDetailList"))	
}

//Method implemenation get_for_pay_list
func (pm *SpecialistPeriodSalaryDetail_Controller_get_for_pay_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	//add filter: bank_payments_ref is null
	pf_cond := sql.FilterCondCollection{&sql.FilterCond{Expression: "(bank_payments_ref IS NULL OR bank_payments_ref->'keys'->>'id' IS NULL)"}}
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistPeriodSalaryDetailList"], &models.SpecialistPeriodSalaryDetailList{}, pf_cond)	
}

//Method implemenation update
func (pm *SpecialistPeriodSalaryDetail_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistPeriodSalaryDetail"], sock.GetPresetFilter("SpecialistPeriodSalaryDetail"))	
}

//Method implemenation documents_to_bank
func (pm *SpecialistPeriodSalaryDetail_Controller_documents_to_bank) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	http_sock, ok := sock.(*httpSrv.HTTPSocket)
	if !ok {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, "SpecialistPeriodSalaryDetail_Controller_documents_to_bank Not HTTPSocket type")
	}
	
	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err := d_store.GetPrimary()
	if err != nil {
		return err
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	args := rfltArgs.Interface().(*models.Object_key_ar)
	if len(args.Ids) == 0 {
		return osbe.NewPublicMethodError(RESP_ER_NO_DOC, RESP_ER_NO_DOC_DESCR)
	}
	
	var params_s strings.Builder
	for i, id:= range args.Ids {
		if i > 0 {
			params_s.WriteString(",")
		}
		params_s.WriteString(fmt.Sprintf("%d", id.Id.GetValue()))		
	}
	
	res_struct := models.BankPayment{}	
	row_fields, fld_ids, err := osbe.MakeStructRowFields(&res_struct, "")
	if err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistPeriodSalaryDetail_Controller_documents_to_bank osbe.MakeStructRowFields(): %v", err))
	}
	q := fmt.Sprintf(`SELECT %s FROM specialist_period_salary_details_docs_to_bank(ARRAY[%s])`, fld_ids, params_s.String())	
	rows, err := conn.Query(context.Background(), q)
	if err != nil {		
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistPeriodSalaryDetail_Controller_documents_to_bank pgx.Conn.Query(): %v", err))
	}
	documents := make([]clbnkexch.BankDocumenter, 0)
	for rows.Next() {		
		if err := rows.Scan(row_fields...); err != nil {		
			return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistPeriodSalaryDetail_Controller_documents_to_bank pgx.Rows.Scan(): %v",err))	
		}
		documents = append(documents,
			&clbnkexch.PPDocument{Num: int(res_struct.Document_num.GetValue()),
				Date: res_struct.Document_date.GetValue(),
				Sum: res_struct.Document_total.GetValue(),
				Payer: &clbnkexch.Firm{Name: res_struct.Payer.GetValue(), Inn: res_struct.Payer_inn.GetValue(), Account: res_struct.Payer_acc.GetValue()},
				PayerBank: &clbnkexch.Bank{Name: res_struct.Payer_bank_acc.GetValue(),
					Place: res_struct.Payer_bank_place.GetValue(),
					Bik: res_struct.Payer_bank_bik.GetValue(),
					Account: res_struct.Payer_acc.GetValue()},
				Receiver: &clbnkexch.Firm{Name: res_struct.Rec.GetValue(), Inn: res_struct.Rec_inn.GetValue(), Account: res_struct.Rec_acc.GetValue()},
				ReceiverBank: &clbnkexch.Bank{Name: res_struct.Rec_bank_acc.GetValue(),
					Place: res_struct.Rec_bank_place.GetValue(),
					Bik: res_struct.Rec_bank_bik.GetValue(),
					Account: res_struct.Rec_acc.GetValue()},			
				PayType: clbnkexch.PAY_TYPE_DIG,
				OplType: "01",
				Order: 5,
				PayComment: res_struct.Document_comment.GetValue(),
			},
		)
	}
	if err := rows.Err(); err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistPeriodSalaryDetail_Controller_documents_to_bank pgx.Rows.Next(): %v",err))
	}
	
	//+event for grid update
	EmitEvent(conn, "SpecialistPeriodSalaryDetail.update", 0)
	
	fl := clbnkexch.NewExchFile([]clbnkexch.DocumentType{clbnkexch.DOCUMENT_TYPE_PP}, documents)
	b, err := fl.Render()
	if err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistPeriodSalaryDetail_documents_to_bank clbnkexch.Render(): %v",err))		
	}
	
	httpSrv.ServeContent(http_sock, &b, "cl_to_bank.txt", httpSrv.MIME_TYPE_txt, time.Now(), httpSrv.CONTENT_DISPOSITION_ATTACHMENT)	
	
	return nil
}

//Method implemenation complete
func (pm *SpecialistPeriodSalaryDetail_Controller_complete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.CompleteOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistPeriodSalaryDetailList"], &models.SpecialistPeriodSalaryDetailList{}, sock.GetPresetFilter("SpecialistPeriodSalaryDetailList"))	
}

