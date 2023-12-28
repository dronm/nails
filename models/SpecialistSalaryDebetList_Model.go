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

type SpecialistSalaryDebetList struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Date_time fields.ValDateTimeTZ `json:"date_time" defOrder:"DESC"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Specialists_ref fields.ValJSON `json:"specialists_ref"`
	Salary_debets_ref fields.ValJSON `json:"salary_debets_ref"`
	Total fields.ValFloat `json:"total"`
}

func (o *SpecialistSalaryDebetList) SetNull() {
	o.Id.SetNull()
	o.Date_time.SetNull()
	o.Specialist_id.SetNull()
	o.Specialists_ref.SetNull()
	o.Salary_debets_ref.SetNull()
	o.Total.SetNull()
}

func NewModelMD_SpecialistSalaryDebetList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistSalaryDebetList{})),
		ID: "SpecialistSalaryDebetList_Model",
		Relation: "specialist_salary_debets_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "total_total", Expr: "sum(total)"},
&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
