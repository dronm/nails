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

type SpecialistRegDialog struct {
	Id fields.ValInt `json:"id"`
	User_operation_id fields.ValText `json:"user_operation_id"`
	Inn fields.ValText `json:"inn"`
	Banks_ref fields.ValJSON `json:"banks_ref"`
	Bank_acc fields.ValText `json:"bank_acc"`
	Name_full fields.ValText `json:"name_full"`
	Tel fields.ValText `json:"tel"`
	Tel_confirmed fields.ValBool `json:"tel_confirmed"`
	Tel_sent fields.ValBool `json:"tel_sent"`
	Email fields.ValText `json:"email"`
	Email_confirmed fields.ValBool `json:"email_confirmed"`
	Email_sent fields.ValBool `json:"email_sent"`
	Address_reg fields.ValText `json:"address_reg"`
	Birthdate fields.ValDate `json:"birthdate"`
	Name fields.ValText `json:"name"`
	Passport fields.ValJSON `json:"passport"`
	Passport_uploaded fields.ValBool `json:"passport_uploaded"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
	Inn_checked fields.ValBool `json:"inn_checked"`
	Inn_fns_ok fields.ValBool `json:"inn_fns_ok"`
	Passport_preview fields.ValJSON `json:"passport_preview"`
	Operation_result fields.ValBool `json:"operation_result"`
}

func (o *SpecialistRegDialog) SetNull() {
	o.Id.SetNull()
	o.User_operation_id.SetNull()
	o.Inn.SetNull()
	o.Banks_ref.SetNull()
	o.Bank_acc.SetNull()
	o.Name_full.SetNull()
	o.Tel.SetNull()
	o.Tel_confirmed.SetNull()
	o.Tel_sent.SetNull()
	o.Email.SetNull()
	o.Email_confirmed.SetNull()
	o.Email_sent.SetNull()
	o.Address_reg.SetNull()
	o.Birthdate.SetNull()
	o.Name.SetNull()
	o.Passport.SetNull()
	o.Passport_uploaded.SetNull()
	o.Date_time.SetNull()
	o.Inn_checked.SetNull()
	o.Inn_fns_ok.SetNull()
	o.Passport_preview.SetNull()
	o.Operation_result.SetNull()
}

func NewModelMD_SpecialistRegDialog() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistRegDialog{})),
		ID: "SpecialistRegDialog_Model",
		Relation: "specialist_regs_dialog",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
