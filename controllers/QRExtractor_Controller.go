package controllers

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * This file contains method descriptions only.
 * Controller implimentation is in QRExtractor_ControllerImp.go file
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
type QRExtractor_Controller struct {
	osbe.Base_Controller
}

func NewController_QRExtractor() *QRExtractor_Controller{
	c := &QRExtractor_Controller{osbe.Base_Controller{ID: "QRExtractor", PublicMethods: make(osbe.PublicMethodCollection)}}	

	//************************** method callback *************************************
	c.PublicMethods["callback"] = &QRExtractor_Controller_callback{
		osbe.Base_PublicMethod{
			ID: "callback",
			Fields: fields.GenModelMD(reflect.ValueOf(models.QRExtractor_callback{})),
		},
	}

	
	return c
}

//************************* callback **********************************************
//Public method: callback
type QRExtractor_Controller_callback struct {
	osbe.Base_PublicMethod
}
//Public method Unmarshal to structure
func (pm *QRExtractor_Controller_callback) Unmarshal(payload []byte) (reflect.Value, error) {
	var res reflect.Value
	argv := &models.QRExtractor_callback_argv{}
		
	if err := json.Unmarshal(payload, argv); err != nil {
		return res, err
	}	
	res = reflect.ValueOf(&argv.Argv).Elem()	
	return res, nil
}







