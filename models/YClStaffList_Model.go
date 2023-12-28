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

type YClStaffList struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name"`
	Specialization fields.ValText `json:"specialization"`
	Avatar fields.ValText `json:"avatar"`
	Avatar_big fields.ValText `json:"avatar_big"`
	Rating fields.ValFloat `json:"rating"`
	Votes_count fields.ValInt `json:"votes_count"`
}

func (o *YClStaffList) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Specialization.SetNull()
	o.Avatar.SetNull()
	o.Avatar_big.SetNull()
	o.Rating.SetNull()
	o.Votes_count.SetNull()
}

func NewModelMD_YClStaffList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(YClStaffList{})),
		ID: "YClStaffList_Model",
		Relation: "ycl_staff_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
