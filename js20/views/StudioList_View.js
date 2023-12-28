/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function StudioList_View(id,options){	

	options = options || {};
	options.HEAD_TITLE = "Студии";

	StudioList_View.superclass.constructor.call(this,id,options);
	
	var model = (options.models && options.models.StudioList_Model)? options.models.StudioList_Model : new StudioList_Model();
	var contr = new Studio_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);
	
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	
	this.addElement(new GridAjx(id+":grid",{
		"model":model,
		"controller":contr,
		"editInline":false,
		"editWinClass":Studio_Form,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"exportFileName" :"Студии"
		}),		
		"popUpMenu":popup_menu,
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:row0",{
					"elements":[
						new GridCellHead(id+":grid:head:name",{
							"value":"Наименование",
							"columns":[
								new GridColumn({"field":model.getField("name")})
							],
							"sortable":true,
							"sort":"asc"							
						})
					]
				})
			]
		}),
		"pagination":new pagClass(id+"_page",
			{"countPerPage":constants.doc_per_page_count.getValue()}),		
		
		"autoRefresh":false,
		"refreshInterval":constants.grid_refresh_interval.getValue()*1000,
		"rowSelect":false,
		"focus":true
	}));	
}
extend(StudioList_View,ViewAjxList);
