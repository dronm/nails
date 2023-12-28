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

type UserProfile struct {
	Id fields.ValInt `json:"id" primaryKey:"true" alias:"Код"`
	Name fields.ValText `json:"name" alias:"Имя"`
	Name_full fields.ValText `json:"name_full"`
	Locale_id fields.ValText `json:"locale_id"`
	Time_zone_locales_ref fields.ValJSON `json:"time_zone_locales_ref"`
	Photo fields.ValJSON `json:"photo"`
}

func (o *UserProfile) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Name_full.SetNull()
	o.Locale_id.SetNull()
	o.Time_zone_locales_ref.SetNull()
	o.Photo.SetNull()
}

func NewModelMD_UserProfile() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(UserProfile{})),
		ID: "UserProfile_Model",
		Relation: "users_profile",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
