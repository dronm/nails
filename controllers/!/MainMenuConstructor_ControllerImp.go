package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 29/11/2023
 *
 * This is MainMenuConstructor controller implimentation file.
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
func (pm *MainMenuConstructor_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["MainMenuConstructor"], &models.MainMenuConstructor_keys{}, sock.GetPresetFilter("MainMenuConstructor"))	
}

//Method implemenation delete
func (pm *MainMenuConstructor_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["MainMenuConstructor"], sock.GetPresetFilter("MainMenuConstructor"))	
}

//Method implemenation get_object
func (pm *MainMenuConstructor_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["MainMenuConstructorDialog"], &models.MainMenuConstructorDialog{}, sock.GetPresetFilter("MainMenuConstructorDialog"))	
}

//Method implemenation get_list
func (pm *MainMenuConstructor_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["MainMenuConstructorList"], &models.MainMenuConstructorList{}, sock.GetPresetFilter("MainMenuConstructorList"))	
}

//Method implemenation update
func (pm *MainMenuConstructor_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["MainMenuConstructor"], sock.GetPresetFilter("MainMenuConstructor"))	
}


