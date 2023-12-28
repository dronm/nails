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

type SpecialityEquipment struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Equipment_types_ref fields.ValJSON `json:"equipment_types_ref" alias:"Вид оборудования"`
	Quant fields.ValInt `json:"quant" alias:"Количество"`
	Measure_unit fields.ValText `json:"measure_unit" alias:"Едииница"`
}

func (o *SpecialityEquipment) SetNull() {
	o.Id.SetNull()
	o.Equipment_types_ref.SetNull()
	o.Quant.SetNull()
	o.Measure_unit.SetNull()
}

func NewModelMD_SpecialityEquipment() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialityEquipment{})),
		ID: "SpecialityEquipment_Model",
		Relation: "",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
