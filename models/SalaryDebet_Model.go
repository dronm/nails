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

type SalaryDebet struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true" defOrder:"ASC"`
}

func (o *SalaryDebet) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
}

func NewModelMD_SalaryDebet() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SalaryDebet{})),
		ID: "SalaryDebet_Model",
		Relation: "salary_debets",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type SalaryDebet_argv struct {
	Argv *SalaryDebet `json:"argv"`	
}

//Keys for delete/get object
type SalaryDebet_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type SalaryDebet_keys_argv struct {
	Argv *SalaryDebet_keys `json:"argv"`	
}

//old keys for update
type SalaryDebet_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Name fields.ValText `json:"name"`
}

type SalaryDebet_old_keys_argv struct {
	Argv *SalaryDebet_old_keys `json:"argv"`	
}

