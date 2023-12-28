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

type Specialist_check_inn struct {
	Inn fields.ValText `json:"inn" required:"true"`
	Captcha_key fields.ValText `json:"captcha_key" required:"true"`
	Operation_id fields.ValText `json:"operation_id" required:"true"`
}

type Specialist_check_inn_argv struct {
	Argv *Specialist_check_inn `json:"argv"`	
}

