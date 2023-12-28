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

type Constant_person_tax struct {
	osbe.ConstantInt
}

func New_Constant_person_tax() *Constant_person_tax {
	return &Constant_person_tax{ osbe.ConstantInt{ID: "person_tax", Autoload: false }}
}

