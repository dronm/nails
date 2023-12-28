/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 *
 */
function SpecialistList_View(id,options){	

	options = options || {};
	options.HEAD_TITLE = "Мастера";

	SpecialistList_View.superclass.constructor.call(this,id,options);
	
	var model = (options.models && options.models. SpecialistList_Model)? options.models. SpecialistList_Model : new SpecialistList_Model();
	var contr = new Specialist_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);
	
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	
	var self = this;
	this.addElement(new GridAjx(id+":grid",{
		"model":model,
		"controller":contr,
		"editInline":false,
		"editWinClass":Specialist_Form,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"cmdInsert":true,
			"cmdDelete":true,
			"cmdEdit":true
			/*"addCustomCommandsAfter":function(commands){
				commands.push(new SpecialistRegGridCmdRegister(id+":grid:cmd:register"));
			}*/
		}),	
		"popUpMenu":popup_menu,
		"filters":null,
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:operation_result",{
					"elements":[
						new GridCellHead(id+":grid:head:name",{
							"value":"ФИО",
							"columns":[
								new GridColumn({
									"field":model.getField("name")
								})
							],
							"sortable":true,
							"sort":"asc"
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
						,new GridCellHead(id+":grid:head:studios_ref",{
							"value":"Салон",
							"columns":[
								new GridColumnRef({
									"field":model.getField("studios_ref"),
									"ctrlClass":StudioEdit,
									"ctrlOptions":{
										"lableCaption":""
									},
									"ctrlBindFieldId":"studio_id"
								})
							]
						})
						,new GridCellHead(id+":grid:head:specialities_ref",{
							"value":"Спец-ть",
							"columns":[
								new GridColumnRef({
									"field":model.getField("specialities_ref"),
									"ctrlClass":SpecialityEdit,
									"ctrlOptions":{
										"lableCaption":""
									},
									"ctrlBindFieldId":"speciality_id"
								})
							]
						})
						
						,new GridCellHead(id+":grid:head:photo",{
							"value":"Фото",
							"colAttrs":{"title":"Открыть файл","style":"cursor:pointer;"},
							"columns":[
								new GridColumnPicture({
									"field":model.getField("photo"),
									"onClick":(function(cont){
										return function(fields, elem){
											var tr = DOMHelper.getParentByTagName(elem, "tr");
											window.getApp().openGridAttachment(cont.getElement("grid").getModel(), tr, "photo", "users", 1);
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
extend(SpecialistList_View,ViewAjxList);


