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

type StudioEquipment struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" alias:"Наименование"`
	Quant fields.ValInt `json:"quant" alias:"Количество"`
	Measure_unit fields.ValText `json:"measure_unit" alias:"Единица"`
	Price fields.ValFloat `json:"price" alias:"Цена"`
	Total fields.ValFloat `json:"total" alias:"Сумма"`
}

func (o *StudioEquipment) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Quant.SetNull()
	o.Measure_unit.SetNull()
	o.Price.SetNull()
	o.Total.SetNull()
}

func NewModelMD_StudioEquipment() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(StudioEquipment{})),
		ID: "StudioEquipment_Model",
		Relation: "",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
