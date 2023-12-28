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

type Specialist struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true"`
	Inn fields.ValText `json:"inn" required:"true"`
	Bank_bik fields.ValText `json:"bank_bik" alias:"БИК банка"`
	Bank_acc fields.ValText `json:"bank_acc" alias:"Банковский счет"`
	Studio_id fields.ValInt `json:"studio_id" required:"true"`
	Birthdate fields.ValDate `json:"birthdate" alias:"Дата рождения"`
	Address_reg fields.ValText `json:"address_reg" alias:"Адрес регистрации"`
	Passport fields.ValJSON `json:"passport" alias:"Паспорт"`
	Equipments fields.ValJSON `json:"equipments" alias:"Список оборудования для одного рабочего места"`
	User_id fields.ValInt `json:"user_id" required:"true"`
	Ycl_staff_id fields.ValInt `json:"ycl_staff_id"`
	Agent_percent fields.ValFloat `json:"agent_percent"`
	Speciality_id fields.ValInt `json:"speciality_id"`
}

func (o *Specialist) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Inn.SetNull()
	o.Bank_bik.SetNull()
	o.Bank_acc.SetNull()
	o.Studio_id.SetNull()
	o.Birthdate.SetNull()
	o.Address_reg.SetNull()
	o.Passport.SetNull()
	o.Equipments.SetNull()
	o.User_id.SetNull()
	o.Ycl_staff_id.SetNull()
	o.Agent_percent.SetNull()
	o.Speciality_id.SetNull()
}

func NewModelMD_Specialist() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(Specialist{})),
		ID: "Specialist_Model",
		Relation: "specialists",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type Specialist_argv struct {
	Argv *Specialist `json:"argv"`	
}

//Keys for delete/get object
type Specialist_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type Specialist_keys_argv struct {
	Argv *Specialist_keys `json:"argv"`	
}

//old keys for update
type Specialist_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Name fields.ValText `json:"name"`
	Inn fields.ValText `json:"inn"`
	Bank_bik fields.ValText `json:"bank_bik" alias:"БИК банка"`
	Bank_acc fields.ValText `json:"bank_acc" alias:"Банковский счет"`
	Studio_id fields.ValInt `json:"studio_id"`
	Birthdate fields.ValDate `json:"birthdate" alias:"Дата рождения"`
	Address_reg fields.ValText `json:"address_reg" alias:"Адрес регистрации"`
	Passport fields.ValJSON `json:"passport" alias:"Паспорт"`
	Equipments fields.ValJSON `json:"equipments" alias:"Список оборудования для одного рабочего места"`
	User_id fields.ValInt `json:"user_id"`
	Ycl_staff_id fields.ValInt `json:"ycl_staff_id"`
	Agent_percent fields.ValFloat `json:"agent_percent"`
	Speciality_id fields.ValInt `json:"speciality_id"`
}

type Specialist_old_keys_argv struct {
	Argv *Specialist_old_keys `json:"argv"`	
}

