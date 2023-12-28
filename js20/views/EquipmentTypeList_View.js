/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function EquipmentTypeList_View(id,options){	

	options.HEAD_TITLE = "Виды оборудования";
	EquipmentTypeList_View.superclass.constructor.call(this,id,options);

	var model = (options.models && options.models.EquipmentType_Model)? options.models.EquipmentType_Model : new EquipmentType_Model();
	var contr = new EquipmentType_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);

	var self = this;
	var filters;
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	this.addElement(new GridAjx(id+":grid",{
		"model":model,
		"controller":contr,
		"editInline":true,
		"editWinClass":null,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"filters": filters,
			"exportFileName":"ВидыОборудования",
			"cmdSearch": !options.detailFilters,
			"cmdAllCommands": !options.detailFilters		
		}),		
		"popUpMenu":popup_menu,
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:operation_result",{
					"elements":[
						new GridCellHead(id+":grid:head:name",{
							"value":"Наименование",
							"columns":[
								new GridColumn({
									"field":model.getField("name")
								})
							],
							"sortable":true,
							"sort":"desc"
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
extend(EquipmentTypeList_View,ViewAjxList);

