/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function YClTransactionDocAllList_View(id,options){	

	if (!options.detail){
		options.HEAD_TITLE = "Документы yclients";
	}
	YClTransactionDocAllList_View.superclass.constructor.call(this,id,options);

	var model = (options.models && options.models.YClTransactionDocAllList_Model)? options.models.YClTransactionDocAllList_Model : new YClTransactionDocAllList_Model();
	var contr = new YClTransaction_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);

	this.m_onSelectDoc = options.onSelectDoc;

	var filters;
	if (!options.detail){
		var period_ctrl = new EditPeriodDate(id+":filter-ctrl-period",{
			"field":new FieldDateTime("date")
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

	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	var self = this;
	this.addElement(new GridAjx(id+":grid",{
		"model":model,
		"keyIds":["document_id"],
		"controller":contr,
		"readPublicMethod":contr.getPublicMethod("get_doc_all_list"),
		"editInline":true,
		"editWinClass":null,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"filters": filters,
			"exportFileName":"Транзакции_yclients",
			"cmdSearch": true,
			"cmdAllCommands": true,
			"cmdInsert":false,
			"cmdEdit":false
		}),
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.YClTransactionDocAllList_Model)? options.detailFilters.YClTransactionDocAllList_Model:null,
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:row0",{
					"elements":[
						options.detail? null : new GridCellHead(id+":grid:head:specialists_ref",{
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
						,new GridCellHead(id+":grid:head:document_id",{
							"value":"ID документа",
							"columns":[
								new GridColumn({
									"field":model.getField("document_id")
								})
							],
							"sortable":true
						})
						,new GridCellHead(id+":grid:head:date",{
							"value":"Дата",
							"columns":[
								new GridColumnDateTime({
									"field":model.getField("date")
								})
							],
							"sortable":true,
							"sort":"desc"
						})
						,new GridCellHead(id+":grid:head:client",{
							"value":"Клиент",
							"columns":[
								new GridColumn({
									"field":model.getField("client")
								})
							]
						})
						,new GridCellHead(id+":grid:head:amount",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"value":"Сумма",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("amount"),
									"precision":"2",
									"length":"15"
								})
							]
						})
						,new GridCellHead(id+":grid:head:len_hour",{
							"value":"Часы",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("len_hour")
								})
							]
						})
						
						,new GridCellHead(id+":grid:head:admin_rate",{
							"value":"Оценка",
							"columns":[
								new GridColumn({
									"field":model.getField("admin_rate")
								})
							]
						})						
						,new GridCellHead(id+":grid:head:photo",{
							"value":"Фото",
							"colAttrs":{"title":"Открыть файл","style":"cursor:pointer;"},
							"columns":[
								new GridColumnPicture({
									"field":model.getField("photo"),
									"onClick":(function(cont){
										return function(fields, elem){
											var tr = DOMHelper.getParentByTagName(elem, "tr");
											window.getApp().openGridAttachment(cont.getElement("grid").getModel(), tr, "photo", "specialist_works", 1);
										}
									})(self)
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
						new GridCell(id+":grid:foot:total_sp1",{
							"colSpan": options.detail? "3" : "4"
						})											
						,new GridCellFoot(id+":grid:foot:tot_amount",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_amount",
							"gridColumn":new GridColumnFloat({"id":"tot_amount"})
						})
						,new GridCellFoot(id+":grid:foot:tot_len_hour",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_len_hour",
							"gridColumn":new GridColumnFloat({"id":"tot_len_hour"})
						})
						
					]
				})		
			]
		}),
		
		"pagination":new pagClass(id+"_page",
			{"countPerPage":constants.doc_per_page_count.getValue()}),		
		
		"autoRefresh":options.detailFilters||options.onSelectDoc? true:false,
		"refreshInterval":constants.grid_refresh_interval.getValue()*1000,
		"rowSelect":false,
		"focus":true
	}));	
	
}
extend(YClTransactionDocAllList_View,ViewAjxList);

