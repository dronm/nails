package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 30/11/2023
 *
 * This is Speciality controller implimentation file.
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
func (pm *Speciality_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["Speciality"], &models.Speciality_keys{}, sock.GetPresetFilter("Speciality"))	
}

//Method implemenation delete
func (pm *Speciality_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["Speciality"], sock.GetPresetFilter("Speciality"))	
}

//Method implemenation get_object
func (pm *Speciality_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialityDialog"], &models.SpecialityDialog{}, sock.GetPresetFilter("SpecialityDialog"))	
}

//Method implemenation get_list
func (pm *Speciality_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialityList"], &models.SpecialityList{}, sock.GetPresetFilter("SpecialityList"))	
}

//Method implemenation update
func (pm *Speciality_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["Speciality"], sock.GetPresetFilter("Speciality"))	
}


