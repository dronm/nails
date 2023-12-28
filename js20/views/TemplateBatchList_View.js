/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function TemplateBatchList_View(id,options){	

	options = options || {};
	options.HEAD_TITLE = "Пакеты шаблонов";

	TemplateBatchList_View.superclass.constructor.call(this,id,options);
	
	var model = (options.models && options.models.TemplateBatch_Model)? options.models.TemplateBatch_Model : new TemplateBatch_Model();
	var contr = new TemplateBatch_Controller();
	
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
			"exportFileName" :"Пакеты шаблонов"
		}),		
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.TemplateBatch_Model)? options.detailFilters.TemplateBatch_Model:null,	
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:row0",{
					"elements":[
						new GridCellHead(id+":grid:head:name",{
							"value":"Наименование",
							"columns":[
								new GridColumn({
									"field":model.getField("name"),
									"master":true,
									"detailViewClass":TemplateBatchItemList_View,
									"detailViewOptions":{
										"detailFilters":{
											"TemplateBatchItemList_Model":[
												{
												"masterFieldId":"id",
												"field":"template_batch_id",
												"sign":"e",
												"val":"0"
												}	
											]
										}													
									}																												
								})
							],
							"sortable":true,
							"sort":"asc"							
						})
						,new GridCellHead(id+":grid:head:template_batch_type",{
							"value":"Предопределенный тип",
							"columns":[
								new EnumGridColumn_template_batch_types({
									"field":model.getField("template_batch_type"),
									"ctrlClass":Enum_template_batch_types,
									"ctrlBindFieldId":"template_batch_type"
									
								})
							]
						})
						/*
						 * это не используется!!
						 * Используются шаблоны набора по салонам
						,new GridCellHead(id+":grid:head:studios_ref",{
							"value":"Студия",
							"colAttrs":{"title":"Набор шаблонов "},
							"columns":[
								new GridColumnRef({
									"field":model.getField("studios_ref"),
									"ctrlClass":StudioEdit,
									"ctrlForm":Studio_Form,
									"ctrlOptions":{
										"labelCaption":""
									},								
									"ctrlBindFieldId":"studio_id"
									
								})
							]
						})
						*/
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
extend(TemplateBatchList_View,ViewAjxList);

