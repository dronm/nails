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

type ValEnum_notif_types struct {
	fields.ValText
}

func (e *ValEnum_notif_types) GetValues() []string {
	return []string{ "new_specialist", "tel_check", "email_check", "docs_for_sign", "signed_docs", "new_account" }
}

//func (e *ValEnum_notif_types) GetDescriptions() map[string]map[string]string {
//	return make(map[string]{ "new_specialist", "tel_check", "email_check", "docs_for_sign", "signed_docs", "new_account" }
//}

