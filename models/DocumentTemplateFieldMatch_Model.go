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

type DocumentTemplateFieldMatch struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Field fields.ValText `json:"field"`
	Cell fields.ValText `json:"cell"`
}

func (o *DocumentTemplateFieldMatch) SetNull() {
	o.Id.SetNull()
	o.Field.SetNull()
	o.Cell.SetNull()
}

func NewModelMD_DocumentTemplateFieldMatch() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(DocumentTemplateFieldMatch{})),
		ID: "DocumentTemplateFieldMatch_Model",
		Relation: "",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
