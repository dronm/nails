package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 13/11/2023
 *
 * This is SpecialistStatus controller implimentation file.
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
func (pm *SpecialistStatus_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistStatus"], &models.SpecialistStatus_keys{}, sock.GetPresetFilter("SpecialistStatus"))	
}

//Method implemenation delete
func (pm *SpecialistStatus_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistStatus"], sock.GetPresetFilter("SpecialistStatus"))	
}

//Method implemenation get_object
func (pm *SpecialistStatus_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistStatus"], &models.SpecialistStatus{}, sock.GetPresetFilter("SpecialistStatus"))	
}

//Method implemenation get_list
func (pm *SpecialistStatus_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistStatus"], &models.SpecialistStatus{}, sock.GetPresetFilter("SpecialistStatus"))	
}

//Method implemenation update
func (pm *SpecialistStatus_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistStatus"], sock.GetPresetFilter("SpecialistStatus"))	
}


