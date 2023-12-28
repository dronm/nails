/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function YClStaffList_View(id,options){	

	options.HEAD_TITLE = "Сотрудники yclients";
	YClStaffList_View.superclass.constructor.call(this,id,options);

	var model = (options.models && options.models.YClStaffList_Model)? options.models.YClStaffList_Model : new YClStaffList_Model();
	var contr = new YClStaff_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);

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
			"exportFileName":"Сотрудники_yclients",
			"cmdSearch": !options.detailFilters,
			"cmdAllCommands": !options.detailFilters,
			"addCustomCommandsAfter":function(commands){
				commands.push(new YClStaffGridCmdApiGetAll(id+":grid:cmd:api_get_all"));
			}
		}),		
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.YClStafflList_Model)? options.detailFilters.YClStaffList_Model:null,
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
						,new GridCellHead(id+":grid:head:name",{
							"value":"Наименование",
							"columns":[
								new GridColumn({
									"field":model.getField("name")
								})
							],
							"sortable":true,
							"sort":"asc"														
						})
						,new GridCellHead(id+":grid:head:specialization",{
							"value":"Специализация",
							"columns":[
								new GridColumn({
									"field":model.getField("specialization")
								})
							]
						})
						,new GridCellHead(id+":grid:head:rating",{
							"value":"Рейтинг",
							"columns":[
								new GridColumn({
									"field":model.getField("rating")
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
		"rowSelect":false,
		"focus":true
	}));	
	
}
extend(YClStaffList_View,ViewAjxList);

