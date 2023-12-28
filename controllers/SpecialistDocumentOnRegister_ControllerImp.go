package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 26/11/2023
 *
 * This is SpecialistDocumentOnRegister controller implimentation file.
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
func (pm *SpecialistDocumentOnRegister_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistDocumentOnRegister"], &models.SpecialistDocumentOnRegister_keys{}, sock.GetPresetFilter("SpecialistDocumentOnRegister"))	
}

//Method implemenation delete
func (pm *SpecialistDocumentOnRegister_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistDocumentOnRegister"], sock.GetPresetFilter("SpecialistDocumentOnRegister"))	
}

//Method implemenation get_object
func (pm *SpecialistDocumentOnRegister_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistDocumentOnRegisterList"], &models.SpecialistDocumentOnRegisterList{}, sock.GetPresetFilter("SpecialistDocumentOnRegisterList"))	
}

//Method implemenation get_list
func (pm *SpecialistDocumentOnRegister_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["SpecialistDocumentOnRegisterList"], &models.SpecialistDocumentOnRegisterList{}, sock.GetPresetFilter("SpecialistDocumentOnRegisterList"))	
}

//Method implemenation update
func (pm *SpecialistDocumentOnRegister_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["SpecialistDocumentOnRegister"], sock.GetPresetFilter("SpecialistDocumentOnRegister"))	
}


