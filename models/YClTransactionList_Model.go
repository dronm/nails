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

type YClTransactionList struct {
	Id fields.ValInt `json:"id" primaryKey:"true"`
	Document_id fields.ValInt `json:"document_id"`
	Date fields.ValDateTimeTZ `json:"date"`
	Amount fields.ValFloat `json:"amount"`
	Seance_length fields.ValFloat `json:"seance_length"`
	Client fields.ValText `json:"client"`
	Client_phone fields.ValText `json:"client_phone"`
	Specialists_ref fields.ValJSON `json:"specialists_ref"`
	Ycl_staff_ref fields.ValJSON `json:"ycl_staff_ref"`
	Specialist_id fields.ValInt `json:"specialist_id"`
}

func (o *YClTransactionList) SetNull() {
	o.Id.SetNull()
	o.Document_id.SetNull()
	o.Date.SetNull()
	o.Amount.SetNull()
	o.Seance_length.SetNull()
	o.Client.SetNull()
	o.Client_phone.SetNull()
	o.Specialists_ref.SetNull()
	o.Ycl_staff_ref.SetNull()
	o.Specialist_id.SetNull()
}

func NewModelMD_YClTransactionList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(YClTransactionList{})),
		ID: "YClTransactionList_Model",
		Relation: "ycl_transactions_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "total_amount", Expr: "sum(amount)"},
&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
