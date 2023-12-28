package models

/**
 * Andrey Mikhalevich 16/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/models/Model.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 */

//Controller method model
import (
		
	"osbe/fields"
)

type PassportVal struct {
	Series fields.ValText `json:"series" required:"true" length="4"`
	Num fields.ValText `json:"num" required:"true" length="6"`
	Issue_body fields.ValText `json:"issue_body" required:"true"`
	Issue_date fields.ValDate `json:"issue_date" required:"true"`
	Dep_code fields.ValText `json:"dep_code" required:"true" length="6"`
}

type BankKey struct {
	Bik string `json:"bik"`
}

type BankRef struct {
	Keys BankKey `json:"keys"`
	Descr string `json:"descr"`
	DataType string `json:"dataType"`
}

type Specialist_register struct {
	Operation_id fields.ValText `json:"operation_id" required:"true"`
	Inn fields.ValText `json:"inn" required:"true"`
	Banks_ref BankRef `json:"banks_ref" required:"true"`
	Bank_acc fields.ValText `json:"bank_acc" required:"true"`
	Name_full fields.ValText `json:"name_full" required:"true"`
	Name fields.ValText `json:"name" required:"false"`
	Email fields.ValText `json:"email" required:"true"`
	Tel fields.ValText `json:"tel" required:"true"`
	Birthdate fields.ValDate `json:"birthdate" required:"true"`
	Address_reg fields.ValText `json:"address_reg" required:"true"`
	Captcha fields.ValText `json:"captcha" required:"true"`
	Passport PassportVal `json:"passport" required:"true"`
	Passport_file_content fields.ValBytea `json:"passport_file_content"`
}

type Specialist_register_argv struct {
	Argv *Specialist_register `json:"argv"`	
}

