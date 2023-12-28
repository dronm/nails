package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 28/11/2023
 *
 * This is SpecialistDocument controller implimentation file.
 *
 */

import (
	"reflect"	
	"fmt"
	"context"
	"os"
	"os/exec"
	"errors"
	"mime/multipart"
	"strings"
	"path/filepath"
	"io"
	"io/ioutil"
	
	"github.com/dronm/ds/pgds"
	"nails/models"
	
	"osbe"
	"osbe/srv"
	"osbe/logger"
	"osbe/socket"
	"osbe/response"
	"osbe/srv/httpSrv"
	
	"osbe/repo/docAttachment"
	"osbe/repo/userOperation"		
	
	//"github.com/jackc/pgx/v4"
)

//Method implemenation insert
func (pm *SpecialistDocument_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistDocument"], &models.SpecialistDocument_keys{}, sock.GetPresetFilter("SpecialistDocument"))	
}

//Method implemenation delete
func (pm *SpecialistDocument_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistDocument"], sock.GetPresetFilter("SpecialistDocument"))	
}

//Method implemenation get_object
func (pm *SpecialistDocument_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistDocumentList"], &models.SpecialistDocumentList{}, sock.GetPresetFilter("SpecialistDocumentList"))	
}

//Method implemenation get_list
func (pm *SpecialistDocument_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistDocumentList"], &models.SpecialistDocumentList{}, sock.GetPresetFilter("SpecialistDocumentList"))	
}

//Method implemenation update
func (pm *SpecialistDocument_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistDocument"], sock.GetPresetFilter("SpecialistDocument"))	
}

//Method implemenation get_for_sign_list
func (pm *SpecialistDocument_Controller_get_for_sign_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistDocumentForSignList"], &models.SpecialistDocumentForSignList{}, sock.GetPresetFilter("SpecialistDocumentForSignList"))	
}

//Method implemenation get_file
func (pm *SpecialistDocument_Controller_get_file) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	args := rfltArgs.Interface().(*models.SpecialistDocument_get_file)
	//fmt.Printf("args= %+v\n", args)
	
	att_pm, err := app.GetMD().Controllers["Attachment"].GetPublicMethod("get_file")
	if err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistDocument_Controller_get_file GetPublicMethod() failed: %v",err))
	}
	att_argv := &docAttachment.Attachment_get_file_argv{Argv: &docAttachment.Attachment_get_file{
		Ref: args.Ref,	
		Content_id: args.Content_id,
		Inline: args.Inline,
	}}
	rflt_rgs := reflect.ValueOf(&att_argv.Argv).Elem()
	if err := att_pm.Run(app, serv, sock, resp, rflt_rgs); err != nil{
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistDocument_Controller_get_file Run() failed: %v",err))
	}
	
	//set opened
	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetPrimary()
	if err_с != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistDocument_Controller_get_file Request.GetPrimary(): %v",err_с))
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	if _, err := conn.Exec(context.Background(),
		`UPDATE specialist_documents SET open_date_time = now()
		WHERE id = $1 AND open_date_time IS NULL`,		
		args.Doc_id.GetValue(),
	); err != nil{
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistDocument_Controller_get_file conn.Exec() failed: %v",err))
	}

	//+event
	EmitEvent(conn, "SpecialistDocument.update", args.Doc_id.GetValue())
	
	return nil
}

//Method implemenation sign
func (pm *SpecialistDocument_Controller_sign) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	http_sock, ok := sock.(*httpSrv.HTTPSocket)
	if !ok {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, "SpecialistDocument_Controller_sign Not HTTPSocket type")
	}

	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetPrimary()
	if err_с != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_print_app Request.GetSecondary(): %v",err_с))
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()

	args := rfltArgs.Interface().(*models.SpecialistDocument_sign)
	
	sess := sock.GetSession()	
	user_id := sess.GetInt(SESS_VAR_ID)
	userOperation.StartUserOperation(conn, "signing", user_id, args.Operation_id.GetValue())
		
	file, file_h, err := http_sock.Request.FormFile("signature[]")
	if err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistDocument_Controller_sign Request.FormFile(): %v",err))
	}	
	mimes := []httpSrv.MIME_TYPE{httpSrv.MIME_TYPE_png}
	if !docAttachment.FileHeaderContainsMimes(file_h, mimes) {
		return osbe.NewPublicMethodError(ER_UNSUPPORTED_MIME_CODE, ER_UNSUPPORTED_MIME_DESCR)
	}
	
	go sign_continue(app.GetDataStorage().(*pgds.PgProvider), app.GetLogger(), app.GetBaseDir(), args.Operation_id.GetValue(), args.Id.GetValue(), file, user_id)
	
	return nil
}

func sign_continue(dStore *pgds.PgProvider, log logger.Logger, baseDir string, operationID string, docID int64, file multipart.File, userID int64) {
	defer file.Close()

	pool_conn, conn_id, err_с := dStore.GetPrimary()
	if err_с != nil {
		log.Errorf("sign_continue dStore.GetPrimary() failed: %v", err_с)				
		//cannot end operation!
		return		
	}
	defer dStore.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	var ret_err error
	defer func(){
		if ret_err != nil {
			log.Errorf("sign_continue failed: %v", ret_err)
			userOperation.EndUserOperationWithError(log, conn, userID, operationID, errors.New(ER_INTERNAL))
		}
	}()
	//Барем документ со всеми заменами и без подписи
	//это attachments с идентификаторм = template_att_id
	//Открываем zip
	//Заменяем файл /word/media/image1.jpeg на новый и все
	
	doc := make([]byte, 0)
	doc_template_name := ""
	ref := docAttachment.Ref_Type{}
	content_info := docAttachment.Content_info_Type{}
	document_att_id := 0
	template_att_id := 0
	template_img_name := ""
	if err := conn.QueryRow(context.Background(),
		`SELECT
			tmpl.name,
			att.ref,
			att.content_info,
			att.content_data,
			sp.document_att_id,
			sp.template_att_id,
			coalesce(tmpl.sign_image_name, 'image1.png')
		FROM specialist_documents AS sp
		LEFT JOIN attachments AS att ON att.id = sp.template_att_id
		LEFT JOIN document_templates AS tmpl ON tmpl.id = sp.document_template_id
		WHERE sp.id = $1
		`,
		docID,
	).Scan(&doc_template_name,		
		&ref,
		&content_info,
		&doc,
		&document_att_id,
		&template_att_id,
		&template_img_name,
	); err != nil {
		ret_err = err
		return
	}

	out_doc, err := os.CreateTemp("", "nails_doc")
	if err != nil {
		ret_err = err
		return
	}
	defer os.Remove(out_doc.Name())
	if _, err := out_doc.Write(doc); err != nil {
		out_doc.Close()
		ret_err = err
		return
	}
	out_doc.Close()
	
	//store signature
	sig_img, err := os.CreateTemp("", "nails_sig")
	if err != nil {
		ret_err = err
		return
	}
	//defer os.Remove(sig_img.Name())
	if _, err = io.Copy(sig_img, file); err != nil {
		sig_img.Close()
		ret_err = err
		return
	}
	sig_img.Close()

	//open as zip
	//substitute jpeg file in /word/media/image1.png
	out_doc_fn, err := ChangeImage(out_doc.Name(), sig_img.Name(), "word/media/" + template_img_name)
	if err != nil {
		ret_err = err
		return
	}
	defer os.Remove(out_doc_fn)

	//out doc to PDF	
	out_dir := filepath.Join(baseDir, docAttachment.CACHE_DIR)
	cmd_args := strings.Split(CONVERT_TO_PDF, " ")
	cmd_args = append(cmd_args, out_dir)
	cmd_args = append(cmd_args, out_doc_fn)
	cmd := exec.Command("soffice", cmd_args...)
	err = cmd.Run()
	if err != nil {		
		ret_err = err
		return
	}
	out_doc_pdf := filepath.Join(out_dir, filepath.Base(out_doc_fn) + ".pdf")
	if _, err := os.Stat(out_doc_pdf); os.IsNotExist(err) {
		ret_err = err
		return
	}
	
	//Переименуем pdf в "правильное" имя для кэша	
	out_pdf_cor_for_cache := docAttachment.GetAttachmentCacheFileName(baseDir, ref.DataType, ref.Keys.Id, doc_template_name + ".pdf")
	if err := os.Rename(out_doc_pdf, out_pdf_cor_for_cache); err != nil {
		ret_err = err				
		return
	}
	
	//pdf preview
	//leave old preview!!!
	/*
	out_pdf_preview_fn := docAttachment.GetPreviewCacheFileName(baseDir, ref.DataType, ref.Keys.Id, content_info.Id)
	if err := docAttachment.GenThumbnail(out_pdf_cor_for_cache, out_pdf_preview_fn, content_info.Name); err != nil {
		ret_err = err
		return
	}
	defer os.Remove(out_pdf_preview_fn)
	
	out_pdf_preview_bt, err := ioutil.ReadFile(out_pdf_preview_fn)
	if err != nil {
		ret_err = err
		return 
	}		
	*/

	out_pdf_bt, err := ioutil.ReadFile(out_pdf_cor_for_cache)
	if err != nil {
		ret_err = err
		return 
	}		
	
	//to database
	
	if _, err := conn.Exec(context.Background(),`BEGIN`); err != nil {
		ret_err = fmt.Errorf("conn.Exec() BEGIN failed: %v", err)
		return
	}
	
	//content_preview = $2
	if _, err := conn.Exec(context.Background(),
		`UPDATE attachments
		SET
			content_data = $1
		WHERE id = $2`,		
		out_pdf_bt,
		document_att_id,
	); err != nil {
		conn.Exec(context.Background(),`ROLLBACK`)
		ret_err = fmt.Errorf("conn.Exec() attachments failed: %v", err)
		return
	}
		
	if _, err := conn.Exec(context.Background(),
		`UPDATE specialist_documents
		SET 
			sign_date_time = now(),
			open_date_time = CASE WHEN open_date_time IS NULL THEN now() ELSE open_date_time END
		WHERE id = $1`,
		docID,		
	); err != nil {
		conn.Exec(context.Background(),`ROLLBACK`)
		ret_err = fmt.Errorf("conn.Exec() specialist_documents failed: %v", err)
		return
	}
	
	//delete template template_att_id	
	
	userOperation.EndUserOperation(log, conn, userID, operationID)
	
	if _, err := conn.Exec(context.Background(),`COMMIT`); err != nil {
		ret_err = fmt.Errorf("conn.Exec() COMMIT failed: %v", err)
		return
	}
	
}




