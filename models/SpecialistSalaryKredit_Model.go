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

type SpecialistSalaryKredit struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Date_time fields.ValDateTimeTZ `json:"date_time" defOrder:"ASC"`
	Specialist_id fields.ValInt `json:"specialist_id" required:"true"`
	Salary_kredit_id fields.ValInt `json:"salary_kredit_id" required:"true"`
	Total fields.ValFloat `json:"total" required:"true"`
}

func (o *SpecialistSalaryKredit) SetNull() {
	o.Id.SetNull()
	o.Date_time.SetNull()
	o.Specialist_id.SetNull()
	o.Salary_kredit_id.SetNull()
	o.Total.SetNull()
}

func NewModelMD_SpecialistSalaryKredit() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistSalaryKredit{})),
		ID: "SpecialistSalaryKredit_Model",
		Relation: "specialist_salary_kredits",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type SpecialistSalaryKredit_argv struct {
	Argv *SpecialistSalaryKredit `json:"argv"`	
}

//Keys for delete/get object
type SpecialistSalaryKredit_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type SpecialistSalaryKredit_keys_argv struct {
	Argv *SpecialistSalaryKredit_keys `json:"argv"`	
}

//old keys for update
type SpecialistSalaryKredit_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Salary_kredit_id fields.ValInt `json:"salary_kredit_id"`
	Total fields.ValFloat `json:"total"`
}

type SpecialistSalaryKredit_old_keys_argv struct {
	Argv *SpecialistSalaryKredit_old_keys `json:"argv"`	
}

