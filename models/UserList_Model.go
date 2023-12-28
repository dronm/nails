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

type UserList struct {
	Id fields.ValInt `json:"id" primaryKey:"true"`
	Name fields.ValText `json:"name" defOrder:"ASC"`
	Name_full fields.ValText `json:"name_full"`
	Role_id enums.ValEnum_role_types `json:"role_id"`
	Photo fields.ValJSON `json:"photo"`
	Tel fields.ValText `json:"tel"`
	Email fields.ValText `json:"email"`
}

func (o *UserList) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Name_full.SetNull()
	o.Role_id.SetNull()
	o.Photo.SetNull()
	o.Tel.SetNull()
	o.Email.SetNull()
}

func NewModelMD_UserList() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(UserList{})),
		ID: "UserList_Model",
		Relation: "users_list",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
