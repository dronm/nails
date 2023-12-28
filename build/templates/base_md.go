package main

/**
 * Andrey Mikhalevich, 18/12/2022
 * This file is part of the OSBE project.
 * It conains MD creation and basic initialization. Once created on project deployment it will not be changed.
 * 
 */

import (
	"osbe"
	"osbe/about"
	"osbe/evnt"
	"osbe/permission"
	osbe_const "osbe/constants"	
	"osbe/stat"
	"osbe/model"
)

func NewMD(projVersion string) *osbe.Metadata{
	md := osbe.NewMetadata()
	md.Version.Value = projVersion
		
	model.Cond_Model_init()
	
	md.Controllers["About"] = about.NewController_About()
	md.Controllers["Event"] = evnt.NewController_Event()
	md.Controllers["Permission"] = permission.NewController_Permission()
	
	md.Controllers["Constant"] = osbe_const.NewController_Constant()
	md.Models["ConstantValue"] = osbe_const.NewModelMD_ConstantValue()
	md.Models["ConstantList"] = osbe_const.NewModelMD_ConstantList()
	
	md.Controllers["SrvStatistics"] = stat.NewSrvStatistics_Controller()

	return md
}
