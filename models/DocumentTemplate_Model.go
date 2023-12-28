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

type DocumentTemplate struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true" alias:"Код"`
	Name fields.ValText `json:"name" alias:"Описание шаблона"`
	Sql_query fields.ValText `json:"sql_query"`
	Fields fields.ValJSON `json:"fields"`
	Need_signing fields.ValBool `json:"need_signing"`
	Sign_image_name fields.ValText `json:"sign_image_name"`
}

func (o *DocumentTemplate) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Sql_query.SetNull()
	o.Fields.SetNull()
	o.Need_signing.SetNull()
	o.Sign_image_name.SetNull()
}

func NewModelMD_DocumentTemplate() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(DocumentTemplate{})),
		ID: "DocumentTemplate_Model",
		Relation: "document_templates",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type DocumentTemplate_argv struct {
	Argv *DocumentTemplate `json:"argv"`	
}

//Keys for delete/get object
type DocumentTemplate_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type DocumentTemplate_keys_argv struct {
	Argv *DocumentTemplate_keys `json:"argv"`	
}

//old keys for update
type DocumentTemplate_old_keys struct {
	Old_id fields.ValInt `json:"old_id" alias:"Код"`
	Id fields.ValInt `json:"id" alias:"Код"`
	Name fields.ValText `json:"name" alias:"Описание шаблона"`
	Sql_query fields.ValText `json:"sql_query"`
	Fields fields.ValJSON `json:"fields"`
	Need_signing fields.ValBool `json:"need_signing"`
	Sign_image_name fields.ValText `json:"sign_image_name"`
}

type DocumentTemplate_old_keys_argv struct {
	Argv *DocumentTemplate_old_keys `json:"argv"`	
}

