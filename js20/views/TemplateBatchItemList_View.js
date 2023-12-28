/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function TemplateBatchItemList_View(id,options){	

	options = options || {};

	TemplateBatchItemList_View.superclass.constructor.call(this,id,options);
	
	var model = (options.models && options.models.TemplateBatchItemList_Model)? options.models.TemplateBatchItemList_Model : new TemplateBatchItemList_Model();
	var contr = new TemplateBatchItem_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);
	
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	
	var filters;
	this.addElement(new GridAjx(id+":grid",{
		"filters": filters,
		"model":model,
		"keyIds":["id"],
		"controller":contr,
		"editInline":true,
		"editWinClass":null,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"exportFileName" :"Шаблоны"
		}),		
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.TemplateBatchItemList_Model)? options.detailFilters.TemplateBatchItemList_Model:null,	
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:row0",{
					"elements":[
						new GridCellHead(id+":grid:head:templates_ref",{
							"value":"Шаблон",
							"columns":[
								new GridColumnRef({
									"field":model.getField("templates_ref"),
									"ctrlClass": DocumentTemplateEdit,
									"ctrlOptions":{
										"labelCaption":""
									},
									"form":DocumentTemplate_Form,
									"ctrlBindFieldId":"template_id"
								})
							]
						})
						,new GridCellHead(id+":grid:head:studios_ref",{
							"value":"Студия",
							"columns":[
								new GridColumnRef({
									"field":model.getField("studios_ref"),
									"ctrlClass": StudioEdit,
									"ctrlOptions":{
										"labelCaption":""
									},
									"form":Studio_Form,
									"ctrlBindFieldId":"studio_id"
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
extend(TemplateBatchItemList_View,ViewAjxList);

