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

type SpecialityList struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true" defOrder:"ASC"`
}

func (o *SpecialityList) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
}

func NewModelMD_SpecialityList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialityList{})),
		ID: "SpecialityList_Model",
		Relation: "specialities_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
