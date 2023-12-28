package controllers

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * This file contains method descriptions only.
 * Controller implimentation is in YClTransaction_ControllerImp.go file
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
type YClTransaction_Controller struct {
	osbe.Base_Controller
}

func NewController_YClTransaction() *YClTransaction_Controller{
	c := &YClTransaction_Controller{osbe.Base_Controller{ID: "YClTransaction", PublicMethods: make(osbe.PublicMethodCollection)}}	
	keys_fields := fields.GenModelMD(reflect.ValueOf(models.YClTransaction_keys{}))
	
	//************************** method insert **********************************
	c.PublicMethods["insert"] = &YClTransaction_Controller_insert{
		osbe.Base_PublicMethod{
			ID: "insert",
			Fields: fields.GenModelMD(reflect.ValueOf(models.YClTransaction{})),
			EventList: osbe.PublicMethodEventList{"YClTransaction.insert"},
		},
	}
	
	//************************** method delete *************************************
	c.PublicMethods["delete"] = &YClTransaction_Controller_delete{
		osbe.Base_PublicMethod{
			ID: "delete",
			Fields: keys_fields,
			EventList: osbe.PublicMethodEventList{"YClTransaction.delete"},
		},
	}
	
	//************************** method update *************************************
	c.PublicMethods["update"] = &YClTransaction_Controller_update{
		osbe.Base_PublicMethod{
			ID: "update",
			Fields: fields.GenModelMD(reflect.ValueOf(models.YClTransaction_old_keys{})),
			EventList: osbe.PublicMethodEventList{"YClTransaction.update"},
		},
	}
	
	//************************** method get_object *************************************
	c.PublicMethods["get_object"] = &YClTransaction_Controller_get_object{
		osbe.Base_PublicMethod{
			ID: "get_object",
			Fields: keys_fields,
		},
	}
	
	//************************** method get_list *************************************
	c.PublicMethods["get_list"] = &YClTransaction_Controller_get_list{
		osbe.Base_PublicMethod{
			ID: "get_list",
			Fields: model.Cond_Model_fields,
		},
	}
	
	//************************** method get_for_work_list *************************************
	c.PublicMethods["get_for_work_list"] = &YClTransaction_Controller_get_for_work_list{
		osbe.Base_PublicMethod{
			ID: "get_for_work_list",
			Fields: model.Cond_Model_fields,
		},
	}

	//************************** method get_doc_all_list *************************************
	c.PublicMethods["get_doc_all_list"] = &YClTransaction_Controller_get_doc_all_list{
		osbe.Base_PublicMethod{
			ID: "get_doc_all_list",
			Fields: model.Cond_Model_fields,
		},
	}

	//************************** method api_get *************************************
	c.PublicMethods["api_get"] = &YClTransaction_Controller_api_get{
		osbe.Base_PublicMethod{
			ID: "api_get",
			Fields: fields.GenModelMD(reflect.ValueOf(models.YClTransaction_api_get{})),
		},
	}
	
	return c
}

type YClTransaction_Controller_keys_argv struct {
	Argv models.YClTransaction_keys `json:"argv"`	
}

//************************* INSERT **********************************************
//Public method: insert
type YClTransaction_Controller_insert struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *YClTransaction_Controller_insert) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.YClTransaction_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* DELETE **********************************************
type YClTransaction_Controller_delete struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *YClTransaction_Controller_delete) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.YClTransaction_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET OBJECT **********************************************
type YClTransaction_Controller_get_object struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *YClTransaction_Controller_get_object) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.YClTransaction_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET LIST **********************************************
//Public method: get_list
type YClTransaction_Controller_get_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *YClTransaction_Controller_get_list) Unmarshal(payload []byte) (reflect.Value, error) {
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
type YClTransaction_Controller_update struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *YClTransaction_Controller_update) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.YClTransaction_old_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* get_for_work_list **********************************************
//Public method: get_for_work_list
type YClTransaction_Controller_get_for_work_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *YClTransaction_Controller_get_for_work_list) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &model.Controller_get_list_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* get_doc_all_list **********************************************
//Public method: get_doc_all_list
type YClTransaction_Controller_get_doc_all_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *YClTransaction_Controller_get_doc_all_list) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &model.Controller_get_list_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}


//************************* api_get **********************************************
//Public method: api_get
type YClTransaction_Controller_api_get struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *YClTransaction_Controller_api_get) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.YClTransaction_api_get_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

