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

type TemplateBatchItem struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Template_batch_id fields.ValInt `json:"template_batch_id" required:"true"`
	Template_id fields.ValInt `json:"template_id" required:"true"`
	Studio_id fields.ValInt `json:"studio_id" alias:"Студия или пусто, если все"`
}

func (o *TemplateBatchItem) SetNull() {
	o.Id.SetNull()
	o.Template_batch_id.SetNull()
	o.Template_id.SetNull()
	o.Studio_id.SetNull()
}

func NewModelMD_TemplateBatchItem() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(TemplateBatchItem{})),
		ID: "TemplateBatchItem_Model",
		Relation: "template_batch_items",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type TemplateBatchItem_argv struct {
	Argv *TemplateBatchItem `json:"argv"`	
}

//Keys for delete/get object
type TemplateBatchItem_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type TemplateBatchItem_keys_argv struct {
	Argv *TemplateBatchItem_keys `json:"argv"`	
}

//old keys for update
type TemplateBatchItem_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Template_batch_id fields.ValInt `json:"template_batch_id"`
	Template_id fields.ValInt `json:"template_id"`
	Studio_id fields.ValInt `json:"studio_id" alias:"Студия или пусто, если все"`
}

type TemplateBatchItem_old_keys_argv struct {
	Argv *TemplateBatchItem_old_keys `json:"argv"`	
}

