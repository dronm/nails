
/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>,2023

 * @class
 * @classdesc
 
 * @requires core/extend.js  
 * @requires controls/GridCmd.js

 * @param {string} id Object identifier
 * @param {namespace} options
*/
function YClTransactionGridCmdApiGet(id,options){
	options = options || {};	

	options.showCmdControl = (options.showCmdControl!=undefined)? options.showCmdControl:true;
	
	options.glyph = "glyphicon-refresh";
	options.caption = " yClients";
	options.title = "Синхронизировать с yClients";
	YClTransactionGridCmdApiGet.superclass.constructor.call(this,id,options);
		
}
extend(YClTransactionGridCmdApiGet, GridCmd);

/* Constants */


/* private members */

/* protected*/


/* public methods */
YClTransactionGridCmdApiGet.prototype.onCommand = function(){
	var self = this;
	(new WindowFormModalBS("YClTransactionRefresh",{
		"dialogWidth":"50%",
		"cmdOk":true,
		"cmdCancel":true,		
		"onClickOk":function(){
			self.onCommandCont(this, );
		},
		"onClickCancel":function(){
			this.close();
		},
		"cmdClose":true,
		"content":new View("YClTransactionRefresh:view", {
			"elements":[
				new EditDate("YClTransactionRefresh:view:date_from", {
					"labelCaption":"Период с:",
					"title":"Период начала получения данных"
				})
				,new EditDate("YClTransactionRefresh:view:date_to", {
					"labelCaption":"Период по:",
					"title":"Период окончания получения данных"
				})
			]
		})
	})).open();

}

YClTransactionGridCmdApiGet.prototype.onCommandCont = function(dlgCont){
	var dlg_v = dlgCont.getContent();
	var operation_id = CommonHelper.md5(DateHelper.time().toString());
	this.startOperationMonitor(operation_id);
	var pm = (new YClTransaction_Controller()).getPublicMethod("api_get");
	pm.setFieldValue("operation_id", operation_id);
	pm.setFieldValue("date_from", dlg_v.getElement("date_from").getValue());	
	pm.setFieldValue("date_to", dlg_v.getElement("date_to").getValue());		
	var self = this;
	pm.run({
		"ok":function(){
			window.showTempNote("Начат обмен с yClients", null, 1000);
			dlgCont.close();
		}
		,"fail":function(resp,erCode,erStr){
			self.unsubscribeFromSrvEvents();
			throw new Error(erStr);
		}		
	})
}

YClTransactionGridCmdApiGet.prototype.startOperationMonitor = function(operationId){	
	var srv = window.getApp().getAppSrv();
	if(srv && srv.connActive()){
		var self = this;		
		this.unsubscribeFromSrvEvents();
		this.subscribeToSrvEvents({
			"events":[
				{"id":"UserOperation." + operationId}
				,{"id":"UserOperation." + operationId}
			]
			,"onEvent":function(json){
				if(json.controllerId == "UserOperation" && json.methodId == operationId){					
					if(json.params.status=="end"){
						self.unsubscribeFromSrvEvents();
						if(!json.params.res){
							window.showTempError("Ошибка получения данных от yClients", null, 5000);
						}else{
							self.m_grid.onRefresh(function(){
								window.showTempOk("Список cинхронизирован с yClients", null, 5000);	
							});							
						}
					}
				}
			}
		});		
	}
}
//Control functions!
YClTransactionGridCmdApiGet.prototype.unsubscribeFromSrvEvents = function(){
	if(this.m_srvEventsId && window.getApp().getAppSrv()){
		console.log("Control.prototype.unsubscribeFromSrvEvents")
		window.getApp().getAppSrv().unsubscribe(this.m_srvEventsId);
		this.m_srvEventsId = undefined;
	}
}

YClTransactionGridCmdApiGet.prototype.subscribeToSrvEvents = function(srvEvents){
	if(srvEvents && srvEvents.events && srvEvents.events.length && srvEvents.onEvent && window.getApp().getAppSrv()){
		this.m_srvEventsId = window.getApp().getAppSrv().subscribe(srvEvents);
	}
}
