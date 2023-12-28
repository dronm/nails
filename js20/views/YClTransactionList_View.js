/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function YClTransactionList_View(id,options){	

	options.HEAD_TITLE = "Транзакции yclients";
	YClTransactionList_View.superclass.constructor.call(this,id,options);

	var model = (options.models && options.models.YClTransactionList_Model)? options.models.YClTransactionList_Model : new YClTransactionList_Model();
	var contr = new YClTransaction_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);

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
	this.addElement(new GridAjx(id+":grid",{
		"model":model,
		"controller":contr,
		"editInline":true,
		"editWinClass":null,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"filters": filters,
			"exportFileName":"Транзакции_yclients",
			"cmdSearch": !options.detailFilters,
			"cmdAllCommands": !options.detailFilters,
			"cmdInsert":false,
			"cmdEdit":false,
			"addCustomCommandsAfter":function(commands){
				commands.push(new YClTransactionGridCmdApiGet(id+":grid:cmd:api_get"));
			}
		}),		
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.YClTransactionlList_Model)? options.detailFilters.YClTransactionList_Model:null,
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:row0",{
					"elements":[
						new GridCellHead(id+":grid:head:id",{
							"value":"ID",
							"columns":[
								new GridColumn({
									"field":model.getField("id")
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
						/*,new GridCellHead(id+":grid:head:client",{
							"value":"Клиент",
							"columns":[
								new GridColumn({
									"field":model.getField("client")
								})
							],
							"sortable":true
						})
						,new GridCellHead(id+":grid:head:client_phone",{
							"value":"Тел.клиента",
							"columns":[
								new GridColumn({
									"field":model.getField("client_phone")
								})
							],
							"sortable":true
						})*/
						,new GridCellHead(id+":grid:head:specialists_ref",{
							"value":"Мастер",
							"columns":[
								new GridColumnRef({
									"field":model.getField("specialists_ref"),
									"form":SpecialistEdit,
									"ctrlBindFieldId":"specialist_id"
								})
							],
							"sortable":true
						})
						,new GridCellHead(id+":grid:head:ycl_staff_ref",{
							"value":"Мастер yclients",
							"columns":[
								new GridColumnRef({
									"field":model.getField("ycl_staff_ref")
								})
							],
							"sortable":true
						})
						,new GridCellHead(id+":grid:head:amount",{
							"value":"Сумма",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("amount"),
									"precision":"2",
									"length":"15"
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
							"colSpan":"7"
						})											
						,new GridCellFoot(id+":grid:foot:tot_amount",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							//"calcOper":"sum",
							//"calcFieldId":"amount",
							"totalFieldId":"total_amount",
							"gridColumn":new GridColumnFloat({"id":"tot_amount"})
						})
					]
				})		
			]
		}),
		
		"pagination":new pagClass(id+"_page",
			{"countPerPage":constants.doc_per_page_count.getValue()}),		
		
		"autoRefresh":options.detailFilters? true:false,
		"refreshInterval":constants.grid_refresh_interval.getValue()*1000,
		"rowSelect":false,
		"focus":true
	}));	
	
}
extend(YClTransactionList_View,ViewAjxList);

