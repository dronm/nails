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

type TemplateBatchItemList struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Template_batch_id fields.ValInt `json:"template_batch_id"`
	Templates_ref fields.ValJSON `json:"templates_ref"`
	Studios_ref fields.ValJSON `json:"studios_ref"`
}

func (o *TemplateBatchItemList) SetNull() {
	o.Id.SetNull()
	o.Template_batch_id.SetNull()
	o.Templates_ref.SetNull()
	o.Studios_ref.SetNull()
}

func NewModelMD_TemplateBatchItemList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(TemplateBatchItemList{})),
		ID: "TemplateBatchItemList_Model",
		Relation: "template_batch_items_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
