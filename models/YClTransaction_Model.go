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

type YClTransaction struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Created_at fields.ValDateTimeTZ `json:"created_at" required:"true"`
	Data fields.ValJSON `json:"data"`
	Specialist_id fields.ValInt `json:"specialist_id" required:"true"`
	Document_id fields.ValInt `json:"document_id"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Amount fields.ValFloat `json:"amount"`
	Seance_length fields.ValInt `json:"seance_length"`
	Record_id fields.ValInt `json:"record_id"`
	Staff_id fields.ValInt `json:"staff_id"`
	Record_inf_updated fields.ValBool `json:"record_inf_updated"`
}

func (o *YClTransaction) SetNull() {
	o.Id.SetNull()
	o.Created_at.SetNull()
	o.Data.SetNull()
	o.Specialist_id.SetNull()
	o.Document_id.SetNull()
	o.Date_time.SetNull()
	o.Amount.SetNull()
	o.Seance_length.SetNull()
	o.Record_id.SetNull()
	o.Staff_id.SetNull()
	o.Record_inf_updated.SetNull()
}

func NewModelMD_YClTransaction() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(YClTransaction{})),
		ID: "YClTransaction_Model",
		Relation: "ycl_transactions",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type YClTransaction_argv struct {
	Argv *YClTransaction `json:"argv"`	
}

//Keys for delete/get object
type YClTransaction_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type YClTransaction_keys_argv struct {
	Argv *YClTransaction_keys `json:"argv"`	
}

//old keys for update
type YClTransaction_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Created_at fields.ValDateTimeTZ `json:"created_at"`
	Data fields.ValJSON `json:"data"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Document_id fields.ValInt `json:"document_id"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Amount fields.ValFloat `json:"amount"`
	Seance_length fields.ValInt `json:"seance_length"`
	Record_id fields.ValInt `json:"record_id"`
	Staff_id fields.ValInt `json:"staff_id"`
	Record_inf_updated fields.ValBool `json:"record_inf_updated"`
}

type YClTransaction_old_keys_argv struct {
	Argv *YClTransaction_old_keys `json:"argv"`	
}

