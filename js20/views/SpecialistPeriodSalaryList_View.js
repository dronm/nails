/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function SpecialistPeriodSalaryList_View(id,options){	

	options = options || {};
	options.HEAD_TITLE = "Зарплата сотрудников";

	SpecialistPeriodSalaryList_View.superclass.constructor.call(this,id,options);
	
	var model = (options.models && options.models.SpecialistPeriodSalaryList_Model)? options.models.SpecialistPeriodSalaryList_Model : new SpecialistPeriodSalaryList_Model();
	var contr = new SpecialistPeriodSalary_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);
	
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	
	var filters;
	if (!options.detail){
		var period_ctrl = new EditPeriodDate(id+":filter-ctrl-period",{
			"field":new FieldDate("period")
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
		"editInline":false,
		"editWinClass":SpecialistPeriodSalary_Form,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"exportFileName" :"Начисления"
		}),		
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.SpecialistPeriodSalaryList_Model)? options.detailFilters.SpecialistPeriodSalaryList_Model:null,	
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:row0",{
					"elements":[
						new GridCellHead(id+":grid:head:studios_ref",{
							"value":"Период",
							"columns":[
								new GridColumnRef({
									"field":model.getField("studios_ref"),
									"ctrlClass":StudioEdit,
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
							"sortable":true,
							"sort":"desc"							
						})
						,new GridCellHead(id+":grid:head:work_total",{
							"value":"Выручка",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("work_total"),
									"precision":"2"
								})
							]
						})
						,new GridCellHead(id+":grid:head:hours",{
							"value":"Часы",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("hours"),
									"precision":"2"
								})
							]
						})
						
						,new GridCellHead(id+":grid:head:debet",{
							"value":"Начисления",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("debet"),
									"precision":"2"
								})
							]
						})
						,new GridCellHead(id+":grid:head:kredit",{
							"value":"Удержания",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("kredit"),
									"precision":"2"
								})
							]
						})
						,new GridCellHead(id+":grid:head:rent_total",{
							"value":"Аренда",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("rent_total"),
									"precision":"2"
								})
							]
						})
						,new GridCellHead(id+":grid:head:work_total_salary",{
							"value":"Для з/п",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("work_total_salary"),
									"precision":"2"
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
							"colSpan":"2"
						})											
						,new GridCellFoot(id+":grid:foot:tot_work_total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_work_total",
							"gridColumn":new GridColumnFloat({"id":"tot_work_total"})
						})
						,new GridCellFoot(id+":grid:foot:tot_hours",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_hours",
							"gridColumn":new GridColumnFloat({"id":"tot_hours"})
						})
						,new GridCellFoot(id+":grid:foot:tot_debet",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_debet",
							"gridColumn":new GridColumnFloat({"id":"tot_debet"})
						})
						,new GridCellFoot(id+":grid:foot:tot_kredit",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_kredit",
							"gridColumn":new GridColumnFloat({"id":"tot_kredit"})
						})
						,new GridCellFoot(id+":grid:foot:tot_rent_total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_rent_total",
							"gridColumn":new GridColumnFloat({"id":"tot_rent_total"})
						})
						,new GridCellFoot(id+":grid:foot:tot_work_total_salary",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_work_total_salary",
							"gridColumn":new GridColumnFloat({"id":"tot_work_total_salary"})
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
	}));	
}
extend(SpecialistPeriodSalaryList_View,ViewAjxList);

