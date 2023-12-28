package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 19/12/2023
 *
 * This is TemplateBatchItem controller implimentation file.
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
func (pm *TemplateBatchItem_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["TemplateBatchItem"], &models.TemplateBatchItem_keys{}, sock.GetPresetFilter("TemplateBatchItem"))	
}

//Method implemenation delete
func (pm *TemplateBatchItem_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["TemplateBatchItem"], sock.GetPresetFilter("TemplateBatchItem"))	
}

//Method implemenation get_object
func (pm *TemplateBatchItem_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["TemplateBatchItemList"], &models.TemplateBatchItemList{}, sock.GetPresetFilter("TemplateBatchItemList"))	
}

//Method implemenation get_list
func (pm *TemplateBatchItem_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["TemplateBatchItemList"], &models.TemplateBatchItemList{}, sock.GetPresetFilter("TemplateBatchItemList"))	
}

//Method implemenation update
func (pm *TemplateBatchItem_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["TemplateBatchItem"], sock.GetPresetFilter("TemplateBatchItem"))	
}


