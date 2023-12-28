package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 23/11/2023
 *
 * This is ConfirmationStatus controller implimentation file.
 *
 */

import (
	"reflect"	
	"fmt"
	"context"
	
	"nails/models"
	
	"github.com/dronm/ds/pgds"
	
	"osbe"
	"osbe/model"
	"osbe/srv"
	"osbe/socket"
	"osbe/response"	
	
	"osbe/repo/userOperation"
	
	"github.com/jackc/pgx/v4"
)

const (
	ER_NO_CONFIRM_SECRET = 1000
	ER_NO_CONFIRM_SECRET_DESCR = "Сообщение не найдено"
)

//Method implemenation set_confirmed
func (pm *ConfirmationStatus_Controller_set_confirmed) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	args := rfltArgs.Interface().(*models.ConfirmationStatus_set_confirmed)
	var secret string
	var ok bool
	secret, ok = (*args)["secret"]
	if !ok {
		for secret, _ = range *args {
			break
		}
	}
	if secret == "" {
		return osbe.NewPublicMethodError(ER_NO_CONFIRM_SECRET, ER_NO_CONFIRM_SECRET_DESCR)
	}
	
	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetPrimary()
	if err_с != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register Request.GetPrimary(): %v",err_с))
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
			
	var field string
	var specialist_reg_name string
	var specialist_reg_id int
	if err := conn.QueryRow(context.Background(),
		`UPDATE confirmation_status set confirmed = TRUE WHERE secret = $1 RETURNING field, ref->>'descr', (ref->'keys'->>'id')::int`,
		secret,
	).Scan(&field, &specialist_reg_name, &specialist_reg_id); err != nil && err != pgx.ErrNoRows {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("ConfirmationStatus_set_confirmed conn.Exec() UPDATE: %v",err))
		
	}else if err != nil && err == pgx.ErrNoRows {
		return osbe.NewPublicMethodError(ER_NO_CONFIRM_SECRET, ER_NO_CONFIRM_SECRET_DESCR)
	}
	
	//admin event
	EmitEvent(conn, "SpecialistReg.insert", 0)
	
	//event
	var user_operation_id string
	if err := conn.QueryRow(context.Background(),
		`SELECT
			coalesce(user_operation_id, '')
		FROM specialist_regs
		WHERE id = $1`,
		specialist_reg_id,
	).Scan(&user_operation_id); err != nil && err != pgx.ErrNoRows {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("ConfirmationStatus_set_confirmed conn.QueryRow() SELECT: %v",err))
	}

	userOperation.ProgressEvent(conn, user_operation_id, field + "_conf", "", true)					
	
	model_data := struct {
		Name_full string `json:"name_full"`
		Field string `json:"field"`
	}{specialist_reg_name, field}
	m := &model.Model{ID: model.ModelID("ConfirmationStatusResult"),
			Rows: make([]model.ModelRow, 1),
	}
	m.Rows[0] = &model_data
	resp.AddModel(m)
	
	return nil
}


