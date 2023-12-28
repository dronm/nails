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

type YClTransactionDocAllList struct {
	Document_id fields.ValInt `json:"document_id" primaryKey:"true"`
	Date fields.ValDateTimeTZ `json:"date" defOrder:"DESC"`
	Amount fields.ValFloat `json:"amount"`
	Client fields.ValText `json:"client"`
	Len_hour fields.ValFloat `json:"len_hour"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Specialists_ref fields.ValJSON `json:"specialists_ref"`
	Photo fields.ValJSON `json:"photo"`
	Admin_rate fields.ValInt `json:"admin_rate"`
}

func (o *YClTransactionDocAllList) SetNull() {
	o.Document_id.SetNull()
	o.Date.SetNull()
	o.Amount.SetNull()
	o.Client.SetNull()
	o.Len_hour.SetNull()
	o.Specialist_id.SetNull()
	o.Specialists_ref.SetNull()
	o.Photo.SetNull()
	o.Admin_rate.SetNull()
}

func NewModelMD_YClTransactionDocAllList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(YClTransactionDocAllList{})),
		ID: "YClTransactionDocAllList_Model",
		Relation: "ycl_transactions_doc_all_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "total_amount", Expr: "sum(amount)"},
&model.AggFunction{Alias: "total_len_hour", Expr: "sum(len_hour)"},
&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
