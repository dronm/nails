package controllers

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * This file contains method descriptions only.
 * Controller implimentation is in SpecialistPeriodSalaryDetail_ControllerImp.go file
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
type SpecialistPeriodSalaryDetail_Controller struct {
	osbe.Base_Controller
}

func NewController_SpecialistPeriodSalaryDetail() *SpecialistPeriodSalaryDetail_Controller{
	c := &SpecialistPeriodSalaryDetail_Controller{osbe.Base_Controller{ID: "SpecialistPeriodSalaryDetail", PublicMethods: make(osbe.PublicMethodCollection)}}	
	keys_fields := fields.GenModelMD(reflect.ValueOf(models.SpecialistPeriodSalaryDetail_keys{}))
	
	//************************** method insert **********************************
	c.PublicMethods["insert"] = &SpecialistPeriodSalaryDetail_Controller_insert{
		osbe.Base_PublicMethod{
			ID: "insert",
			Fields: fields.GenModelMD(reflect.ValueOf(models.SpecialistPeriodSalaryDetail{})),
			EventList: osbe.PublicMethodEventList{"SpecialistPeriodSalaryDetail.insert"},
		},
	}
	
	//************************** method delete *************************************
	c.PublicMethods["delete"] = &SpecialistPeriodSalaryDetail_Controller_delete{
		osbe.Base_PublicMethod{
			ID: "delete",
			Fields: keys_fields,
			EventList: osbe.PublicMethodEventList{"SpecialistPeriodSalaryDetail.delete"},
		},
	}
	
	//************************** method update *************************************
	c.PublicMethods["update"] = &SpecialistPeriodSalaryDetail_Controller_update{
		osbe.Base_PublicMethod{
			ID: "update",
			Fields: fields.GenModelMD(reflect.ValueOf(models.SpecialistPeriodSalaryDetail_old_keys{})),
			EventList: osbe.PublicMethodEventList{"SpecialistPeriodSalaryDetail.update"},
		},
	}
	
	//************************** method get_object *************************************
	c.PublicMethods["get_object"] = &SpecialistPeriodSalaryDetail_Controller_get_object{
		osbe.Base_PublicMethod{
			ID: "get_object",
			Fields: keys_fields,
		},
	}
	
	//************************** method get_list *************************************
	c.PublicMethods["get_list"] = &SpecialistPeriodSalaryDetail_Controller_get_list{
		osbe.Base_PublicMethod{
			ID: "get_list",
			Fields: model.Cond_Model_fields,
		},
	}
	
	//************************** method get_for_pay_list *************************************
	c.PublicMethods["get_for_pay_list"] = &SpecialistPeriodSalaryDetail_Controller_get_for_pay_list{
		osbe.Base_PublicMethod{
			ID: "get_for_pay_list",
			Fields: model.Cond_Model_fields,
		},
	}

	//************************** method documents_to_bank *************************************
	c.PublicMethods["documents_to_bank"] = &SpecialistPeriodSalaryDetail_Controller_documents_to_bank{
		osbe.Base_PublicMethod{
			ID: "documents_to_bank",
			Fields: fields.GenModelMD(reflect.ValueOf(models.Object_key_ar{})),
		},		
	}

	//************************** method complete *************************************
	c.PublicMethods["complete"] = &SpecialistPeriodSalaryDetail_Controller_complete{
		osbe.Base_PublicMethod{
			ID: "complete",
			Fields: fields.GenModelMD(reflect.ValueOf(models.SpecialistPeriodSalaryDetail_complete{})),
		},
	}

	return c
}

type SpecialistPeriodSalaryDetail_Controller_keys_argv struct {
	Argv models.SpecialistPeriodSalaryDetail_keys `json:"argv"`	
}

//************************* INSERT **********************************************
//Public method: insert
type SpecialistPeriodSalaryDetail_Controller_insert struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *SpecialistPeriodSalaryDetail_Controller_insert) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistPeriodSalaryDetail_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* DELETE **********************************************
type SpecialistPeriodSalaryDetail_Controller_delete struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *SpecialistPeriodSalaryDetail_Controller_delete) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistPeriodSalaryDetail_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET OBJECT **********************************************
type SpecialistPeriodSalaryDetail_Controller_get_object struct {
	osbe.Base_PublicMethod
}

//Public method Unmarshal to structure
func (pm *SpecialistPeriodSalaryDetail_Controller_get_object) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistPeriodSalaryDetail_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* GET LIST **********************************************
//Public method: get_list
type SpecialistPeriodSalaryDetail_Controller_get_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistPeriodSalaryDetail_Controller_get_list) Unmarshal(payload []byte) (reflect.Value, error) {
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
type SpecialistPeriodSalaryDetail_Controller_update struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistPeriodSalaryDetail_Controller_update) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistPeriodSalaryDetail_old_keys_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* get_for_pay_list **********************************************
//Public method: get_for_pay_list
type SpecialistPeriodSalaryDetail_Controller_get_for_pay_list struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistPeriodSalaryDetail_Controller_get_for_pay_list) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &model.Controller_get_list_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//***************************** documents_to_bank ***************************************************************
//Public method: set_person_confirmed
type SpecialistPeriodSalaryDetail_Controller_documents_to_bank struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistPeriodSalaryDetail_Controller_documents_to_bank) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &models.Object_key_ar_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}


//************************** COMPLETE ********************************************************
//Public method: complete
type SpecialistPeriodSalaryDetail_Controller_complete struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *SpecialistPeriodSalaryDetail_Controller_complete) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.SpecialistPeriodSalaryDetail_complete_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}
