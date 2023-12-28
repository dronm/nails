package controllers

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * This file contains method descriptions only.
 * Controller implimentation is in Mgateway_ControllerImp.go file
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
type Mgateway_Controller struct {
	osbe.Base_Controller
}

func NewController_Mgateway() *Mgateway_Controller{
	c := &Mgateway_Controller{osbe.Base_Controller{ID: "Mgateway", PublicMethods: make(osbe.PublicMethodCollection)}}	
	
	//************************** method callback *************************************
	c.PublicMethods["callback"] = &Mgateway_Controller_callback{
		osbe.Base_PublicMethod{
			ID: "callback",
			Fields: fields.GenModelMD(reflect.ValueOf(models.Mgateway_callback{})),
		},
	}

	//************************** method get_qr *************************************
	c.PublicMethods["get_qr"] = &Mgateway_Controller_get_qr{
		osbe.Base_PublicMethod{
			ID: "get_qr",
		},
	}
	
	return c
}


//************************* callback **********************************************
//Public method: callback
type Mgateway_Controller_callback struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *Mgateway_Controller_callback) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.Mgateway_callback_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}

//************************* callback **********************************************
//Public method: callback
type Mgateway_Controller_get_qr struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *Mgateway_Controller_get_qr) Unmarshal(payload []byte) (reflect.Value, error) {
	res := reflect.ValueOf(nil)
	return res, nil
}


