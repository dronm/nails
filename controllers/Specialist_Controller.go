package controllers

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * This file contains method descriptions only.
 * Controller implimentation is in Specialist_ControllerImp.go file
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
type Specialist_Controller struct {
	osbe.Base_Controller
}

func NewController_Specialist() *Specialist_Controller{
	c := &Specialist_Controller{osbe.Base_Controller{ID: "Specialist", PublicMethods: make(osbe.PublicMethodCollection)}}	
	keys_fields := fields.GenModelMD(reflect.ValueOf(models.Specialist_keys{}))
	
	//************************** method insert **********************************
	c.PublicMethods["insert"] = &Specialist_Controller_insert{
		osbe.Base_PublicMethod{
			ID: "insert",
			Fields: fields.GenModelMD(reflect.ValueOf(models.Specialist{})),
			EventList: osbe.PublicMethodEventList{"Specialist.insert"},
		},
	}
	
	//************************** method delete *************************************
	c.PublicMethods["delete"] = &Specialist_Controller_delete{
		osbe.Base_PublicMethod{
			ID: "delete",
			Fields: keys_fields,
			EventList: osbe.PublicMethodEventList{"Specialist.delete"},
		},
	}
	
	//************************** method update *************************************
	c.PublicMethods["update"] = &Specialist_Controller_update{
		osbe.Base_PublicMethod{
			ID: "update",
			Fields: fields.GenModelMD(reflect.ValueOf(models.Specialist_old_keys{})),
			EventList: osbe.PublicMethodEventList{"Specialist.update"},
		},
	}
	
	//************************** method get_object *************************************
	c.PublicMethods["get_object"] = &Specialist_Controller_get_object{
		osbe.Base_PublicMethod{
			ID: "get_object",
			Fields: keys_fields,
		},
	}
	
	//************************** method get_list *************************************
	c.PublicMethods["get_list"] = &Specialist_Controller_get_list{
		osbe.Base_PublicMethod{
			ID: "get_list",
			Fields: model.Cond_Model_fields,
		},
	}
	
	//************************** method register *************************************
	c.PublicMethods["register"] = &Specialist_Controller_register{
		osbe.Base_PublicMethod{
			ID: "register",
			Fields: fields.GenModelMD(reflect.ValueOf(models.Specialist_register{})),
		},
	}
	
	//************************** method complete *************************************
	c.PublicMethods["complete"] = &Specialist_Controller_complete{
		osbe.Base_PublicMethod{
			ID: "complete",
			Fields: fields.GenModelMD(reflect.ValueOf(models.Specialist_complete{})),
		},
	}
			
	//************************** method send_reg_documents *************************************
	c.PublicMethods["send_reg_documents"] = &Specialist_Controller_send_reg_documents{
		osbe.Base_PublicMethod{
			ID: "send_reg_documents",
			Fields: fields.GenModelMD(reflect.ValueOf(models.DocAsync_id{})),
		},
	}

	//************************** method send_salary_documents *************************************
	c.PublicMethods["send_salary_documents"] = &Specialist_Controller_send_salary_documents{
		osbe.Base_PublicMethod{
			ID: "send_salary_documents",
			Fields: fields.GenModelMD(reflect.ValueOf(models.DocAsync_id{})),
		},
	}
	
	return c
}

type Specialist_Controller_keys_argv struct {
	Argv models.Specialist_keys `json:"argv"`	
}

//************************* INSERT **********************************************
//Public method: insert
type Specialist_Controller_insert struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *Specialist_Controller_insert) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.Specialist_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* DELETE **********************************************
type Specialist_Controller_delete struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *Specialist_Controller_delete) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.Specialist_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET OBJECT **********************************************
type Specialist_Controller_get_object struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *Specialist_Controller_get_object) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.Specialist_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET LIST **********************************************
//Public method: get_list
type Specialist_Controller_get_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *Specialist_Controller_get_list) Unmarshal(payload []byte) (reflect.Value, error) {
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
type Specialist_Controller_update struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *Specialist_Controller_update) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.Specialist_old_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* register **********************************************
//Public method: register
type Specialist_Controller_register struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *Specialist_Controller_register) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.Specialist_register_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************** COMPLETE ********************************************************
//Public method: complete
type Specialist_Controller_complete struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *Specialist_Controller_complete) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.Specialist_complete_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* send_reg_documents **********************************************
//Public method: send_reg_documents
type Specialist_Controller_send_reg_documents struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *Specialist_Controller_send_reg_documents) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.DocAsync_id_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* send_salary_documents **********************************************
//Public method: send_salary_documents
type Specialist_Controller_send_salary_documents struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *Specialist_Controller_send_salary_documents) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.DocAsync_id_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}


