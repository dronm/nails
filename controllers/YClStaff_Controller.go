package controllers

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * This file contains method descriptions only.
 * Controller implimentation is in YClStaff_ControllerImp.go file
 *
 */

import (
	"reflect"	
	"encoding/json"
	
	"nails/models"
	
	"osbe"
	"osbe/fields"
	"osbe/model"
)

//Controller
type YClStaff_Controller struct {
	osbe.Base_Controller
}

func NewController_YClStaff() *YClStaff_Controller{
	c := &YClStaff_Controller{osbe.Base_Controller{ID: "YClStaff", PublicMethods: make(osbe.PublicMethodCollection)}}	
	keys_fields := fields.GenModelMD(reflect.ValueOf(models.YClStaff_keys{}))
	
	//************************** method insert **********************************
	c.PublicMethods["insert"] = &YClStaff_Controller_insert{
		osbe.Base_PublicMethod{
			ID: "insert",
			Fields: fields.GenModelMD(reflect.ValueOf(models.YClStaff{})),
			EventList: osbe.PublicMethodEventList{"YClStaff.insert"},
		},
	}
	
	//************************** method delete *************************************
	c.PublicMethods["delete"] = &YClStaff_Controller_delete{
		osbe.Base_PublicMethod{
			ID: "delete",
			Fields: keys_fields,
			EventList: osbe.PublicMethodEventList{"YClStaff.delete"},
		},
	}
	
	//************************** method update *************************************
	c.PublicMethods["update"] = &YClStaff_Controller_update{
		osbe.Base_PublicMethod{
			ID: "update",
			Fields: fields.GenModelMD(reflect.ValueOf(models.YClStaff_old_keys{})),
			EventList: osbe.PublicMethodEventList{"YClStaff.update"},
		},
	}
	
	//************************** method get_object *************************************
	c.PublicMethods["get_object"] = &YClStaff_Controller_get_object{
		osbe.Base_PublicMethod{
			ID: "get_object",
			Fields: keys_fields,
		},
	}
	
	//************************** method get_list *************************************
	c.PublicMethods["get_list"] = &YClStaff_Controller_get_list{
		osbe.Base_PublicMethod{
			ID: "get_list",
			Fields: model.Cond_Model_fields,
		},
	}
	
	//************************** method api_get_all *************************************
	c.PublicMethods["api_get_all"] = &YClStaff_Controller_api_get_all{
		osbe.Base_PublicMethod{
			ID: "api_get_all",
			Fields: fields.GenModelMD(reflect.ValueOf(models.YClStaff_api_get_all{})),
		},
	}			
	
	return c
}

type YClStaff_Controller_keys_argv struct {
	Argv models.YClStaff_keys `json:"argv"`	
}

//************************* INSERT **********************************************
//Public method: insert
type YClStaff_Controller_insert struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *YClStaff_Controller_insert) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.YClStaff_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* DELETE **********************************************
type YClStaff_Controller_delete struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *YClStaff_Controller_delete) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.YClStaff_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET OBJECT **********************************************
type YClStaff_Controller_get_object struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *YClStaff_Controller_get_object) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.YClStaff_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET LIST **********************************************
//Public method: get_list
type YClStaff_Controller_get_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *YClStaff_Controller_get_list) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &model.Controller_get_list_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* UPDATE **********************************************
//Public method: update
type YClStaff_Controller_update struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *YClStaff_Controller_update) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.YClStaff_old_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* api_get_all **********************************************
//Public method: api_get_all
type YClStaff_Controller_api_get_all struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *YClStaff_Controller_api_get_all) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.YClStaff_api_get_all_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

