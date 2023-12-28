package controllers

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * This file contains method descriptions only.
 * Controller implimentation is in Speciality_ControllerImp.go file
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
type Speciality_Controller struct {
	osbe.Base_Controller
}

func NewController_Speciality() *Speciality_Controller{
	c := &Speciality_Controller{osbe.Base_Controller{ID: "Speciality", PublicMethods: make(osbe.PublicMethodCollection)}}	
	keys_fields := fields.GenModelMD(reflect.ValueOf(models.Speciality_keys{}))
	
	//************************** method insert **********************************
	c.PublicMethods["insert"] = &Speciality_Controller_insert{
		osbe.Base_PublicMethod{
			ID: "insert",
			Fields: fields.GenModelMD(reflect.ValueOf(models.Speciality{})),
			EventList: osbe.PublicMethodEventList{"Speciality.insert"},
		},
	}
	
	//************************** method delete *************************************
	c.PublicMethods["delete"] = &Speciality_Controller_delete{
		osbe.Base_PublicMethod{
			ID: "delete",
			Fields: keys_fields,
			EventList: osbe.PublicMethodEventList{"Speciality.delete"},
		},
	}
	
	//************************** method update *************************************
	c.PublicMethods["update"] = &Speciality_Controller_update{
		osbe.Base_PublicMethod{
			ID: "update",
			Fields: fields.GenModelMD(reflect.ValueOf(models.Speciality_old_keys{})),
			EventList: osbe.PublicMethodEventList{"Speciality.update"},
		},
	}
	
	//************************** method get_object *************************************
	c.PublicMethods["get_object"] = &Speciality_Controller_get_object{
		osbe.Base_PublicMethod{
			ID: "get_object",
			Fields: keys_fields,
		},
	}
	
	//************************** method get_list *************************************
	c.PublicMethods["get_list"] = &Speciality_Controller_get_list{
		osbe.Base_PublicMethod{
			ID: "get_list",
			Fields: model.Cond_Model_fields,
		},
	}
	
			
	
	return c
}

type Speciality_Controller_keys_argv struct {
	Argv models.Speciality_keys `json:"argv"`	
}

//************************* INSERT **********************************************
//Public method: insert
type Speciality_Controller_insert struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *Speciality_Controller_insert) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.Speciality_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* DELETE **********************************************
type Speciality_Controller_delete struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *Speciality_Controller_delete) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.Speciality_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET OBJECT **********************************************
type Speciality_Controller_get_object struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *Speciality_Controller_get_object) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.Speciality_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET LIST **********************************************
//Public method: get_list
type Speciality_Controller_get_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *Speciality_Controller_get_list) Unmarshal(payload []byte) (reflect.Value, error) {
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
type Speciality_Controller_update struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *Speciality_Controller_update) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.Speciality_old_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

