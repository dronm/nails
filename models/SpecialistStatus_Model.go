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
	"nails/enums"	
	"osbe/fields"
	"osbe/model"
)

type SpecialistStatus struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Specialist_id fields.ValInt `json:"specialist_id" required:"true"`
	Status_type enums.ValEnum_specialist_status_types `json:"status_type" required:"true"`
	Date_time fields.ValDateTimeTZ `json:"date_time" required:"true"`
}

func (o *SpecialistStatus) SetNull() {
	o.Id.SetNull()
	o.Specialist_id.SetNull()
	o.Status_type.SetNull()
	o.Date_time.SetNull()
}

func NewModelMD_SpecialistStatus() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(SpecialistStatus{})),
		ID: "SpecialistStatus_Model",
		Relation: "specialist_statuses",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type SpecialistStatus_argv struct {
	Argv *SpecialistStatus `json:"argv"`	
}

//Keys for delete/get object
type SpecialistStatus_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type SpecialistStatus_keys_argv struct {
	Argv *SpecialistStatus_keys `json:"argv"`	
}

//old keys for update
type SpecialistStatus_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Specialist_id fields.ValInt `json:"specialist_id"`
	Status_type enums.ValEnum_specialist_status_types `json:"status_type"`
	Date_time fields.ValDateTimeTZ `json:"date_time"`
}

type SpecialistStatus_old_keys_argv struct {
	Argv *SpecialistStatus_old_keys `json:"argv"`	
}

