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

type UserDialog struct {
	Id fields.ValInt `json:"id" primaryKey:"true" alias:"Код"`
	Name fields.ValText `json:"name" alias:"Имя"`
	Name_full fields.ValText `json:"name_full"`
	Role_id fields.ValText `json:"role_id"`
	Create_dt fields.ValDateTimeTZ `json:"create_dt"`
	Banned fields.ValBool `json:"banned"`
	Time_zone_locales_ref fields.ValJSON `json:"time_zone_locales_ref"`
	Locale_id fields.ValText `json:"locale_id"`
	Photo fields.ValJSON `json:"photo"`
}

func (o *UserDialog) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Name_full.SetNull()
	o.Role_id.SetNull()
	o.Create_dt.SetNull()
	o.Banned.SetNull()
	o.Time_zone_locales_ref.SetNull()
	o.Locale_id.SetNull()
	o.Photo.SetNull()
}

func NewModelMD_UserDialog() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(UserDialog{})),
		ID: "UserDialog_Model",
		Relation: "users_dialog",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
