/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function SpecialistSalaryKreditList_View(id,options){	

	options = options || {};
	
	if (!options.detail){
		options.HEAD_TITLE = "Удержания сотрудников";
	}
	SpecialistSalaryKreditList_View.superclass.constructor.call(this,id,options);
	
	var model = (options.models && options.models.SpecialistSalaryKreditList_Model)? options.models.SpecialistSalaryKreditList_Model : new SpecialistSalaryKreditList_Model();
	var contr = new SpecialistSalaryKredit_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);
	
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	
	var filters;
	if (!options.detail){
		var period_ctrl = new EditPeriodDate(id+":filter-ctrl-period",{
			"field":new FieldDateTime("date")
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
	this.addElement(new GridAjx(id+":grid",{
		"filters": filters,
		"model":model,
		"controller":contr,
		"editInline":true,
		"editWinClass":null,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"exportFileName" :"Удержания"
		}),		
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.SpecialistSalaryKreditList_Model)? options.detailFilters.SpecialistSalaryKreditList_Model:null,	
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:row0",{
					"elements":[
						new GridCellHead(id+":grid:head:date_time",{
							"value":"Дата",
							"columns":[
								new GridColumnDateTime({
									"field":model.getField("date_time")
								})
							],
							"sortable":true,
							"sort":"desc"							
						})
						,options.detail? null : new GridCellHead(id+":grid:head:specialists_ref",{
							"value":"Сотрудник",
							"columns":[
								new GridColumnRef({
									"field":model.getField("specialists_ref"),
									"ctrlClass":SpecialistEdit,
									"ctrlBindFieldId":"specialist_id",
									"ctrlOptions":{
										"labelCaption":""
									}
								})
							]
						})
						
						,new GridCellHead(id+":grid:head:salary_kredits_ref",{
							"value":"Удержание",
							"columns":[
								new GridColumnRef({
									"field":model.getField("salary_kredits_ref"),
									"ctrlClass":SalaryKreditEdit,
									"ctrlBindFieldId":"salary_kredit_id",
									"ctrlOptions":{
										"labelCaption":""
									}
								})
							]
						})
						
						,new GridCellHead(id+":grid:head:total",{
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
							"colSpan": options.detail? "2" : "3"
						})											
						,new GridCellFoot(id+":grid:foot:tot_total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							//"calcOper":"sum",
							//"calcFieldId":"total",
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
	}));	
}
extend(SpecialistSalaryKreditList_View,ViewAjxList);

