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

type SpecialistPeriodSalaryDetailList struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Specialist_period_salary_id fields.ValInt `json:"specialist_period_salary_id"`
	Line_num fields.ValInt `json:"line_num" alias:"Номер" defOrder:"ASC"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Specialists_ref fields.ValJSON `json:"specialists_ref"`
	Period fields.ValDate `json:"period"`
	Period_descr fields.ValText `json:"period_descr"`
	Studio_id fields.ValInt `json:"studio_id"`
	Studios_ref fields.ValJSON `json:"studios_ref"`
	Hours fields.ValInt `json:"hours" alias:"Часов"`
	Agent_percent fields.ValFloat `json:"agent_percent"`
	Work_total fields.ValFloat `json:"work_total" alias:"Выручка"`
	Work_total_salary fields.ValFloat `json:"work_total_salary" alias:"Выручка для зп"`
	Debet fields.ValFloat `json:"debet" alias:"Доп.начисления"`
	Kredit fields.ValFloat `json:"kredit" alias:"Цена за аренду"`
	Rent_price fields.ValFloat `json:"rent_price" alias:"Сумма за аренду"`
	Rent_total fields.ValFloat `json:"rent_total" alias:"Заработано"`
	Total fields.ValFloat `json:"total" alias:"К перечислению"`
	Receipt_total fields.ValFloat `json:"receipt_total" alias:"Сумма по чекам по данной строке"`
	Receipt_photos fields.ValJSON `json:"receipt_photos" alias:"Привью чеков"`
	Receipt_checked fields.ValBool `json:"receipt_checked" alias:"Чек проверен ФНС"`
	Receipt_error fields.ValText `json:"receipt_error" alias:"Текст ошибки проверки чека"`
	Receipt_href fields.ValText `json:"receipt_href" alias:"Ссылка на чек"`
	Bank_payments_ref fields.ValJSON `json:"bank_payments_ref" alias:"Платежное поручение"`
	Descr fields.ValText `json:"descr" alias:"Представление"`
}

func (o *SpecialistPeriodSalaryDetailList) SetNull() {
	o.Id.SetNull()
	o.Specialist_period_salary_id.SetNull()
	o.Line_num.SetNull()
	o.Specialist_id.SetNull()
	o.Specialists_ref.SetNull()
	o.Period.SetNull()
	o.Period_descr.SetNull()
	o.Studio_id.SetNull()
	o.Studios_ref.SetNull()
	o.Hours.SetNull()
	o.Agent_percent.SetNull()
	o.Work_total.SetNull()
	o.Work_total_salary.SetNull()
	o.Debet.SetNull()
	o.Kredit.SetNull()
	o.Rent_price.SetNull()
	o.Rent_total.SetNull()
	o.Total.SetNull()
	o.Receipt_total.SetNull()
	o.Receipt_photos.SetNull()
	o.Receipt_checked.SetNull()
	o.Receipt_error.SetNull()
	o.Receipt_href.SetNull()
	o.Bank_payments_ref.SetNull()
	o.Descr.SetNull()
}

func NewModelMD_SpecialistPeriodSalaryDetailList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistPeriodSalaryDetailList{})),
		ID: "SpecialistPeriodSalaryDetailList_Model",
		Relation: "specialist_period_salary_details_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "total_hours", Expr: "sum(hours)"},
&model.AggFunction{Alias: "total_work_total", Expr: "sum(work_total)"},
&model.AggFunction{Alias: "total_debet", Expr: "sum(debet)"},
&model.AggFunction{Alias: "total_kredit", Expr: "sum(kredit)"},
&model.AggFunction{Alias: "total_rent_total", Expr: "sum(rent_total)"},
&model.AggFunction{Alias: "total_total", Expr: "sum(total)"},
&model.AggFunction{Alias: "total_receipt_total", Expr: "sum(receipt_total)"},
&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
