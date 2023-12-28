/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function SpecialistPeriodSalaryDetailForPayList_View(id,options){	

	options = options || {};
	if (!options.detail){
		options.HEAD_TITLE = "Формирование платежных поручений";
	}

	SpecialistPeriodSalaryDetailForPayList_View.superclass.constructor.call(this,id,options);
	
	var model = (options.models && options.models.SpecialistPeriodSalaryDetailList_Model)? options.models.SpecialistPeriodSalaryDetailList_Model : new SpecialistPeriodSalaryDetailList_Model();
	
	var contr = new SpecialistPeriodSalaryDetail_Controller();
	var pm = contr.getPublicMethod("get_for_pay_list");
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);
	
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	
	var role = window.getApp().getServVar("role_id");
	
	var filters;
	var grid = new GridAjx(id+":grid",{
		"filters": filters,
		"model":model,
		"controller":contr,
		"readPublicMethod":pm,
		"editInline":null,
		"editWinClass":null,		
		"editViewClass":null,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"cmdInsert": false,
			"cmdEdit": false,
			"exportFileName" :"Зарплата",
			"addCustomCommandsAfter":function(commands){
				commands.push(new SpecialistPeriodSalaryDetailForPayGridCmdMakeDocs(id+":grid:cmd:makeDocs"));
			}
		}),		
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.SpecialistPeriodSalaryDetailList_Model)? options.detailFilters.SpecialistPeriodSalaryDetailList_Model:null,	
		"head":new GridHead(id+"-grid:head",{
			"elements":[				
				new GridRow(id+":grid:head:row0",{
					"elements":[
						new GridCellHeadMark(id+":grid:head:mark")
						,new GridCellHead(id+":grid:head:specialists_ref",{
							"value":"Мастер",
							"columns":[
								new GridColumnRef({
									"field":model.getField("specialists_ref"),
									"ctrlClass":SpecialistEdit,
									"ctrlBindFieldId":"specialist_id",
									"ctrlOptions":{
										"labelCaption":""
									}
								})
							],
							"sortable":true
						})
						,options.detailFilters? null : new GridCellHead(id+":grid:head:studios_ref",{
							"value":"Студия",
							"columns":[
								new GridColumnRef({
									"field":model.getField("studios_ref"),
									"ctrlClass":SpecialistEdit,
									"ctrlBindFieldId":"studio_id",
									"ctrlOptions":{
										"labelCaption":""
									}									
								})
							],
							"sortable":true
						})												
						,new GridCellHead(id+":grid:head:period_descr",{
							"value":"Период",
							"columns":[
								new GridColumn({
									"field":model.getField("period_descr"),
									"ctrlClass":EditDate,
									"ctrlBindFieldId":"period"
								})
							],
							"sortable":true
						})
						
						,new GridCellHead(id+":grid:head:total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"value":"Сумма",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("total"),
									"precision":"2"
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
							"colSpan":  "4"
						})
						,new GridCellFoot(id+":grid:foot:tot_total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_total",
							"gridColumn":new GridColumnFloat({"id":"tot_total"})
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
	});	
	this.addElement(grid);	
}
extend(SpecialistPeriodSalaryDetailForPayList_View,ViewAjxList);

