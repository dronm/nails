package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 24/11/2023
 *
 * This is DocumentTemplate controller implimentation file.
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
func (pm *DocumentTemplate_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["DocumentTemplate"], &models.DocumentTemplate_keys{}, sock.GetPresetFilter("DocumentTemplate"))	
}

//Method implemenation delete
func (pm *DocumentTemplate_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["DocumentTemplate"], sock.GetPresetFilter("DocumentTemplate"))	
}

//Method implemenation get_object
func (pm *DocumentTemplate_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["DocumentTemplateDialog"], &models.DocumentTemplateDialog{}, sock.GetPresetFilter("DocumentTemplateDialog"))	
}

//Method implemenation get_list
func (pm *DocumentTemplate_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["DocumentTemplateList"], &models.DocumentTemplateList{}, sock.GetPresetFilter("DocumentTemplateList"))	
}

//Method implemenation update
func (pm *DocumentTemplate_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["DocumentTemplate"], sock.GetPresetFilter("DocumentTemplate"))	
}


