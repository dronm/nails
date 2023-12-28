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

type Speciality struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true"`
	Equipments fields.ValJSON `json:"equipments" alias:"Оборудование"`
	Agent_percent fields.ValFloat `json:"agent_percent"`
}

func (o *Speciality) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Equipments.SetNull()
	o.Agent_percent.SetNull()
}

func NewModelMD_Speciality() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(Speciality{})),
		ID: "Speciality_Model",
		Relation: "specialities",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type Speciality_argv struct {
	Argv *Speciality `json:"argv"`	
}

//Keys for delete/get object
type Speciality_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type Speciality_keys_argv struct {
	Argv *Speciality_keys `json:"argv"`	
}

//old keys for update
type Speciality_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Name fields.ValText `json:"name"`
	Equipments fields.ValJSON `json:"equipments" alias:"Оборудование"`
	Agent_percent fields.ValFloat `json:"agent_percent"`
}

type Speciality_old_keys_argv struct {
	Argv *Speciality_old_keys `json:"argv"`	
}

