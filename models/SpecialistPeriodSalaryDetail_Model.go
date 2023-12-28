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

type SpecialistPeriodSalaryDetail struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Specialist_period_salary_id fields.ValInt `json:"specialist_period_salary_id" required:"true"`
	Line_num fields.ValInt `json:"line_num" alias:"Номер"`
	Specialist_id fields.ValInt `json:"specialist_id" required:"true"`
	Studio_id fields.ValInt `json:"studio_id"`
	Period fields.ValDate `json:"period"`
	Hours fields.ValInt `json:"hours" alias:"Часов"`
	Agent_percent fields.ValFloat `json:"agent_percent"`
	Work_total fields.ValFloat `json:"work_total" alias:"Выручка"`
	Work_total_salary fields.ValFloat `json:"work_total_salary" alias:"Выручка для зп"`
	Debet fields.ValFloat `json:"debet" alias:"Доп.начисления"`
	Kredit fields.ValFloat `json:"kredit" alias:"Цена за аренду"`
	Rent_price fields.ValFloat `json:"rent_price" alias:"Сумма за аренду"`
	Rent_total fields.ValFloat `json:"rent_total" alias:"Заработано"`
	Total fields.ValFloat `json:"total" alias:"К перечислению"`
}

func (o *SpecialistPeriodSalaryDetail) SetNull() {
	o.Id.SetNull()
	o.Specialist_period_salary_id.SetNull()
	o.Line_num.SetNull()
	o.Specialist_id.SetNull()
	o.Studio_id.SetNull()
	o.Period.SetNull()
	o.Hours.SetNull()
	o.Agent_percent.SetNull()
	o.Work_total.SetNull()
	o.Work_total_salary.SetNull()
	o.Debet.SetNull()
	o.Kredit.SetNull()
	o.Rent_price.SetNull()
	o.Rent_total.SetNull()
	o.Total.SetNull()
}

func NewModelMD_SpecialistPeriodSalaryDetail() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistPeriodSalaryDetail{})),
		ID: "SpecialistPeriodSalaryDetail_Model",
		Relation: "specialist_period_salary_details",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type SpecialistPeriodSalaryDetail_argv struct {
	Argv *SpecialistPeriodSalaryDetail `json:"argv"`	
}

//Keys for delete/get object
type SpecialistPeriodSalaryDetail_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type SpecialistPeriodSalaryDetail_keys_argv struct {
	Argv *SpecialistPeriodSalaryDetail_keys `json:"argv"`	
}

//old keys for update
type SpecialistPeriodSalaryDetail_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Specialist_period_salary_id fields.ValInt `json:"specialist_period_salary_id"`
	Line_num fields.ValInt `json:"line_num" alias:"Номер"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Studio_id fields.ValInt `json:"studio_id"`
	Period fields.ValDate `json:"period"`
	Hours fields.ValInt `json:"hours" alias:"Часов"`
	Agent_percent fields.ValFloat `json:"agent_percent"`
	Work_total fields.ValFloat `json:"work_total" alias:"Выручка"`
	Work_total_salary fields.ValFloat `json:"work_total_salary" alias:"Выручка для зп"`
	Debet fields.ValFloat `json:"debet" alias:"Доп.начисления"`
	Kredit fields.ValFloat `json:"kredit" alias:"Цена за аренду"`
	Rent_price fields.ValFloat `json:"rent_price" alias:"Сумма за аренду"`
	Rent_total fields.ValFloat `json:"rent_total" alias:"Заработано"`
	Total fields.ValFloat `json:"total" alias:"К перечислению"`
}

type SpecialistPeriodSalaryDetail_old_keys_argv struct {
	Argv *SpecialistPeriodSalaryDetail_old_keys `json:"argv"`	
}

