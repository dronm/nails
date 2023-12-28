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

type SpecialistReg_get_by_operation_id struct {
	Operation_id fields.ValText `json:"operation_id" required:"true"`
}

type SpecialistReg_get_by_operation_id_argv struct {
	Argv *SpecialistReg_get_by_operation_id `json:"argv"`	
}

