
/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>,2023

 * @class
 * @classdesc
 
 * @requires core/extend.js  
 * @requires controls/GridCmd.js

 * @param {string} id Object identifier
 * @param {namespace} options
*/
function YClStaffGridCmdApiGetAll(id,options){
	options = options || {};	

	options.showCmdControl = (options.showCmdControl!=undefined)? options.showCmdControl:true;
	
	options.glyph = "glyphicon-refresh";
	options.caption = " yClients";
	options.title = "Синхронизировать с yClients";
	YClStaffGridCmdApiGetAll.superclass.constructor.call(this,id,options);
		
}
extend(YClStaffGridCmdApiGetAll, GridCmd);

/* Constants */


/* private members */

/* protected*/


/* public methods */
YClStaffGridCmdApiGetAll.prototype.onCommand = function(){
	var operation_id = CommonHelper.md5(DateHelper.time().toString());
	this.startOperationMonitor(operation_id);
	var pm = (new YClStaff_Controller()).getPublicMethod("api_get_all");
	pm.setFieldValue("operation_id", operation_id);	
	var self = this;
	pm.run({
		"ok":function(){
			window.showTempNote("Начат обмен с yClients", null, 1000);
		}
		,"fail":function(resp,erCode,erStr){
			self.unsubscribeFromSrvEvents();
			throw new Error(erStr);
		}		
	})
}

YClStaffGridCmdApiGetAll.prototype.startOperationMonitor = function(operationId){	
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
YClStaffGridCmdApiGetAll.prototype.unsubscribeFromSrvEvents = function(){
	if(this.m_srvEventsId && window.getApp().getAppSrv()){
		console.log("Control.prototype.unsubscribeFromSrvEvents")
		window.getApp().getAppSrv().unsubscribe(this.m_srvEventsId);
		this.m_srvEventsId = undefined;
	}
}

YClStaffGridCmdApiGetAll.prototype.subscribeToSrvEvents = function(srvEvents){
	if(srvEvents && srvEvents.events && srvEvents.events.length && srvEvents.onEvent && window.getApp().getAppSrv()){
		this.m_srvEventsId = window.getApp().getAppSrv().subscribe(srvEvents);
	}
}
