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

type YClStaff struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name"`
	Data fields.ValJSON `json:"data"`
}

func (o *YClStaff) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Data.SetNull()
}

func NewModelMD_YClStaff() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(YClStaff{})),
		ID: "YClStaff_Model",
		Relation: "ycl_staff",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type YClStaff_argv struct {
	Argv *YClStaff `json:"argv"`	
}

//Keys for delete/get object
type YClStaff_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type YClStaff_keys_argv struct {
	Argv *YClStaff_keys `json:"argv"`	
}

//old keys for update
type YClStaff_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Name fields.ValText `json:"name"`
	Data fields.ValJSON `json:"data"`
}

type YClStaff_old_keys_argv struct {
	Argv *YClStaff_old_keys `json:"argv"`	
}

