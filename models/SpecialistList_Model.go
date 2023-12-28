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

type SpecialistList struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true"`
	Inn fields.ValText `json:"inn"`
	Tel fields.ValText `json:"tel"`
	Email fields.ValText `json:"email"`
	Banks_ref fields.ValJSON `json:"banks_ref"`
	Bank_acc fields.ValText `json:"bank_acc" alias:"Банковский счет"`
	Studio_id fields.ValInt `json:"studio_id" sysCol:"true"`
	Studios_ref fields.ValJSON `json:"studios_ref"`
	Photo fields.ValJSON `json:"photo"`
	Specialities_ref fields.ValJSON `json:"specialities_ref"`
}

func (o *SpecialistList) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Inn.SetNull()
	o.Tel.SetNull()
	o.Email.SetNull()
	o.Banks_ref.SetNull()
	o.Bank_acc.SetNull()
	o.Studio_id.SetNull()
	o.Studios_ref.SetNull()
	o.Photo.SetNull()
	o.Specialities_ref.SetNull()
}

func NewModelMD_SpecialistList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistList{})),
		ID: "SpecialistList_Model",
		Relation: "specialists_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
