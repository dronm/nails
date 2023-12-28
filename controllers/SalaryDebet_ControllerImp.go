package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 12/12/2023
 *
 * This is SalaryDebet controller implimentation file.
 *
 */

import (
	"reflect"	
	
	"nails/models"
	
	"osbe"
	"osbe/srv"
	"osbe/socket"
	"osbe/response"	
	
	//"github.com/jackc/pgx/v4"
)

//Method implemenation insert
func (pm *SalaryDebet_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SalaryDebet"], &models.SalaryDebet_keys{}, sock.GetPresetFilter("SalaryDebet"))	
}

//Method implemenation delete
func (pm *SalaryDebet_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SalaryDebet"], sock.GetPresetFilter("SalaryDebet"))	
}

//Method implemenation get_object
func (pm *SalaryDebet_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["SalaryDebet"], &models.SalaryDebet{}, sock.GetPresetFilter("SalaryDebet"))	
}

//Method implemenation get_list
func (pm *SalaryDebet_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SalaryDebet"], &models.SalaryDebet{}, sock.GetPresetFilter("SalaryDebet"))	
}

//Method implemenation update
func (pm *SalaryDebet_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SalaryDebet"], sock.GetPresetFilter("SalaryDebet"))	
}


