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

type SpecialistPeriodSalaryList struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Period fields.ValDate `json:"period" defOrder:"DESC"`
	Studio_id fields.ValInt `json:"studio_id"`
	Studios_ref fields.ValJSON `json:"studios_ref"`
	Period_descr fields.ValText `json:"period_descr"`
	Work_total fields.ValFloat `json:"work_total" alias:"Вырчка"`
	Work_total_salary fields.ValFloat `json:"work_total_salary" alias:"Вырчка для з/п"`
	Hours fields.ValInt `json:"hours" alias:"Всего часов"`
	Debet fields.ValFloat `json:"debet" alias:"Всего доп начислений"`
	Kredit fields.ValFloat `json:"kredit" alias:"Всего доп удержаний"`
	Rent_total fields.ValFloat `json:"rent_total" alias:"Всего доп удержаний"`
	Total fields.ValFloat `json:"total" alias:"К перечислению"`
}

func (o *SpecialistPeriodSalaryList) SetNull() {
	o.Id.SetNull()
	o.Date_time.SetNull()
	o.Period.SetNull()
	o.Studio_id.SetNull()
	o.Studios_ref.SetNull()
	o.Period_descr.SetNull()
	o.Work_total.SetNull()
	o.Work_total_salary.SetNull()
	o.Hours.SetNull()
	o.Debet.SetNull()
	o.Kredit.SetNull()
	o.Rent_total.SetNull()
	o.Total.SetNull()
}

func NewModelMD_SpecialistPeriodSalaryList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistPeriodSalaryList{})),
		ID: "SpecialistPeriodSalaryList_Model",
		Relation: "specialist_period_salaries_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "total_total", Expr: "sum(total)"},
&model.AggFunction{Alias: "total_work_total", Expr: "sum(work_total)"},
&model.AggFunction{Alias: "total_work_total_salary", Expr: "sum(work_total_salary)"},
&model.AggFunction{Alias: "total_hours", Expr: "sum(hours)"},
&model.AggFunction{Alias: "total_debet", Expr: "sum(debet)"},
&model.AggFunction{Alias: "total_kredit", Expr: "sum(kredit)"},
&model.AggFunction{Alias: "total_rent_total", Expr: "sum(rent_total)"},
&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
