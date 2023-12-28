package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 12/12/2023
 *
 * This is SpecialistSalaryDebet controller implimentation file.
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
func (pm *SpecialistSalaryDebet_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistSalaryDebet"], &models.SpecialistSalaryDebet_keys{}, sock.GetPresetFilter("SpecialistSalaryDebet"))	
}

//Method implemenation delete
func (pm *SpecialistSalaryDebet_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistSalaryDebet"], sock.GetPresetFilter("SpecialistSalaryDebet"))	
}

//Method implemenation get_object
func (pm *SpecialistSalaryDebet_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistSalaryDebetList"], &models.SpecialistSalaryDebetList{}, sock.GetPresetFilter("SpecialistSalaryDebetList"))	
}

//Method implemenation get_list
func (pm *SpecialistSalaryDebet_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistSalaryDebetList"], &models.SpecialistSalaryDebetList{}, sock.GetPresetFilter("SpecialistSalaryDebetList"))	
}

//Method implemenation update
func (pm *SpecialistSalaryDebet_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistSalaryDebet"], sock.GetPresetFilter("SpecialistSalaryDebet"))	
}


