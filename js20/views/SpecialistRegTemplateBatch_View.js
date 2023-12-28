/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023

 * @extends View
 * @requires core/extend.js
 * @requires controls/View.js     

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {object} options
 */
function SpecialistRegTemplateBatch_View(id,options){
	options = options || {};	
	
	var self = this;
	options.addElement = function(){
		this.addElement(new SpecialistEdit(id + ":specialist",{
		}));
		
		this.addElement(new ButtonCmd(id + ":cmdSend", {
			"caption":"Отправить регистрационный пакет",
			"title":"Отправить мастеру регистрационный пакет документов на подписание",
			"onClick":function(){
				self.cmdSend();
			}
		}));
		
	}
	
	SpecialistRegTemplateBatch_View.superclass.constructor.call(this,id,options);
}
//ViewObjectAjx,ViewAjxList
extend(SpecialistRegTemplateBatch_View, View);

/* Constants */


/* private members */

/* protected*/


/* public methods */
SpecialistRegTemplateBatch_View.prototype.cmdSend = function(){
	var ref = this.getElement("specialist").getValue();
	if(!ref || ref.isNull()){
		return;
	}
	var operation_id = CommonHelper.md5(DateHelper.time().toString());
	this.startOperationMonitor(operation_id);
	var pm = (new Specialist_Controller()).getPublicMethod("send_reg_documents");
	pm.setFieldValue("operation_id", operation_id);
	pm.setFieldValue("id", ref.getKey("id"));
	pm.run({
		"ok":function(){
			window.showTempNote("Начало формирования документов",null,5000);
		}
	});
}

SpecialistRegTemplateBatch_View.prototype.startOperationMonitor = function(operation_id){	
	var srv = window.getApp().getAppSrv();
	if(srv && srv.connActive()){
		var self = this;		
		this.unsubscribeFromSrvEvents();
		this.subscribeToSrvEvents({
			"events":[
				{"id":"UserOperation." + operation_id}
				,{"id":"UserOperation." + operation_id}
			]
			,"onEvent":function(json){
				if(json.controllerId == "UserOperation" && json.methodId == operation_id){
					if(json.params.status=="end"){
						if(!json.params.res){
							window.showTempError("Ошибки формирования документов", null, 5000);
						}else{
							window.showTempOk("Документы отправлены",null,5000);
						}
					}
				}
			}
		});		
	}
}

