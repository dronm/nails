package enums

/**
 * Andrey Mikhalevich 16/12/21
 * This file is part of the OSBE framework
 * 
 * This file is generated from the template build/templates/enum.go.tmpl
 * All direct modification will be lost with the next build.
 * Edit template instead.
*/

import (
	"osbe/fields"
)

type ValEnum_role_types struct {
	fields.ValText
}

func (e *ValEnum_role_types) GetValues() []string {
	return []string{ "admin", "specialist", "accountant" }
}

//func (e *ValEnum_role_types) GetDescriptions() map[string]map[string]string {
//	return make(map[string]{ "admin", "specialist", "accountant" }
//}

