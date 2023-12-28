package controllers

//Разные функции по работе с шаблонами документов

import(
	"context"
	"os"
	"io"
	"net/http"
	"io/ioutil"
	"fmt"
	"bytes"
	"errors"
	
	"nails/nails_config"
	
	"github.com/dronm/ds/pgds"
	
	"osbe"
	
	"osbe/repo/userOperation"
	"osbe/repo/docAttachment"		
	
	"github.com/jackc/pgx/v4"
)

//Продолжение регистрации асинхронно
//Сформировать пакет документов, отправить на почту
//Отправить сообщение о регистрации (последним, чтобы не пришло раньше регистрации)
func specialist_reg_register_continue(app osbe.Applicationer, adminOperationID, userPwd string, userID int64, specialistID int64, specialistRegID int64) {	
	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetPrimary()
	if err_с != nil {
		app.GetLogger().Errorf("specialist_reg_register_continue d_store.GetPrimary() failed: %v", err_с)	
		return
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	userOperation.StartUserOperation(conn, SEND_DOCS_OPERATION_ID, userID, adminOperationID)
	
	//delete all previously written documents + attachments	
	if _, err := conn.Exec(context.Background(),
		`DELETE FROM specialist_documents WHERE specialist_id = $1`,
		specialistID,
	); err != nil {
		app.GetLogger().Errorf("specialist_reg_register_continue DELETE failed: %v", err)	
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, adminOperationID, err)
		return
	}
		
	pdf_list, pdf_name_list, err := make_reg_documents(app, conn, specialistID) //specialist.go
	if err != nil {
		app.GetLogger().Errorf("specialist_reg_register_continue make_reg_documents() failed: %v", err)	
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, adminOperationID, err)
		return
	}
	
	//reassign attachments (passport) from reg to specialist
	if _, err := conn.Exec(context.Background(),	
		`UPDATE attachments
		SET ref = jsonb_build_object(
				'keys', jsonb_build_object('id', $1::int),
				'dataType', 'specialists'				
			)
		WHERE (ref->'keys'->>'id')::int = $2 AND ref->>'dataType' = 'specialist_regs'`,		
		specialistID,
		specialistRegID,
	); err != nil {
		if _, err := conn.Exec(context.Background(), `ROLLBACK`); err != nil {
			app.GetLogger().Errorf("SpecialistReg_Controller_register conn.Exec() ROLLBACK UPDATE attachments faled: %v",err)
		}
		app.GetLogger().Errorf("specialist_reg_register_continue conn.Exec() UPDATE attachments failed: %v", err)			
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, adminOperationID, err)
		return				
	}
	
	//delete temp registration
	var user_operation_id string //register fl operation
	if err := conn.QueryRow(context.Background(),	
		`DELETE FROM specialist_regs WHERE id = $1 RETURNING user_operation_id`,		
		specialistRegID,
	).Scan(&user_operation_id); err != nil {
		if _, err := conn.Exec(context.Background(), `ROLLBACK`); err != nil {
			app.GetLogger().Errorf("SpecialistReg_Controller_register conn.Exec() ROLLBACK DELETE specialist_regs faled: %v",err)
		}	
		app.GetLogger().Errorf("specialist_reg_register_continue conn.Exec() DELETE specialist_regs failed: %v", err)			
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, adminOperationID, err)
		return				
	}
	
	if _, err := conn.Exec(context.Background(),`COMMIT`); err != nil{
		app.GetLogger().Errorf("specialist_reg_register_continue COMMIT attachments failed: %v", err)			
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, adminOperationID, err)
		return				
	}		

	//Отправим сообщения ФЛ о регистрации + пакет документов
	if err := NotifNewAccount(conn, specialistID, userPwd, pdf_list, pdf_name_list); err != nil { //notifs.go
		app.GetLogger().Errorf("specialist_reg_register_continue NotifNewAccount() failed: %v", err)			
	}

	//Перекиним аватарки из yclients
	avatar := ""
	avatar_big := ""
	if err := conn.QueryRow(context.Background(),	
		`SELECT
			l.avatar
			,l.avatar_big
		FROM ycl_staff_list AS l
		WHERE l.id = (SELECT ycl_staff_id FROM specialists WHERE id = $1)
		`,
		specialistID,
	).Scan(&avatar, &avatar_big); err != nil {
		app.GetLogger().Errorf("specialist_reg_register_continue conn.QueryRow() SELECT ycl_staff_list failed: %v", err)			
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, adminOperationID, err)
		return				
	}
	avatar_buf := bytes.Buffer{}
	avatar_big_buf := bytes.Buffer{}
	
	//errors are not critical	
	resp_avatar, err := http.Get(avatar)
	if err != nil {
		app.GetLogger().Errorf("specialist_reg_register_continue http.Get() avatar failed: %v", err)			
	}else{
		defer resp_avatar.Body.Close()
		if _, err = io.Copy(&avatar_buf, resp_avatar.Body); err != nil {
			app.GetLogger().Errorf("specialist_reg_register_continue io.Copy() avatar failed: %v", err)			
		}
	}
	resp_avatar_big, err := http.Get(avatar_big)
	if err != nil {
		app.GetLogger().Errorf("specialist_reg_register_continue http.Get() avatar_big failed: %v", err)			
	}else{
		defer resp_avatar_big.Body.Close()
		if _, err = io.Copy(&avatar_big_buf, resp_avatar_big.Body); err != nil {
			app.GetLogger().Errorf("specialist_reg_register_continue io.Copy() avatar_big failed: %v", err)			
		}
	}
	ref := docAttachment.NewRef_Type(userID, "users")
	file_info := docAttachment.NewContent_info_Type("photo", "photo")
	if err := docAttachment.StoreAttachment(conn, ref, file_info, avatar_big_buf.Bytes(), avatar_buf.Bytes()); err != nil {
		//not critical
		app.GetLogger().Errorf("specialist_reg_register_continue docAttachment.StoreAttachment() failed: %v", err)			
	}
	

	//end admin operation
	userOperation.EndUserOperation(app.GetLogger(), conn, userID, adminOperationID)
	
	//end user (fl) operation
	userOperation.EndUserOperation(app.GetLogger(), conn, 0, user_operation_id)	
}

//Функция подстановки данных в шаблон, конвертации документа в pdf
//Возвращает путь к файлу pdf
func convert_doc(baseDir string, conf *nails_config.NailsConfig, doc *TemplateDocument, conn *pgx.Conn, specialistID int64) (string, error) {
	_, out_doc, out_pdf, err := ConvertDocFromTemplate(baseDir, conf, doc, conn, true) //convert_doc_from_tmpl
	if err != nil {
		return "", err
	}
	defer os.Remove(out_doc)
	
	out_doc_bt, err := ioutil.ReadFile(out_doc) //В базу складываем замены, потом добавим подпись
	if err != nil {
		return "", errors.New(fmt.Sprintf("ioutil.ReadFile() on out_docf file failed: %v", err))
	}		
	
	ref := &docAttachment.Ref_Type{Keys: docAttachment.AttachmentKey{Id: docAttachment.HttpInt(specialistID)}, DataType: "specialists"}	
	//doc.TemplateName,
	out_doc_fi := &docAttachment.Content_info_Type{Id: genUniqID(12),
			Name: doc.TemplateName + ".docx",
			Size: int64(len(out_doc_bt)),			
		}	
	//doc.TemplateName + ".pdf"
	out_pdf_fi := &docAttachment.Content_info_Type{Id: genUniqID(12), Name: doc.TemplateName + ".pdf"}		
	out_pdf_preview_fn := docAttachment.GetPreviewCacheFileName(baseDir, ref.DataType, ref.Keys.Id, out_pdf_fi.Id)
	if err := docAttachment.GenThumbnail(out_pdf, out_pdf_preview_fn, out_pdf_fi.Name); err != nil {
		return "", errors.New(fmt.Sprintf("docAttachment.GenThumbnail() on out_pdf failed: %v", err))
	}
	out_pdf_bt, err := ioutil.ReadFile(out_pdf)
	if err != nil {
		return "", errors.New(fmt.Sprintf("ioutil.ReadFile() on out_pdf file failed: %v", err))
	}		
	out_pdf_prw_bt, err := ioutil.ReadFile(out_pdf_preview_fn)	
	if err != nil {
		return "", errors.New(fmt.Sprintf("ioutil.ReadFile() on out_pdf preview file failed: %v", err))
	}
	
	//----------------------------- Складываем все в базу данных
	var out_doc_att_id int
	var out_pdf_att_id int
	
	if _, err := conn.Exec(context.Background(), `BEGIN`); err != nil {
		return "", errors.New(fmt.Sprintf("conn.Exec() BEGIN failed: %v", err))
	}
	
	//out_doc
	if err := conn.QueryRow(context.Background(),
		`INSERT INTO attachments
		(ref, content_info, content_data)
		VALUES ($1, $2, $3)
		RETURNING id
		`,
			ref,
			out_doc_fi,	
			out_doc_bt,
	).Scan(&out_doc_att_id); err != nil {	
		conn.Exec(context.Background(), `ROLLBACK`)
		return "", errors.New(fmt.Sprintf("conn.QueryRow() on out_doc file failed: %v", err))	
	}

	//out_pdf + preview
	if err := conn.QueryRow(context.Background(),
		`INSERT INTO attachments
		(ref, content_info, content_data, content_preview)
		VALUES ($1, $2, $3, $4)
		RETURNING id
		`,
			ref,
			out_pdf_fi,	
			out_pdf_bt,
			out_pdf_prw_bt,
	).Scan(&out_pdf_att_id); err != nil {	
		conn.Exec(context.Background(), `ROLLBACK`)
		return "", errors.New(fmt.Sprintf("conn.QueryRow() on out_doc.pdf file failed: %v", err))	
	}
	
	if _, err := conn.Exec(context.Background(),
		`INSERT INTO specialist_documents
		(specialist_id, template_att_id, document_att_id, need_signing, name, document_template_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		`,
		specialistID,
		out_doc_att_id,
		out_pdf_att_id,
		doc.NeedSigning,
		doc.TemplateName,
		doc.TemplateID,
	); err != nil {
		conn.Exec(context.Background(), `ROLLBACK`)
		return "", errors.New(fmt.Sprintf("conn.Exec() INSERT specialist_documents failed: %v", err))	
	}

	if _, err := conn.Exec(context.Background(), `COMMIT`); err != nil {
		return "", errors.New(fmt.Sprintf("conn.Exec() COMMIT failed: %v", err))
	}
	
	return out_pdf, nil	
}


