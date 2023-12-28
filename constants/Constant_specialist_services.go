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

type Constant_specialist_services struct {
	osbe.ConstantText
}

func New_Constant_specialist_services() *Constant_specialist_services {
	return &Constant_specialist_services{ osbe.ConstantText{ID: "specialist_services", Autoload: false }}
}

