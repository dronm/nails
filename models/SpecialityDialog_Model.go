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

type SpecialityDialog struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true"`
	Equipments fields.ValJSON `json:"equipments" alias:"Оборудование"`
	Agent_percent fields.ValFloat `json:"agent_percent"`
}

func (o *SpecialityDialog) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Equipments.SetNull()
	o.Agent_percent.SetNull()
}

func NewModelMD_SpecialityDialog() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialityDialog{})),
		ID: "SpecialityDialog_Model",
		Relation: "specialities_dialog",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
