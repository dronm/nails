/** Copyright (c) 2023
	Andrey Mikhalevich, Katren ltd.
 */
function SignEdit(id,options){
	options = options || {};
	
	options.className = "form-group";
	options.addElement = function(){
		var self = this;
		this.addElement(new Control(id+":text","SPAN",{
			"value": "Образец подписи",
		}));
		
		this.addElement(new ButtonCtrl(id+":openSign",{
			"glyph": "glyphicon-edit",
			"title": "Прикрепить подпись",
			"onClick":function(){
				self.openSign();
			}
		}));
	}
	
	SignEdit.superclass.constructor.call(this,id,options);	
}
	
extend(SignEdit, View);//EditModalDialog

SignEdit.prototype.openSign = function(){
	//calc dialog width
	var self = this;
	(new WindowFormModalBS(this.getId()+":form",{
		"cmdCancel":this.m_cmdCancel,
		"controlCancelCaption":"Отмена",
		"controlCancelTitle":"Отменить создание подписи, закрыть форму",
		"cmdOk":true,
		"controlOkCaption":"Ок",
		"controlOkTitle":"Записать подпись",
		"onClickCancel":function(){
			self.closeSelect(false);
		},		
		"onClickOk":function(){
			self.closeSelect(true);
		},				
		"content":new Sign_View(this.getId()+":form:body:Sign_View", {}),
		"contentHead":"Образец подписи",
		"dialogWidth":"80%"
	})).open();
}

