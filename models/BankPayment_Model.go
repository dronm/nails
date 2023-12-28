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

type BankPayment struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Date_time fields.ValDateTimeTZ `json:"date_time" alias:"Дата создания"`
	Document_date fields.ValDateTimeTZ `json:"document_date" alias:"Дата документа"`
	Document_num fields.ValInt `json:"document_num" alias:"Номер документа"`
	Document_total fields.ValFloat `json:"document_total" alias:"Сумма документа"`
	Document_comment fields.ValText `json:"document_comment" alias:"Назначение платежа"`
	Payer fields.ValText `json:"payer" alias:"Наименование плательщика"`
	Payer_acc fields.ValText `json:"payer_acc" alias:"Расчетный счет плательщика"`
	Payer_inn fields.ValText `json:"payer_inn" alias:"ИНН получателя"`
	Payer_bank_acc fields.ValText `json:"payer_bank_acc" alias:"Кор. счет банка плательщика"`
	Payer_bank_bik fields.ValText `json:"payer_bank_bik" alias:"Бик банка плательщика"`
	Payer_bank fields.ValText `json:"payer_bank" alias:"Банк плательщика"`
	Payer_bank_place fields.ValText `json:"payer_bank_place" alias:"Город банка плательщика"`
	Rec fields.ValText `json:"rec" alias:"Наименование получателя"`
	Rec_inn fields.ValText `json:"rec_inn" alias:"ИНН получателя"`
	Rec_acc fields.ValText `json:"rec_acc" alias:"Расчетный счет получателя"`
	Rec_bank_acc fields.ValText `json:"rec_bank_acc" alias:"Кор. счет банка получателя"`
	Rec_bank_bik fields.ValText `json:"rec_bank_bik" alias:"Бик банка получателя"`
	Rec_bank fields.ValText `json:"rec_bank" alias:"Банк получателя"`
	Rec_bank_place fields.ValText `json:"rec_bank_place" alias:"Город банка получателя"`
	Specialist_id fields.ValInt `json:"specialist_id" required:"true"`
	Specialist_period_salary_detail_id fields.ValInt `json:"specialist_period_salary_detail_id" required:"true"`
}

func (o *BankPayment) SetNull() {
	o.Id.SetNull()
	o.Date_time.SetNull()
	o.Document_date.SetNull()
	o.Document_num.SetNull()
	o.Document_total.SetNull()
	o.Document_comment.SetNull()
	o.Payer.SetNull()
	o.Payer_acc.SetNull()
	o.Payer_inn.SetNull()
	o.Payer_bank_acc.SetNull()
	o.Payer_bank_bik.SetNull()
	o.Payer_bank.SetNull()
	o.Payer_bank_place.SetNull()
	o.Rec.SetNull()
	o.Rec_inn.SetNull()
	o.Rec_acc.SetNull()
	o.Rec_bank_acc.SetNull()
	o.Rec_bank_bik.SetNull()
	o.Rec_bank.SetNull()
	o.Rec_bank_place.SetNull()
	o.Specialist_id.SetNull()
	o.Specialist_period_salary_detail_id.SetNull()
}

func NewModelMD_BankPayment() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(BankPayment{})),
		ID: "BankPayment_Model",
		Relation: "bank_payments",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type BankPayment_argv struct {
	Argv *BankPayment `json:"argv"`	
}

//Keys for delete/get object
type BankPayment_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type BankPayment_keys_argv struct {
	Argv *BankPayment_keys `json:"argv"`	
}

//old keys for update
type BankPayment_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Date_time fields.ValDateTimeTZ `json:"date_time" alias:"Дата создания"`
	Document_date fields.ValDateTimeTZ `json:"document_date" alias:"Дата документа"`
	Document_num fields.ValInt `json:"document_num" alias:"Номер документа"`
	Document_total fields.ValFloat `json:"document_total" alias:"Сумма документа"`
	Document_comment fields.ValText `json:"document_comment" alias:"Назначение платежа"`
	Payer fields.ValText `json:"payer" alias:"Наименование плательщика"`
	Payer_acc fields.ValText `json:"payer_acc" alias:"Расчетный счет плательщика"`
	Payer_inn fields.ValText `json:"payer_inn" alias:"ИНН получателя"`
	Payer_bank_acc fields.ValText `json:"payer_bank_acc" alias:"Кор. счет банка плательщика"`
	Payer_bank_bik fields.ValText `json:"payer_bank_bik" alias:"Бик банка плательщика"`
	Payer_bank fields.ValText `json:"payer_bank" alias:"Банк плательщика"`
	Payer_bank_place fields.ValText `json:"payer_bank_place" alias:"Город банка плательщика"`
	Rec fields.ValText `json:"rec" alias:"Наименование получателя"`
	Rec_inn fields.ValText `json:"rec_inn" alias:"ИНН получателя"`
	Rec_acc fields.ValText `json:"rec_acc" alias:"Расчетный счет получателя"`
	Rec_bank_acc fields.ValText `json:"rec_bank_acc" alias:"Кор. счет банка получателя"`
	Rec_bank_bik fields.ValText `json:"rec_bank_bik" alias:"Бик банка получателя"`
	Rec_bank fields.ValText `json:"rec_bank" alias:"Банк получателя"`
	Rec_bank_place fields.ValText `json:"rec_bank_place" alias:"Город банка получателя"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Specialist_period_salary_detail_id fields.ValInt `json:"specialist_period_salary_detail_id"`
}

type BankPayment_old_keys_argv struct {
	Argv *BankPayment_old_keys `json:"argv"`	
}

