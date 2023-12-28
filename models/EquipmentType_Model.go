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

type EquipmentType struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true" defOrder:"ASC"`
}

func (o *EquipmentType) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
}

func NewModelMD_EquipmentType() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(EquipmentType{})),
		ID: "EquipmentType_Model",
		Relation: "equipment_types",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type EquipmentType_argv struct {
	Argv *EquipmentType `json:"argv"`	
}

//Keys for delete/get object
type EquipmentType_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type EquipmentType_keys_argv struct {
	Argv *EquipmentType_keys `json:"argv"`	
}

//old keys for update
type EquipmentType_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Name fields.ValText `json:"name"`
}

type EquipmentType_old_keys_argv struct {
	Argv *EquipmentType_old_keys `json:"argv"`	
}

