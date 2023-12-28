/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function SpecialityList_View(id,options){	

	options.HEAD_TITLE = "Специальности";
	SpecialityList_View.superclass.constructor.call(this,id,options);

	var model = (options.models && options.models.SpecialityList_Model)? options.models.SpecialityList_Model : new SpecialityList_Model();
	var contr = new Speciality_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);

	var self = this;
	var filters;
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	this.addElement(new GridAjx(id+":grid",{
		"model":model,
		"controller":contr,
		"editInline":false,
		"editWinClass":Speciality_Form,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"filters": filters,
			"exportFileName":"Специальности",
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
extend(SpecialityList_View,ViewAjxList);

