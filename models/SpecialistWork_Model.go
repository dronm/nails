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

type SpecialistWork struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Specialist_id fields.ValInt `json:"specialist_id" required:"true"`
	Studio_id fields.ValInt `json:"studio_id" required:"true"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Admin_rate fields.ValInt `json:"admin_rate"`
	Ycl_document_id fields.ValInt `json:"ycl_document_id"`
}

func (o *SpecialistWork) SetNull() {
	o.Id.SetNull()
	o.Specialist_id.SetNull()
	o.Studio_id.SetNull()
	o.Date_time.SetNull()
	o.Admin_rate.SetNull()
	o.Ycl_document_id.SetNull()
}

func NewModelMD_SpecialistWork() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistWork{})),
		ID: "SpecialistWork_Model",
		Relation: "specialist_works",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type SpecialistWork_argv struct {
	Argv *SpecialistWork `json:"argv"`	
}

//Keys for delete/get object
type SpecialistWork_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type SpecialistWork_keys_argv struct {
	Argv *SpecialistWork_keys `json:"argv"`	
}

//old keys for update
type SpecialistWork_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Studio_id fields.ValInt `json:"studio_id"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Admin_rate fields.ValInt `json:"admin_rate"`
	Ycl_document_id fields.ValInt `json:"ycl_document_id"`
}

type SpecialistWork_old_keys_argv struct {
	Argv *SpecialistWork_old_keys `json:"argv"`	
}

