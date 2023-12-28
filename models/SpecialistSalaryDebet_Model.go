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

type SpecialistSalaryDebet struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Date_time fields.ValDateTimeTZ `json:"date_time" defOrder:"ASC"`
	Specialist_id fields.ValInt `json:"specialist_id" required:"true"`
	Salary_debet_id fields.ValInt `json:"salary_debet_id" required:"true"`
	Total fields.ValFloat `json:"total" required:"true"`
}

func (o *SpecialistSalaryDebet) SetNull() {
	o.Id.SetNull()
	o.Date_time.SetNull()
	o.Specialist_id.SetNull()
	o.Salary_debet_id.SetNull()
	o.Total.SetNull()
}

func NewModelMD_SpecialistSalaryDebet() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistSalaryDebet{})),
		ID: "SpecialistSalaryDebet_Model",
		Relation: "specialist_salary_debets",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type SpecialistSalaryDebet_argv struct {
	Argv *SpecialistSalaryDebet `json:"argv"`	
}

//Keys for delete/get object
type SpecialistSalaryDebet_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type SpecialistSalaryDebet_keys_argv struct {
	Argv *SpecialistSalaryDebet_keys `json:"argv"`	
}

//old keys for update
type SpecialistSalaryDebet_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Salary_debet_id fields.ValInt `json:"salary_debet_id"`
	Total fields.ValFloat `json:"total"`
}

type SpecialistSalaryDebet_old_keys_argv struct {
	Argv *SpecialistSalaryDebet_old_keys `json:"argv"`	
}

