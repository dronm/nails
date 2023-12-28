package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 13/11/2023
 *
 * This is Specialist controller implimentation file.
 *
 */

import (
	"reflect"	
	"fmt"
	"context"
	"mime/multipart"
	
	"nails/models"
	
	"github.com/dronm/ds/pgds"
	
	"osbe"
	"osbe/srv"
	"osbe/socket"
	"osbe/response"	
	"osbe/srv/httpSrv"
	"osbe/model"
	
	cpt "osbe/repo/captcha"
	vld_bank "osbe/repo/validation/rusBank"
	vld_tel "osbe/repo/validation/rusTel"
	vld_email "osbe/repo/validation/email"
	"osbe/repo/docAttachment"
	"osbe/repo/userOperation"
	
	"github.com/jackc/pgx/v4"
)

const (
	SPECIALIST_CHECK_CAPTCHA_ID = "registration"	
	SPECIALIST_CHECK_CAPTCHA_W = 240
	SPECIALIST_CHECK_CAPTCHA_H = 80
	SPECIALIST_CHECK_CAPTCHA_CNT = 6
	
	SEND_REG_DOCS_OPER = "send_reg_docs"
	SEND_SAL_DOCS_OPER = "send_salary_docs"
	
	RESP_ER_REGISTER_CAPTCHA_CODE = 1000
	
	ER_UNSUPPORTED_MIME_CODE = 1001
	ER_UNSUPPORTED_MIME_DESCR = "Неподдерживаемый тип файла"
	
	RESP_ER_REGISTER_TEL_CODE = 1001
	RESP_ER_REGISTER_TEL_DESCR = "Неверный номер телефона"
	
	RESP_ER_REGISTER_EMAIL_CODE = 1002
	RESP_ER_REGISTER_EMAIL_DESCR = "Неверный адрес"
	
	RESP_ER_REGISTER_BIK_CODE = 1003
	RESP_ER_REGISTER_BIK_DESCR = "Неверный БИК банка"

	RESP_ER_REGISTER_ACC_CODE = 1004
	RESP_ER_REGISTER_ACC_DESCR = "Неверный расчетный счет"

	RESP_ER_REGISTER_EXISTS_CODE = 1005
	RESP_ER_REGISTER_EXISTS_DESCR = "С самозанятым с таким ИНН уже заключен контракт"

	RESP_ER_REGISTER_TEL_EXISTS_CODE = 1006
	RESP_ER_REGISTER_TEL_EXISTS_DESCR = "С самозанятым с таким номером телефона уже зарегистрирован"
	
	REG_OPERATION_ID = "register"
	
	ER_MAIL_SEND_INTERNAL = "Ошибка отправки почты: внутренняя ошибка сервера"
	ER_SMS_SEND_INTERNAL = "Ошибка отправки sms: внутренняя ошибка сервера"
	ER_INTERNAL = "Внутренняя ошибка сервера"
)

//Method implemenation insert
func (pm *Specialist_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["Specialist"], &models.Specialist_keys{}, sock.GetPresetFilter("Specialist"))	
}

//Method implemenation delete
func (pm *Specialist_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["Specialist"], sock.GetPresetFilter("Specialist"))	
}

//Method implemenation get_object
func (pm *Specialist_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistDialog"], &models.SpecialistDialog{}, sock.GetPresetFilter("SpecialistDialog"))	
}

//Method implemenation get_list
func (pm *Specialist_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistList"], &models.SpecialistList{}, sock.GetPresetFilter("SpecialistList"))	
}

//Method implemenation update
func (pm *Specialist_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["Specialist"], sock.GetPresetFilter("Specialist"))	
}

//Method implemenation complete
func (pm *Specialist_Controller_complete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.CompleteOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistList"], &models.SpecialistList{}, sock.GetPresetFilter("SpecialistList"))	
}

//Method implemenation register
func (pm *Specialist_Controller_register) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	args := rfltArgs.Interface().(*models.Specialist_register)
	args.Name.SetValue(args.Tel.GetValue())
	sess := sock.GetSession()
	
	var ret_error error
	var file multipart.File
	var file_h *multipart.FileHeader
	defer (func(){
		if ret_error != nil {
			if file != nil {
				file.Close()
			}
			cpt.AddNewCaptcha(sess, app.GetLogger(), resp, SPECIALIST_CHECK_CAPTCHA_ID, SPECIALIST_CHECK_CAPTCHA_W, SPECIALIST_CHECK_CAPTCHA_H, SPECIALIST_CHECK_CAPTCHA_CNT)
		}
	})()

	//1) captcha check
	if ok, err := cpt.CaptchaVerify(sess, app.GetLogger(), []byte(args.Captcha.GetValue()), SPECIALIST_CHECK_CAPTCHA_ID); !ok || err != nil {
		if err != nil {
			ret_error = osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register CaptchaVerify(): %v",err))
		}else{
			ret_error = osbe.NewPublicMethodError(RESP_ER_REGISTER_CAPTCHA_CODE, RESP_ER_CAPTCHA_DESCR) //from user
		}
		return ret_error
	}
		
	http_sock, ok := sock.(*httpSrv.HTTPSocket)
	if !ok {
		ret_error = osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, "Attachment_Controller_add_file Not HTTPSocket type")
		return ret_error
	}
	//		
	var err error
	file, file_h, err = http_sock.Request.FormFile("passport_file_content[]")
	if err != nil {
		ret_error = osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register Request.FormFile(): %v",err))
		return ret_error
	}	
	mimes := []httpSrv.MIME_TYPE{httpSrv.MIME_TYPE_pdf,
		httpSrv.MIME_TYPE_png,
		httpSrv.MIME_TYPE_jpg,
		httpSrv.MIME_TYPE_jpeg,
		httpSrv.MIME_TYPE_jpe,
	}
	if !docAttachment.FileHeaderContainsMimes(file_h, mimes) {
		ret_error = osbe.NewPublicMethodError(ER_UNSUPPORTED_MIME_CODE, ER_UNSUPPORTED_MIME_DESCR)
		return ret_error
	}

	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetPrimary()
	if err_с != nil {
		ret_error = osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register Request.GetPrimary(): %v",err_с))
		return ret_error
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	//tel
	if !vld_tel.Check(args.Tel.GetValue()) {
		ret_error = osbe.NewPublicMethodError(RESP_ER_REGISTER_TEL_CODE, RESP_ER_REGISTER_TEL_DESCR)
		return ret_error
	}

	//email
	if !vld_email.Check(args.Email.GetValue()) {
		ret_error = osbe.NewPublicMethodError(RESP_ER_REGISTER_EMAIL_CODE, RESP_ER_REGISTER_EMAIL_DESCR)
		return ret_error
	}
	
	
	//bik
	if !vld_bank.CheckBik(args.Banks_ref.Keys.Bik, conn, context.Background()) {
		ret_error = osbe.NewPublicMethodError(RESP_ER_REGISTER_BIK_CODE, RESP_ER_REGISTER_BIK_DESCR)
		return ret_error		
	}
	//acc
	if !vld_bank.CheckAcc(args.Banks_ref.Keys.Bik, args.Bank_acc.GetValue()) {
		ret_error = osbe.NewPublicMethodError(RESP_ER_REGISTER_ACC_CODE, RESP_ER_REGISTER_ACC_DESCR)
		return ret_error		
	}
	
	//user name (tel) should be unique
	found := false
	if err := conn.QueryRow(context.Background(),
		`SELECT TRUE FROM users AS u WHERE u.name = $1`,
		args.Tel.GetValue(),
	).Scan(&found); err != nil && err != pgx.ErrNoRows {		
		ret_error = osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register  pgx.QueryRow() SELECT name check failed: %v",err))
		return ret_error
	}
	if found {
		ret_error = osbe.NewPublicMethodError(RESP_ER_REGISTER_TEL_EXISTS_CODE, RESP_ER_REGISTER_TEL_EXISTS_DESCR)
		return ret_error
	}
	
	//inn present, not registered
	scan_model := &models.SpecialistDialog{}
	fields, field_ids, err := osbe.MakeStructRowFields(scan_model, "sp.")
	if err != nil {
		ret_error = osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register osbe.MakeStructRowFields(): %v",err))
		return ret_error		
	}
	if err := conn.QueryRow(context.Background(),
		fmt.Sprintf(`SELECT %s
		FROM specialists_dialog AS sp
		WHERE sp.inn = $1`, field_ids),
		args.Inn.GetValue(),
	).Scan(fields...); err != nil && err != pgx.ErrNoRows {		
		ret_error = osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register  pgx.QueryRow() specialists_dialog failed: %v",err))
		return ret_error
		
	}else if err == nil {	
		if scan_model.Last_status_type.GetValue() == "contract_signed" || scan_model.Last_status_type.GetValue() == "contract_signing" {
			//inn exists, contract signed
			ret_error = osbe.NewPublicMethodError(RESP_ER_REGISTER_EXISTS_CODE, RESP_ER_REGISTER_EXISTS_DESCR)
			return ret_error
		}
		//inn exists, contract not signed
		m := &model.Model{ID: model.ModelID("SpecialistDialog"), Rows: make([]model.ModelRow, 1)}	
		m.Rows[0] = scan_model
		resp.AddModel(m)	
		return nil		
	}
	
	if _, err := conn.Exec(context.Background(), "BEGIN"); err != nil {
		ret_error = osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register  conn.Exec() BEGIN failed: %v",err))
		return ret_error
	}
	userOperation.StartUserOperation(conn, REG_OPERATION_ID, 0, args.Operation_id.GetValue())		
	
	var specialist_reg_id int
	if _, err := conn.Exec(context.Background(),		
		`DELETE FROM specialist_regs WHERE user_operation_id=$1 OR inn = $2`,
		args.Operation_id.GetValue(),
		args.Inn.GetValue(),			
	); err != nil {
		conn.Exec(context.Background(), "ROLLBACK")
		ret_error = osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register  conn.Exec() DELETE specialist_regs failed: %v",err))
		return ret_error			
	}
	if err := conn.QueryRow(context.Background(), 
		`INSERT INTO specialist_regs
		(user_operation_id, 
		inn,
		banks_ref,
		bank_acc,
		name_full,
		name,
		email,
		tel,
		birthdate,
		address_reg,
		passport)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			jsonb_build_object(
				'series', $11::text,
				'num', $12::text,
				'issue_body', $13::text,
				'issue_date', $14::date,
				'dep_code', $15::text
			)
		) RETURNING id`,
		args.Operation_id.GetValue(),
		args.Inn.GetValue(),
		args.Banks_ref,
		args.Bank_acc.GetValue(),
		args.Name_full.GetValue(),
		args.Name.GetValue(),
		args.Email.GetValue(),
		args.Tel.GetValue(),
		args.Birthdate.GetValue(),
		args.Address_reg.GetValue(),
		args.Passport.Series.GetValue(),
		args.Passport.Num.GetValue(),
		args.Passport.Issue_body.GetValue(),
		args.Passport.Issue_date.GetValue(),
		args.Passport.Dep_code.GetValue(),
	).Scan(&specialist_reg_id); err != nil {
		conn.Exec(context.Background(), "ROLLBACK")
		ret_error = osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register  conn.QueryRow() INSERT specialist_regs failed: %v",err))
		return ret_error			
	}

	//attachment
	/*
	file_buf := new(bytes.Buffer)
	file_buf.ReadFrom(file)	
	ref := docAttachment.Ref_Type{DataType: "specialist_regs", Keys: docAttachment.AttachmentKey{Id: docAttachment.HttpInt(specialist_reg_id)}}
	cont_info := docAttachment.Content_info_Type{Name: file_h.Filename, Id: args.Operation_id.GetValue(), Size: int64(file_buf.Len())}
	if err := docAttachment.StoreAttachment(conn, &ref, &cont_info, file_buf.Bytes(), []byte{}); err != nil {
		ret_error = osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register  docAttachment.StoreAttachment() failed: %v",err))
		return ret_error
	}
	*/
	
	if _, err := conn.Exec(context.Background(), "COMMIT"); err != nil {
		ret_error = osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register  conn.Exec() COMMIT failed: %v",err))
		return ret_error
	}
	
	//event to admin
	EmitEvent(conn, "SpecialistReg.insert", 0)
	
	go register_specialist(app, d_store, args, file, file_h, specialist_reg_id)
	
	return nil
}

//async call
func (pm *Specialist_Controller_send_reg_documents) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	args := rfltArgs.Interface().(*models.DocAsync_id)	

	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetPrimary()
	if err_с != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_send_reg_documents Request.GetPrimary(): %v",err_с))
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	sess := sock.GetSession()	
	user_id := sess.GetInt("USER_ID")
	userOperation.StartUserOperation(conn, SEND_REG_DOCS_OPER, user_id, args.Operation_id.GetValue())		
		
	go send_reg_documents(app, args.Id.GetValue(), args.Operation_id.GetValue(), user_id) //specialist_utils
	return nil
}

//async call
func (pm *Specialist_Controller_send_salary_documents) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	args := rfltArgs.Interface().(*models.DocAsync_id)	

	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetPrimary()
	if err_с != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_send_salary_documents Request.GetPrimary(): %v",err_с))
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	sess := sock.GetSession()	
	user_id := sess.GetInt("USER_ID")
	userOperation.StartUserOperation(conn, SEND_SAL_DOCS_OPER, user_id, args.Operation_id.GetValue())		
		
	go send_salary_documents(app, args.Id.GetValue(), args.Operation_id.GetValue(), user_id) //specialist_utils
	return nil
}

