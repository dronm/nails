/** Copyright (c) 2022
 *	Andrey Mikhalevich, Katren ltd.
 */
function FirmList_View(id,options){	

	options = options || {};
	options.HEAD_TITLE = "Организации";

	FirmList_View.superclass.constructor.call(this,id,options);
	
	var model = (options.models && options.models.FirmList_Model)? options.models.FirmList_Model : new FirmList_Model();
	var contr = new Firm_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);
	
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	
	this.addElement(new GridAjx(id+":grid",{
		"model":model,
		"controller":contr,
		"editInline":false,
		"editWinClass":Firm_Form,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"exportFileName" :"Организации"
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
						}),
						new GridCellHead(id+":grid:head:inn",{
							"value":"ИНН",
							"columns":[
								new GridColumn({"field":model.getField("inn")})
							],
							"sortable":true
						}),
						new GridCellHead(id+":grid:head:ogrn",{
							"value":"ОГРН",
							"columns":[
								new GridColumn({"field":model.getField("ogrn")})
							],
							"sortable":true
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
extend(FirmList_View,ViewAjxList);
