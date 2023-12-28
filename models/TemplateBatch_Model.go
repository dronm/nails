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
	"nails/enums"	
	"osbe/fields"
	"osbe/model"
)

type TemplateBatch struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" alias:"Наименование пакета" defOrder:"ASC"`
	Template_batch_type enums.ValEnum_template_batch_types `json:"template_batch_type"`
	Studio_id fields.ValInt `json:"studio_id" alias:"Студия или пусто, если все"`
}

func (o *TemplateBatch) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Template_batch_type.SetNull()
	o.Studio_id.SetNull()
}

func NewModelMD_TemplateBatch() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(TemplateBatch{})),
		ID: "TemplateBatch_Model",
		Relation: "template_batches",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type TemplateBatch_argv struct {
	Argv *TemplateBatch `json:"argv"`	
}

//Keys for delete/get object
type TemplateBatch_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type TemplateBatch_keys_argv struct {
	Argv *TemplateBatch_keys `json:"argv"`	
}

//old keys for update
type TemplateBatch_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Name fields.ValText `json:"name" alias:"Наименование пакета"`
	Template_batch_type enums.ValEnum_template_batch_types `json:"template_batch_type"`
	Studio_id fields.ValInt `json:"studio_id" alias:"Студия или пусто, если все"`
}

type TemplateBatch_old_keys_argv struct {
	Argv *TemplateBatch_old_keys `json:"argv"`	
}

