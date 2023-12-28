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

type YClStaff_api_get_all struct {
	Operation_id fields.ValText `json:"operation_id" required:"true"`
}

type YClStaff_api_get_all_argv struct {
	Argv *YClStaff_api_get_all `json:"argv"`	
}

