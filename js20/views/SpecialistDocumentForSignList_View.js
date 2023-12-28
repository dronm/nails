/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function SpecialistDocumentForSignList_View(id,options){	

	options.HEAD_TITLE = "Документы на подпись";
	SpecialistDocumentForSignList_View.superclass.constructor.call(this,id,options);

	var model = (options.models && options.models.SpecialistDocumentForSignList_Model)? options.models.SpecialistDocumentForSignList_Model : new SpecialistDocumentForSignList_Model();
	var contr = new SpecialistDocument_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);

	var is_specialist = (window.getApp().getServVar("role_id") == "specialist");
	
	var self = this;
	var filters;
	if (!options.detail && !is_specialist){
		var period_ctrl = new EditPeriodDate(id+":filter-ctrl-period",{
			"field":new FieldDateTime("date_time")
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
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	this.addElement(new GridAjx(id+":grid",{
		"model":model,
		"controller":contr,
		"readPublicMethod":contr.getPublicMethod("get_for_sign_list"),
		"editInline":true,
		"editWinClass":null,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"cmdInsert":false,
			"cmdEdit":false,
			"filters": filters,
			"exportFileName":"Документы",
			"cmdSearch": !options.detailFilters && !is_specialist,
			"cmdAllCommands": !options.detailFilters && !is_specialist,
			"addCustomCommandsAfter":function(commands){
				commands.push(new SpecialistDocumentGridCmdSign(id+":grid:cmd:sign"));
			}
		}),
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.SpecialistDocumentlList_Model)? options.detailFilters.SpecialistDocumentList_Model:null,
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:operation_result",{
					"elements":[
						new GridCellHead(id+":grid:head:date_time",{
							"value":"Дата",
							"columns":[
								new GridColumnDate({
									"field":model.getField("date_time")
								})
							],
							"sortable":true,
							"sort":"desc"
						})						
						,new GridCellHead(id+":grid:head:name",{
							"value":"Наименование",
							"columns":[
								new GridColumn({
									"field":model.getField("name")
								})
							],
							"sortable":true
						})						
						
						,(options.detailFilters || options.detail || is_specialist)? null : new GridCellHead(id+":grid:head:specialists_ref",{
							"value":"Мастер",
							"columns":[
								new GridColumnRef({
									"field":model.getField("specialists_ref"),
									"form":Specialist_Form,
									"ctrlClass":SpecialistEdit,
									"ctrlBindFieldId":"specialist_id"
								})
							],
							"sortable":true
						})	
						,new GridCellHead(id+":grid:head:document_att_ref",{
							"value":"Документ",
							"colAttrs":{"title":"Открыть файл","style":"cursor:pointer;"},
							"columns":[
								new GridColumnPicture({
									"field":model.getField("document_att_ref"),
									"onClick":(function(cont){
										return function(fields, elem){
											//alert(fields.specialist_id.getValue());
																						
											//window.getApp().openGridAttachment(cont.getElement("grid").getModel(), tr, "document_att_ref", "specialists", 1);
											var tr = DOMHelper.getParentByTagName(elem, "tr");
											var gridModel = cont.getElement("grid").getModel();
											var previewFieldId = "document_att_ref";
											var attDataType = "specialists";
											var inline = 1;
											var ind = 0;
											var ref_key_val = "specialist_id";
											//!!! ref = specialists. keys->id=SPECIALIST_ID!!!
											var keys = CommonHelper.unserialize(tr.getAttribute("keys"));
											var ind = gridModel.locate(keys,true);
											if(ind == undefined || ind < 0 || !gridModel.getRow(ind)){
												return;
											}
											var content_id;
											var preview = gridModel.getFieldValue(previewFieldId);
											if(preview && preview.length && preview[0].id){
												content_id = preview[0].id;
												
											}else if(preview && preview.id){
												content_id = preview.id;
											}
											var pm = (new SpecialistDocument_Controller()).getPublicMethod("get_file");
											pm.setFieldValue("ref", CommonHelper.serialize(new RefType({"keys":{"id": gridModel.getFieldValue(ref_key_val)}, "dataType": attDataType})));
											pm.setFieldValue("content_id", content_id);
											pm.setFieldValue("inline", inline? "1":"0");
											pm.setFieldValue("doc_id", gridModel.getFieldValue("id"));
											if(inline){
												if(ind==undefined){
													ind = 0;
												}
												var offset = ind * 20;
												var h = $( window ).width()/3*2;
												var left = $( window ).width()/2;
												var w = left - 20;	
												window.getApp().openHrefDownload(window, pm, "ViewXML", "location=0,menubar=0,status=0,titlebar=0,top="+(50+offset)+",left="+(left+offset)+",width="+w+",height="+h);
											}else{
												//download
												pm.download();	
												window.showTempNote("Загрузка файла",null,3000);
											}
											//TODO отметка об открытии и обновить грид
										}
									})(self)
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
		"rowSelect":(window.getWidthType()=="sm"),
		"focus":true
	}));	
	
}
extend(SpecialistDocumentForSignList_View,ViewAjxList);

