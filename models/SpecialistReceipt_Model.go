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

type SpecialistReceipt struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Date_time fields.ValDateTimeTZ `json:"date_time" alias:"Дата загрузки"`
	Document_date_time fields.ValDateTimeTZ `json:"document_date_time" alias:"Дата документа"`
	Document_total fields.ValFloat `json:"document_total" alias:"Сумма документа"`
	Document_parsed fields.ValBool `json:"document_parsed" alias:"Сумма проверен ФНС"`
	Specialist_id fields.ValInt `json:"specialist_id" required:"true"`
	Specialist_period_salary_detail_id fields.ValInt `json:"specialist_period_salary_detail_id" required:"true"`
	Document_error fields.ValText `json:"document_error" alias:"Ошибка проверки сека ФНС"`
	Qrextr_request_id fields.ValText `json:"qrextr_request_id" alias:"ИД сервиса qrextr"`
	Operation_id fields.ValText `json:"operation_id"`
	Document_href fields.ValText `json:"document_href"`
}

func (o *SpecialistReceipt) SetNull() {
	o.Id.SetNull()
	o.Date_time.SetNull()
	o.Document_date_time.SetNull()
	o.Document_total.SetNull()
	o.Document_parsed.SetNull()
	o.Specialist_id.SetNull()
	o.Specialist_period_salary_detail_id.SetNull()
	o.Document_error.SetNull()
	o.Qrextr_request_id.SetNull()
	o.Operation_id.SetNull()
	o.Document_href.SetNull()
}

func NewModelMD_SpecialistReceipt() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistReceipt{})),
		ID: "SpecialistReceipt_Model",
		Relation: "specialist_receipts",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type SpecialistReceipt_argv struct {
	Argv *SpecialistReceipt `json:"argv"`	
}

//Keys for delete/get object
type SpecialistReceipt_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type SpecialistReceipt_keys_argv struct {
	Argv *SpecialistReceipt_keys `json:"argv"`	
}

//old keys for update
type SpecialistReceipt_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Date_time fields.ValDateTimeTZ `json:"date_time" alias:"Дата загрузки"`
	Document_date_time fields.ValDateTimeTZ `json:"document_date_time" alias:"Дата документа"`
	Document_total fields.ValFloat `json:"document_total" alias:"Сумма документа"`
	Document_parsed fields.ValBool `json:"document_parsed" alias:"Сумма проверен ФНС"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Specialist_period_salary_detail_id fields.ValInt `json:"specialist_period_salary_detail_id"`
	Document_error fields.ValText `json:"document_error" alias:"Ошибка проверки сека ФНС"`
	Qrextr_request_id fields.ValText `json:"qrextr_request_id" alias:"ИД сервиса qrextr"`
	Operation_id fields.ValText `json:"operation_id"`
	Document_href fields.ValText `json:"document_href"`
}

type SpecialistReceipt_old_keys_argv struct {
	Argv *SpecialistReceipt_old_keys `json:"argv"`	
}

