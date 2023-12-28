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

type SpecialistDocument_sign struct {
	Id fields.ValInt `json:"id" required:"true"`
	Signature fields.ValBytea `json:"signature"`
	Operation_id fields.ValText `json:"operation_id" required:"true"`
}

type SpecialistDocument_sign_argv struct {
	Argv *SpecialistDocument_sign `json:"argv"`	
}

