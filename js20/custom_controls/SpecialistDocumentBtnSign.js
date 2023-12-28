/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023

 * @extends ButtonCmd
 * @requires core/extend.js
 * @requires controls/ButtonCmd.js     

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {object} options
 */
function SpecialistDocumentBtnSign(id,options){
	options = options || {};	
	
	options.caption = " Подписать";
	options.glyph = "glyphicon-pencil";
	options.title = "Подписать документ";
	
	var self = this;
	options.onClick = function(){
		self.signDocument();
	}
	
	this.m_cmd = options.cmd;
	this.m_getId = options.getId;
	
	SpecialistDocumentBtnSign.superclass.constructor.call(this,id,options);
}
//ViewObjectAjx,ViewAjxList
extend(SpecialistDocumentBtnSign,ButtonCmd);

/* Constants */

/* private members */

/* protected*/


/* public methods */
SpecialistDocumentBtnSign.prototype.startOperationMonitor = function(){	
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
						self.unsubscribeFromSrvEvents();
						if(!json.params.res){
							var er = "Ошибки при подписании документа";
							window.showTempError(er,null,5000);
						}else{
							window.showTempOk("Документ подписан", null, 5000);
							//уменьшим в меню
							var n = document.getElementById("unsigned_docs_cnt");
							if(n) {
								var v = parseInt(DOMHelper.getText(n), 10);
								v--;
								if(!v){
									DOMHelper.hide(n);
									DOMHelper.setText(n, 0);
								}else{
									DOMHelper.setText(n, v);
									DOMHelper.show(n);
								}
							}
							self.m_grid.onRefresh();
							
						}
					}
				}
			}
		});		
	}
}


SpecialistDocumentBtnSign.prototype.signDocumentCont = function(docId, dlgCont){
	this.m_operationId = CommonHelper.md5(DateHelper.time().toString());
	this.startOperationMonitor();
	var file = dlgCont.getContent().getSignatureFile();
	var pm = (new SpecialistDocument_Controller()).getPublicMethod("sign");
	pm.setFieldValue("id", docId);
	pm.setFieldValue("signature", [file]);
	pm.setFieldValue("operation_id", this.m_operationId);
	var self = this;	
	pm.run({
		"ok":function(resp){
			window.showTempNote("Выполняется подписание документа", null, 1000);
			dlgCont.close();
		},
		"fail":function(){
			self.unsubscribeFromSrvEvents();
		}
	})
}

SpecialistDocumentBtnSign.prototype.signDocument = function(){
	var id;
	if(this.m_cmd){
		this.m_grid.setModelToCurrentRow();
		id = this.m_grid.getModel().getFieldValue("id");
	}
	else{
		id = this.m_getId();
	}
	if(!id){
		throw new Error('Не выбран документ');
	}
	
	var sqare_d = Math.min(window.innerWidth, window.innerHeight)
	
	var view = new Sign_View(this.getId()+":form:body:Sign_View", {});
	//open sign form
	var self = this;
	(new WindowFormModalBS(this.getId()+":form",{
		"cmdCancel":this.m_cmdCancel,
		"controlCancelCaption":"Отмена",
		"controlCancelTitle":"Отменить создание подписи, закрыть форму",
		"cmdOk":true,
		"controlOkCaption":"Ок",
		"controlOkTitle":"Записать подпись",
		"onClickCancel":function(){
			if(this.close){
				this.close();
			}
		},		
		"onClickOk":function(){			
			self.signDocumentCont(id, this);
		},				
		"content":view,
		"contentHead":"Поставьте подпись",
		"dialogWidth":sqare_d + "px"
	})).open();
	
	var offset = document.getElementById(this.getId() + ":form:head").offsetHeight + document.getElementById(this.getId() + ":form:footer").offsetHeight;	
	var canvas = sqare_d - offset;
	
	view.setCanvasSize(canvas);
	
}

