/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 *
 */
function SpecialistRegList_View(id,options){	

	options = options || {};
	options.HEAD_TITLE = "Регистрация самозанятых";

	SpecialistRegList_View.superclass.constructor.call(this,id,options);
	
	var model = (options.models && options.models.SpecialistRegList_Model)? options.models. SpecialistRegList_Model : new SpecialistRegList_Model();
	var contr = new SpecialistReg_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);
	
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	
	var self = this;
	this.addElement(new GridAjx(id+":grid",{
		"model":model,
		"keyIds":["id"],
		"controller":contr,
		"editInline":true,
		"editWinClass":SpecialistReg_Form,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"cmdInsert":false,
			"cmdDelete":true,
			"cmdEdit":false,
			"addCustomCommandsAfter":function(commands){
				commands.push(new SpecialistRegGridCmdRegister(id+":grid:cmd:register"));
				commands.push(new SpecialistRegGridCmdPrintApp(id+":grid:cmd:printApp"));
			}
		}),	
		"onEventSetCellOptions": function(opts){
			var m = this.getModel();
			var col_id = opts.gridColumn.getId();
			//tel
			if(col_id == "tel" && m.getFieldValue("tel_confirmed")){
				opts.className = opts.className || "";
				opts.className += (opts.className=="")? "":" ";
				opts.className += "reg_field_confirmed";
				opts.title = "Номер подтвержден.";
				
			}else if(col_id == "tel" && m.getFieldValue("tel_sent")){
				opts.className = opts.className || "";
				opts.className += (opts.className=="")? "":" ";
				opts.className += "reg_field_sent";
				opts.title = "Отправлено сообщение, не подтверждено";
				
			}else if(col_id == "tel" ){
				opts.className = opts.className || "";
				opts.className += (opts.className=="")? "":" ";
				opts.className += "reg_field_error";
				opts.title = "Ошибка отправки сообщения";
				
			}
			
			//email
			if(col_id == "email" && m.getFieldValue("email_confirmed")){
				opts.className = opts.className || "";
				opts.className += (opts.className=="")? "":" ";
				opts.className += "reg_field_confirmed";
				opts.title = "Адрес подтвержден.";
				
			}else if(col_id == "email" && m.getFieldValue("email_sent")){
				opts.className = opts.className || "";
				opts.className += (opts.className=="")? "":" ";
				opts.className += "reg_field_sent";
				opts.title = "Отправлено сообщение, не подтверждено";
				
			}else if(col_id == "email" ){
				opts.className = opts.className || "";
				opts.className += (opts.className=="")? "":" ";
				opts.className += "reg_field_error";
				opts.title = "Ошибка отправки сообщения";				
			}
			
			//inn
			if(col_id == "inn" && m.getFieldValue("inn_fns_ok")){
				opts.className = opts.className || "";
				opts.className += (opts.className=="")? "":" ";
				opts.className += "reg_field_confirmed";
				opts.title = "Проверен по данным ФНС.";
				
			}else if(col_id == "inn" && m.getFieldValue("inn_checked")){
				opts.className = opts.className || "";
				opts.className += (opts.className=="")? "":" ";
				opts.className += "reg_field_sent";
				opts.title = "Не проверен ФНС";
				
			}else if(col_id == "inn" ){
				opts.className = opts.className || "";
				opts.className += (opts.className=="")? "":" ";
				opts.className += "reg_field_error";
				opts.title = "Ошибка проверки ФНС";				
			}
			
		},		
		"popUpMenu":popup_menu,
		"filters":null,
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
						,new GridCellHead(id+":grid:head:name_full",{
							"value":"ФИО",
							"columns":[
								new GridColumn({
									"field":model.getField("name_full")
								})
							],
							"sortable":true
						})
						,new GridCellHead(id+":grid:head:inn",{
							"value":"ИНН",
							"columns":[
								new GridColumn({
									"field":model.getField("inn")
								})
							],
							"sortable":true
						})
						,new GridCellHead(id+":grid:head:tel",{
							"value":"Телефон",
							"columns":[
								new GridColumnPhone({
									"field":model.getField("tel")
								})
							],
							"sortable":true
						})
						,new GridCellHead(id+":grid:head:email",{
							"value":"Электронная почта",
							"columns":[
								new GridColumnEmail({
									"field":model.getField("email")
								})
							],
							"sortable":true
						})
						,new GridCellHead(id+":grid:head:birthdate",{
							"value":"Дата рождения",
							"columns":[
								new GridColumnDate({
									"field":model.getField("birthdate")
								})
							],
							"sortable":true
						})
						,new GridCellHead(id+":grid:head:banks_ref",{
							"value":"Банк",
							"columns":[
								new GridColumnRef({
									"field":model.getField("banks_ref"),
									"ctrlClass":BankEditRef,
									"ctrlOptions":{
										"lableCaption":""
									},
									"ctrlBindFieldId":"banks_ref"
								})
							]
						})
						,new GridCellHead(id+":grid:head:bank_acc",{
							"value":"Счет",
							"columns":[
								new GridColumn({
									"field":model.getField("bank_acc")
								})
							]
						})
						
						,new GridCellHead(id+":grid:head:passport_preview",{
							"value":"Паспорт",
							"colAttrs":{"title":"Открыть файл","style":"cursor:pointer;"},
							"columns":[
								new GridColumnPicture({
									"field":model.getField("passport_preview"),
									"onClick":(function(cont){
										return function(fields, elem){
											var tr = DOMHelper.getParentByTagName(elem, "tr");
											window.getApp().openGridAttachment(cont.getElement("grid").getModel(), tr, "passport_preview", "specialist_regs", 1);
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
		
		"autoRefresh":false,
		"refreshInterval":constants.grid_refresh_interval.getValue()*1000,
		"rowSelect":false,
		"focus":true
	}));	
	


}
extend(SpecialistRegList_View,ViewAjxList);


