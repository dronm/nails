package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 21/12/2023
 *
 * This is BankPayment controller implimentation file.
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
func (pm *BankPayment_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["BankPayment"], &models.BankPayment_keys{}, sock.GetPresetFilter("BankPayment"))	
}

//Method implemenation delete
func (pm *BankPayment_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["BankPayment"], sock.GetPresetFilter("BankPayment"))	
}

//Method implemenation get_object
func (pm *BankPayment_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["BankPaymentList"], &models.BankPaymentList{}, sock.GetPresetFilter("BankPaymentList"))	
}

//Method implemenation get_list
func (pm *BankPayment_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["BankPaymentList"], &models.BankPaymentList{}, sock.GetPresetFilter("BankPaymentList"))	
}

//Method implemenation update
func (pm *BankPayment_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["BankPayment"], sock.GetPresetFilter("BankPayment"))	
}


