/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function SpecialistDocumentOnRegisterList_View(id,options){	

	options.HEAD_TITLE = "Документы для регистрации мастера";
	SpecialistDocumentOnRegisterList_View.superclass.constructor.call(this,id,options);

	var model = (options.models && options.models.SpecialistDocumentOnRegisterList_Model)? options.models.SpecialistDocumentOnRegisterList_Model : new SpecialistDocumentOnRegisterList_Model();
	var contr = new SpecialistDocumentOnRegister_Controller();
	
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
			"exportFileName":"ДокументыДляРегистрацииМастера",
			"cmdSearch": !options.detailFilters,
			"cmdAllCommands": !options.detailFilters		
		}),		
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.SpecialistDocumentOnRegisterlList_Model)? options.detailFilters.SpecialistDocumentOnRegisterList_Model:null,
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:row0",{
					"elements":[
						new GridCellHead(id+":grid:head:document_templates_ref",{
							"value":"Шаблон",
							"columns":[
								new GridColumnRef({
									"field":model.getField("document_templates_ref"),
									"ctrlBindFieldId":"document_template_id",									
									"ctrlClass":DocumentTemplateEdit,
									"ctrlOptions":{
										"labelCaption":""
									},
									"searchOptions":{
										"searchType":"on_match",
										"typeChange":false
									}							
								})
							]
						})	
						,new GridCellHead(id+":grid:head:need_signing",{
							"value":"Нужна подпись",
							"columns":[
								new GridColumnBool({
									"field":model.getField("need_signing")
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
extend(SpecialistDocumentOnRegisterList_View,ViewAjxList);

