/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023

 * @extends ViewObjectAjx
 * @requires js20/core/extend.js
 * @requires js20/controls/ViewObjectAjx.js     

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {Object} options
 */
function NotifTemplate_View(id,options){	

	options = options || {};
	
	options.model = (options && options.models && options.models.NotifTemplate_Model)? options.models.NotifTemplate_Model : new NotifTemplate_Model();
	options.controller = options.controller || new NotifTemplate_Controller();
	
	var self = this;
	
	var is_admin = (window.getApp().getServVar("role_id")=="admin");
	
	options.addElement = function(){
		this.addElement(new EditString(id+":notif_provider",{
			"labelCaption":"Провайдер:",
			"maxLength":"20",
			"required":true
		}));	

		
		this.addElement(new Enum_notif_types(id+":notif_type",{			
			"labelCaption":"Тип сообщения:",
			"enabled":is_admin,
			"required":true
		}));	
	
		this.addElement(new EditText(id+":template",{
			"labelCaption":"Шаблон:",
			"required":true,
			"focus":true
		}));	

		this.addElement(new EditText(id+":comment_text",{
			"labelCaption": "Комментарий:",
			"required":true,
			"enabled":is_admin
		}));	

		//********* fields grid ***********************
		var model = new ReportTemplateField_Model();
			
		this.addElement(new GridAjx(id+":fields",{
			"model":model,
			"keyIds":["id"],
			"controller":new ReportTemplateField_Controller({"clientModel":model}),
			"editInline":true,
			"editWinClass":null,
			"popUpMenu":new PopUpMenu(),
			"commands":new GridCmdContainerAjx(id+":fields:cmd",{
				"cmdSearch":false,
				"cmdExport":false
			}),
			"head":new GridHead(id+":fields:head",{
				"elements":[
					new GridRow(id+":fields:head:row0",{
						"elements":[
							new GridCellHead(id+":fields:head:id",{
								"value":"Идентификатор",
								"columns":[
									new GridColumn({
										"field":model.getField("id"),
										"ctrlClass":EditString,
										"maxLength":"50",
										"ctrlOptions":{
											"required":true
										}
																										
									})
								]
							}),					
							new GridCellHead(id+":fields:head:descr",{
								"value":"Описание",
								"columns":[
									new GridColumn({
										"field":model.getField("descr"),
										"ctrlClass":EditString,
										"maxLength":"250",
										"ctrlOptions":{
											"required":true
										}
																										
									})
								]
							})				
						]
					})
				]
			}),
			"pagination":null,				
			"autoRefresh":false,
			"refreshInterval":0,
			"rowSelect":true,
			"focus":true		
		}));
		
		var model_pv = new ReportTemplateField_Model();
		this.addElement(new GridAjx(id+":provider_values",{
			"model":model_pv,
			"keyIds":["id"],
			"controller":new ReportTemplateField_Controller({"clientModel":model_pv}),
			"editInline":true,
			"editWinClass":null,
			"popUpMenu":new PopUpMenu(),
			"commands":new GridCmdContainerAjx(id+":provider_values:cmd",{
				"cmdSearch":false,
				"cmdExport":false
			}),
			"head":new GridHead(id+":provider_values:head",{
				"elements":[
					new GridRow(id+":provider_values:head:row0",{
						"elements":[
							new GridCellHead(id+":provider_values:head:id",{
								"value":"Идентификатор",
								"columns":[
									new GridColumn({
										"field":model_pv.getField("id"),
										"ctrlClass":EditString,
										"maxLength":"50",
										"ctrlOptions":{
											"required":true
										}
																										
									})
								]
							}),					
							new GridCellHead(id+":provider_values:head:descr",{
								"value":"Значение",
								"columns":[
									new GridColumn({
										"field":model_pv.getField("descr"),
										"ctrlClass":EditString,
										"maxLength":"250",
										"ctrlOptions":{
											"required":true
										}
																										
									})
								]
							})				
						]
					})
				]
			}),
			"pagination":null,				
			"autoRefresh":false,
			"refreshInterval":0,
			"rowSelect":true,
			"focus":true		
		}));
		
	}
	
	
	NotifTemplate_View.superclass.constructor.call(this,id,options);
	
	
	//****************************************************	
	
	//read
	var read_b = [
		new DataBinding({"control":this.getElement("notif_type"),"fieldId":"notif_type"})
		,new DataBinding({"control":this.getElement("notif_provider")})
		,new DataBinding({"control":this.getElement("template")})
		,new DataBinding({"control":this.getElement("comment_text")})
		,new DataBinding({"control":this.getElement("fields"),"fieldId":"fields"})
		,new DataBinding({"control":this.getElement("provider_values"),"fieldId":"provider_values"})
	];
	this.setDataBindings(read_b);
	
	//write
	var write_b = [
		new CommandBinding({"control":this.getElement("notif_type"),"fieldId":"notif_type"})
		,new CommandBinding({"control":this.getElement("notif_provider")})
		,new CommandBinding({"control":this.getElement("template")})
		,new CommandBinding({"control":this.getElement("comment_text")})
		,new CommandBinding({"control":this.getElement("fields"),"fieldId":"fields"})
		,new CommandBinding({"control":this.getElement("provider_values"),"fieldId":"provider_values"})
	];
	this.setWriteBindings(write_b);
	
}
extend(NotifTemplate_View,ViewObjectAjx);

