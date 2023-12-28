/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023
 
 * @extends ViewObjectAjx.js
 * @requires core/extend.js  
 * @requires controls/ViewObjectAjx.js 
 
 * @class
 * @classdesc
	
 * @param {string} id view identifier
 * @param {object} options
 */	
function SpecialistDialog_View(id,options){	

	options = options || {};
	
	options.controller = new Specialist_Controller();
	options.model = (options.models&&options.models.SpecialistDialog_Model)? options.models.SpecialistDialog_Model : new SpecialistDialog_Model();
	
	//var valid_tel_pref = CommonHelper.unserialize(window.getApp().getServVar("validTelPrefixes"));
	
	var self = this;
	options.addElement = function(){
		//inn
		var er_ctrl = new ErrorControl(id+":inn:error");
		var info_ctrl = new EditInfo(id+":inn:info", {
			"templateOptions":{
				"TEXT_CLASS":"text-info",
				"PICT_CLASS":"icon-checkmark"
			}
		});
		var success_ctrl = new EditInfo(id+":inn:success", {
			"templateOptions":{
				"TEXT_CLASS":"text-success",
				"PICT_CLASS":"icon-checkmark"
			}
		});
		this.addElement(new EditINNSelfEmployed(id+":inn",{
			"required":true,
			"labelCaption":"ИНН:",		
			"title:": "ИНН самозанятого.",
			"errorControl": er_ctrl,
			"infoControlElements": [er_ctrl, info_ctrl, success_ctrl],
			"events":{
				"input": function(e){
					var ctrl = self.getElement("inn");
					var v = ctrl.getValue();
					if(v && v.length == 12 && !ctrl.validate()){
						return;
					}
					self.getElement("inn").setValid();
				}
			},
			"onClear":function(){
				self.getElement("inn").setValid();
			}			
		}));	

		//banks_ref
		var er_ctrl = new ErrorControl(id+":banks_ref:error");
		var info_ctrl = new EditInfo(id+":banks_ref:info", {
			"templateOptions":{
				"TEXT_CLASS":"text-info",
				"PICT_CLASS":"icon-checkmark"
			}
		});
		this.addElement(new BankEditRef(id+":banks_ref",{
			"acCount": 5,
			"required":true,
			"labelCaption":"Банк:",
			"cmdInsert":false,
			"cmdOpen":false,
			"cmdSelect":false,
			"placeholder": "БИК банка",
			"title:": "Банк самозанятого",
			"selectFormatFunction":function(f){
				return f.bik.getValue();
			},
			"errorControl": er_ctrl,
			"infoControlElements": [er_ctrl, info_ctrl],
			"onSelect":function(f){
				self.getElement("banks_ref").setValid();
				var bik = f.bik.getValue();
				self.getElement("bank_acc").setBik(bik);
				var descr = f["name"].getValue();
				var kor = f["korshet"].getValue();
				if(kor){
					descr+= " "+kor;
				}
				var gor = f["gor"].getValue();
				if(gor){
					descr+= " "+gor;
				}
				self.getElement("banks_ref").getInfoControls().getElement("info").setValue(descr);
				self.checkBankAcc();
			},
			"onClear":function(){
				self.getElement("bank_acc").setBik(undefined);
				self.getElement("banks_ref").getInfoControls().getElement("info").clear();
				self.checkBankAcc();
			}
		}));
		this.addElement(new BankAccEdit(id+":bank_acc",{
			"required":true,
			"labelCaption":"Расчетный счет:",
			"placeholder": "Расчетный счет",
			"title": "Расчетный счет самозанятого",
			"events":{
				"input": function(){
					self.checkBankAcc();					
				}
			},
			"onClear":function(){
				self.getElement("bank_acc").setValid();
			}
		}));
		
		//name
		this.addElement(new FIOEdit(id+":name",{
			"required":true,
			"labelCaption": "ФИО:",
			"placeholder": "ФИО",
			"title": "ФИО самозанятого",
			"maxLength":200,
			"events":{
				"input":function(){
					self.getElement("name_full").setValid();					
				},
				"blur":function(e){
					var fio = e.target.value;
					if(fio && fio.length > 0){
						self.getElement("name_full").validate();
					}					
				}
			},
			"onClear":function(){
				self.getElement("name_full").setValid();
			}			
		}));

		//birthdate
		this.addElement(new EditDate(id+":birthdate",{
			"required":true,
			"labelCaption": "Дата рождения:",
			"placeholder": "Дата рождения",
			"title": "Дата рождения",
			"events":{
				"input":function(){
					self.getElement("birthdate").setValid();
				}
			}
		}));
		
		//address_reg
		var er_ctrl = new ErrorControl(id+":address_reg:error");
		var info_ctrl = new EditInfo(id+":address_reg:info", {
			"templateOptions":{
				"TEXT_CLASS":"text-info",
				"PICT_CLASS":"icon-checkmark"
			}
		});
		this.addElement(new EditString(id+":address_reg",{
			"required":true,
			"labelCaption": "Адрес регистрации:",
			"placeholder": "Адрес регистрации",
			"title": "Адрес регистрации самозанятого",
			"maxLength":1500,
			"events":{
				"input": function(e){
					self.getElement("address_reg").setValid();
				}
			},
			"errorControl": er_ctrl,
			"infoControlElements": [er_ctrl, info_ctrl]
		}));

		//email
		var er_ctrl = new ErrorControl(id+":email:error");
		var info_ctrl = new EditInfo(id+":email:info", {
			"templateOptions":{
				"TEXT_CLASS":"text-info",
				"PICT_CLASS":"icon-checkmark"
			}
		});
		var success_ctrl = new EditInfo(id+":email:success", {
			"templateOptions":{
				"TEXT_CLASS":"text-success",
				"PICT_CLASS":"icon-checkmark"
			}
		});
		this.addElement(new EditEmailInlineValidation(id+":email",{
			"required":true,
			"disabled":true,
			"labelCaption": "Электронная почта:",
			"placeholder": "Электронная почта",
			"title": "Электронная почта. Бедет выполнена проверка адреса электронной почты.",
			"errorControl": er_ctrl,
			"infoControlElements": [er_ctrl, info_ctrl, success_ctrl]
		}));

		//tel
		var er_ctrl = new ErrorControl(id+":tel:error");
		var info_ctrl = new EditInfo(id+":tel:info", {
			"templateOptions":{
				"TEXT_CLASS":"text-info",
				"PICT_CLASS":"icon-checkmark"
			}
		});
		var success_ctrl = new EditInfo(id+":tel:success", {
			"templateOptions":{
				"TEXT_CLASS":"text-success",
				"PICT_CLASS":"icon-checkmark"
			}
		});
		this.addElement(new EditPhone(id+":tel",{
			"labelCaption": "Мобильный телефон:",
			"disabled":true,
			"required":true,
			"placeholder": "Мобильный телефон",
			"title": "Мобильный телефон",
			//"regExpression": "^" + valid_tel_pref.join("|") + "[0-9]+$",
			"errorControl": er_ctrl,
			"infoControlElements": [er_ctrl, info_ctrl, success_ctrl],
			"events":{
				"input": function(e){
					var v = e.target.value;
					if(v && v.length == 16){
						self.getElement("tel").validate();
					}else{
						self.getElement("tel").setValid();
					}
				}
			},
			"onClear":function(){
				self.getElement("tel").setValid();
			}			
		}));

		//studio
		this.addElement(new StudioEdit(id+":studios_ref",{
			"required":true,
			"labelCaption":"Салон:",
			"title:": "Салон самозанятого"
		}));
		
		//studio
		this.addElement(new UserEditRef(id+":users_ref",{
			"required":true,
			"labelCaption":"Аккаунт:",
			"title:": "Аккаунт самозанятого"
		}));

		//passport
		this.addElement(new PassportEdit(id+":passport",{
			"required":true,
			"labelCaption":"Паспорт:",
			"picture":false
		}));
		
		//equipments
		this.addElement(new StudioEquipmentGrid(id+":equipments",{
		}));

		this.addElement(new SpecialistDocumentList_View(id+":specialist_documents",{
			"detail":true
		}));
				
		//this.addElement(new SpecialistWorkList_View(id+":specialist_works",{
		this.addElement(new YClTransactionDocAllList_View(id+":specialist_transactions",{			
			"detail":true
		}));		
		
		this.addElement(new EditFile(id+":passport_preview",{
			"maxWidth":"100",
			"maxHeight":"100",
			"multipleFiles":false
			,"showHref":false
			,"showPic":true
			,"onDeleteFile":function(fileId,callBack){
				self.m_attachManager.deleteAttachment(fileId,callBack);
			}
			,"onFileAdded":function(fileId){
				//self.addAttachment(fileId);
				self.m_attachManager.addAttachment(fileId);
			}
			,"onDownload":function(fileId){
				self.m_attachManager.downloadAttachment(fileId);
			}
			,"allowedFileExtList":"pdf"
		}));	
		
		this.addElement(new EditFloat(id+":agent_percent",{
			"required": true,
			"precision":"2",
			"labelCaption": "Агентское вознаграждение:",
			"title":"Процент агентского вознаграждения"
		}));	

		this.addElement(new SpecialityEdit(id+":specialities_ref",{
			"required": true
		}));	
		
		//grid
		this.addElement(new SpecialistSalaryDebetList_View(id+":specialist_salary_debets",{
			"detail":true
		}));
		this.addElement(new SpecialistSalaryKreditList_View(id+":specialist_salary_kredits",{
			"detail":true
		}));
		
		this.addElement(new SpecialistPeriodSalaryDetailList_View(id+":specialist_period_salary",{
			"detail":true
		}));		
		
	}
	
	SpecialistDialog_View.superclass.constructor.call(this,id,options);
	
	this.m_attachManager = new AttachmentManager({
		"view": this,
		"dataType": "specialists",
		"attachmentViewName": "passport_preview"
	});
	
	//****************************************************
	//read
	this.setDataBindings([
		new DataBinding({"control":this.getElement("name")})
		,new DataBinding({"control":this.getElement("inn")})
		,new DataBinding({"control":this.getElement("studios_ref"), "fieldId": "studios_ref"})
		,new DataBinding({"control":this.getElement("banks_ref"), "fieldId": "banks_ref"})
		,new DataBinding({"control":this.getElement("bank_acc")})		
		,new DataBinding({"control":this.getElement("birthdate")})
		,new DataBinding({"control":this.getElement("address_reg")})
		,new DataBinding({"control":this.getElement("passport")})
		,new DataBinding({"control":this.getElement("equipments")})
		,new DataBinding({"control":this.getElement("users_ref"), "fieldId": "users_ref"})
		,new DataBinding({"control":this.getElement("passport_preview")})
		,new DataBinding({"control":this.getElement("email")})
		,new DataBinding({"control":this.getElement("tel")})
		,new DataBinding({"control":this.getElement("agent_percent")})
		,new DataBinding({"control":this.getElement("specialities_ref"), "fieldId": "specialities_ref"})
	]);
	
	//write
	this.setWriteBindings([
		new CommandBinding({"control":this.getElement("name")})
		,new CommandBinding({"control":this.getElement("inn")})
		,new CommandBinding({"control":this.getElement("banks_ref"), "fieldId":"bank_bik"})
		,new CommandBinding({"control":this.getElement("bank_acc")})
		,new CommandBinding({"control":this.getElement("studios_ref"), "fieldId":"studio_id"})
		,new CommandBinding({"control":this.getElement("birthdate")})
		,new CommandBinding({"control":this.getElement("address_reg")})
		,new CommandBinding({"control":this.getElement("passport")})
		,new CommandBinding({"control":this.getElement("users_ref"), "fieldId":"user_id"})
		,new CommandBinding({"control":this.getElement("agent_percent")})
		,new CommandBinding({"control":this.getElement("specialities_ref"), "fieldId":"speciality_id"})
	]);
	
	this.addDetailDataSet({
		"control":this.getElement("specialist_documents").getElement("grid"),
		"controlFieldId": ["specialist_id"],
		"value": [function(){
			return self.m_model.getFieldValue("id");
		}]
	});		

	this.addDetailDataSet({
		"control":this.getElement("specialist_transactions").getElement("grid"),
		"controlFieldId": ["specialist_id"],
		"value": [function(){
			return self.m_model.getFieldValue("id");
		}]
	});		

	this.addDetailDataSet({
		"control":this.getElement("specialist_salary_debets").getElement("grid"),
		"controlFieldId": ["specialist_id"],
		"value": [function(){
			return self.m_model.getFieldValue("id");
		}]
	});		
	this.addDetailDataSet({
		"control":this.getElement("specialist_salary_kredits").getElement("grid"),
		"controlFieldId": ["specialist_id"],
		"value": [function(){
			return self.m_model.getFieldValue("id");
		}]
	});		
	this.addDetailDataSet({
		"control":this.getElement("specialist_period_salary").getElement("grid"),
		"controlFieldId": ["specialist_id"],
		"value": [function(){
			return self.m_model.getFieldValue("id");
		}]
	});		
}
extend(SpecialistDialog_View, ViewObjectAjx);


SpecialistDialog_View.prototype.checkBankAcc = function(){	
	var v = document.getElementById(this.getId() + ":bank_acc").value;
	var bank_ctrl = this.getElement("bank_acc");
	if (v.length == 20 && bank_ctrl.m_bik && bank_ctrl.m_bik.length){
		if(!bank_ctrl.validate()){
			return;
		}
	}else{
		bank_ctrl.setValid();
	}
}

