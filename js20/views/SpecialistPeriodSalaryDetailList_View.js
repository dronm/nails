/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function SpecialistPeriodSalaryDetailList_View(id,options){	

	options = options || {};
	if (!options.detail){
		options.HEAD_TITLE = "Зарплата сотрудников";
	}

	SpecialistPeriodSalaryDetailList_View.superclass.constructor.call(this,id,options);
	
	var model = (options.models && options.models.SpecialistPeriodSalaryDetailList_Model)? options.models.SpecialistPeriodSalaryDetailList_Model : new SpecialistPeriodSalaryDetailList_Model();
	var contr = new SpecialistPeriodSalaryDetail_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);
	
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	
	var role = window.getApp().getServVar("role_id");
	
	var filters;
	if (!options.detail){
		var period_ctrl = new EditPeriodDate(id+":filter-ctrl-period",{
			"field":new FieldDate("period")
		});
	
		filters = {
			"period":{
				"binding":new CommandBinding({
					"control":period_ctrl,
					"field":period_ctrl.getField()
				}),
				"bindings":[
					{"binding":new CommandBinding({
						"control":period_ctrl.getControlFrom(),
						"field":period_ctrl.getField()
						}),
					"sign":"ge"
					},
					{"binding":new CommandBinding({
						"control":period_ctrl.getControlTo(),
						"field":period_ctrl.getField()
						}),
					"sign":"le"
					}
				]
			}		
		}
	}
	var self = this;
	this.addElement(new GridAjx(id+":grid",{
		"filters": filters,
		"model":model,
		"controller":contr,
		"editInline":false,
		"editWinClass":WindowFormModalBS,		
		"editViewClass":SpecialistPeriodSalaryDetailDialog_View,
		"editViewOptions":function(){
			return {
				"dialogWidth":"80%"
				//"specialist_period_salary_id":self.m_mainView.getElement("id").getValue()
			}
		},
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"cmdInsert": !options.detail && role!="specialist",
			"cmdEdit": !options.detail && role!="specialist",
			"exportFileName" :"Зарплата"
		}),		
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.SpecialistPeriodSalaryDetailList_Model)? options.detailFilters.SpecialistPeriodSalaryDetailList_Model:null,	
		"head":new GridHead(id+"-grid:head",{
			"elements":[				
				new GridRow(id+":grid:head:row0",{
					"elements":[
						(options.detail || role=="specialist")? null : new GridCellHead(id+":grid:head:line_num",{
							"value":"№",
							"columns":[
								new GridColumn({
									"field":model.getField("line_num"),
									"ctrlClass":EditNum
								})
							],
							"sortable":true,
							"sort":"desc"							
						})
						,(options.detail || role=="specialist")? null : new GridCellHead(id+":grid:head:specialists_ref",{
							"value":"Мастер",
							"columns":[
								new GridColumnRef({
									"field":model.getField("specialists_ref"),
									"ctrlClass":SpecialistEdit,
									"ctrlBindFieldId":"specialist_id",
									"ctrlOptions":{
										"labelCaption":""
									}
								})
							],
							"sortable":true
						})
						/*,options.detailFilters? null : new GridCellHead(id+":grid:head:studios_ref",{
							"value":"Студия",
							"columns":[
								new GridColumnRef({
									"field":model.getField("studios_ref"),
									"ctrlClass":SpecialistEdit,
									"ctrlBindFieldId":"studio_id",
									"ctrlOptions":{
										"labelCaption":""
									}									
								})
							],
							"sortable":true
						})						
						*/
						,!options.detail? null : options.detailFilters? null : new GridCellHead(id+":grid:head:period_descr",{
							"value":"Период",
							"columns":[
								new GridColumn({
									"field":model.getField("period_descr"),
									"ctrlClass":EditDate,
									"ctrlBindFieldId":"period"
								})
							],
							"sortable":true
						})
						
						,new GridCellHead(id+":grid:head:work_total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"value":"Выручка",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("work_total"),
									"precision":"2"
								})
							]
						})
						,new GridCellHead(id+":grid:head:debet",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"value":"Доп.нач",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("debet"),
									"precision":"2"
								})
							]
						})
						,new GridCellHead(id+":grid:head:kredit",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"value":"Доп.уд",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("kredit"),
									"precision":"2"
								})
							]
						})
						,new GridCellHead(id+":grid:head:rent_price",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"value":"Цена аренда",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("rent_price"),
									"precision":"2"
								})
							]
						})
						,new GridCellHead(id+":grid:head:hours",{
							"value":"Часы",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("hours"),
									"precision":"2"
								})
							]
						})
						
						,new GridCellHead(id+":grid:head:rent_total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"value":"Сумма аренда",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("rent_total"),
									"precision":"2"
								})
							]
						})
						,new GridCellHead(id+":grid:head:agent_percent",{
							"value":"%",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("agent_percent"),
									"precision":"2"
								})
							]
						})
						
						,new GridCellHead(id+":grid:head:work_total_salary",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"value":"Сумма для з/п",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("work_total_salary"),
									"precision":"2"
								})
							]
						})
						
						,new GridCellHead(id+":grid:head:total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"value":"Сумма",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("total"),
									"precision":"2"
								})
							]
						})
						//TODO: add style if total==receipt_total
						,new GridCellHead(id+":grid:head:receipt_total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"value":"Чек",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("receipt_total"),
									"precision":"2"
								})
							]
						})
						//TODO: add style if receipt_checked==FALSE
						,new GridCellHead(id+":grid:head:receipt_error",{
							"value":"Ошибка",
							"columns":[
								new GridColumn({
									"field":model.getField("receipt_error")
								})
							]
						})
						,new GridCellHead(id+":grid:head:receipt_href",{
							"value":"Ссылка",
							"columns":[
								new GridColumn({
									"field":model.getField("receipt_href")
								})
							]
						})						
						//TODO:Если не закрыт - возможность загрузки
						,new GridCellHead(id+":grid:head:receipt_photos",{
							"value":"Файл",
							"columns":[
								new GridColumn({
									//"field":model.getField("receipt_photos"),
									"cellOptions":function(column,row){
										return self.onSetPhotoOpts(this, column, row);
									}
								})
							]
						})
						,new GridCellHead(id+":grid:head:bank_payments_ref",{
							"value":"п/п",
							"columns":[
								new GridColumnRef({
									"field":model.getField("bank_payments_ref")
								})
							]
						})
						
					]
				})
			]
		}),
		"foot":new GridFoot(id+"grid:foot",{
			"autoCalc":true,			
			"elements":[
				new GridRow(id+":grid:foot:row0",{
					"elements":[
						(options.detail || role=="specialist")? null : new GridCell(id+":grid:foot:total_sp1",{
							"colSpan":  "3"
						})
						,new GridCellFoot(id+":grid:foot:tot_work_total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_work_total",
							"gridColumn":new GridColumnFloat({"id":"tot_work_total"})
						})
						,new GridCellFoot(id+":grid:foot:tot_debet",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_debet",
							"gridColumn":new GridColumnFloat({"id":"tot_debet"})
						})
						,new GridCellFoot(id+":grid:foot:tot_kredit",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_kredit",
							"gridColumn":new GridColumnFloat({"id":"tot_kredit"})
						})
						,new GridCellFoot(id+":grid:foot:tot_kredit",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_kredit",
							"gridColumn":new GridColumnFloat({"id":"tot_kredit"})
						})
						,new GridCell(id+":grid:foot:total_sp2",{
							"colSpan":"1"
						})											
						,new GridCellFoot(id+":grid:foot:tot_hours",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_hours",
							"gridColumn":new GridColumnFloat({"id":"tot_hours"})
						})
						,new GridCellFoot(id+":grid:foot:tot_rent_total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_rent_total",
							"gridColumn":new GridColumnFloat({"id":"tot_rent_total"})
						})
						,new GridCell(id+":grid:foot:total_sp3",{
							"colSpan":"1"
						})											
						,new GridCellFoot(id+":grid:foot:tot_work_total_salary",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_work_total_salary",
							"gridColumn":new GridColumnFloat({"id":"tot_work_total_salary"})
						})
						,new GridCellFoot(id+":grid:foot:tot_total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_total",
							"gridColumn":new GridColumnFloat({"id":"tot_total"})
						})
						,new GridCellFoot(id+":grid:foot:tot_receipt_total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_receipt_total",
							"gridColumn":new GridColumnFloat({"id":"tot_receipt_total"})
						})
						,new GridCell(id+":grid:foot:total_sp4",{
							"colSpan":"2"
						})											
					]
				})		
			]
		}),
		"pagination":new pagClass(id+"_page",
			{"countPerPage":constants.doc_per_page_count.getValue()}),		
		
		"autoRefresh":options.detailFilters? true:false,
		"refreshInterval":constants.grid_refresh_interval.getValue()*1000,
		"rowSelect":(window.getWidthType()=="sm"),
		"focus":true
	}));	
}
extend(SpecialistPeriodSalaryDetailList_View,ViewAjxList);

SpecialistPeriodSalaryDetailList_View.prototype.addReceipt = function(keyId, fl){
	var oper_id = CommonHelper.md5(DateHelper.time().toString());
	this.startOperationMonitor(oper_id);
	var pm = (new SpecialistReceipt_Controller()).getPublicMethod("add");
	pm.setFieldValue("operation_id", oper_id);
	pm.setFieldValue("specialist_period_salary_detail_id", keyId);
	pm.setFieldValue("file_content", [fl]);
	var self = this;
	pm.run();
}

SpecialistPeriodSalaryDetailList_View.prototype.onSetPhotoOpts = function(gridCont, column, row){
	var self = this;
	var res = {"elements":[]};
	var editable = (gridCont.m_model.getFieldValue("total") != gridCont.m_model.getFieldValue("receipt_total"));

	var ins_id = gridCont.m_model.getFieldValue("id");
	
	var att = gridCont.m_model.getFieldValue("receipt_photos");
	if(att && att.length){
		for(var i = 0; i < att.length; i++){
			if(!att[i]){
				continue;
			}
			var att_i = i;
			res.elements.push(new Control(null,"IMG",{
				"attrs":{
					"src": "data:image/png;base64, "+att[att_i].dataBase64,
					"width": "50px",
					"height": "50px",
					"title": att[att_i]["name"]+" ("+CommonHelper.byteFormat(att[att_i]["size"],2)+")",
					"style":"cursor:pointer;"
				},
				"events":{
					"click":(function(keyId, fileId){
						return function(e){
							self.getAttachment(keyId, fileId);
						}
					})(att[att_i].receipt_id, att[att_i].id)
				}
			}));
			
			//delete
			if(editable){
				res.elements.push(new Control(null,"I",{
					"attrs":{
						"class":"glyphicon glyphicon-trash",
						"title": "Удалить скан",
						"style":"cursor:pointer;"
					},
					"events":{
						"click":(function(receiptId){
							return function(e){
								self.delAttachmentByReceipt(receiptId);
							}
						})(att[att_i].receipt_id)
					}
				}));
			}
		}
	}
	
	if(editable){		
		var ctrl = new ButtonCmd(null,{
			"glyph":"glyphicon-plus",
			"attrs":{"style": "margin-right: 10px;"},
			"caption":" Чек",
			"title": "Добавить чек",
			"onClick":(function(keyId){												
				return function(){
					window.getApp().addPhoto(function(fl){
						self.addReceipt(keyId, fl);
					});					
				}
			})(ins_id)
		});										
		ctrl.m_row = row;		
		res.elements.push(ctrl);		
	}

	return res;
}

SpecialistPeriodSalaryDetailList_View.prototype.delAttachmentByReceipt = function(receiptId){	
	var self = this;
	WindowQuestion.show({
		"no":false,
		"callBackYes":function(){
			var pm = (new SpecialistReceipt_Controller()).getPublicMethod("delete");
			pm.setFieldValue("id", receiptId);
			pm.run({
				"ok":function(){
					self.getElement("grid").onRefresh(function(){
						window.showTempNote("Чек удален", null, 2000);
					});
				}
			})
		},
		"text":"Удалить чек?"
	});
}

SpecialistPeriodSalaryDetailList_View.prototype.getAttachment = function(keyId, fileId){	
	var pm = (new Attachment_Controller()).getPublicMethod("get_file");
	pm.setFieldValue("ref", CommonHelper.serialize(new RefType({"keys": {"id": keyId}, "dataType": "specialist_receipts"})));
	pm.setFieldValue("content_id", fileId);
	pm.setFieldValue("inline",1);	
	var offset = 0;
	var h = $( window ).width()/3*2;
	var left = $( window ).width()/2;
	var w = left - 20;	
	pm.openHref("ViewXML", "location=0,menubar=0,status=0,titlebar=0,top="+(50+offset)+",left="+(left+offset)+",width="+w+",height="+h);
}

//
SpecialistPeriodSalaryDetailList_View.prototype.startOperationMonitor = function(operationId){	
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
					if(json.params.status=="progress"){
						self.getElement("grid").onRefresh(function(){
							window.showTempNote("Чек загружен, выполняется проверка ФНС", null,2000);	
						});						
						
					}else if(json.params.status=="end"){
						self.unsubscribeFromSrvEvents();
						if(!json.params.res){
							self.getElement("grid").onRefresh(function(){
								window.showTempError("Ошибка при проверке чека", null, 5000);
							});
						}else{
							self.getElement("grid").onRefresh(function(){
								window.showTempNote("Проверка чека выполнена", null, 5000);
							});
						}
					}
				}
			}
		});		
	}
}

