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

type Constant_email struct {
	osbe.ConstantJSON
}

func New_Constant_email() *Constant_email {
	return &Constant_email{ osbe.ConstantJSON{ID: "email", Autoload: false }}
}

