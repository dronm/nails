package models

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/models/Model.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 */

import (
	"reflect"	
		
	"osbe/fields"
	"osbe/model"
)

type SpecialistDocumentOnRegister struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Document_template_id fields.ValInt `json:"document_template_id" required:"true"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Need_signing fields.ValBool `json:"need_signing"`
}

func (o *SpecialistDocumentOnRegister) SetNull() {
	o.Id.SetNull()
	o.Document_template_id.SetNull()
	o.Date_time.SetNull()
	o.Need_signing.SetNull()
}

func NewModelMD_SpecialistDocumentOnRegister() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistDocumentOnRegister{})),
		ID: "SpecialistDocumentOnRegister_Model",
		Relation: "specialist_documents_on_register",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type SpecialistDocumentOnRegister_argv struct {
	Argv *SpecialistDocumentOnRegister `json:"argv"`	
}

//Keys for delete/get object
type SpecialistDocumentOnRegister_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type SpecialistDocumentOnRegister_keys_argv struct {
	Argv *SpecialistDocumentOnRegister_keys `json:"argv"`	
}

//old keys for update
type SpecialistDocumentOnRegister_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Document_template_id fields.ValInt `json:"document_template_id"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Need_signing fields.ValBool `json:"need_signing"`
}

type SpecialistDocumentOnRegister_old_keys_argv struct {
	Argv *SpecialistDocumentOnRegister_old_keys `json:"argv"`	
}

