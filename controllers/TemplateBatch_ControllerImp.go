package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 19/12/2023
 *
 * This is TemplateBatch controller implimentation file.
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
func (pm *TemplateBatch_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["TemplateBatch"], &models.TemplateBatch_keys{}, sock.GetPresetFilter("TemplateBatch"))	
}

//Method implemenation delete
func (pm *TemplateBatch_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["TemplateBatch"], sock.GetPresetFilter("TemplateBatch"))	
}

//Method implemenation get_object
func (pm *TemplateBatch_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["TemplateBatch"], &models.TemplateBatch{}, sock.GetPresetFilter("TemplateBatch"))	
}

//Method implemenation get_list
func (pm *TemplateBatch_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["TemplateBatch"], &models.TemplateBatch{}, sock.GetPresetFilter("TemplateBatch"))	
}

//Method implemenation update
func (pm *TemplateBatch_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["TemplateBatch"], sock.GetPresetFilter("TemplateBatch"))	
}


