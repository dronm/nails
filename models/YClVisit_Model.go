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

type YClVisit struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Created_at fields.ValDateTimeTZ `json:"created_at" required:"true"`
	Data fields.ValJSON `json:"data"`
	Specialist_id fields.ValInt `json:"specialist_id" required:"true"`
}

func (o *YClVisit) SetNull() {
	o.Id.SetNull()
	o.Created_at.SetNull()
	o.Data.SetNull()
	o.Specialist_id.SetNull()
}

func NewModelMD_YClVisit() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(YClVisit{})),
		ID: "YClVisit_Model",
		Relation: "ycl_visits",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type YClVisit_argv struct {
	Argv *YClVisit `json:"argv"`	
}

//Keys for delete/get object
type YClVisit_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type YClVisit_keys_argv struct {
	Argv *YClVisit_keys `json:"argv"`	
}

//old keys for update
type YClVisit_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Created_at fields.ValDateTimeTZ `json:"created_at"`
	Data fields.ValJSON `json:"data"`
	Specialist_id fields.ValInt `json:"specialist_id"`
}

type YClVisit_old_keys_argv struct {
	Argv *YClVisit_old_keys `json:"argv"`	
}

