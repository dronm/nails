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

type ConfirmationStatus struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Ref fields.ValJSON `json:"ref" required:"true" alias:"Объект проверки,"`
	Field fields.ValText `json:"field" required:"true"`
	Secret fields.ValText `json:"secret" required:"true"`
	Confirmed fields.ValBool `json:"confirmed" required:"true"`
	Tries fields.ValInt `json:"tries" required:"true"`
	Try_date_time fields.ValDateTimeTZ `json:"try_date_time"`
	Date_time fields.ValDateTimeTZ `json:"date_time" required:"true"`
}

func (o *ConfirmationStatus) SetNull() {
	o.Id.SetNull()
	o.Ref.SetNull()
	o.Field.SetNull()
	o.Secret.SetNull()
	o.Confirmed.SetNull()
	o.Tries.SetNull()
	o.Try_date_time.SetNull()
	o.Date_time.SetNull()
}

func NewModelMD_ConfirmationStatus() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(ConfirmationStatus{})),
		ID: "ConfirmationStatus_Model",
		Relation: "confirmation_status",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type ConfirmationStatus_argv struct {
	Argv *ConfirmationStatus `json:"argv"`	
}

//Keys for delete/get object
type ConfirmationStatus_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type ConfirmationStatus_keys_argv struct {
	Argv *ConfirmationStatus_keys `json:"argv"`	
}

//old keys for update
type ConfirmationStatus_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Ref fields.ValJSON `json:"ref" alias:"Объект проверки,"`
	Field fields.ValText `json:"field"`
	Secret fields.ValText `json:"secret"`
	Confirmed fields.ValBool `json:"confirmed"`
	Tries fields.ValInt `json:"tries"`
	Try_date_time fields.ValDateTimeTZ `json:"try_date_time"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
}

type ConfirmationStatus_old_keys_argv struct {
	Argv *ConfirmationStatus_old_keys `json:"argv"`	
}

