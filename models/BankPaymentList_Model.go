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

type BankPaymentList struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Date_time fields.ValDateTimeTZ `json:"date_time" alias:"Дата создания"`
	Document_date fields.ValDateTimeTZ `json:"document_date" alias:"Дата документа"`
	Document_num fields.ValInt `json:"document_num" alias:"Номер документа" defOrder:"DESC"`
	Document_total fields.ValFloat `json:"document_total" alias:"Сумма документа"`
	Document_comment fields.ValText `json:"document_comment" alias:"Назначение платежа"`
	Payer_acc fields.ValText `json:"payer_acc" alias:"Расчетный счет плательщика"`
	Payer_bank_acc fields.ValText `json:"payer_bank_acc" alias:"Кор. счет банка плательщика"`
	Payer_bank_bik fields.ValText `json:"payer_bank_bik" alias:"Бик банка плательщика"`
	Payer_bank fields.ValText `json:"payer_bank" alias:"Банк плательщика"`
	Payer_bank_place fields.ValText `json:"payer_bank_place" alias:"Город банка плательщика"`
	Rec_acc fields.ValText `json:"rec_acc" alias:"Расчетный счет получателя"`
	Rec_bank_acc fields.ValText `json:"rec_bank_acc" alias:"Кор. счет банка получателя"`
	Rec_bank_bik fields.ValText `json:"rec_bank_bik" alias:"Бик банка получателя"`
	Rec_bank fields.ValText `json:"rec_bank" alias:"Банк получателя"`
	Rec_bank_place fields.ValText `json:"rec_bank_place" alias:"Город банка получателя"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Specialists_ref fields.ValJSON `json:"specialists_ref"`
	Specialist_period_salary_detail_id fields.ValInt `json:"specialist_period_salary_detail_id"`
	Specialist_period_salary_details_ref fields.ValJSON `json:"specialist_period_salary_details_ref"`
}

func (o *BankPaymentList) SetNull() {
	o.Id.SetNull()
	o.Date_time.SetNull()
	o.Document_date.SetNull()
	o.Document_num.SetNull()
	o.Document_total.SetNull()
	o.Document_comment.SetNull()
	o.Payer_acc.SetNull()
	o.Payer_bank_acc.SetNull()
	o.Payer_bank_bik.SetNull()
	o.Payer_bank.SetNull()
	o.Payer_bank_place.SetNull()
	o.Rec_acc.SetNull()
	o.Rec_bank_acc.SetNull()
	o.Rec_bank_bik.SetNull()
	o.Rec_bank.SetNull()
	o.Rec_bank_place.SetNull()
	o.Specialist_id.SetNull()
	o.Specialists_ref.SetNull()
	o.Specialist_period_salary_detail_id.SetNull()
	o.Specialist_period_salary_details_ref.SetNull()
}

func NewModelMD_BankPaymentList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(BankPaymentList{})),
		ID: "BankPaymentList_Model",
		Relation: "bank_payments_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "total_document_total", Expr: "sum(document_total)"},
&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
