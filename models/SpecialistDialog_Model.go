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

type SpecialistDialog struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true"`
	Inn fields.ValText `json:"inn" required:"true"`
	Banks_ref fields.ValJSON `json:"banks_ref"`
	Bank_acc fields.ValText `json:"bank_acc" alias:"Банковский счет"`
	Studio_id fields.ValInt `json:"studio_id" sysCol:"true"`
	Studios_ref fields.ValJSON `json:"studios_ref"`
	Birthdate fields.ValDate `json:"birthdate" alias:"Дата рождения"`
	Address_reg fields.ValText `json:"address_reg" alias:"Адрес регистрации"`
	Passport fields.ValJSON `json:"passport" alias:"Паспорт"`
	Last_status_type fields.ValText `json:"last_status_type" alias:"Текущий статус"`
	Equipments fields.ValJSON `json:"equipments" alias:"Список оборудования для одного рабочего места"`
	Users_ref fields.ValJSON `json:"users_ref"`
	Passport_preview fields.ValJSON `json:"passport_preview"`
	Tel fields.ValText `json:"tel"`
	Email fields.ValText `json:"email"`
	Agent_percent fields.ValFloat `json:"agent_percent"`
	Specialities_ref fields.ValJSON `json:"specialities_ref"`
}

func (o *SpecialistDialog) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Inn.SetNull()
	o.Banks_ref.SetNull()
	o.Bank_acc.SetNull()
	o.Studio_id.SetNull()
	o.Studios_ref.SetNull()
	o.Birthdate.SetNull()
	o.Address_reg.SetNull()
	o.Passport.SetNull()
	o.Last_status_type.SetNull()
	o.Equipments.SetNull()
	o.Users_ref.SetNull()
	o.Passport_preview.SetNull()
	o.Tel.SetNull()
	o.Email.SetNull()
	o.Agent_percent.SetNull()
	o.Specialities_ref.SetNull()
}

func NewModelMD_SpecialistDialog() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistDialog{})),
		ID: "SpecialistDialog_Model",
		Relation: "specialists_dialog",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
