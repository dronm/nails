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

type Firm struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true"`
	Inn fields.ValText `json:"inn"`
	Name_full fields.ValText `json:"name_full" alias:"Полное наименование"`
	Legal_address fields.ValText `json:"legal_address" alias:"Адрес юридический"`
	Post_address fields.ValText `json:"post_address" alias:"Адрес почтовый"`
	Kpp fields.ValText `json:"kpp" alias:"КПП"`
	Ogrn fields.ValText `json:"ogrn"`
	Okpo fields.ValText `json:"okpo"`
	Okved fields.ValText `json:"okved"`
	Bank_acc fields.ValText `json:"bank_acc" alias:"Расчетный счет"`
	Bank_bik fields.ValText `json:"bank_bik" alias:"Банк"`
}

func (o *Firm) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Inn.SetNull()
	o.Name_full.SetNull()
	o.Legal_address.SetNull()
	o.Post_address.SetNull()
	o.Kpp.SetNull()
	o.Ogrn.SetNull()
	o.Okpo.SetNull()
	o.Okved.SetNull()
	o.Bank_acc.SetNull()
	o.Bank_bik.SetNull()
}

func NewModelMD_Firm() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(Firm{})),
		ID: "Firm_Model",
		Relation: "firms",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type Firm_argv struct {
	Argv *Firm `json:"argv"`	
}

//Keys for delete/get object
type Firm_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type Firm_keys_argv struct {
	Argv *Firm_keys `json:"argv"`	
}

//old keys for update
type Firm_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Name fields.ValText `json:"name"`
	Inn fields.ValText `json:"inn"`
	Name_full fields.ValText `json:"name_full" alias:"Полное наименование"`
	Legal_address fields.ValText `json:"legal_address" alias:"Адрес юридический"`
	Post_address fields.ValText `json:"post_address" alias:"Адрес почтовый"`
	Kpp fields.ValText `json:"kpp" alias:"КПП"`
	Ogrn fields.ValText `json:"ogrn"`
	Okpo fields.ValText `json:"okpo"`
	Okved fields.ValText `json:"okved"`
	Bank_acc fields.ValText `json:"bank_acc" alias:"Расчетный счет"`
	Bank_bik fields.ValText `json:"bank_bik" alias:"Банк"`
}

type Firm_old_keys_argv struct {
	Argv *Firm_old_keys `json:"argv"`	
}

