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

type DocumentTemplateDialog struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true" alias:"Код"`
	Name fields.ValText `json:"name" alias:"Описание шаблона"`
	Sql_query fields.ValText `json:"sql_query"`
	Fields fields.ValJSON `json:"fields"`
	File_preview fields.ValJSON `json:"file_preview"`
	Need_signing fields.ValBool `json:"need_signing"`
	Sign_image_name fields.ValText `json:"sign_image_name"`
}

func (o *DocumentTemplateDialog) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Sql_query.SetNull()
	o.Fields.SetNull()
	o.File_preview.SetNull()
	o.Need_signing.SetNull()
	o.Sign_image_name.SetNull()
}

func NewModelMD_DocumentTemplateDialog() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(DocumentTemplateDialog{})),
		ID: "DocumentTemplateDialog_Model",
		Relation: "document_templates_dialog",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
