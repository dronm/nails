package controllers

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * This file contains method descriptions only.
 * Controller implimentation is in SpecialistDocumentOnRegister_ControllerImp.go file
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
type SpecialistDocumentOnRegister_Controller struct {
	osbe.Base_Controller
}

func NewController_SpecialistDocumentOnRegister() *SpecialistDocumentOnRegister_Controller{
	c := &SpecialistDocumentOnRegister_Controller{osbe.Base_Controller{ID: "SpecialistDocumentOnRegister", PublicMethods: make(osbe.PublicMethodCollection)}}	
	keys_fields := fields.GenModelMD(reflect.ValueOf(models.SpecialistDocumentOnRegister_keys{}))
	
	//************************** method insert **********************************
	c.PublicMethods["insert"] = &SpecialistDocumentOnRegister_Controller_insert{
		osbe.Base_PublicMethod{
			ID: "insert",
			Fields: fields.GenModelMD(reflect.ValueOf(models.SpecialistDocumentOnRegister{})),
			EventList: osbe.PublicMethodEventList{"SpecialistDocumentOnRegister.insert"},
		},
	}
	
	//************************** method delete *************************************
	c.PublicMethods["delete"] = &SpecialistDocumentOnRegister_Controller_delete{
		osbe.Base_PublicMethod{
			ID: "delete",
			Fields: keys_fields,
			EventList: osbe.PublicMethodEventList{"SpecialistDocumentOnRegister.delete"},
		},
	}
	
	//************************** method update *************************************
	c.PublicMethods["update"] = &SpecialistDocumentOnRegister_Controller_update{
		osbe.Base_PublicMethod{
			ID: "update",
			Fields: fields.GenModelMD(reflect.ValueOf(models.SpecialistDocumentOnRegister_old_keys{})),
			EventList: osbe.PublicMethodEventList{"SpecialistDocumentOnRegister.update"},
		},
	}
	
	//************************** method get_object *************************************
	c.PublicMethods["get_object"] = &SpecialistDocumentOnRegister_Controller_get_object{
		osbe.Base_PublicMethod{
			ID: "get_object",
			Fields: keys_fields,
		},
	}
	
	//************************** method get_list *************************************
	c.PublicMethods["get_list"] = &SpecialistDocumentOnRegister_Controller_get_list{
		osbe.Base_PublicMethod{
			ID: "get_list",
			Fields: model.Cond_Model_fields,
		},
	}
	
			
	
	return c
}

type SpecialistDocumentOnRegister_Controller_keys_argv struct {
	Argv models.SpecialistDocumentOnRegister_keys `json:"argv"`	
}

//************************* INSERT **********************************************
//Public method: insert
type SpecialistDocumentOnRegister_Controller_insert struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *SpecialistDocumentOnRegister_Controller_insert) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistDocumentOnRegister_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* DELETE **********************************************
type SpecialistDocumentOnRegister_Controller_delete struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *SpecialistDocumentOnRegister_Controller_delete) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistDocumentOnRegister_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET OBJECT **********************************************
type SpecialistDocumentOnRegister_Controller_get_object struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *SpecialistDocumentOnRegister_Controller_get_object) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistDocumentOnRegister_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET LIST **********************************************
//Public method: get_list
type SpecialistDocumentOnRegister_Controller_get_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistDocumentOnRegister_Controller_get_list) Unmarshal(payload []byte) (reflect.Value, error) {
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
type SpecialistDocumentOnRegister_Controller_update struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistDocumentOnRegister_Controller_update) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistDocumentOnRegister_old_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

