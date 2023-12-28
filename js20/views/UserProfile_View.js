/* Copyright (c) 2016, 2021
 *	Andrey Mikhalevich, Katren ltd.
 */
function UserProfile_View(id,options){	

	options = options || {};
	
	options.cmdOkAsync = false;
	options.cmdOk = false;
	options.cmdCancel = false;
	
	options.templateOptions = options.templateOptions || {};
	
	var self = this;
	options.addElement = function(){
		this.addElement(new EditEmail(id+":name",{
			"labelCaption":"Логин:",
			"events":{
				"keyup":function(){
					self.getControlSave().setEnabled(true);
				}
			}
			
		}));	

		this.addElement(new EditPassword(id+":pwd",{
			"labelCaption":"Пароль:",
			"events":{
				"keyup":function(){
					self.checkPass();	
				}
			}
		}));	
		this.addElement(new EditPassword(id+":pwd_confirm",{
			"labelCaption":"Подтверждение пароля:",
			"events":{
				"keyup":function(){
					self.checkPass();	
				}
			}
		}));	

		this.addElement(new EditString(id+":name_full",{
			"maxLength":"500",
			"labelCaption":"ФИО:",
			"events":{
				"keyup":function(){
					self.getControlSave().setEnabled(true);	
				}
			}
		}));
		
		/*this.addElement(new Enum_locales(id+":locale_id",{
			"labelCaption":"Локаль:",
			"required":true,
			"events":{
				"keyup":function(){
					self.getControlSave().setEnabled(true);	
				}
			}			
		}));*/

		this.addElement(new TimeZoneLocaleEdit(id+":time_zone_locales_ref",{
			"events":{
				"change":function(){
					self.getControlSave().setEnabled(true);	
				}
			}			
		}));
		this.addElement(new EditFile(id+":photo",{
			"maxWidth":"250",
			"maxHeight":"250",
			"multipleFiles":false
			,"showHref":false
			,"showPic":true
			,"fileId":"photo"
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
		
	}
	
	UserProfile_View.superclass.constructor.call(this,id,options);

	this.m_attachManager = new AttachmentManager({
		"view": this,
		"dataType": "users",
		"attachmentViewName": "photo"
	});

	//****************************************************
	var contr = new User_Controller();
	
	//read
	this.setReadPublicMethod(contr.getPublicMethod("get_profile"));
	this.m_model = options.models.UserProfile_Model;
	
	//read
	this.setDataBindings([
		new DataBinding({"control":this.getElement("id")})
		,new DataBinding({"control":this.getElement("name")})
		//,new DataBinding({"control":this.getElement("locale_id")})
		,new DataBinding({"control":this.getElement("time_zone_locales_ref")})
		,new DataBinding({"control":this.getElement("name_full")})
		,new DataBinding({"control":this.getElement("photo")})
	]);

	//write
	this.setController(contr);
	this.getCommand(this.CMD_OK).setBindings([
		new CommandBinding({"control":this.getElement("id")})
		,new CommandBinding({"control":this.getElement("name")})
		//,new CommandBinding({"control":this.getElement("locale_id")})
		,new CommandBinding({"control":this.getElement("time_zone_locales_ref"),"fieldId":"time_zone_locale_id"})
		,new CommandBinding({"control":this.getElement("name_full")})
		,new CommandBinding({"control":this.getElement("pwd")})
	]);
	
	this.getControlSave().setEnabled(false);
}
extend(UserProfile_View,ViewObjectAjx);


UserProfile_View.prototype.TXT_PWD_ER = "Пароли не совпадают";


UserProfile_View.prototype.checkPass = function(){
	var pwd = this.getElement("pwd").getValue();
	if (pwd && pwd.length){
		var pwd_conf = this.getElement("pwd_confirm").getValue();
		if (pwd_conf && pwd_conf.length && pwd!=pwd_conf){
			this.getElement("pwd_confirm").setNotValid(this.TXT_PWD_ER);
			this.getControlSave().setEnabled(false);
		}
		else if (pwd_conf && pwd_conf.length){
			this.getElement("pwd_confirm").setValid();
			if (!this.getControlSave().getEnabled()){
				this.getControlSave().setEnabled(true);
			}
		}
		else if ((!pwd_conf || !pwd_conf.length) && this.getControlSave().getEnabled()){
			this.getControlSave().setEnabled(false);
		}
	}
}

/*
UserProfile_View.prototype.onSaveOk = function(resp){
	UserProfile_View.superclass.onSaveOk.call(this,resp);
	
	window.showNote(this.NOTE_DATA_SAVED,null,3000);
}
*/
