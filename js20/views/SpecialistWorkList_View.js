/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function SpecialistWorkList_View(id,options){	

	if(!options.detail){
		options.HEAD_TITLE = "Работы";
	}
	SpecialistWorkList_View.superclass.constructor.call(this,id,options);

	var model = (options.models && options.models.SpecialistWorkList_Model)? options.models.SpecialistWorkList_Model : new SpecialistWorkList_Model();
	var contr = new SpecialistWork_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);

	var role = window.getApp().getServVar("role_id");

	var self = this;
	var filters;
	if (!options.detail){
		var period_ctrl = new EditPeriodDate(id+":filter-ctrl-period",{
			"field":new FieldDateTime("date_time")
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
			"cmdInsert": !options.detail && role!="specialist",
			"cmdEdit": !options.detail && role!="specialist",
			"filters": filters,
			"exportFileName":"Работы",
			"cmdSearch": !options.detailFilters,
			"cmdAllCommands": !options.detailFilters,
			"addCustomCommandsAfter": (options.detail || role=="specialist")? null : function(commands){
				commands.push(new SpecialistWorkGridCmdRate(id+":grid:cmd:rate"));
			}
		}),		
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.SpecialistWorklList_Model)? options.detailFilters.SpecialistWorkList_Model:null,
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:operation_result",{
					"elements":[
						new GridCellHead(id+":grid:head:date_time",{
							"value":"Дата",
							"columns":[
								new GridColumnDateTime({
									"field":model.getField("date_time"),
									"ctrlClass":EditDateTime,
									"ctrlBindFieldId":"date_time"
								})
							],
							"sortable":true,
							"sort":"desc"
						})						
						,(options.detail || role=="specialist")? null : new GridCellHead(id+":grid:head:studios_ref",{
							"value":"Салон",
							"columns":[
								new GridColumnRef({
									"field":model.getField("studios_ref"),
									"form":Studio_Form,
									"ctrlClass":StudioEdit,
									"ctrlBindFieldId":"studio_id",
									"ctrlOptions":{
										"labelCaption":""
									}
								})
							],
							"sortable":true
						})
						,(options.detailFilters || options.detail|| role=="specialist")? null : new GridCellHead(id+":grid:head:specialists_ref",{
							"value":"Мастер",
							"columns":[
								new GridColumnRef({
									"field":model.getField("specialists_ref"),
									"form":Specialist_Form,
									"ctrlClass":SpecialistEdit,
									"ctrlBindFieldId":"specialist_id",
									"ctrlOptions":{
										"labelCaption":""
									}
								})
							],
							"sortable":true
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
						,new GridCellHead(id+":grid:head:amount",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"value":"Выручка",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("amount"),
									"precision":"2"
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
		"pagination":new pagClass(id+"_page",
			{"countPerPage":constants.doc_per_page_count.getValue()}),		
		
		"autoRefresh":options.detailFilters? true:false,
		"refreshInterval":constants.grid_refresh_interval.getValue()*1000,
		"rowSelect":(window.getWidthType()=="sm"),
		"focus":true
	}));	
	
}
extend(SpecialistWorkList_View,ViewAjxList);

