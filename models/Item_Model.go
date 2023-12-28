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

type Item struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true" defOrder:"ASC"`
	Comment_text fields.ValText `json:"comment_text"`
	Price fields.ValFloat `json:"price"`
}

func (o *Item) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Comment_text.SetNull()
	o.Price.SetNull()
}

func NewModelMD_Item() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(Item{})),
		ID: "Item_Model",
		Relation: "items",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type Item_argv struct {
	Argv *Item `json:"argv"`	
}

//Keys for delete/get object
type Item_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type Item_keys_argv struct {
	Argv *Item_keys `json:"argv"`	
}

//old keys for update
type Item_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Name fields.ValText `json:"name"`
	Comment_text fields.ValText `json:"comment_text"`
	Price fields.ValFloat `json:"price"`
}

type Item_old_keys_argv struct {
	Argv *Item_old_keys `json:"argv"`	
}

