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

type SpecialistWorkList struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Specialists_ref fields.ValJSON `json:"specialists_ref"`
	Studio_id fields.ValInt `json:"studio_id"`
	Studios_ref fields.ValJSON `json:"studios_ref"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Photo fields.ValJSON `json:"photo"`
	Admin_rate fields.ValInt `json:"admin_rate"`
	Amount fields.ValFloat `json:"amount"`
	Hours fields.ValFloat `json:"hours"`
}

func (o *SpecialistWorkList) SetNull() {
	o.Id.SetNull()
	o.Specialist_id.SetNull()
	o.Specialists_ref.SetNull()
	o.Studio_id.SetNull()
	o.Studios_ref.SetNull()
	o.Date_time.SetNull()
	o.Photo.SetNull()
	o.Admin_rate.SetNull()
	o.Amount.SetNull()
	o.Hours.SetNull()
}

func NewModelMD_SpecialistWorkList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistWorkList{})),
		ID: "SpecialistWorkList_Model",
		Relation: "specialist_works_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
