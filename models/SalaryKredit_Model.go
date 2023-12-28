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

type SalaryKredit struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true" defOrder:"ASC"`
}

func (o *SalaryKredit) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
}

func NewModelMD_SalaryKredit() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SalaryKredit{})),
		ID: "SalaryKredit_Model",
		Relation: "salary_kredits",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type SalaryKredit_argv struct {
	Argv *SalaryKredit `json:"argv"`	
}

//Keys for delete/get object
type SalaryKredit_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type SalaryKredit_keys_argv struct {
	Argv *SalaryKredit_keys `json:"argv"`	
}

//old keys for update
type SalaryKredit_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Name fields.ValText `json:"name"`
}

type SalaryKredit_old_keys_argv struct {
	Argv *SalaryKredit_old_keys `json:"argv"`	
}

