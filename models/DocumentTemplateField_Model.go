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

type DocumentTemplateField struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Field fields.ValText `json:"field"`
	Comment_text fields.ValText `json:"comment_text"`
}

func (o *DocumentTemplateField) SetNull() {
	o.Id.SetNull()
	o.Field.SetNull()
	o.Comment_text.SetNull()
}

func NewModelMD_DocumentTemplateField() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(DocumentTemplateField{})),
		ID: "DocumentTemplateField_Model",
		Relation: "",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
