/* Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function SpecialistRegDialog_View(id,options){	

	options = options || {};
	options.HEAD_TITLE = "Карточка физического лица";
	options.cmdSave = true;
	options.controller = new SpecialistReg_Controller();
	options.model = (options.models&&options.models.SpecialistRegDialog_Model)? options.models.SpecialistRegDialog_Model : new SpecialistRegDialog_Model();

	var self = this;
	options.addElement = function(){
		this.addElement(new EditINNSelfEmployed(id+":inn",{
			"required":true,
			"maxLength":12,
			"labelCaption":"ИНН самозанятого:",		
			"title:": "ИНН самозанятого. Бедет выполнена проверка по данным ФНС на текущую дату."
		}));	
		
		this.addElement(new BankEditRef(id+":banks_ref",{
			"acCount": 5,
			"required":true,
			"labelCaption":"Банк:",
			"cmdInsert":false,
			"cmdOpen":false,
			"cmdSelect":false,
			"placeholder": "БИК банка",
			"title:": "Банк самозанятого"
		}));
		this.addElement(new BankAccEdit(id+":bank_acc",{
			"required":true,
			"labelCaption": "Расчетный счет:",
			"placeholder": "Расчетный счет",
			"title": "Расчетный счет самозанятого"
		}));
		this.addElement(new FIOEdit(id+":name_full",{
			"required":true,
			"labelCaption": "ФИО самозанятого:",
			"placeholder": "ФИО",
			"title": "ФИО самозанятого",
			"maxLength":200
		}));
		this.addElement(new EditString(id+":name",{
			"required":true,
			"labelCaption": "Логин для авторизации:",
			"placeholder": "Логин для авторизации",
			"title": "Логин для авторизации",
			"maxLength":100,
			"minLength":6
		}));
		this.addElement(new EditPhone(id+":tel",{
			"labelCaption": "Мобильный телефон:",
			"required":true,
			"placeholder": "Мобильный телефон",
			"title": "Мобильный телефон. Бедет выполнена проверка номера."
		}));
		this.addElement(new EditEmailInlineValidation(id+":email",{
			"required":true,
			"labelCaption": "Электронная почта:",
			"placeholder": "Электронная почта",
			"title": "Электронная почта. Бедет выполнена проверка адреса электронной почты."
		}));
		this.addElement(new PassportEdit(id+":passport",{
			"editContClassName": cont_cl,
			"required":true
		}));
		this.addElement(new EditString(id+":address_reg",{
			"required":true,
			"labelCaption": "Адрес регистрации:",
			"placeholder": "Адрес регистрации",
			"title": "Адрес регистрации самозанятого",
			"maxLength":1500
		}));
		this.addElement(new EditDate(id+":birthdate",{
			"required":true,
			"labelCaption": "Дата рождения:",
			"placeholder": "Дата рождения",
			"title": "Дата рождения"
		}));
		
	}	
	
	UserDialog_View.superclass.constructor.call(this,id,options);
	
	//****************************************************	
	
	//read
	var r_bd = [
		new DataBinding({"control":this.getElement("name")})
		,new DataBinding({"control":this.getElement("name_full")})
		,new DataBinding({"control":this.getElement("role"), "field":this.m_model.getField("role_id")})
	];
	this.setDataBindings(r_bd);
	
	//write
	var wr_b = [
		new CommandBinding({"control":this.getElement("name")})
		,new CommandBinding({"control":this.getElement("name_full")})
		,new CommandBinding({"control":this.getElement("role"),"fieldId":"role_id"})
	];
	this.setWriteBindings(wr_b);	
	
}
extend(SpecialistRegDialog_View,ViewObjectAjx);


