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

type YClTransactionDocList struct {
	Document_id fields.ValInt `json:"document_id" primaryKey:"true"`
	Date fields.ValDateTimeTZ `json:"date"`
	Amount fields.ValFloat `json:"amount"`
	Client fields.ValText `json:"client"`
	Seance_length fields.ValInt `json:"seance_length"`
	Specialist_id fields.ValInt `json:"specialist_id"`
}

func (o *YClTransactionDocList) SetNull() {
	o.Document_id.SetNull()
	o.Date.SetNull()
	o.Amount.SetNull()
	o.Client.SetNull()
	o.Seance_length.SetNull()
	o.Specialist_id.SetNull()
}

func NewModelMD_YClTransactionDocList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(YClTransactionDocList{})),
		ID: "YClTransactionDocList_Model",
		Relation: "ycl_transactions_doc_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
