/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function BankPaymentList_View(id,options){	

	options = options || {};
	options.HEAD_TITLE = "Платежные поручения";

	BankPaymentList_View.superclass.constructor.call(this,id,options);
	
	var model = (options.models && options.models.BankPaymentList_Model)? options.models.BankPaymentList_Model : new BankPaymentList_Model();
	var contr = new BankPayment_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);
	
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	
	var filters;
	if (!options.detail){
		var period_ctrl = new EditPeriodDate(id+":filter-ctrl-period",{
			"field":new FieldDate("document_date")
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
		"keyIds":["id"],
		"controller":contr,
		"editInline":false,
		"editWinClass":BankPayment_Form,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"exportFileName" :"Платежные поручения"
		}),		
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.BankPaymentList_Model)? options.detailFilters.BankPaymentList_Model:null,	
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:row0",{
					"elements":[
						new GridCellHead(id+":grid:head:document_num",{
							"value":"Дата",
							"columns":[
								new GridColumn({
									"field":model.getField("document_num")
								})
							],
							"sortable":true,
							"sort":"desc"
						})
						,new GridCellHead(id+":grid:head:document_date",{
							"value":"Дата",
							"columns":[
								new GridColumnDate({
									"field":model.getField("document_date")
								})
							],
							"sortable":true
						})
						,new GridCellHead(id+":grid:head:document_total",{
							"value":"Сумма",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("document_total"),
									"precision":"2"
								})
							]
						})						
						,new GridCellHead(id+":grid:head:specialists_ref",{
							"value":"Мастер",
							"columns":[
								new GridColumnRef({
									"field":model.getField("specialists_ref"),
									"ctrlClass":SpecialistEdit,
									"ctrlOptions":{
										"lableCaption":""
									},
									"ctrlBindFieldId":"specialist_id"
								})
							]
						})
						,new GridCellHead(id+":grid:head:specialist_period_salary_details_ref",{
							"value":"Документ",
							"columns":[
								new GridColumnRef({
									"field":model.getField("specialist_period_salary_details_ref"),
									"ctrlEdit":false
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
						,new GridCellFoot(id+":grid:foot:tot_document_total",{
							"attrs":{"align":"right", "nowrap":"nowrap"},
							"totalFieldId":"total_document_total",
							"gridColumn":new GridColumnFloat({"id":"tot_document_total"})
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
extend(BankPaymentList_View,ViewAjxList);

