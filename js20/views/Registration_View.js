/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023
 
 * @class
 * @classdesc
	
 * @param {string} id view identifier
 * @param {namespace} options
 * @param {namespace} options.models All data models
 * @param {namespace} options.variantStorage {name,model}
 * 
 * 720203973455
 * 42307810967101966244
 * 047102651
 */	
function Registration_View(id,options){	

	options = options || {};
	
	var app = window.getApp();
	var reg_captcha = CommonHelper.unserialize(app.getServVar("regCaptcha"));
	var valid_tel_pref = CommonHelper.unserialize(app.getServVar("validTelPrefixes"));

	var self = this;
	options.addElement = function(){
		var cont_cl = "input-group " + window.getBsCol(12);
		this.addElement(new Captcha(id+":captcha",{
			"errorControl": new ErrorControl(id+":captcha:error"),
			"required":true,
			"captchaId": reg_captcha.id,
			"captchaWidth": reg_captcha.width,
			"captchaHeight": reg_captcha.height,
			"captchaCount": reg_captcha.count,
			"templateOptions":{
				"IMG_CONT_CLASS": ""
			},
			"keyEvents":{
				"input": function(){
					self.getElement("captcha").setValid();
				}
			}
		}));	
		
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
			"editContClassName": cont_cl,
			"cmdSmClear":true,
			"required":true,
			"labelCaption":"",		
			"title:": "ИНН самозанятого. Бедет выполнена проверка по данным ФНС на текущую дату.",
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
			"editContClassName": cont_cl,
			"cmdSmClear":true,
			"acCount": 5,
			"required":true,
			"labelCaption":"",
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
				//self.getElement("banks_ref").getErrorControl().setValue(descr);
				self.getElement("banks_ref").getInfoControls().getElement("info").setValue(descr);
				self.checkBankAcc();
			},
			"onClear":function(){
				self.getElement("bank_acc").setBik(undefined);
				//self.getElement("banks_ref").getErrorControl().setValue("");
				self.getElement("banks_ref").getInfoControls().getElement("info").clear();
				self.checkBankAcc();
			}
		}));
		this.addElement(new BankAccEdit(id+":bank_acc",{
			"editContClassName": cont_cl,
			"cmdSmClear":true,
			"required":true,
			"labelCaption": "",
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
		this.addElement(new FIOEdit(id+":name_full",{
			"editContClassName": cont_cl,
			"cmdSmClear":true,
			"required":true,
			"labelCaption": "",
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
		/*
		this.addElement(new EditString(id+":name",{
			"editContClassName": cont_cl,
			"required":true,
			"labelCaption": "",
			"placeholder": "Логин для авторизации",
			"title": "Логин для авторизации",
			"maxLength":100,
			"minLength":6,
			"events":{
				"input":function(){
					self.getElement("name").setValid();					
				}
			},
			"onClear":function(){
				self.getElement("name").setValid();
			}									
		}));
		*/
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
			"editContClassName": cont_cl,
			"cmdSmClear":true,
			"labelCaption": "",
			"required":true,
			"placeholder": "Мобильный телефон",
			"title": "Мобильный телефон. Бедет выполнена проверка номера.",
			"regExpression": "^" + valid_tel_pref.join("|") + "[0-9]+$",
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
			"editContClassName": cont_cl,
			"cmdSmClear":true,
			"required":true,
			"labelCaption": "",
			"placeholder": "Электронная почта",
			"title": "Электронная почта. Бедет выполнена проверка адреса электронной почты.",
			"errorControl": er_ctrl,
			"infoControlElements": [er_ctrl, info_ctrl, success_ctrl]
		}));
		
		this.addElement(new PassportEdit(id+":passport",{
			"editContClassName": cont_cl,
			"required":true
		}));
		
		//address
		var er_ctrl = new ErrorControl(id+":address_reg:error");
		var info_ctrl = new EditInfo(id+":address_reg:info", {
			"templateOptions":{
				"TEXT_CLASS":"text-info",
				"PICT_CLASS":"icon-checkmark"
			}
		});
		this.addElement(new EditString(id+":address_reg",{
			"editContClassName": cont_cl,
			"cmdSmClear":true,
			"required":true,
			"labelCaption": "",
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
		
		//birthdate
		this.addElement(new EditDateInlineValidation(id+":birthdate",{
			"editContClassName": cont_cl,
			"cmdSmClear":true,
			"required":true,
			"labelCaption": "",
			"placeholder": "Дата рождения",
			"title": "Дата рождения"
			/*
			"events":{
				"input":function(){
					self.getElement("birthdate").setValid();
				}
			}*/
		}));
		
		this.addElement(new ButtonCmd(id+":submit",{
			"onClick":function(){
				self.onSubmit();				
			}
		}));
		
		this.addElement(new ButtonCmd(id+":test",{
			"caption":"TEST",
			"onClick":function(){
				self.getElement("inn").setValue("720203973455");
				self.getElement("banks_ref").setValue(new RefType({"keys":{"bik":"044012000"}, "descr":"044012000"}));
				self.getElement("bank_acc").setValue("42307810967101966244");
				self.getElement("name_full").setValue("Михалевич Андрей Александрович");
				self.getElement("tel").setValue("9222695251");
				self.getElement("email").setValue("katren_shd@rambler.ru");
				self.getElement("address_reg").setValue("Тюмень");
				self.getElement("birthdate").setValue("01/01/1970");
				//self.getElement("name").setValue("dron_m");
				self.getElement("passport").setValue({
					"series":"1111",
					"num":"777777",
					"issue_body":"Отделение УВД по Центральному району г. Тюмени",
					"issue_date":"2023-01-11",
					"dep_code":"720001"
				});
			}
		}));		
		
	}
	
	Registration_View.superclass.constructor.call(this, id, options);
	
	this.setCommands({
		"submit": new Command("submit", {
			"control":this.getElement("submit"),
			"publicMethod": (new Specialist_Controller()).getPublicMethod("register"),
			"bindings":[
				new CommandBinding({"control":this.getElement("inn")})			
				,new CommandBinding({"control":this.getElement("bank_acc")})
				,new CommandBinding({"control":this.getElement("name_full")})
				//,new CommandBinding({"control":this.getElement("name")})
				,new CommandBinding({"control":this.getElement("email")})
				,new CommandBinding({"control":this.getElement("tel")})
				,new CommandBinding({"control":this.getElement("birthdate")})
				,new CommandBinding({"control":this.getElement("address_reg")})
				,new CommandBinding({"control":this.getElement("captcha")})
				,new CommandBinding({"control":this.getElement("passport")})
				,new CommandBinding({"func": function(pm){					
					pm.setFieldValue("operation_id", self.m_operationId);
					pm.setFieldValue("banks_ref", CommonHelper.serialize(this.getElement("banks_ref").getValue()));
					pm.setFieldValue("passport_file_content", this.getElement("passport").m_docFileNode.files);					
				}})
			]		
		})
	});
	
	this.getStoredOperation();
}
extend(Registration_View, ViewAjx);//

Registration_View.prototype.CAPTCHA_LEN = 6;

Registration_View.prototype.m_orgModified;
Registration_View.prototype.m_saveOrgInt;
Registration_View.prototype.m_orgAttrModified = {};
Registration_View.prototype.m_userAttrModified = {};
Registration_View.prototype.m_orgFound = false;
Registration_View.prototype.m_orgExists = false;

Registration_View.prototype.SAVE_INT_DELAY = 5000;

Registration_View.prototype.storeOperationId = function(){	
	if(window["localStorage"]){
		localStorage.setItem('operationId', this.m_operationId);
	}
}

Registration_View.prototype.unsetOperationId = function(){	
	if(window["localStorage"]){
		localStorage.removeItem('operationId');
	}
}
Registration_View.prototype.getOperationId = function(){	
	var operationId;
	if(window["localStorage"]){
		operationId = localStorage.getItem('operationId');
	}
	return operationId;
}

Registration_View.prototype.getStoredOperation = function(){	
	this.m_operationId = this.getOperationId();
	if(!this.m_operationId||this.m_operationId=="null"){
		return;
	}
	this.getElement("captcha").setVisible(false);	
	//fetch from server
	var pm = (new SpecialistReg_Controller()).getPublicMethod("get_by_operation_id");
	pm.setFieldValue("operation_id", this.m_operationId);
	var self = this;
	pm.run({
		"ok":function(resp){
			var m = resp.getModel("SpecialistRegDialog_Model");
			if(m && m.getNextRow()){
				self.getElement("banks_ref").setValue(m.getFieldValue("banks_ref"));
				self.getElement("bank_acc").setValue(m.getFieldValue("bank_acc"));
				self.getElement("inn").setValue(m.getFieldValue("inn"));
				self.getElement("name_full").setValue(m.getFieldValue("name_full"));
				//self.getElement("name").setValue(m.getFieldValue("name"));
				self.getElement("email").setValue(m.getFieldValue("email"));
				self.getElement("tel").setValue(m.getFieldValue("tel"));
				self.getElement("birthdate").setValue(m.getFieldValue("birthdate"));
				self.getElement("address_reg").setValue(m.getFieldValue("address_reg"));
				self.getElement("passport").setValue(m.getFieldValue("passport"));
				
				if(m.getFieldValue("inn_checked")){
					self.setInnFnsChecked(m.getFieldValue("inn_fns_ok"), null);
				}
				if(m.getFieldValue("tel_confirmed")){
					self.setTelConfirmed();
				}else{
					self.setTelChecked(m.getFieldValue("tel_sent"), null);
				}
				if(m.getFieldValue("email_confirmed")){
					self.setEmailConfirmed();
				}else{
					self.setEmailChecked(m.getFieldValue("email_sent"), null);
				}
				self.setPassportUploaded(m.getFieldValue("passport_uploaded"), null);
				
				var op_res = m.getFieldValue("operation_result");
				if(op_res === false){
					self.setSubmitError("Ошибка проверки данных");
					self.getElement("captcha").refresh();
				}else{
					self.setCommentChecking();
				}				
			}
		}
		,"fail":function(resp,erCode,erStr){
			if(erCode=="1001"){
				//deleted already
				self.unsubscribeFromSrvEvents();
				self.unsetOperationId();
				//go to login page
				self.m_operationId = undefined;				
				self.getElement("captcha").setVisible(true);
			}else{
				//throw new Error(erStr);
				self.setSubmitError(erStr);
			}
		}
	});

}

Registration_View.prototype.setInnFnsChecked = function(res, erText){	
	if(!res){
		this.getElement("inn").setNotValid(erText || "Не найден по данным ФНС");
	}else{
		this.getElement("inn").getInfoControls().getElement("success").setText("Проверен по данным ФНС");
	}
}

Registration_View.prototype.setTelConfirmed = function(){	
	this.getElement("tel").getInfoControls().getElement("success").setText("Подтвержден");
}

Registration_View.prototype.setTelChecked = function(res, erText){	
	if(!res){
		this.getElement("tel").setNotValid(erText || "Ошибка отправки сообщения");
	}else{
		this.getElement("tel").getInfoControls().getElement("success").setText("Отправлено сообщение");
	}
}

Registration_View.prototype.setEmailChecked = function(res, erText){	
	if(!res){
		this.getElement("email").setNotValid(erText || "Ошибка отправки сообщения");
	}else{
		this.getElement("email").getInfoControls().getElement("success").setText("Отправлено сообщение");
	}
}
Registration_View.prototype.setEmailConfirmed = function(){	
	this.getElement("email").getInfoControls().getElement("success").setText("Подтвержден");
}

Registration_View.prototype.setPassportUploaded = function(res, erText){	
	if(!res){
		this.getElement("passport").setNotValid(erText);
	}
}

Registration_View.prototype.toDOM = function(p){	
	Registration_View.superclass.toDOM.call(this, p)

	this.getElement("submit").setVisible(true);
	this.getElement("inn").focus();
	
	//event server
	var ws_host = window.location.host;
	var ws_sep = ws_host.indexOf(":");
	var app = window.getApp();
	app.m_appSrv = new AppSrv({
		"host":(ws_sep>=0)? ws_host.substring(0,ws_sep) : ws_host,
		"port": app.getServVar("wsPort"),
		"appId": app.getServVar("app_id"),
		"token": app.getServVar("token"),
		"tokenExpires": app.getServVar("tokenExpires")
	});
	app.m_appSrv.connect();
	if(this.m_operationId){
		var self = this;
		this.m_operStartInterval = setInterval(function(){
			if(app.m_appSrv.connActive()){
				clearInterval(self.m_operStartInterval);
				self.startOperationMonitor();
			}
		}, 1000);
	}	
}

Registration_View.prototype.onSubmit = function(){	

var c = this.getElement("passport");
	this.m_operationId = CommonHelper.md5(DateHelper.time().toString());
	var info_n = document.getElementById("Registration:info");
	DOMHelper.hide(info_n);
	DOMHelper.delAllChildren(info_n);
	DOMHelper.delClass(info_n, "alert-danger");
	DOMHelper.delClass(info_n, "alert-info");
	window.setGlobalWait(true);
	this.startOperationMonitor();
	
	var self = this;
	this.execCommand("submit",
		function(resp){
			self.setCommentChecking();			
			self.storeOperationId();						
		}
		,function(resp,errCode,errStr){
			self.setSubmitError(errStr);
			if(resp && resp.modelExists("Captcha_Model")){
				//new captcha
				self.getElement("captcha").setFromResp(resp);				
			}			
		}
		,function(){
			window.setGlobalWait(false);
		}
	);
}


Registration_View.prototype.setSubmitError = function(errStr){
	var info_n = document.getElementById("Registration:info");
	DOMHelper.addClass(info_n, "alert-danger");
	DOMHelper.setText(info_n, errStr);
	DOMHelper.show(info_n);	
	this.setEnabled(true);
	this.getElement("captcha").setVisible(true);
	this.getElement("submit").setVisible(true);		
}

Registration_View.prototype.setCommentChecking = function(){
	this.setTempEnabled("submit");
	this.setEnabled(false);
	var info_n = document.getElementById("Registration:info");
	DOMHelper.addClass(info_n, "alert-info");			
	DOMHelper.setText(info_n, "Выполняется проверка");
	DOMHelper.show(info_n);	
	this.getElement("captcha").setVisible(false);
	this.getElement("submit").setVisible(false);		
}

Registration_View.prototype.startOperationMonitor = function(){	
	var srv = window.getApp().getAppSrv();
	if(srv && srv.connActive()){
		var self = this;		
		this.unsubscribeFromSrvEvents();
		this.subscribeToSrvEvents({
			"events":[
				{"id":"UserOperation." + this.m_operationId}
				,{"id":"UserOperation." + this.m_operationId}
			]
			,"onEvent":function(json){
				if(json.controllerId == "UserOperation" && json.methodId == self.m_operationId){
					if(json.params.status=="progress"){
						if(json.params.f=="tel"){
							self.setTelChecked(json.params.res, json.params.c);

						}else if(json.params.f=="tel_conf"){
							self.setTelConfirmed();
							
						}else if(json.params.f=="email"){
							self.setEmailChecked(json.params.res, json.params.c);
							
						}else if(json.params.f=="email_conf"){
							self.setEmailConfirmed();
							
							
						}else if(json.params.f=="inn"){
							self.setInnFnsChecked(json.params.res, json.params.c);
							
						}else if(json.params.f=="passport_file_content"){
							self.setPassportUploaded(json.params.res, json.params.c);
						}
					
					}else if(json.params.status=="end"){
						if(!json.params.res){
							var er = "Ошибки при проверке самозанятого";
							window.showTempError(er,null,5000);
							self.setSubmitError(er);
							self.getElement("captcha").refresh();
						}else{
							var info_n = document.getElementById("Registration:info");
							DOMHelper.delAllChildren(info_n);
							DOMHelper.delClass(info_n, "alert-danger");
							DOMHelper.delClass(info_n, "alert-info");							
							DOMHelper.addClass(info_n, "alert-success");			
							DOMHelper.setText(info_n, "Проверка самозанятого выполнена. Сообщение отпрвлено на почту и в месенджер.");
							DOMHelper.show(info_n);	
							
							window.showTempNote("Проверка самозанятого выполнена",null,5000);
						}
					}
				}
			}
		});		
	}
}

Registration_View.prototype.checkBankAcc = function(){	
	var v = document.getElementById("Registration:bank_acc").value;
	var bank_ctrl = this.getElement("bank_acc");
	if (v.length == 20 && bank_ctrl.m_bik && bank_ctrl.m_bik.length){
		if(!bank_ctrl.validate()){
			return;
		}
	}else{
		bank_ctrl.setValid();
	}
}

