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

type FirmList struct {
	Id fields.ValInt `json:"id" primaryKey:"true"`
	Name fields.ValText `json:"name" defOrder:"ASC"`
	Inn fields.ValText `json:"inn"`
	Ogrn fields.ValText `json:"ogrn"`
}

func (o *FirmList) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Inn.SetNull()
	o.Ogrn.SetNull()
}

func NewModelMD_FirmList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(FirmList{})),
		ID: "FirmList_Model",
		Relation: "firms_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
