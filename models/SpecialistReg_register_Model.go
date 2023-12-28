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

type SpecialistReg_register struct {
	Id fields.ValInt `json:"id" required:"true"`
	Studio_id fields.ValInt `json:"studio_id" required:"true"`
	Ycl_staff_id fields.ValInt `json:"ycl_staff_id" required:"true"`
	Operation_id fields.ValText `json:"operation_id" required:"true"`
}

type SpecialistReg_register_argv struct {
	Argv *SpecialistReg_register `json:"argv"`	
}

