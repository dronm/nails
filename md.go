package main

/**
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/md.go.tmpl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 */

import (
	"nails/controllers"
	"nails/models"
	"nails/constants"

	"osbe"
)

func initMD(md *osbe.Metadata){	
	md.Controllers["User"] = controllers.NewController_User()
	md.Models["User"] = models.NewModelMD_User()
	md.Models["UserDialog"] = models.NewModelMD_UserDialog()
	md.Models["UserList"] = models.NewModelMD_UserList()
	md.Models["UserProfile"] = models.NewModelMD_UserProfile()
	
	md.Controllers["Firm"] = controllers.NewController_Firm()
	md.Models["Firm"] = models.NewModelMD_Firm()
	md.Models["FirmDialog"] = models.NewModelMD_FirmDialog()
	md.Models["FirmList"] = models.NewModelMD_FirmList()

	md.Controllers["Studio"] = controllers.NewController_Studio()
	md.Models["Studio"] = models.NewModelMD_Studio()
	md.Models["StudioDialog"] = models.NewModelMD_StudioDialog()
	md.Models["StudioList"] = models.NewModelMD_StudioList()

	md.Controllers["Specialist"] = controllers.NewController_Specialist()
	md.Models["Specialist"] = models.NewModelMD_Specialist()
	md.Models["SpecialistDialog"] = models.NewModelMD_SpecialistDialog()
	md.Models["SpecialistList"] = models.NewModelMD_SpecialistList()

	md.Controllers["SpecialistStatus"] = controllers.NewController_SpecialistStatus()
	md.Models["SpecialistStatus"] = models.NewModelMD_SpecialistStatus()

	md.Controllers["SpecialistReg"] = controllers.NewController_SpecialistReg()
	md.Models["SpecialistReg"] = models.NewModelMD_SpecialistReg()
	md.Models["SpecialistRegList"] = models.NewModelMD_SpecialistRegList()
	md.Models["SpecialistRegDialog"] = models.NewModelMD_SpecialistRegDialog()

	md.Controllers["DocumentTemplate"] = controllers.NewController_DocumentTemplate()
	md.Models["DocumentTemplate"] = models.NewModelMD_DocumentTemplate()
	md.Models["DocumentTemplateDialog"] = models.NewModelMD_DocumentTemplateDialog()
	md.Models["DocumentTemplateList"] = models.NewModelMD_DocumentTemplateList()

	//md.Controllers["SpecialistDocumentOnRegister"] = controllers.NewController_SpecialistDocumentOnRegister()
	//md.Models["SpecialistDocumentOnRegister"] = models.NewModelMD_SpecialistDocumentOnRegister()
	//md.Models["SpecialistDocumentOnRegisterList"] = models.NewModelMD_SpecialistDocumentOnRegisterList()

	md.Controllers["SpecialistDocument"] = controllers.NewController_SpecialistDocument()
	md.Models["SpecialistDocument"] = models.NewModelMD_SpecialistDocument()
	md.Models["SpecialistDocumentList"] = models.NewModelMD_SpecialistDocumentList()
	md.Models["SpecialistDocumentForSignList"] = models.NewModelMD_SpecialistDocumentForSignList()

	md.Controllers["SpecialistWork"] = controllers.NewController_SpecialistWork()
	md.Models["SpecialistWork"] = models.NewModelMD_SpecialistWork()
	md.Models["SpecialistWorkList"] = models.NewModelMD_SpecialistWorkList()
	md.Models["SpecialistWorkForRateList"] = models.NewModelMD_SpecialistWorkForRateList()

	md.Controllers["EquipmentType"] = controllers.NewController_EquipmentType()
	md.Models["EquipmentType"] = models.NewModelMD_EquipmentType()

	md.Controllers["Speciality"] = controllers.NewController_Speciality()
	md.Models["Speciality"] = models.NewModelMD_Speciality()
	md.Models["SpecialityList"] = models.NewModelMD_SpecialityList()
	md.Models["SpecialityDialog"] = models.NewModelMD_SpecialityDialog()

	md.Controllers["Item"] = controllers.NewController_Item()
	md.Models["Item"] = models.NewModelMD_Item()

	md.Controllers["YClStaff"] = controllers.NewController_YClStaff()
	md.Models["YClStaff"] = models.NewModelMD_YClStaff()
	md.Models["YClStaffList"] = models.NewModelMD_YClStaffList()

	md.Controllers["YClTransaction"] = controllers.NewController_YClTransaction()
	md.Models["YClTransaction"] = models.NewModelMD_YClTransaction()
	md.Models["YClTransactionList"] = models.NewModelMD_YClTransactionList()
	md.Models["YClTransactionDocList"] = models.NewModelMD_YClTransactionDocList()
	md.Models["YClTransactionDocAllList"] = models.NewModelMD_YClTransactionDocAllList()

	md.Controllers["SalaryDebet"] = controllers.NewController_SalaryDebet()
	md.Models["SalaryDebet"] = models.NewModelMD_SalaryDebet()
	md.Controllers["SalaryKredit"] = controllers.NewController_SalaryKredit()
	md.Models["SalaryKredit"] = models.NewModelMD_SalaryKredit()

	md.Controllers["SpecialistSalaryDebet"] = controllers.NewController_SpecialistSalaryDebet()
	md.Models["SpecialistSalaryDebet"] = models.NewModelMD_SpecialistSalaryDebet()
	md.Models["SpecialistSalaryDebetList"] = models.NewModelMD_SpecialistSalaryDebetList()

	md.Controllers["SpecialistSalaryKredit"] = controllers.NewController_SpecialistSalaryKredit()
	md.Models["SpecialistSalaryKredit"] = models.NewModelMD_SpecialistSalaryKredit()
	md.Models["SpecialistSalaryKreditList"] = models.NewModelMD_SpecialistSalaryKreditList()

	md.Controllers["SpecialistPeriodSalary"] = controllers.NewController_SpecialistPeriodSalary()
	md.Models["SpecialistPeriodSalary"] = models.NewModelMD_SpecialistPeriodSalary()
	md.Models["SpecialistPeriodSalaryList"] = models.NewModelMD_SpecialistPeriodSalaryList()

	md.Controllers["SpecialistPeriodSalaryDetail"] = controllers.NewController_SpecialistPeriodSalaryDetail()
	md.Models["SpecialistPeriodSalaryDetail"] = models.NewModelMD_SpecialistPeriodSalaryDetail()
	md.Models["SpecialistPeriodSalaryDetailList"] = models.NewModelMD_SpecialistPeriodSalaryDetailList()

	md.Controllers["SpecialistReceipt"] = controllers.NewController_SpecialistReceipt()
	md.Models["SpecialistReceipt"] = models.NewModelMD_SpecialistReceipt()

	md.Controllers["TemplateBatch"] = controllers.NewController_TemplateBatch()
	md.Models["TemplateBatch"] = models.NewModelMD_TemplateBatch()

	md.Controllers["TemplateBatchItem"] = controllers.NewController_TemplateBatchItem()
	md.Models["TemplateBatchItem"] = models.NewModelMD_TemplateBatchItem()
	md.Models["TemplateBatchItemList"] = models.NewModelMD_TemplateBatchItemList()

	md.Controllers["BankPayment"] = controllers.NewController_BankPayment()
	md.Models["BankPayment"] = models.NewModelMD_BankPayment()
	md.Models["BankPaymentList"] = models.NewModelMD_BankPaymentList()

	md.Controllers["ConfirmationStatus"] = controllers.NewController_ConfirmationStatus()
	md.Controllers["Mgateway"] = controllers.NewController_Mgateway()
	md.Controllers["QRExtractor"] = controllers.NewController_QRExtractor()

	//constants
	md.Constants["grid_refresh_interval"] = constants.New_Constant_grid_refresh_interval()
	md.Constants["doc_per_page_count"] = constants.New_Constant_doc_per_page_count()
	md.Constants["person_tax"] = constants.New_Constant_person_tax()
	md.Constants["email"] = constants.New_Constant_email()				
	md.Constants["join_contract"] = constants.New_Constant_join_contract()				
}
