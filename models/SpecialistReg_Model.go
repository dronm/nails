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

type SpecialistReg struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	User_operation_id fields.ValText `json:"user_operation_id" required:"true"`
	Name_full fields.ValText `json:"name_full" required:"true"`
	Name fields.ValText `json:"name" required:"true"`
	Inn fields.ValText `json:"inn" required:"true"`
	Email fields.ValText `json:"email" required:"true"`
	Email_sent fields.ValBool `json:"email_sent"`
	Tel fields.ValText `json:"tel" required:"true"`
	Tel_sent fields.ValBool `json:"tel_sent"`
	Banks_ref fields.ValJSON `json:"banks_ref"`
	Studio_id fields.ValInt `json:"studio_id"`
	Birthdate fields.ValDate `json:"birthdate" alias:"Дата рождения"`
	Address_reg fields.ValText `json:"address_reg" alias:"Адрес регистрации"`
	Passport fields.ValJSON `json:"passport" alias:"Паспорт"`
	Passport_uploaded fields.ValBool `json:"passport_uploaded"`
	Date_time fields.ValDateTimeTZ `json:"date_time" required:"true"`
	Inn_checked fields.ValBool `json:"inn_checked" required:"true"`
	Inn_fns_ok fields.ValBool `json:"inn_fns_ok" required:"true"`
}

func (o *SpecialistReg) SetNull() {
	o.Id.SetNull()
	o.User_operation_id.SetNull()
	o.Name_full.SetNull()
	o.Name.SetNull()
	o.Inn.SetNull()
	o.Email.SetNull()
	o.Email_sent.SetNull()
	o.Tel.SetNull()
	o.Tel_sent.SetNull()
	o.Banks_ref.SetNull()
	o.Studio_id.SetNull()
	o.Birthdate.SetNull()
	o.Address_reg.SetNull()
	o.Passport.SetNull()
	o.Passport_uploaded.SetNull()
	o.Date_time.SetNull()
	o.Inn_checked.SetNull()
	o.Inn_fns_ok.SetNull()
}

func NewModelMD_SpecialistReg() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistReg{})),
		ID: "SpecialistReg_Model",
		Relation: "specialist_regs",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type SpecialistReg_argv struct {
	Argv *SpecialistReg `json:"argv"`	
}

//Keys for delete/get object
type SpecialistReg_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type SpecialistReg_keys_argv struct {
	Argv *SpecialistReg_keys `json:"argv"`	
}

//old keys for update
type SpecialistReg_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	User_operation_id fields.ValText `json:"user_operation_id"`
	Name_full fields.ValText `json:"name_full"`
	Name fields.ValText `json:"name"`
	Inn fields.ValText `json:"inn"`
	Email fields.ValText `json:"email"`
	Email_sent fields.ValBool `json:"email_sent"`
	Tel fields.ValText `json:"tel"`
	Tel_sent fields.ValBool `json:"tel_sent"`
	Banks_ref fields.ValJSON `json:"banks_ref"`
	Studio_id fields.ValInt `json:"studio_id"`
	Birthdate fields.ValDate `json:"birthdate" alias:"Дата рождения"`
	Address_reg fields.ValText `json:"address_reg" alias:"Адрес регистрации"`
	Passport fields.ValJSON `json:"passport" alias:"Паспорт"`
	Passport_uploaded fields.ValBool `json:"passport_uploaded"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Inn_checked fields.ValBool `json:"inn_checked"`
	Inn_fns_ok fields.ValBool `json:"inn_fns_ok"`
}

type SpecialistReg_old_keys_argv struct {
	Argv *SpecialistReg_old_keys `json:"argv"`	
}

