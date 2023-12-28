/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2022

 * @extends ButtonCmd
 * @requires core/extend.js
 * @requires controls/ButtonCmd.js     

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {object} options
 */
function SpecialistRegBtnRegister(id,options){
	options = options || {};	
	
	options.caption = " " + this.BTN_CAP;
	options.glyph = "glyphicon-user";
	options.title = this.BTN_TITLE;
	
	var self = this;
	options.onClick = function(){
		self.register();
	}
	
	this.m_cmd = options.cmd;
	this.m_getId = options.getId;
	
	SpecialistRegBtnRegister.superclass.constructor.call(this,id,options);
}
//ViewObjectAjx,ViewAjxList
extend(SpecialistRegBtnRegister,ButtonCmd);

/* Constants */
SpecialistRegBtnRegister.prototype.BTN_TITLE = "Сформировать пакет документов, отправить самозанятому";
SpecialistRegBtnRegister.prototype.BTN_CAP = "Начать сотрудничество";

/* private members */

/* protected*/


/* public methods */
SpecialistRegBtnRegister.prototype.registerCont = function(formCont, fl){
	var v = formCont.getContent();
	if(!v.validate()){
		return;
	}
	this.m_operationId = CommonHelper.md5(DateHelper.time().toString());
	var pm = (new SpecialistReg_Controller()).getPublicMethod("register");
	pm.setFieldValue("id", fl.id);
	pm.setFieldValue("studio_id", v.getElement("studio").getValue().getKey());
	pm.setFieldValue("ycl_staff_id", v.getElement("ycl_staff").getValue().getKey());
	pm.setFieldValue("operation_id", this.m_operationId);
	this.startOperationMonitor(fl);
	
	var self = this;
	pm.run({
		"ok":function(resp){
			formCont.close();
			window.showTempNote("Выполняется операция регистрации " + fl.name_full, null, 5000);
		}
	});	
}

SpecialistRegBtnRegister.prototype.register = function(){
	var id, name_full;
	if(this.m_cmd){
		this.m_grid.setModelToCurrentRow();
		id = this.m_grid.getModel().getFieldValue("id");
		name_full = this.m_grid.getModel().getFieldValue("name_full");
	}
	else{
		id = this.m_getId();
	}

	var self = this;
	(new WindowFormModalBS("specialistRegRegister",{
		"dialogWidth":"30%",
		"cmdOk":true,
		"cmdCancel":true,
		"cmdCancelCaption":"Отмена",
		"cmdCancelTitle":"Отменить начало сотрудничества, закрыть форму",
		"controlOkCaption": this.BTN_CAP,
		"controlOkTitle": this.BTN_TITLE,
		"contentHead": name_full,
		"onClickOk":function(){
			self.registerCont(this, {"id":id, "name_full":name_full});
		},
		"onClickCancel":function(){
			this.close();
		},
		"cmdClose":true,
		"content":new SpecialistRegRegister_View("specialistRegRegister:view")
	})).open();

}

SpecialistRegBtnRegister.prototype.startOperationMonitor = function(fl){	
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
					if(json.params.status=="end"){
						if(!json.params.res){
							var er = "Ошибки при формирования пакета документов";
							window.showTempError(er, null, 5000);
						}else{
							window.showTempOk("Замозанятый " + fl.name_full+ " зарегистрирован. Пакет документов отправлен.", null, 5000);
							if(self.m_grid){
								self.m_grid.onRefresh();
							}
						}
					}
				}
			}
		});		
	}
}

