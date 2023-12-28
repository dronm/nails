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

type Studio struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true"`
	Firm_id fields.ValInt `json:"firm_id" required:"true"`
	Equipments fields.ValJSON `json:"equipments" alias:"Список оборудования для одного рабочего места"`
	Hour_rent_price fields.ValFloat `json:"hour_rent_price"`
}

func (o *Studio) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Firm_id.SetNull()
	o.Equipments.SetNull()
	o.Hour_rent_price.SetNull()
}

func NewModelMD_Studio() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(Studio{})),
		ID: "Studio_Model",
		Relation: "studios",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type Studio_argv struct {
	Argv *Studio `json:"argv"`	
}

//Keys for delete/get object
type Studio_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type Studio_keys_argv struct {
	Argv *Studio_keys `json:"argv"`	
}

//old keys for update
type Studio_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Name fields.ValText `json:"name"`
	Firm_id fields.ValInt `json:"firm_id"`
	Equipments fields.ValJSON `json:"equipments" alias:"Список оборудования для одного рабочего места"`
	Hour_rent_price fields.ValFloat `json:"hour_rent_price"`
}

type Studio_old_keys_argv struct {
	Argv *Studio_old_keys `json:"argv"`	
}

