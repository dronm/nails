package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 18/12/2023
 *
 * This is QRExtractor controller implimentation file.
 *
 */

import (
	"reflect"	
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
	"context"
	"errors"
	"time"
	"math"
	
	"github.com/dronm/fnsnpd"
	"github.com/dronm/ds/pgds"
	
	"nails/nails_config"
	//"nails/models"
	
	"osbe/repo/userOperation"
	
	"osbe"
	"osbe/logger"
	"osbe/srv"
	"osbe/srv/httpSrv"
	"osbe/socket"
	"osbe/response"	
	
	"github.com/jackc/pgx/v4"
)

//Method implemenation callback
func (pm *QRExtractor_Controller_callback) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	http_sock, ok := sock.(*httpSrv.HTTPSocket)
	if !ok {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, "Attachment_Controller_add_file Not HTTPSocket type")
	}
	r := http_sock.Request
	
	auth := r.Header.Get("Authorization")

	if len(auth) < len(AUTH_PREF) || strings.ToLower(auth[0:len(AUTH_PREF)]) != AUTH_PREF {
		app.GetLogger().Errorf("QRExtractor_Controller_callback: unauthorized access %s", r.RemoteAddr)
		http_sock.Response.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	//b64.StdEncoding.DecodeString(
	token := auth[len(AUTH_PREF):]
	conf := app.GetConfig().(*nails_config.NailsConfig)
	if token != conf.QRExtractorParams.CallbackKey {
		app.GetLogger().Errorf("QRExtractor_Controller_callback: unauthorized access, wrong token %s", r.RemoteAddr)
		http_sock.Response.WriteHeader(http.StatusUnauthorized)
		return nil
	}
	
	request_id := r.Header.Get("request_id")
	
	qr_url, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.GetLogger().Errorf("QRExtractor_Controller_callback: ioutil.ReadAll() failed %v", err)
		http_sock.Response.WriteHeader(http.StatusBadGateway)
		return nil		
	}
	document_href := string(qr_url)
	
	dStore := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := dStore.GetPrimary()
	if err_с != nil {
		app.GetLogger().Errorf("QRExtractor_Controller_callback receipt_add_continue dStore.GetPrimary() failed: %v", err_с)
		return nil
	}
	defer dStore.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	db_rec_details := struct{
		id int
		user_id int64
		operation_id string
		specialist_inn string
		total float32
		period time.Time
		max_rec_date time.Time
		firm_inn string
		services string
	}{}
	
	if err := conn.QueryRow(context.Background(),
		`SELECT
			rct.id,
			op.user_id,
			rct.operation_id,
			sp.inn,
			coalesce(sal.total,0) - coalesce((
				SELECT
					sum(prev_rec.document_total)
				FROM specialist_receipts AS prev_rec
				WHERE prev_rec.specialist_period_salary_detail_id = rct.specialist_period_salary_detail_id
					AND prev_rec.document_parsed AND coalesce(prev_rec.document_error,'') = ''
			), 0),
			sal.period,
			sal.period::date + '2 months'::interval - '1 day'::interval AS max_rec_date,
			coalesce(firms.inn, ''),
			coalesce(lower(const_specialist_services_val()), '')
		FROM specialist_receipts AS rct
		LEFT JOIN specialists AS sp ON sp.id = rct.specialist_id
		LEFT JOIN studios AS stud ON stud.id = sp.studio_id
		LEFT JOIN firms ON firms.id = stud.firm_id
		LEFT JOIN specialist_period_salary_details AS sal ON sal.id = rct.specialist_period_salary_detail_id
		LEFT JOIN user_operations AS op ON op.operation_id = rct.operation_id
		WHERE rct.qrextr_request_id = $1`,
		request_id,
	).Scan(&db_rec_details.id,
		&db_rec_details.user_id,	
		&db_rec_details.operation_id,	
		&db_rec_details.specialist_inn,
		&db_rec_details.total,
		&db_rec_details.period,
		&db_rec_details.max_rec_date,
		&db_rec_details.firm_inn,
		&db_rec_details.services,
	); err != nil {
		end_qr_check(app.GetLogger(),
			conn,
			request_id,
			"Завершено с ошибкой",
			fmt.Sprintf("QRExtractor_Controller_callback conn.QueryRow() failed: %v, qrextr_request_id=", err, request_id),
			db_rec_details.user_id,
			db_rec_details.operation_id,
			document_href,
			db_rec_details.id,
			0,
			time.Time{})
		return nil
	}
app.GetLogger().Debugf("request_id=%s", request_id)		
app.GetLogger().Debugf("QRUrl=%s", document_href)	

	if document_href == "" {
		end_qr_check(app.GetLogger(),
			conn,
			request_id,
			"QR код не найден",
			"",
			db_rec_details.user_id,
			db_rec_details.operation_id,
			document_href,
			db_rec_details.id,
			0,
			time.Time{})		
		return nil
	}

	receipt, err := fnsnpd.NewCheckFlFromUrl(document_href)
	if err != nil {
		end_qr_check(app.GetLogger(),
			conn,
			request_id,
			"Завершено с ошибкой",
			fmt.Sprintf("QRExtractor_Controller_callback fnsnpd.NewCheckFlFromUrl() failed: %v", err),
			db_rec_details.user_id,
			db_rec_details.operation_id,
			document_href,
			db_rec_details.id,
			0,
			time.Time{})
		return nil
	}
	
	//checking
	if receipt.BuyerInn != db_rec_details.firm_inn {
		end_qr_check(app.GetLogger(),
			conn,
			request_id,
			fmt.Sprintf("Неверный ИНН покупателя в чеке: %s, нужен: %s", receipt.BuyerInn, db_rec_details.firm_inn),
			"",
			db_rec_details.user_id,
			db_rec_details.operation_id,
			document_href,
			db_rec_details.id,
			receipt.Total,
			receipt.Date)		
		return nil
	} 

	if receipt.Inn != db_rec_details.specialist_inn {
		end_qr_check(app.GetLogger(),
			conn,
			request_id,
			fmt.Sprintf("Неверный ИНН мастера в чеке: %s, нужен: %s", receipt.Inn, db_rec_details.specialist_inn),
			"",
			db_rec_details.user_id,
			db_rec_details.operation_id,
			document_href,
			db_rec_details.id,
			receipt.Total,
			receipt.Date)		
		return nil
	} 

	if receipt.Total > db_rec_details.total {
		end_qr_check(app.GetLogger(),
			conn,
			request_id,
			fmt.Sprintf("Неверная сумма в чеке: %f, остаток: %f", math.Floor(float64(receipt.Total)*100)/100, math.Floor(float64(db_rec_details.total)*100)/100),
			"",
			db_rec_details.user_id,
			db_rec_details.operation_id,
			document_href,
			db_rec_details.id,
			receipt.Total,
			receipt.Date)
		return nil
	} 
	
	//Видимо взять месяц з/п и следующий т.е.
	// КонМесяца(db_rec_details.period + 1 месяц) <= receipt.Date >= db_rec_details.period
	if receipt.Date.After(db_rec_details.max_rec_date) || receipt.Date.Before(db_rec_details.period) {
		end_qr_check(app.GetLogger(),
			conn,
			request_id,
			fmt.Sprintf("Неверная дата чека: %s, должна быть: %s - %s", receipt.Date.Format("01/02/2006"),
				db_rec_details.period.Format("01/02/2006"), db_rec_details.max_rec_date.Format("01/02/2006")),
			"",
			db_rec_details.user_id,
			db_rec_details.operation_id,
			document_href,
			db_rec_details.id,
			receipt.Total,
			receipt.Date)		
		return nil
	} 
	
	if len(receipt.Items) == 0 {
		end_qr_check(app.GetLogger(),
			conn,
			request_id,
			"Не найдены товары/услуги в чеке",
			"",
			db_rec_details.user_id,
			db_rec_details.operation_id,
			document_href,
			db_rec_details.id,
			receipt.Total,
			receipt.Date)		
		return nil
	}
	 for _, it := range receipt.Items {
		if !strings.Contains(db_rec_details.services, strings.ToLower(it.Name)) {
			end_qr_check(app.GetLogger(),
				conn,
				request_id,
				fmt.Sprintf("Неверная услуга в чеке: %s, должно быть одно из: %s", it.Name, db_rec_details.services),
				"",
				db_rec_details.user_id,
				db_rec_details.operation_id,
				document_href,
				db_rec_details.id,
				receipt.Total,
				receipt.Date)		
			
			return nil
		}
	}	

app.GetLogger().Debugf("receipt=%+v", receipt)	
	
	return nil
}

func end_qr_check(lg logger.Logger, conn *pgx.Conn, requestID string, operErrText, logErrText string, userID int64, operationID string, qrURL string, dbRecID int, recTotal float32, recDate time.Time) {
	var check_err_pt *string
	if operErrText != "" {
		if logErrText == "" {
			logErrText = operErrText
		}
		lg.Error(logErrText)
		userOperation.EndUserOperationWithError(lg, conn, userID, operationID, errors.New(operErrText))	
		check_err_pt = &operErrText
		
	}else{
		userOperation.EndUserOperation(lg, conn, userID, operationID)	
	}

	if _, err := conn.Exec(context.Background(),
		`UPDATE specialist_receipts
		SET
			document_date_time = $2,
			document_total = $3,
			document_parsed = TRUE,
			document_href = $4,
			document_error = $5
		WHERE id = $1`,
		dbRecID,		
		recDate,
		recTotal,
		qrURL,
		check_err_pt,
	); err != nil {
		lg.Errorf("QRExtractor_Controller_callback conn.Exec() UPDATE failed: %v, qrextr_request_id=%s", err, requestID)
	}					
	
}



