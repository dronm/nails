package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 15/12/2023
 *
 * This is SpecialistReceipt controller implimentation file.
 *
 */

import (
	"reflect"	
	"context"
	"fmt"
	"errors"
	"mime/multipart"
	"net/http"
	"io"
	"bytes"
	b64 "encoding/base64"
	
	"nails/models"
	
	"nails/nails_config"
	"github.com/dronm/ds/pgds"
	
	"osbe"
	"osbe/srv"
	"osbe/socket"
	"osbe/response"	

	"osbe/repo/docAttachment"
	"osbe/repo/userOperation"
	"osbe/srv/httpSrv"
	
	"github.com/jackc/pgx/v4"
)

//Method implemenation insert
func (pm *SpecialistReceipt_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistReceipt"], &models.SpecialistReceipt_keys{}, sock.GetPresetFilter("SpecialistReceipt"))	
}

//Method implemenation delete
func (pm *SpecialistReceipt_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistReceipt"], sock.GetPresetFilter("SpecialistReceipt"))	
}

//Method implemenation get_object
func (pm *SpecialistReceipt_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistReceipt"], &models.SpecialistReceipt{}, sock.GetPresetFilter("SpecialistReceipt"))	
}

//Method implemenation get_list
func (pm *SpecialistReceipt_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistReceipt"], &models.SpecialistReceipt{}, sock.GetPresetFilter("SpecialistReceipt"))	
}

//Method implemenation update
func (pm *SpecialistReceipt_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistReceipt"], sock.GetPresetFilter("SpecialistReceipt"))	
}

//Method implemenation add
func (pm *SpecialistReceipt_Controller_add) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	http_sock, ok := sock.(*httpSrv.HTTPSocket)
	if !ok {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, "SpecialistReceipt_Controller_add Not HTTPSocket type")
	}
	
	file, file_h, err := http_sock.Request.FormFile("file_content[]")
	if err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReceipt_Controller_add Request.FormFile(): %v",err))
	}	
	mimes := []httpSrv.MIME_TYPE{httpSrv.MIME_TYPE_pdf,
		httpSrv.MIME_TYPE_png,
		httpSrv.MIME_TYPE_jpg,
		httpSrv.MIME_TYPE_jpeg,
		httpSrv.MIME_TYPE_jpe,
	}
	if !docAttachment.FileHeaderContainsMimes(file_h, mimes) {
		return osbe.NewPublicMethodError(ER_UNSUPPORTED_MIME_CODE, ER_UNSUPPORTED_MIME_DESCR)
	}
	
	args := rfltArgs.Interface().(*models.SpecialistReceipt_add)
	sess := sock.GetSession()
	user_id := sess.GetInt(SESS_VAR_ID)
	oper_id := args.Operation_id.GetValue()
	
	go receipt_add_continue(app, app.GetDataStorage().(*pgds.PgProvider), oper_id, args.Specialist_period_salary_detail_id.GetValue(), user_id, file, file_h)
	
	return nil
}

func receipt_add_continue(app osbe.Applicationer, dStore *pgds.PgProvider, operationID string, salDetailID int64, userID int64, file multipart.File, file_h *multipart.FileHeader) {
	pool_conn, conn_id, err_с := dStore.GetPrimary()
	if err_с != nil {
		app.GetLogger().Errorf("SpecialistReceipt_Controller_add receipt_add_continue dStore.GetPrimary() failed: %v", err_с)
		return
	}
	defer dStore.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	userOperation.StartUserOperation(conn, "add_receipt", userID, operationID)			
	
	//preapre attachment	
	ref := docAttachment.Ref_Type{DataType: "specialist_receipts"}
	cont_info := docAttachment.Content_info_Type{Name: file_h.Filename, Id: operationID}		
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		app.GetLogger().Errorf("SpecialistReceipt_Controller_add receipt_add_continue io.Copy() failed: %v", err)
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, errors.New("Завершено с ошибкой"))	
		return		
	}		
	file_bt := buf.Bytes()	
	cont_info.Size = int64(buf.Len())
		
	//call qrextr service it returns requestID
	conf := app.GetConfig().(*nails_config.NailsConfig)
	qr_extr_b := bytes.NewBuffer(nil)
	qr_extr_w := multipart.NewWriter(qr_extr_b)
	part, err := qr_extr_w.CreateFormFile("image", file_h.Filename)
	if err != nil {
		app.GetLogger().Errorf("SpecialistReceipt_Controller_add receipt_add_continue w.CreateFormFile() failed: %v", err)
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, errors.New("Завершено с ошибкой"))	
		return		
	}
	file.Seek(0, io.SeekStart)
	if _, err := io.Copy(part, file); err != nil {
		app.GetLogger().Errorf("SpecialistReceipt_Controller_add receipt_add_continue io.Copy() failed: %v", err)
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, errors.New("Завершено с ошибкой"))	
		return		
	}
	qr_extr_w.Close()

	url := conf.QRExtractorParams.Host
	if len(url) == 0 {
		app.GetLogger().Error("SpecialistReceipt_Controller_add conf.QRExtractorParams.Host not set")
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, errors.New("Завершено с ошибкой"))	
		return		
	}
	if url[len(url) - 1 : len(url)] != "/" {
		url += "/"
	}
	req, err := http.NewRequest("POST", url + "add", qr_extr_b)
	if err != nil {
		app.GetLogger().Errorf("SpecialistReceipt_Controller_add receipt_add_continue http.NewRequest() failed: %v", err)
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, errors.New("Завершено с ошибкой"))	
		return		
	}
	req.Header.Set("Content-Type", qr_extr_w.FormDataContentType())
	req.Header.Set("Authorization", "Basic " + b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", conf.QRExtractorParams.AppName, conf.QRExtractorParams.Pwd))))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		app.GetLogger().Errorf("SpecialistReceipt_Controller_add receipt_add_continue client.Do() failed: %v", err)
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, errors.New("Завершено с ошибкой"))	
		return		
	}
	if resp.StatusCode != http.StatusOK {
		app.GetLogger().Errorf("SpecialistReceipt_Controller_add receipt_add_continue client.Do() StatusCode: %d", resp.StatusCode)
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, errors.New("Завершено с ошибкой"))	
		return		
	}
	defer resp.Body.Close()
	qrextr_request_id, err := io.ReadAll(resp.Body)
	if err != nil {
		app.GetLogger().Errorf("SpecialistReceipt_Controller_add receipt_add_continue io.ReadAll() failed: %v", err)
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, errors.New("Завершено с ошибкой"))	
		return				
	}
	
	//
	rec_id := 0
	if err := conn.QueryRow(context.Background(),
		`INSERT INTO specialist_receipts (specialist_id, specialist_period_salary_detail_id, qrextr_request_id, operation_id)
		SELECT
			d.specialist_id,
			d.id,
			$2,
			$3
		FROM specialist_period_salary_details AS d
		WHERE d.id = $1
		RETURNING id`,
		salDetailID,
		qrextr_request_id,
		operationID,
	).Scan(&rec_id); err != nil && err != pgx.ErrNoRows {		
		app.GetLogger().Errorf("SpecialistReceipt_Controller_add receipt_add_continue conn.QueryRow() failed: %v", err)
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, errors.New("Завершено с ошибкой"))	
		return
	}
	
	ref.Keys = docAttachment.AttachmentKey{Id: docAttachment.HttpInt(rec_id)}
	preview_bt, err := docAttachment.GenAttachmentThumbnail(app.GetBaseDir(), ref.DataType, ref.Keys.Id, &cont_info, buf)
	if err != nil {
		app.GetLogger().Errorf("SpecialistReceipt_Controller_add receipt_add_continue GenAttachmentThumbnail() failed: %v", err)
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, errors.New("Завершено с ошибкой"))	
		return
	}

	if err := docAttachment.StoreAttachment(conn, &ref, &cont_info, file_bt, preview_bt); err != nil {
		app.GetLogger().Errorf("SpecialistReceipt_Controller_add receipt_add_continue StoreAttachment() failed: %v", err)
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, errors.New("Завершено с ошибкой"))	
		return
	}
	
	userOperation.ProgressEvent(conn, operationID, "file", "Чек загружен, выполняется проверка ФНС", true)					
	
	//continue in QRExtractor
}

