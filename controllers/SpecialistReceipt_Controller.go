package controllers

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * This file contains method descriptions only.
 * Controller implimentation is in SpecialistReceipt_ControllerImp.go file
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
type SpecialistReceipt_Controller struct {
	osbe.Base_Controller
}

func NewController_SpecialistReceipt() *SpecialistReceipt_Controller{
	c := &SpecialistReceipt_Controller{osbe.Base_Controller{ID: "SpecialistReceipt", PublicMethods: make(osbe.PublicMethodCollection)}}	
	keys_fields := fields.GenModelMD(reflect.ValueOf(models.SpecialistReceipt_keys{}))
	
	//************************** method insert **********************************
	c.PublicMethods["insert"] = &SpecialistReceipt_Controller_insert{
		osbe.Base_PublicMethod{
			ID: "insert",
			Fields: fields.GenModelMD(reflect.ValueOf(models.SpecialistReceipt{})),
			EventList: osbe.PublicMethodEventList{"SpecialistReceipt.insert"},
		},
	}
	
	//************************** method delete *************************************
	c.PublicMethods["delete"] = &SpecialistReceipt_Controller_delete{
		osbe.Base_PublicMethod{
			ID: "delete",
			Fields: keys_fields,
			EventList: osbe.PublicMethodEventList{"SpecialistReceipt.delete"},
		},
	}
	
	//************************** method update *************************************
	c.PublicMethods["update"] = &SpecialistReceipt_Controller_update{
		osbe.Base_PublicMethod{
			ID: "update",
			Fields: fields.GenModelMD(reflect.ValueOf(models.SpecialistReceipt_old_keys{})),
			EventList: osbe.PublicMethodEventList{"SpecialistReceipt.update"},
		},
	}
	
	//************************** method get_object *************************************
	c.PublicMethods["get_object"] = &SpecialistReceipt_Controller_get_object{
		osbe.Base_PublicMethod{
			ID: "get_object",
			Fields: keys_fields,
		},
	}
	
	//************************** method get_list *************************************
	c.PublicMethods["get_list"] = &SpecialistReceipt_Controller_get_list{
		osbe.Base_PublicMethod{
			ID: "get_list",
			Fields: model.Cond_Model_fields,
		},
	}
	
	//************************** method add *************************************
	c.PublicMethods["add"] = &SpecialistReceipt_Controller_add{
		osbe.Base_PublicMethod{
			ID: "add",
			Fields: fields.GenModelMD(reflect.ValueOf(models.SpecialistReceipt_add{})),
		},
	}
			
	
	return c
}

type SpecialistReceipt_Controller_keys_argv struct {
	Argv models.SpecialistReceipt_keys `json:"argv"`	
}

//************************* INSERT **********************************************
//Public method: insert
type SpecialistReceipt_Controller_insert struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *SpecialistReceipt_Controller_insert) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistReceipt_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* DELETE **********************************************
type SpecialistReceipt_Controller_delete struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *SpecialistReceipt_Controller_delete) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistReceipt_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET OBJECT **********************************************
type SpecialistReceipt_Controller_get_object struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *SpecialistReceipt_Controller_get_object) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistReceipt_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET LIST **********************************************
//Public method: get_list
type SpecialistReceipt_Controller_get_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistReceipt_Controller_get_list) Unmarshal(payload []byte) (reflect.Value, error) {
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
type SpecialistReceipt_Controller_update struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistReceipt_Controller_update) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistReceipt_old_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* add **********************************************
//Public method: add
type SpecialistReceipt_Controller_add struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistReceipt_Controller_add) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistReceipt_add_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}


