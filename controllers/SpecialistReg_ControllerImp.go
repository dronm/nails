package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 21/11/2023
 *
 * This is SpecialistReg controller implimentation file.
 *
 */

import (
	"reflect"	
	"context"
	"fmt"
	"os"
	"io/ioutil"
	"errors"
	"time"
	
	"nails/models"
	"nails/nails_config"
	
	"github.com/dronm/ds/pgds"
	
	"osbe"
	"osbe/model"
	"osbe/srv"
	"osbe/socket"
	"osbe/response"	
	"osbe/srv/httpSrv"
	
	"github.com/jackc/pgx/v4"
)

const (
	ER_REGISTERED = 1001
	ER_REGISTERED_DESCR = "Уже зарегистрирован или неверный идентификатор"	
	
	SEND_DOCS_OPERATION_ID = "register"
)

//Method implemenation insert
func (pm *SpecialistReg_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistReg"], &models.SpecialistReg_keys{}, sock.GetPresetFilter("SpecialistReg"))	
}

//Method implemenation delete
func (pm *SpecialistReg_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistReg"], sock.GetPresetFilter("SpecialistReg"))	
}

//Method implemenation get_object
func (pm *SpecialistReg_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistRegDialog"], &models.SpecialistRegDialog{}, sock.GetPresetFilter("SpecialistRegDialog"))	
}

//Method implemenation get_list
func (pm *SpecialistReg_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistRegList"], &models.SpecialistRegList{}, sock.GetPresetFilter("SpecialistRegList"))	
}

//Method implemenation update
func (pm *SpecialistReg_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistReg"], sock.GetPresetFilter("SpecialistReg"))	
}

//Method implemenation get_by_operation_id
func (pm *SpecialistReg_Controller_get_by_operation_id) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetSecondary("")
	if err_с != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register Request.GetPrimary(): %v",err_с))
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	args := rfltArgs.Interface().(*models.SpecialistReg_get_by_operation_id)
	
	model_md := app.GetMD().Models["SpecialistRegDialog"]
	scan_model := &models.SpecialistRegDialog{}
	fields, field_ids, err := osbe.MakeStructRowFields(scan_model, "")
	if err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_get_by_operation_id osbe.MakeStructRowFields(): %v",err))
	}
	if err := conn.QueryRow(context.Background(),
		fmt.Sprintf(`SELECT %s
			FROM specialist_regs_dialog
			WHERE user_operation_id = $1`, field_ids),
		args.Operation_id.GetValue(),
	).Scan(fields...); err != nil && err != pgx.ErrNoRows {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_get_by_operation_id conn.QueryRow() SELECT: %v",err))
		
	}else if err != nil && err == pgx.ErrNoRows {
		return osbe.NewPublicMethodError(ER_REGISTERED, ER_REGISTERED_DESCR)
	}
	m := &model.Model{ID: model.ModelID(model_md.ID),
			Rows: make([]model.ModelRow, 1),
			Metadata: model_md,
	}
	m.Rows[0] = scan_model
	resp.AddModel(m)
	
	return nil
}

func (pm *SpecialistReg_Controller_register) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetPrimary()
	if err_с != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_register Request.GetPrimary(): %v",err_с))
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
		
	var user_id int64
	var specilalist_id int64	
	var contact_id int64
	
	args := rfltArgs.Interface().(*models.SpecialistReg_register)
	pwd, hash := GenUserPwd()
	
	if _, err := conn.Exec(context.Background(), `BEGIN`); err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_register conn.Exec() BEGIN faled: %v",err))
	}
	
	//add user account	
	if err := conn.QueryRow(context.Background(),
		`INSERT INTO users
		(name, name_full, role_id, pwd, time_zone_locale_id, locale_id)
		SELECT
			reg.name,
			reg.name_full,
			'specialist',
			$1,
			(SELECT id FROM time_zone_locales WHERE name='Asia/Yekaterinburg'),
			'ru'
		FROM specialist_regs AS reg
		WHERE reg.id = $2
		ON CONFLICT (lower(name)) DO UPDATE SET
			name_full = excluded.name_full,
			role_id = 'specialist',
			pwd = $1
		RETURNING id`,
		hash,
		args.Id.GetValue(),
	).Scan(&user_id); err != nil {
		if _, err := conn.Exec(context.Background(), `ROLLBACK`); err != nil {
			app.GetLogger().Errorf("SpecialistReg_Controller_register conn.Exec() ROLLBACK INSERT INTO users faled: %v",err)
		}
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_register conn.QueryRow() INSERT INTO users faled: %v",err))
	}	
	
	//add specialist
	//speciality_id+agent_percent set in trigger!
	if err := conn.QueryRow(context.Background(),
		`INSERT INTO specialists
		(name, inn, bank_bik, bank_acc, studio_id, birthdate, address_reg, passport, equipments, user_id, ycl_staff_id)
		SELECT
			reg.name_full,
			reg.inn,
			reg.banks_ref->'keys'->>'bik',
			reg.bank_acc,
			$1,
			reg.birthdate,
			reg.address_reg,
			reg.passport,
			(SELECT st.equipments FROM studios AS st WHERE st.id = $1),
			$2,
			$3
		FROM specialist_regs AS reg
		WHERE reg.id = $4
		ON CONFLICT (inn) DO UPDATE SET studio_id = $1
		RETURNING id`,
		args.Studio_id.GetValue(),
		user_id,
		args.Ycl_staff_id.GetValue(),
		args.Id.GetValue(),		
	).Scan(&specilalist_id); err != nil {
		if _, err := conn.Exec(context.Background(), `ROLLBACK`); err != nil {
			app.GetLogger().Errorf("SpecialistReg_Controller_register conn.Exec() ROLLBACK INSERT INTO specialists faled: %v",err)
		}
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_register conn.QueryRow() INSERT INTO specialists faled: %v",err))
	}	
	
	//add specialist status	
	if _,err := conn.Exec(context.Background(),
		`DELETE FROM specialist_statuses WHERE specialist_id = $1`,
		specilalist_id,
	); err != nil {
		if _, err := conn.Exec(context.Background(), `ROLLBACK`); err != nil {
			app.GetLogger().Errorf("SpecialistReg_Controller_register conn.Exec() ROLLBACK DELETE FROM specialist_statuses faled: %v",err)
		}
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_register conn.Exec() DELETE FROM specialist_statuses faled: %v",err))
	}	
	if _,err := conn.Exec(context.Background(),
		`INSERT INTO specialist_statuses
		(specialist_id, status_type)
		VALUES ($1, 'contract_signing'::specialist_status_types)
		`,
		specilalist_id,
	); err != nil {
		if _, err := conn.Exec(context.Background(), `ROLLBACK`); err != nil {
			app.GetLogger().Errorf("SpecialistReg_Controller_register conn.Exec() ROLLBACK INSERT INTO specialist_statuses faled: %v",err)
		}
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_register conn.Exec() INSERT INTO specialist_statuses faled: %v",err))
	}	
	
	//add contact
	if err := conn.QueryRow(context.Background(),
		`INSERT INTO contacts
		(name, email, tel,
		email_confirmed, tel_confirmed)
		SELECT 
			sp.name,
			sp.email,
			sp.tel,
			TRUE,
			TRUE
		FROM specialist_regs AS sp
		WHERE sp.id = $1
		ON CONFLICT (tel) DO UPDATE SET name = excluded.name
		RETURNING id`,
		args.Id.GetValue(),
	).Scan(&contact_id); err != nil {
		if _, err := conn.Exec(context.Background(), `ROLLBACK`); err != nil {
			app.GetLogger().Errorf("SpecialistReg_Controller_register conn.Exec() ROLLBACK INSERT INTO contacts faled: %v",err)
		}
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_register conn.Exec() INSERT INTO contacts faled: %v",err))
	}		
	//entity_contacts specialists
	if _,err := conn.Exec(context.Background(),
		`INSERT INTO entity_contacts
		(entity_type, entity_id, contact_id)
		VALUES ('specialists', $1, $2)
		ON CONFLICT(entity_type, entity_id, contact_id) DO UPDATE SET contact_id = excluded.contact_id
		`,
		specilalist_id,
		contact_id,
	); err != nil {
		if _, err := conn.Exec(context.Background(), `ROLLBACK`); err != nil {
			app.GetLogger().Errorf("SpecialistReg_Controller_register conn.Exec() ROLLBACK INSERT INTO entity_contacts specialists faled: %v",err)
		}
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_register conn.Exec() INSERT INTO entity_contacts specialists faled: %v",err))
	}	
	//entity_contacts users
	if _,err := conn.Exec(context.Background(),
		`INSERT INTO entity_contacts
		(entity_type, entity_id, contact_id)
		VALUES ('users', $1, $2)
		ON CONFLICT(entity_type, entity_id, contact_id) DO UPDATE SET contact_id = excluded.contact_id
		`,
		user_id,
		contact_id,
	); err != nil {
		if _, err := conn.Exec(context.Background(), `ROLLBACK`); err != nil {
			app.GetLogger().Errorf("SpecialistReg_Controller_register conn.Exec() ROLLBACK INSERT INTO entity_contacts users faled: %v",err)
		}
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_register conn.Exec() INSERT INTO entity_contacts users faled: %v",err))
	}	
	
	if _, err := conn.Exec(context.Background(), `COMMIT`); err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_register conn.Exec() COMMIT faled: %v",err))
	}
	
	//notify specialist about new account, documents
	//document_templates.go
	go specialist_reg_register_continue(app, args.Operation_id.GetValue(), pwd, user_id, specilalist_id, args.Id.GetValue())
	
	return nil
}

//Печать Заявления
func (pm *SpecialistReg_Controller_print_app) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	sock_http, ok := sock.(*httpSrv.HTTPSocket)
	if !ok {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, "SpecialistReg_Controller_print_app sock must be *HTTPSocket")
	}
	
	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetSecondary("")
	if err_с != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_print_app Request.GetSecondary(): %v",err_с))
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()

	args := rfltArgs.Interface().(*models.Doc_id)

	//
	doc := TemplateDocument{}
	if err := conn.QueryRow(context.Background(),
		`SELECT
			att.id,
			att.ref,
			att.content_info->>'id',
			document_templates_spec_reg_exec_query($1::int, tmpl.sql_query::text)
			
		FROM document_templates AS tmpl
		LEFT JOIN attachments AS att ON (att.ref->'keys'->>'id')::int = tmpl.id AND att.ref->>'dataType' = 'document_templates'
		WHERE tmpl.name = 'Заявление'
		LIMIT 1`,
		args.Id.GetValue(),
	).Scan(&doc.AttID,
		&doc.Ref,
		&doc.ContentID,
		&doc.Values,
	); err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_print_app conn.Query() failed: %v",err))
	}	
	conf := app.GetConfig().(*nails_config.NailsConfig)
	_, out_doc, out_pdf, err := ConvertDocFromTemplate(app.GetBaseDir(), conf, &doc, conn, true)
	if err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistReg_Controller_print_app convert_doc() failed: %v",err))
	}
	os.Remove(out_doc)
	defer os.Remove(out_pdf)
	
	out_pdf_bt, err := ioutil.ReadFile(out_pdf)
	if err != nil {
		return errors.New(fmt.Sprintf("ioutil.ReadFile() on out_pdf file failed: %v", err))
	}		
	
	httpSrv.ServeContent(sock_http, &out_pdf_bt, "Заявление.pdf", "", time.Time{}, httpSrv.CONTENT_DISPOSITION_INLINE)		
	
	return nil
}

