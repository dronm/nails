/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023

 * @extends ViewAjx
 * @requires js20/core/extend.js
 * @requires js20/controls/ViewAjx.js     

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {namespace} options
 */
function DocumentTemplateList_View(id,options){
	options = options || {};	
	
	options.HEAD_TITLE = "Шаблоны документов";
	DocumentTemplateList_View.superclass.constructor.call(this,id,options);
	
	var model = (options && options.models && options.models.DocumentTemplateList_Model)? options.models.DocumentTemplateList_Model : new DocumentTemplateList_Model();
	var contr = new DocumentTemplate_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);
	
	var pagClass = window.getApp().getPaginationClass();
	
	var popup_menu = new PopUpMenu();
	
	var is_admin = (window.getApp().getServVar("role_id")=="admin");
	
	var self = this;
	this.addElement(new GridAjx(id+":grid",{
		"model":model,
		"keyIds":["id"],
		"controller":contr,
		"editInline":false,
		"editWinClass":DocumentTemplate_Form,
		"popUpMenu":popup_menu,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"cmdInsert":is_admin,
			"cmdCopy":is_admin,
			"cmdDelete":is_admin,
			"cmdEdit":is_admin,
			"exportFileName" :"Шаблоны"
		}),
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:row0",{
					"elements":[
						new GridCellHead(id+":grid:head:name",{
							"value":"Наименование",
							"columns":[
								new GridColumn({"field":model.getField("name")})
							]
						}),
						
						new GridCellHead(id+":grid:head:file_preview",{
							"value":"Файл",
							"colAttrs":{"title":"Открыть файл","style":"cursor:pointer;"},
							"columns":[
								new GridColumnPicture({
									"field":model.getField("file_preview"),
									"onClick":(function(cont){
										return function(fields, elem){
											var tr = DOMHelper.getParentByTagName(elem, "tr");
											window.getApp().openGridAttachment(cont.getElement("grid").getModel(), tr, "file_preview", "document_templates", 0);
										}
									})(self)
									
								})
							]
						})
						,new GridCellHead(id+":grid:head:need_signing",{
							"value":"Требовать подпись",
							"columns":[
								new GridColumnBool({
									"field":model.getField("need_signing")
								})
							]
						}),
						
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
extend(DocumentTemplateList_View,ViewAjx);

/* Constants */


/* private members */

/* protected*/


/* public methods */

