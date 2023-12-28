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

type SpecialistDocumentForSignList struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	User_id fields.ValInt `json:"user_id"`
	Specialists_ref fields.ValJSON `json:"specialists_ref"`
	Document_att_ref fields.ValJSON `json:"document_att_ref"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Opened fields.ValBool `json:"opened"`
	Name fields.ValText `json:"name" alias:"Описание, представление"`
}

func (o *SpecialistDocumentForSignList) SetNull() {
	o.Id.SetNull()
	o.Specialist_id.SetNull()
	o.User_id.SetNull()
	o.Specialists_ref.SetNull()
	o.Document_att_ref.SetNull()
	o.Date_time.SetNull()
	o.Opened.SetNull()
	o.Name.SetNull()
}

func NewModelMD_SpecialistDocumentForSignList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistDocumentForSignList{})),
		ID: "SpecialistDocumentForSignList_Model",
		Relation: "specialist_documents_for_sign_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
