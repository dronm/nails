package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 05/12/2023
 *
 * This is YClTransaction controller implimentation file.
 *
 */

import (
	"reflect"	
	
	"nails/models"
	
	"osbe"
	"osbe/srv"
	"osbe/socket"
	"osbe/response"	
	
	"github.com/dronm/ds/pgds"
	//"github.com/jackc/pgx/v4"
)

//Method implemenation insert
func (pm *YClTransaction_Controller_insert) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.InsertOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["YClTransaction"], &models.YClTransaction_keys{}, sock.GetPresetFilter("YClTransaction"))	
}

//Method implemenation delete
func (pm *YClTransaction_Controller_delete) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.DeleteOnArgKeys(app, pm, resp, sock, rfltArgs, app.GetMD().Models["YClTransaction"], sock.GetPresetFilter("YClTransaction"))	
}

//Method implemenation get_object
func (pm *YClTransaction_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, app.GetMD().Models["YClTransactionList"], &models.YClTransactionList{}, sock.GetPresetFilter("YClTransactionList"))	
}

//Method implemenation get_list
func (pm *YClTransaction_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["YClTransactionList"], &models.YClTransactionList{}, sock.GetPresetFilter("YClTransactionList"))	
}

//Method implemenation get_for_work_list
func (pm *YClTransaction_Controller_get_for_work_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["YClTransactionDocList"], &models.YClTransactionDocList{}, sock.GetPresetFilter("YClTransactionDocList"))	
}

//Method implemenation get_doc_all_list
func (pm *YClTransaction_Controller_get_doc_all_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, app.GetMD().Models["YClTransactionDocAllList"], &models.YClTransactionDocAllList{}, sock.GetPresetFilter("YClTransactionDocAllList"))	
}

//Method implemenation update
func (pm *YClTransaction_Controller_update) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.UpdateOnArgs(app, pm, resp, sock, rfltArgs, app.GetMD().Models["YClTransaction"], sock.GetPresetFilter("YClTransaction"))	
}

//Method implemenation api_get
func (pm *YClTransaction_Controller_api_get) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	args := rfltArgs.Interface().(*models.YClTransaction_api_get)
	
	sess := sock.GetSession()	
	user_id := sess.GetInt(SESS_VAR_ID)
	go YClients.UpdateTransactions(args.Operation_id.GetValue(), user_id, args.Date_from.GetValue(), args.Date_to.GetValue(), app.GetDataStorage().(*pgds.PgProvider), app.GetLogger())
	
	return nil
}


