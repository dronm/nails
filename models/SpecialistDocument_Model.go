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

type SpecialistDocument struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Specialist_id fields.ValInt `json:"specialist_id" required:"true"`
	Template_att_id fields.ValInt `json:"template_att_id" required:"true"`
	Document_att_id fields.ValInt `json:"document_att_id" required:"true"`
	Sign_img fields.ValBytea `json:"sign_img" alias:"Изображение подписи"`
	Date_time fields.ValDateTimeTZ `json:"date_time" required:"true"`
	Sign_date_time fields.ValDateTimeTZ `json:"sign_date_time"`
	Open_date_time fields.ValDateTimeTZ `json:"open_date_time"`
	Need_signing fields.ValBool `json:"need_signing"`
	Name fields.ValText `json:"name" alias:"Описание, представление"`
	Document_template_id fields.ValInt `json:"document_template_id" required:"true"`
}

func (o *SpecialistDocument) SetNull() {
	o.Id.SetNull()
	o.Specialist_id.SetNull()
	o.Template_att_id.SetNull()
	o.Document_att_id.SetNull()
	o.Sign_img.SetNull()
	o.Date_time.SetNull()
	o.Sign_date_time.SetNull()
	o.Open_date_time.SetNull()
	o.Need_signing.SetNull()
	o.Name.SetNull()
	o.Document_template_id.SetNull()
}

func NewModelMD_SpecialistDocument() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistDocument{})),
		ID: "SpecialistDocument_Model",
		Relation: "specialist_documents",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type SpecialistDocument_argv struct {
	Argv *SpecialistDocument `json:"argv"`	
}

//Keys for delete/get object
type SpecialistDocument_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type SpecialistDocument_keys_argv struct {
	Argv *SpecialistDocument_keys `json:"argv"`	
}

//old keys for update
type SpecialistDocument_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Template_att_id fields.ValInt `json:"template_att_id"`
	Document_att_id fields.ValInt `json:"document_att_id"`
	Sign_img fields.ValBytea `json:"sign_img" alias:"Изображение подписи"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Sign_date_time fields.ValDateTimeTZ `json:"sign_date_time"`
	Open_date_time fields.ValDateTimeTZ `json:"open_date_time"`
	Need_signing fields.ValBool `json:"need_signing"`
	Name fields.ValText `json:"name" alias:"Описание, представление"`
	Document_template_id fields.ValInt `json:"document_template_id"`
}

type SpecialistDocument_old_keys_argv struct {
	Argv *SpecialistDocument_old_keys `json:"argv"`	
}

