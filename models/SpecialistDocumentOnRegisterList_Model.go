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

type SpecialistDocumentOnRegisterList struct {
	Id fields.ValInt `json:"id" primaryKey:"true"`
	Document_templates_ref fields.ValJSON `json:"document_templates_ref"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Need_signing fields.ValBool `json:"need_signing"`
}

func (o *SpecialistDocumentOnRegisterList) SetNull() {
	o.Id.SetNull()
	o.Document_templates_ref.SetNull()
	o.Date_time.SetNull()
	o.Need_signing.SetNull()
}

func NewModelMD_SpecialistDocumentOnRegisterList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistDocumentOnRegisterList{})),
		ID: "SpecialistDocumentOnRegisterList_Model",
		Relation: "specialist_documents_on_register_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
