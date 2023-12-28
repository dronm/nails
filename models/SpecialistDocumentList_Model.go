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

type SpecialistDocumentList struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Specialists_ref fields.ValJSON `json:"specialists_ref"`
	Document_att_ref fields.ValJSON `json:"document_att_ref"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Sign_date_time fields.ValDateTimeTZ `json:"sign_date_time"`
	Signed fields.ValBool `json:"signed"`
	Open_date_time fields.ValDateTimeTZ `json:"open_date_time"`
	Opened fields.ValBool `json:"opened"`
	Need_signing fields.ValBool `json:"need_signing"`
	Name fields.ValText `json:"name" alias:"Описание, представление"`
}

func (o *SpecialistDocumentList) SetNull() {
	o.Id.SetNull()
	o.Specialist_id.SetNull()
	o.Specialists_ref.SetNull()
	o.Document_att_ref.SetNull()
	o.Date_time.SetNull()
	o.Sign_date_time.SetNull()
	o.Signed.SetNull()
	o.Open_date_time.SetNull()
	o.Opened.SetNull()
	o.Need_signing.SetNull()
	o.Name.SetNull()
}

func NewModelMD_SpecialistDocumentList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistDocumentList{})),
		ID: "SpecialistDocumentList_Model",
		Relation: "specialist_documents_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
