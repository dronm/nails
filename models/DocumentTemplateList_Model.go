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

type DocumentTemplateList struct {
	Id fields.ValInt `json:"id" primaryKey:"true"`
	Name fields.ValText `json:"name"`
	File_preview fields.ValJSON `json:"file_preview"`
	Need_signing fields.ValBool `json:"need_signing"`
}

func (o *DocumentTemplateList) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.File_preview.SetNull()
	o.Need_signing.SetNull()
}

func NewModelMD_DocumentTemplateList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(DocumentTemplateList{})),
		ID: "DocumentTemplateList_Model",
		Relation: "document_templates_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
