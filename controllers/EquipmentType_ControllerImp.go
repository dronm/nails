package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 30/11/2023
 *
 * This is EquipmentType controller implimentation file.
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
func (pm *EquipmentType_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["EquipmentType"], &models.EquipmentType_keys{}, sock.GetPresetFilter("EquipmentType"))	
}

//Method implemenation delete
func (pm *EquipmentType_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["EquipmentType"], sock.GetPresetFilter("EquipmentType"))	
}

//Method implemenation get_object
func (pm *EquipmentType_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["EquipmentType"], &models.EquipmentType{}, sock.GetPresetFilter("EquipmentType"))	
}

//Method implemenation get_list
func (pm *EquipmentType_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["EquipmentType"], &models.EquipmentType{}, sock.GetPresetFilter("EquipmentType"))	
}

//Method implemenation update
func (pm *EquipmentType_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["EquipmentType"], sock.GetPresetFilter("EquipmentType"))	
}

//Method implemenation complete
func (pm *EquipmentType_Controller_complete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.CompleteOnArgs(app, resp, rfltArgs, app.GetMD().Models["EquipmentType"], &models.EquipmentType{}, sock.GetPresetFilter("EquipmentType"))	
}

