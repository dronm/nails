package controllers

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * Controller implimentation should be in User_ControllerImp.go file
 *
 */

import (
	"encoding/json"
	"reflect"
	
	"nails/models"
	
	"osbe"	
	"osbe/fields"
	"osbe/model"
)

//Controller
type User_Controller struct {
	osbe.Base_Controller
}

func NewController_User() *User_Controller{
	c := &User_Controller{osbe.Base_Controller{ID: "User", PublicMethods: make(osbe.PublicMethodCollection)}}
	
	keys_fields := fields.GenModelMD(reflect.ValueOf(models.User_keys{}))
	
	//************************** method insert **********************************
	c.PublicMethods["insert"] = &User_Controller_insert{
		osbe.Base_PublicMethod{
			ID: "insert",
			Fields: fields.GenModelMD(reflect.ValueOf(models.User{})),
			EventList: osbe.PublicMethodEventList{"User.insert"},
		},				
	}
	
	//************************** method delete *************************************
	c.PublicMethods["delete"] = &User_Controller_delete{
		osbe.Base_PublicMethod{
			ID: "delete",
			Fields: keys_fields,
			EventList: osbe.PublicMethodEventList{"User.delete"},
		},		
	}

	//************************** method update *************************************
	c.PublicMethods["update"] = &User_Controller_update{
		osbe.Base_PublicMethod{
			ID: "update",
			Fields: fields.GenModelMD(reflect.ValueOf(models.User_old_keys{})),
			EventList: osbe.PublicMethodEventList{"User.update"},
		},		
	}
	
	//************************** method get_object *************************************
	c.PublicMethods["get_object"] = &User_Controller_get_object{
		osbe.Base_PublicMethod{
			ID: "get_object",
			Fields: keys_fields,
		},
	}

	//************************** method get_list *************************************
	c.PublicMethods["get_list"] = &User_Controller_get_list{
		osbe.Base_PublicMethod{
			ID: "get_list",
			Fields: model.Cond_Model_fields,
		},		
	}

	//************************** method reset_pwd *************************************
	c.PublicMethods["reset_pwd"] = &User_Controller_reset_pwd{
		osbe.Base_PublicMethod{
			ID: "reset_pwd",
			Fields: fields.GenModelMD(reflect.ValueOf(models.User_reset_pwd{})),
		},		
	}

	//************************** method login *************************************
	c.PublicMethods["login"] = &User_Controller_login{
		osbe.Base_PublicMethod{
			ID: "login",
			Fields: fields.GenModelMD(reflect.ValueOf(models.User_login{})),
		},		
	}

	//************************** method login_token *************************************
	c.PublicMethods["login_token"] = &User_Controller_login_token{
		osbe.Base_PublicMethod{
			ID: "login_token",
			Fields: fields.GenModelMD(reflect.ValueOf(models.User_login_token{})),
		},		
	}

	//************************** method logout *************************************
	c.PublicMethods["logout"] = &User_Controller_logout{
		osbe.Base_PublicMethod{
			ID: "logout",
		},		
	}
	
	//************************** method get_profile *************************************
	c.PublicMethods["get_profile"] = &User_Controller_get_profile{
		osbe.Base_PublicMethod{
			ID: "get_profile",
		},		
	}
	
	//************************** method complete *************************************
	c.PublicMethods["complete"] = &User_Controller_complete{
		osbe.Base_PublicMethod{
			ID: "complete",
			Fields: fields.GenModelMD(reflect.ValueOf(models.User_complete{})),
		},		
	}
	
	//************************** method password_recover *************************************
	c.PublicMethods["password_recover"] = &User_Controller_password_recover{
		osbe.Base_PublicMethod{
			ID: "password_recover",
			Fields: fields.GenModelMD(reflect.ValueOf(models.User_password_recover{})),
		},		
	}
	
	return c
}

//**************************************************************************************
//Public method: insert
type User_Controller_insert struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *User_Controller_insert) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &models.User_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}

	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}

//**************************************************************************************
//Public method: delete
type User_Controller_delete struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *User_Controller_delete) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &models.User_keys_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}

//Public method: get_object
type User_Controller_get_object struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *User_Controller_get_object) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &models.User_keys_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}

//**************************************************************************************
//Public method: get_list
type User_Controller_get_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *User_Controller_get_list) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &model.Controller_get_list_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}


//**************************************************************************************
//Public method: update
type User_Controller_update struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *User_Controller_update) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &models.User_old_keys_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}


//Public method: reset_pwd
type User_Controller_reset_pwd struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *User_Controller_reset_pwd) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &models.User_reset_pwd_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}


//Public method: login
type User_Controller_login struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *User_Controller_login) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &models.User_login_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}


//Public method: logout
type User_Controller_logout struct {
	osbe.Base_PublicMethod
}
func (pm *User_Controller_logout) Unmarshal(payload []byte) (res reflect.Value, err error) {
	return
}

//Public method: login_token
type User_Controller_login_token struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *User_Controller_login_token) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &models.User_login_token_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}

//********************************************************************************************
//Public method: get_profile
type User_Controller_get_profile struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *User_Controller_get_profile) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &models.User_keys_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}

//********************************************************************************************
//Public method: complete
type User_Controller_complete struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *User_Controller_complete) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &models.User_complete_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}

//
//Public method: password_recover
type User_Controller_password_recover struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *User_Controller_password_recover) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &models.User_password_recover_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}

