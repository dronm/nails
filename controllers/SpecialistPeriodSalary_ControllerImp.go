package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 12/12/2023
 *
 * This is SpecialistPeriodSalary controller implimentation file.
 *
 */

import (
	"reflect"	
	"context"
	"fmt"
	
	"nails/models"
	
	"osbe"
	"osbe/srv"
	"osbe/socket"
	"osbe/response"	
	
	"github.com/dronm/ds/pgds"
	//"github.com/jackc/pgx/v4"
)

//Method implemenation insert
func (pm *SpecialistPeriodSalary_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistPeriodSalary"], &models.SpecialistPeriodSalary_keys{}, sock.GetPresetFilter("SpecialistPeriodSalary"))	
}

//Method implemenation delete
func (pm *SpecialistPeriodSalary_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistPeriodSalary"], sock.GetPresetFilter("SpecialistPeriodSalary"))	
}

//Method implemenation get_object
func (pm *SpecialistPeriodSalary_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistPeriodSalaryList"], &models.SpecialistPeriodSalaryList{}, sock.GetPresetFilter("SpecialistPeriodSalaryList"))	
}

//Method implemenation get_list
func (pm *SpecialistPeriodSalary_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistPeriodSalaryList"], &models.SpecialistPeriodSalaryList{}, sock.GetPresetFilter("SpecialistPeriodSalaryList"))	
}

//Method implemenation update
func (pm *SpecialistPeriodSalary_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistPeriodSalary"], sock.GetPresetFilter("SpecialistPeriodSalary"))	
}

//fill_doc
func (pm *SpecialistPeriodSalary_Controller_fill_doc) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetPrimary()
	if err_с != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistPeriodSalary_Controller_fill_doc Request.GetPrimary(): %v",err_с))
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()

	args := rfltArgs.Interface().(*models.Doc_id)

	//
	if _, err := conn.Exec(context.Background(),
		`SELECT specialist_period_salaries_fill($1::int)`,
		args.Id.GetValue(),
	); err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistPeriodSalary_Controller_fill_doc conn.Exec() failed: %v",err))
	}	
	
	return nil
}

