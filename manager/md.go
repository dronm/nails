package main

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/md.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 */

import (
	"/controllers"
	"/constants"

	"osbe"
	"osbe/model"
)

func initMD(md *osbe.Metadata){
	md.Controllers["BankPayment"] = controllers.NewController_BankPayment()
	md.Models["BankPayment"] = models.NewModelMD_BankPayment()md.Models["BankPayment"] = models.NewModelMD_BankPayment()

}
