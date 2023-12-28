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
function NotifTemplateList_View(id,options){
	options = options || {};	
	
	options.HEAD_TITLE = "Шаблоны сообщений";
	NotifTemplateList_View.superclass.constructor.call(this,id,options);
	
	var model = (options && options.models && options.models.NotifTemplateList_Model)? options.models.NotifTemplateList_Model : new NotifTemplateList_Model();
	var contr = new NotifTemplate_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);
	
	var pagClass = window.getApp().getPaginationClass();
	
	var popup_menu = new PopUpMenu();
	
	var is_admin = (window.getApp().getServVar("role_id")=="admin");
	
	this.addElement(new GridAjx(id+":grid",{
		"model":model,
		"keyIds":["id"],
		"controller":contr,
		"editInline":false,
		"editWinClass":NotifTemplate_Form,
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
						new GridCellHead(id+":grid:head:notif_provider",{
							"value":"Тип доставки",
							"columns":[
								new GridColumn({"field":model.getField("notif_provider")})
							]
						}),
						
						new GridCellHead(id+":grid:head:notif_type",{
							"value":"Тип сообщения",
							"columns":[
								new EnumGridColumn_notif_types({"field":model.getField("notif_type")})
							]
						}),
					
						new GridCellHead(id+":grid:head:template",{
							"value":"Шаблон",
							"columns":[
								new GridColumn({"field":model.getField("template")})
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
extend(NotifTemplateList_View,ViewAjx);

/* Constants */


/* private members */

/* protected*/


/* public methods */

