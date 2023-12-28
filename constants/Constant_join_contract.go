package constants

/**
 * Andrey Mikhalevich 20/12/21
 * This file is part of the OSBE framework
 * 
 * This file is generated from the template build/templates/constant.go.tmpl
 * All direct modification will be lost with the next build.
 * Edit template instead.
*/

import (
	"osbe"
)

type Constant_join_contract struct {
	osbe.ConstantText
}

func New_Constant_join_contract() *Constant_join_contract {
	return &Constant_join_contract{ osbe.ConstantText{ID: "join_contract", Autoload: false }}
}

