package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 28/11/2023
 *
 * This is SpecialistWork controller implimentation file.
 *
 */

import (
	"reflect"	
	"context"
	"fmt"
	"strings"
	
	"github.com/dronm/ds/pgds"
	
	"nails/models"
	
	"osbe"
	"osbe/srv"
	"osbe/socket"
	"osbe/response"	
		
	//"github.com/jackc/pgx/v4"
)

//Method implemenation insert
func (pm *SpecialistWork_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistWork"], &models.SpecialistWork_keys{}, sock.GetPresetFilter("SpecialistWork"))	
}

//Method implemenation delete
func (pm *SpecialistWork_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistWork"], sock.GetPresetFilter("SpecialistWork"))	
}

//Method implemenation get_object
func (pm *SpecialistWork_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistWorkList"], &models.SpecialistWorkList{}, sock.GetPresetFilter("SpecialistWorkList"))	
}

//Method implemenation get_list
func (pm *SpecialistWork_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistWorkList"], &models.SpecialistWorkList{}, sock.GetPresetFilter("SpecialistWorkList"))	
}

//Method implemenation get_for_rate_list
func (pm *SpecialistWork_Controller_get_for_rate_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistWorkForRateList"], &models.SpecialistWorkForRateList{}, sock.GetPresetFilter("SpecialistWorkForRateList"))	
}

//Method implemenation update
func (pm *SpecialistWork_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistWork"], sock.GetPresetFilter("SpecialistWork"))	
}

//Method implemenation set_admin_rates
func (pm *SpecialistWork_Controller_set_admin_rates) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	pool_conn, conn_id, err_с := d_store.GetPrimary()
	if err_с != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("specialist_reg_register_continue d_store.GetPrimary() failed: %v", err_с))
	}
	defer d_store.Release(pool_conn, conn_id)
	conn := pool_conn.Conn()
	
	args := rfltArgs.Interface().(*models.SpecialistWork_set_admin_rates)
	tmpl := `UPDATE specialist_works SET admin_rate=%d WHERE id=%d`
	var q strings.Builder
	for i, rate := range args.Rates {
		if i > 0 {
			q.WriteString("; ")
		}
		q.WriteString(fmt.Sprintf(tmpl, rate.Admin_rate, rate.Id))
	}
fmt.Println(q.String())	
	if _, err := conn.Exec(context.Background(),
		q.String(),
	); err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SpecialistWork_Controller_set_admin_rates conn.Exec() faled: %v",err))
	}
	return nil
}

