package controllers

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * This file contains method descriptions only.
 * Controller implimentation is in ConfirmationStatus_ControllerImp.go file
 *
 */

import (
	"reflect"	
	"encoding/json"
	
	"nails/models"
	
	"osbe"
	"osbe/fields"
	
)

//Controller
type ConfirmationStatus_Controller struct {
	osbe.Base_Controller
}

func NewController_ConfirmationStatus() *ConfirmationStatus_Controller{
	c := &ConfirmationStatus_Controller{osbe.Base_Controller{ID: "ConfirmationStatus", PublicMethods: make(osbe.PublicMethodCollection)}}	
	
	//************************** method set_confirmed *************************************
	c.PublicMethods["set_confirmed"] = &ConfirmationStatus_Controller_set_confirmed{
		osbe.Base_PublicMethod{
			ID: "set_confirmed",
			Fields: fields.GenModelMD(reflect.ValueOf(models.ConfirmationStatus_set_confirmed{})),
		},
	}
	
	return c
}

//************************* set_confirmed **********************************************
//Public method: set_confirmed
type ConfirmationStatus_Controller_set_confirmed struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *ConfirmationStatus_Controller_set_confirmed) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.ConfirmationStatus_set_confirmed_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}


