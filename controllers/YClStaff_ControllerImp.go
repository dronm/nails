package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 05/12/2023
 *
 * This is YClStaff controller implimentation file.
 *
 */

import (
	"reflect"	
	
	"nails/models"
	
	"osbe"
	"osbe/srv"
	"osbe/socket"
	"osbe/response"	
	
	"github.com/dronm/ds/pgds"
	//"github.com/jackc/pgx/v4"
)

//Method implemenation insert
func (pm *YClStaff_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["YClStaff"], &models.YClStaff_keys{}, sock.GetPresetFilter("YClStaff"))	
}

//Method implemenation delete
func (pm *YClStaff_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["YClStaff"], sock.GetPresetFilter("YClStaff"))	
}

//Method implemenation get_object
func (pm *YClStaff_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["YClStaffList"], &models.YClStaffList{}, sock.GetPresetFilter("YClStaffList"))	
}

//Method implemenation get_list
func (pm *YClStaff_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["YClStaffList"], &models.YClStaffList{}, sock.GetPresetFilter("YClStaffList"))	
}

//Method implemenation update
func (pm *YClStaff_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["YClStaff"], sock.GetPresetFilter("YClStaff"))	
}

//Method implemenation api_get_all
func (pm *YClStaff_Controller_api_get_all) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	args := rfltArgs.Interface().(*models.YClStaff_api_get_all)
	
	sess := sock.GetSession()	
	user_id := sess.GetInt(SESS_VAR_ID)
	go YClients.UpdateStaff(args.Operation_id.GetValue(), user_id, app.GetDataStorage().(*pgds.PgProvider), app.GetLogger())
	
	return nil
}


