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

type User struct {
	Id fields.ValInt `json:"id" primaryKey:"true" autoInc:"true"`
	Name fields.ValText `json:"name" required:"true"`
	Name_full fields.ValText `json:"name_full"`
	Role_id enums.ValEnum_role_types `json:"role_id" required:"true"`
	Pwd fields.ValText `json:"pwd"`
	Create_dt fields.ValDateTimeTZ `json:"create_dt"`
	Banned fields.ValBool `json:"banned"`
	Time_zone_locale_id fields.ValInt `json:"time_zone_locale_id"`
	Locale_id enums.ValEnum_locales `json:"locale_id"`
}

func (o *User) SetNull() {
	o.Id.SetNull()
	o.Name.SetNull()
	o.Name_full.SetNull()
	o.Role_id.SetNull()
	o.Pwd.SetNull()
	o.Create_dt.SetNull()
	o.Banned.SetNull()
	o.Time_zone_locale_id.SetNull()
	o.Locale_id.SetNull()
}

func NewModelMD_User() *model.ModelMD{
	return &model.ModelMD{Fields: fields.GenModelMD(reflect.ValueOf(User{})),
		ID: "User_Model",
		Relation: "users",
		AggFunctions: []*model.AggFunction{
			&model.AggFunction{Alias: "totalCount", Expr: "count(*)"},
		},
	}
}
//for insert
type User_argv struct {
	Argv *User `json:"argv"`	
}

//Keys for delete/get object
type User_keys struct {
	Id fields.ValInt `json:"id"`
	Mode string `json:"mode" openMode:"true"` //open mode insert|copy|edit
}
type User_keys_argv struct {
	Argv *User_keys `json:"argv"`	
}

//old keys for update
type User_old_keys struct {
	Old_id fields.ValInt `json:"old_id"`
	Id fields.ValInt `json:"id"`
	Name fields.ValText `json:"name"`
	Name_full fields.ValText `json:"name_full"`
	Role_id enums.ValEnum_role_types `json:"role_id"`
	Pwd fields.ValText `json:"pwd"`
	Create_dt fields.ValDateTimeTZ `json:"create_dt"`
	Banned fields.ValBool `json:"banned"`
	Time_zone_locale_id fields.ValInt `json:"time_zone_locale_id"`
	Locale_id enums.ValEnum_locales `json:"locale_id"`
}

type User_old_keys_argv struct {
	Argv *User_old_keys `json:"argv"`	
}

