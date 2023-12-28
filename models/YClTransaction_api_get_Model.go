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

type YClTransaction_api_get struct {
	Operation_id fields.ValText `json:"operation_id" required:"true"`
	Date_from fields.ValDate `json:"date_from" required:"true"`
	Date_to fields.ValDate `json:"date_to" required:"true"`
}

type YClTransaction_api_get_argv struct {
	Argv *YClTransaction_api_get `json:"argv"`	
}

