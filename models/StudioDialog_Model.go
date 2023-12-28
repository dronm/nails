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

type StudioDialog struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true"`
	Firm_id fields.ValInt `json:"firm_id" sysCol:"true"`
	Firms_ref fields.ValJSON `json:"firms_ref"`
	Equipments fields.ValJSON `json:"equipments" alias:"Список оборудования для одного рабочего места"`
	Hour_rent_price fields.ValFloat `json:"hour_rent_price"`
}

func (o *StudioDialog) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Firm_id.SetNull()
	o.Firms_ref.SetNull()
	o.Equipments.SetNull()
	o.Hour_rent_price.SetNull()
}

func NewModelMD_StudioDialog() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(StudioDialog{})),
		ID: "StudioDialog_Model",
		Relation: "studios_dialog",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
