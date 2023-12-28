package controllers

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * This file contains method descriptions only.
 * Controller implimentation is in SpecialistWork_ControllerImp.go file
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
type SpecialistWork_Controller struct {
	osbe.Base_Controller
}

func NewController_SpecialistWork() *SpecialistWork_Controller{
	c := &SpecialistWork_Controller{osbe.Base_Controller{ID: "SpecialistWork", PublicMethods: make(osbe.PublicMethodCollection)}}	
	keys_fields := fields.GenModelMD(reflect.ValueOf(models.SpecialistWork_keys{}))
	
	//************************** method insert **********************************
	c.PublicMethods["insert"] = &SpecialistWork_Controller_insert{
		osbe.Base_PublicMethod{
			ID: "insert",
			Fields: fields.GenModelMD(reflect.ValueOf(models.SpecialistWork{})),
			EventList: osbe.PublicMethodEventList{"SpecialistWork.insert"},
		},
	}
	
	//************************** method delete *************************************
	c.PublicMethods["delete"] = &SpecialistWork_Controller_delete{
		osbe.Base_PublicMethod{
			ID: "delete",
			Fields: keys_fields,
			EventList: osbe.PublicMethodEventList{"SpecialistWork.delete"},
		},
	}
	
	//************************** method update *************************************
	c.PublicMethods["update"] = &SpecialistWork_Controller_update{
		osbe.Base_PublicMethod{
			ID: "update",
			Fields: fields.GenModelMD(reflect.ValueOf(models.SpecialistWork_old_keys{})),
			EventList: osbe.PublicMethodEventList{"SpecialistWork.update"},
		},
	}
	
	//************************** method get_object *************************************
	c.PublicMethods["get_object"] = &SpecialistWork_Controller_get_object{
		osbe.Base_PublicMethod{
			ID: "get_object",
			Fields: keys_fields,
		},
	}
	
	//************************** method get_list *************************************
	c.PublicMethods["get_list"] = &SpecialistWork_Controller_get_list{
		osbe.Base_PublicMethod{
			ID: "get_list",
			Fields: model.Cond_Model_fields,
		},
	}
	
	//************************** method get_for_rate_list *************************************
	c.PublicMethods["get_for_rate_list"] = &SpecialistWork_Controller_get_for_rate_list{
		osbe.Base_PublicMethod{
			ID: "get_for_rate_list",
			Fields: model.Cond_Model_fields,
		},
	}
			
	//************************** method set_admin_rates *************************************
	c.PublicMethods["set_admin_rates"] = &SpecialistWork_Controller_set_admin_rates{
		osbe.Base_PublicMethod{
			ID: "set_admin_rates",
			Fields: fields.GenModelMD(reflect.ValueOf(models.SpecialistWork_set_admin_rates{})),
		},
	}
	
	return c
}

type SpecialistWork_Controller_keys_argv struct {
	Argv models.SpecialistWork_keys `json:"argv"`	
}

//************************* INSERT **********************************************
//Public method: insert
type SpecialistWork_Controller_insert struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *SpecialistWork_Controller_insert) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistWork_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* DELETE **********************************************
type SpecialistWork_Controller_delete struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *SpecialistWork_Controller_delete) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistWork_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET OBJECT **********************************************
type SpecialistWork_Controller_get_object struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *SpecialistWork_Controller_get_object) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistWork_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET LIST **********************************************
//Public method: get_list
type SpecialistWork_Controller_get_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistWork_Controller_get_list) Unmarshal(payload []byte) (reflect.Value, error) {
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
type SpecialistWork_Controller_update struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistWork_Controller_update) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistWork_old_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* get_for_rate_list **********************************************
//Public method: get_for_rate_list
type SpecialistWork_Controller_get_for_rate_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistWork_Controller_get_for_rate_list) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &model.Controller_get_list_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}


//************************* set_admin_rates **********************************************
//Public method: set_admin_rates
type SpecialistWork_Controller_set_admin_rates struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistWork_Controller_set_admin_rates) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistWork_set_admin_rates_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}


