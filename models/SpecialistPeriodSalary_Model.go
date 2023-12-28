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

type SpecialistPeriodSalary struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Period fields.ValDate `json:"period" required:"true" defOrder:"DESC"`
	Total fields.ValFloat `json:"total"`
	Work_total fields.ValFloat `json:"work_total" alias:"Вырчка"`
	Work_total_salary fields.ValFloat `json:"work_total_salary" alias:"Вырчка для з/п"`
	Hours fields.ValInt `json:"hours" alias:"Всего часов"`
	Debet fields.ValFloat `json:"debet" alias:"Всего доп начислений"`
	Kredit fields.ValFloat `json:"kredit" alias:"Всего доп удержаний"`
	Rent_total fields.ValFloat `json:"rent_total" alias:"Всего доп удержаний"`
	Studio_id fields.ValInt `json:"studio_id" required:"true"`
}

func (o *SpecialistPeriodSalary) SetNull() {
	o.Id.SetNull()
	o.Date_time.SetNull()
	o.Period.SetNull()
	o.Total.SetNull()
	o.Work_total.SetNull()
	o.Work_total_salary.SetNull()
	o.Hours.SetNull()
	o.Debet.SetNull()
	o.Kredit.SetNull()
	o.Rent_total.SetNull()
	o.Studio_id.SetNull()
}

func NewModelMD_SpecialistPeriodSalary() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistPeriodSalary{})),
		ID: "SpecialistPeriodSalary_Model",
		Relation: "specialist_period_salaries",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type SpecialistPeriodSalary_argv struct {
	Argv *SpecialistPeriodSalary `json:"argv"`	
}

//Keys for delete/get object
type SpecialistPeriodSalary_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type SpecialistPeriodSalary_keys_argv struct {
	Argv *SpecialistPeriodSalary_keys `json:"argv"`	
}

//old keys for update
type SpecialistPeriodSalary_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Period fields.ValDate `json:"period"`
	Total fields.ValFloat `json:"total"`
	Work_total fields.ValFloat `json:"work_total" alias:"Вырчка"`
	Work_total_salary fields.ValFloat `json:"work_total_salary" alias:"Вырчка для з/п"`
	Hours fields.ValInt `json:"hours" alias:"Всего часов"`
	Debet fields.ValFloat `json:"debet" alias:"Всего доп начислений"`
	Kredit fields.ValFloat `json:"kredit" alias:"Всего доп удержаний"`
	Rent_total fields.ValFloat `json:"rent_total" alias:"Всего доп удержаний"`
	Studio_id fields.ValInt `json:"studio_id"`
}

type SpecialistPeriodSalary_old_keys_argv struct {
	Argv *SpecialistPeriodSalary_old_keys `json:"argv"`	
}

