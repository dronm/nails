package controllers

import(
	"fmt"
	"sync"
	"context"
	"errors"
	"mime/multipart"
	
	"nails/models"
	"nails/nails_config"
	
	"github.com/dronm/fnsnpd"	
	"github.com/dronm/ds/pgds"
	
	"osbe"
	
	notif "osbe/repo/notif"
	"osbe/repo/docAttachment"
	"osbe/repo/userOperation"
	
	"github.com/jackc/pgx/v4"
	
)

//Регистрация
func register_specialist(app osbe.Applicationer, dStore *pgds.PgProvider, args *models.Specialist_register, file multipart.File, file_h *multipart.FileHeader, specialistRegId int) {
	defer file.Close()
	
	var wg sync.WaitGroup
	inn_fns_ok := false
	inn_checked := false
	passport_uploaded := false
	email_sent := false
	tel_sent := false

	//FL fns
	wg.Add(1)
	go func() {
		defer wg.Done()
		pool_conn, conn_id, err_с := dStore.GetPrimary()
		if err_с != nil {
			app.GetLogger().Errorf("dStore.GetPrimary() inn failed: %v", err_с)	
			return
		}
		defer dStore.Release(pool_conn, conn_id)
		conn := pool_conn.Conn()
		
		error_s := ""
		defer (func(){
			userOperation.ProgressEvent(conn, args.Operation_id.GetValue(), "inn", error_s, inn_fns_ok)					
		})()
		ch := fnsnpd.PersonCheckerFNS.AddCheck(args.Inn.GetValue())
		inn_fns_ok = <-ch //wait	
		if !inn_fns_ok {
			error_s = "Не найден в ФНС"
		}
		inn_checked = true
	}()

	//attachment + sumbnail
	wg.Add(1)
	go func() {
		defer wg.Done()
		pool_conn, conn_id, err_с := dStore.GetPrimary()
		if err_с != nil {
			app.GetLogger().Errorf("dStore.GetPrimary() attachment failed: %v", err_с)	
			return
		}
		defer dStore.Release(pool_conn, conn_id)
		conn := pool_conn.Conn()
		
		ref := docAttachment.Ref_Type{DataType: "specialist_regs", Keys: docAttachment.AttachmentKey{Id: docAttachment.HttpInt(specialistRegId)}}
		cont_info := docAttachment.Content_info_Type{Name: file_h.Filename, Id: "passport"}		
		
		error_s := ""
		defer (func(){
			userOperation.ProgressEvent(conn, args.Operation_id.GetValue(), "passport_file_content", error_s, passport_uploaded)					
		})()
		_, err := docAttachment.AddFileThumbnailToDb(conn, app.GetBaseDir(), file, &cont_info, &ref);
		if err != nil {
			app.GetLogger().Errorf("docAttachment.AddFileThumbnailToDb failed: %v", err)				
			error_s = ER_INTERNAL
			return
		}
		passport_uploaded = true
	}()

	//send email
	wg.Add(1)
	go func() {
		defer wg.Done()

		pool_conn, conn_id, err_с := dStore.GetPrimary()
		if err_с != nil {
			app.GetLogger().Errorf("dStore.GetPrimary() email failed: %v", err_с)	
			return
		}
		defer dStore.Release(pool_conn, conn_id)
		conn := pool_conn.Conn()
		
		error_s := ""
		defer (func(){
			userOperation.ProgressEvent(conn, args.Operation_id.GetValue(), "email", error_s, email_sent)					
		})()
		
		email_key := genUniqID(7)		
		msg_email, err := notif.NewEmailMessageFromSQL(conn, "email_check", nil, nil, []interface{}{args.Email.GetValue(), args.Name_full.GetValue(), email_key})
		if err != nil {
			app.GetLogger().Errorf("notif.NewEmailMessageFromSQL() failed: %v", err)	
			error_s = ER_INTERNAL
			return
		}		
		notif_msg := msg_email.NewNotif("email_check")
		resp, err := NailsNotifier.Send([]*notif.NotifMessage{notif_msg})
		if err != nil {
			app.GetLogger().Errorf("notif.Send() email failed: %v", err)	
			error_s = ER_INTERNAL
			return
			
		}else if resp[0].Error != "" {
			error_s = resp[0].Error
			return
		}
		
		if _,err := conn.Exec(context.Background(),
			`INSERT INTO confirmation_status (ref, field, secret)
			VALUES (
				jsonb_build_object(
					'keys', jsonb_build_object('id', $1::int),
					'dataType', 'specialist_regs',
					'descr', $2::text
				),
				'email',
				$3
			)`,
			specialistRegId,
			args.Name_full.GetValue(),
			email_key,
		); err != nil {
			app.GetLogger().Errorf("conn.Exec() confirmation_status email failed: %v", err)	
			error_s = ER_INTERNAL
			return			
		}
		email_sent = true
	}()

	//send sms/wa message
	wg.Add(1)
	go func() {
		defer wg.Done()

		pool_conn, conn_id, err_с := dStore.GetPrimary()
		if err_с != nil {
			app.GetLogger().Errorf("dStore.GetPrimary() sms/wa failed: %v", err_с)	
			return
		}
		defer dStore.Release(pool_conn, conn_id)
		conn := pool_conn.Conn()
		
		error_s := ""
		defer (func(){
			userOperation.ProgressEvent(conn, args.Operation_id.GetValue(), "tel", error_s, tel_sent)					
		})()
		
		tel_key := genUniqNumberID(5)
		if err := NotifTelCheck(conn, args.Tel.GetValue(), tel_key); err != nil {
			app.GetLogger().Errorf("NotifTelCheck() failed: %v", err)	
			error_s = ER_INTERNAL
			return
		}
		if _, err := conn.Exec(context.Background(),
			`INSERT INTO confirmation_status (ref, field, secret)
			VALUES (
				jsonb_build_object(
					'keys', jsonb_build_object('id', $1::int),
					'dataType', 'specialist_regs',
					'descr', $2::text
				),
				'tel',
				$3
			)`,
			specialistRegId,
			args.Name_full.GetValue(),
			tel_key,
		); err != nil {
			app.GetLogger().Errorf("conn.Exec() confirmation_status tel failed: %v", err)	
			error_s = ER_SMS_SEND_INTERNAL
			return			
		}

		tel_sent = true
	}()

	wg.Wait()
	
	//end operation
	pool_conn, conn_id, err_с := dStore.GetPrimary()
	if err_с != nil {
		app.GetLogger().Errorf("dStore.GetPrimary() end operation failed: %v", err_с)	
		return
	}
	defer dStore.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	if _, err := conn.Exec(context.Background(),
		`UPDATE specialist_regs
		SET
			inn_checked = $1,
			inn_fns_ok = $2,
			passport_uploaded = $3,
			email_sent = $4,
			tel_sent = $5
		WHERE id = $6`,
		inn_checked,
		inn_fns_ok,
		passport_uploaded,
		email_sent,
		tel_sent,
		specialistRegId,
	); err != nil {
		app.GetLogger().Errorf("conn.Exec() specialist_regs UPDATE inn failed: %v", err)	
	}
				
	if !inn_fns_ok || !passport_uploaded || !email_sent && !tel_sent {		
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, 0, args.Operation_id.GetValue(), errors.New("Завершено с ошибкой"))
	}else{
		//to admin - wa message
		//refresh grid
		EmitEvent(conn, "SpecialistReg.update", int64(specialistRegId))
		
		if err := NotifNewSpecialist(conn); err != nil {
			//log only
			app.GetLogger().Errorf("NotifNewSpecialist() failed: %v", err)	
		}
	}
}


//Отправка пакета регистрационных документов
func send_reg_documents(app osbe.Applicationer, specialistID int64, operationID string, userID int64) {
	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetPrimary()
	if err_с != nil {
		app.GetLogger().Errorf("send_reg_documents d_store.GetPrimary() failed: %v", err_с)	
		return
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	_, _, err := make_reg_documents(app, conn, specialistID)
	if err != nil{
		app.GetLogger().Errorf("send_reg_documents make_reg_documents() failed: %v", err)	
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, err)
		return
	}
	
	mag_type:= "docs_for_sign"
	//wa meassage
	wa_msg, err := notif.NewWAMessageFromSQL(conn, "wa_" + mag_type, []interface{}{specialistID})	
	if err != nil {
		app.GetLogger().Errorf("send_reg_documents notif.NewWAMessageFromSQL() failed: %v", err)	
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, err)
		return
	}	
	
	//email message
	/*
	Такое не отправляем письмом, только в ЛК!	 
	e_msg, err := notif.NewEmailMessageFromSQL(conn, "email_" + mag_type, pdf_list, alias_list, []interface{}{specialistID})	
	if err != nil {
		app.GetLogger().Errorf("send_reg_documents notif.NewEmailMessageFromSQL() failed: %v", err)	
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, err)
		return
	}
	, e_msg.NewNotif(mag_type)
	*/
	//send all
	if err := NotifSend([]*notif.NotifMessage{wa_msg.NewNotif(mag_type)}); err != nil {
		app.GetLogger().Errorf("send_reg_documents NotifSend() failed: %v", err)	
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, err)
		return
	}
	userOperation.EndUserOperation(app.GetLogger(), conn, userID, operationID)
}

//Отправка пакета документов по з/п
func send_salary_documents(app osbe.Applicationer, salaryDetailID int64, operationID string, userID int64) {
	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetPrimary()
	if err_с != nil {
		app.GetLogger().Errorf("send_salary_documents d_store.GetPrimary() failed: %v", err_с)	
		return
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	var specialist_id int64
	if err := conn.QueryRow(context.Background(),
		`SELECT
			specialist_id
		FROM specialist_period_salary_details
		WHERE id = $1`,
		salaryDetailID,
	).Scan(&specialist_id); err != nil {
		app.GetLogger().Errorf("send_salary_documents conn.QueryRow() failed: %v", err)	
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, err)
		return
	}
	
	query := `SELECT
			att.id,
			att.ref,
			coalesce(att.content_info->>'id', ''),
			coalesce(tmpl.need_signing, FALSE),
			coalesce(tmpl.name, ''),
			tmpl.id,
			coalesce(document_templates_spec_sal_exec_query($1::int, tmpl.sql_query::text), '{}'::jsonb)
		FROM template_batches AS tmpl_b
		LEFT JOIN template_batch_items AS tmpl_b_it ON tmpl_b_it.template_batch_id = tmpl_b.id
			AND tmpl_b_it.studio_id IS NULL
		LEFT JOIN document_templates AS tmpl ON tmpl.id = tmpl_b_it.template_id
		LEFT JOIN attachments AS att ON (att.ref->'keys'->>'id')::int = tmpl.id AND att.ref->>'dataType' = 'document_templates'
		WHERE tmpl_b.template_batch_type = 'specialist_salary'::template_batch_types`
		
	if _, _, err := make_documents(app, conn, specialist_id, query, []interface{}{salaryDetailID}); err != nil{
		app.GetLogger().Errorf("send_salary_documents make_documents() failed: %v", err)	
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, err)
		return
	}
	
	mag_type:= "docs_for_sign"
	//wa meassage
	wa_msg, err := notif.NewWAMessageFromSQL(conn, "wa_" + mag_type, []interface{}{specialist_id})	
	if err != nil {
		app.GetLogger().Errorf("send_salary_documents notif.NewWAMessageFromSQL() failed: %v", err)	
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, err)
		return
	}	
	if err := NotifSend([]*notif.NotifMessage{wa_msg.NewNotif(mag_type)}); err != nil {
		app.GetLogger().Errorf("send_salary_documents NotifSend() failed: %v", err)	
		userOperation.EndUserOperationWithError(app.GetLogger(), conn, userID, operationID, err)
		return
	}
	userOperation.EndUserOperation(app.GetLogger(), conn, userID, operationID)
}

//returns file list, alias list - convinient structs for sending
func make_reg_documents(app osbe.Applicationer, conn *pgx.Conn, specialistID int64) ([]string, []string, error){
	query := `SELECT
			att.id,
			att.ref,
			coalesce(att.content_info->>'id', ''),
			coalesce(tmpl.need_signing, FALSE),
			coalesce(tmpl.name, ''),
			tmpl.id,
			coalesce(document_templates_spec_reg_exec_query($1::int, tmpl.sql_query::text), '{}'::jsonb)
		FROM template_batches AS tmpl_b
		LEFT JOIN specialists AS sp ON sp.id = $1
		LEFT JOIN template_batch_items AS tmpl_b_it ON tmpl_b_it.template_batch_id = tmpl_b.id
			AND (tmpl_b_it.studio_id IS NULL OR tmpl_b_it.studio_id=sp.studio_id)
		LEFT JOIN document_templates AS tmpl ON tmpl.id = tmpl_b_it.template_id
		LEFT JOIN attachments AS att ON (att.ref->'keys'->>'id')::int = tmpl.id AND att.ref->>'dataType' = 'document_templates'
		WHERE tmpl_b.template_batch_type = 'specialist_registration'::template_batch_types`
	return make_documents(app, conn, specialistID, query, []interface{}{specialistID})
}

//returns file list, alias list - convinient structs for sending
func make_documents(app osbe.Applicationer, conn *pgx.Conn, specialistID int64, query string, queryParams []interface{}) ([]string, []string, error){
	rows, err := conn.Query(context.Background(),
		query,
		queryParams...,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("make_reg_documents conn.Query() template_batches failed: %v", err)
	}	
	
	conv_docs := make([]*TemplateDocument, 0)	
	for rows.Next() {
		conv_docum := &TemplateDocument{}
		if err := rows.Scan(&conv_docum.AttID,
				&conv_docum.Ref,
				&conv_docum.ContentID,
				&conv_docum.NeedSigning,
				&conv_docum.TemplateName,
				&conv_docum.TemplateID,
				&conv_docum.Values,
		); err != nil {
			return nil, nil, fmt.Errorf("make_reg_documents rows.Next() template_batches failed: %v", err)
		}
		conv_docs = append(conv_docs, conv_docum)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("make_reg_documents rows.Err() template_batches failed: %v", err)
	}
	rows.Close()

	//converting
	file_list := make([]string, len(conv_docs))
	alias_list := make([]string, len(conv_docs))
	conf := app.GetConfig().(*nails_config.NailsConfig)
	base_dir := app.GetBaseDir()
	for i,doc := range conv_docs {		
		out_pdf, err := convert_doc(base_dir, conf, doc, conn, specialistID)
		if err != nil {
			return nil, nil, fmt.Errorf("make_reg_documents convert_doc() failed: %v", err)
		}
		file_list[i] = out_pdf
		alias_list[i] = doc.TemplateName + ".pdf"
	}
	
	return file_list, alias_list, nil
}
