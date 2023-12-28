package controllers

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * This file contains method descriptions only.
 * Controller implimentation is in EquipmentType_ControllerImp.go file
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
type EquipmentType_Controller struct {
	osbe.Base_Controller
}

func NewController_EquipmentType() *EquipmentType_Controller{
	c := &EquipmentType_Controller{osbe.Base_Controller{ID: "EquipmentType", PublicMethods: make(osbe.PublicMethodCollection)}}	
	keys_fields := fields.GenModelMD(reflect.ValueOf(models.EquipmentType_keys{}))
	
	//************************** method insert **********************************
	c.PublicMethods["insert"] = &EquipmentType_Controller_insert{
		osbe.Base_PublicMethod{
			ID: "insert",
			Fields: fields.GenModelMD(reflect.ValueOf(models.EquipmentType{})),
			EventList: osbe.PublicMethodEventList{"EquipmentType.insert"},
		},
	}
	
	//************************** method delete *************************************
	c.PublicMethods["delete"] = &EquipmentType_Controller_delete{
		osbe.Base_PublicMethod{
			ID: "delete",
			Fields: keys_fields,
			EventList: osbe.PublicMethodEventList{"EquipmentType.delete"},
		},
	}
	
	//************************** method update *************************************
	c.PublicMethods["update"] = &EquipmentType_Controller_update{
		osbe.Base_PublicMethod{
			ID: "update",
			Fields: fields.GenModelMD(reflect.ValueOf(models.EquipmentType_old_keys{})),
			EventList: osbe.PublicMethodEventList{"EquipmentType.update"},
		},
	}
	
	//************************** method get_object *************************************
	c.PublicMethods["get_object"] = &EquipmentType_Controller_get_object{
		osbe.Base_PublicMethod{
			ID: "get_object",
			Fields: keys_fields,
		},
	}
	
	//************************** method get_list *************************************
	c.PublicMethods["get_list"] = &EquipmentType_Controller_get_list{
		osbe.Base_PublicMethod{
			ID: "get_list",
			Fields: model.Cond_Model_fields,
		},
	}
	
	//************************** method complete *************************************
	c.PublicMethods["complete"] = &EquipmentType_Controller_complete{
		osbe.Base_PublicMethod{
			ID: "complete",
			Fields: fields.GenModelMD(reflect.ValueOf(models.EquipmentType_complete{})),
		},
	}
			
	
	return c
}

type EquipmentType_Controller_keys_argv struct {
	Argv models.EquipmentType_keys `json:"argv"`	
}

//************************* INSERT **********************************************
//Public method: insert
type EquipmentType_Controller_insert struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *EquipmentType_Controller_insert) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.EquipmentType_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* DELETE **********************************************
type EquipmentType_Controller_delete struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *EquipmentType_Controller_delete) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.EquipmentType_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET OBJECT **********************************************
type EquipmentType_Controller_get_object struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *EquipmentType_Controller_get_object) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.EquipmentType_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET LIST **********************************************
//Public method: get_list
type EquipmentType_Controller_get_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *EquipmentType_Controller_get_list) Unmarshal(payload []byte) (reflect.Value, error) {
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
type EquipmentType_Controller_update struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *EquipmentType_Controller_update) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.EquipmentType_old_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************** COMPLETE ********************************************************
//Public method: complete
type EquipmentType_Controller_complete struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *EquipmentType_Controller_complete) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.EquipmentType_complete_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}
