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

type FirmDialog struct {
	Id fields.ValInt `json:"id" primaryKey:"true"`
	Name fields.ValText `json:"name"`
	Inn fields.ValText `json:"inn"`
	Name_full fields.ValText `json:"name_full"`
	Legal_address fields.ValText `json:"legal_address"`
	Post_address fields.ValText `json:"post_address"`
	Kpp fields.ValText `json:"kpp"`
	Ogrn fields.ValText `json:"ogrn"`
	Okpo fields.ValText `json:"okpo"`
	Okved fields.ValText `json:"okved"`
	Bank_acc fields.ValText `json:"bank_acc"`
	Banks_ref fields.ValJSON `json:"banks_ref"`
}

func (o *FirmDialog) SetNull() {
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
	o.Banks_ref.SetNull()
}

func NewModelMD_FirmDialog() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(FirmDialog{})),
		ID: "FirmDialog_Model",
		Relation: "firms_dialog",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
