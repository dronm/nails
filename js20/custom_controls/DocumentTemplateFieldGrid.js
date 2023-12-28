/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023

 * @extends GridAjx
 * @requires core/extend.js  

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {Object} options
 * @param {string} options.className
 */
function DocumentTemplateFieldGrid(id,options){
	var model = new DocumentTemplateField_Model({
		"sequences":{"id":0}
	});

	var cells = [
		new GridCellHead(id+":head:field",{
			"value":"Поле запроса",
			"columns":[
				new GridColumn({
					"field":model.getField("field"),
					"ctrlClass":EditString,
					"ctrlOptions":{
						"maxLength":100
					}					
				})
			]
		})
		,new GridCellHead(id+":head:comment_text",{
			"value":"Описание",
			"columns":[
				new GridColumn({
					"field":model.getField("comment_text"),
					"ctrlClass":EditString,
					"ctrlOptions":{
						"maxLength":500
					}					
				})
			]
		})
		
	];

	options = {
		"model":model,
		"keyIds":["id"],
		"controller":new DocumentTemplateField_Controller({"clientModel":model}),
		"editInline":true,
		"editWinClass":null,
		"popUpMenu":new PopUpMenu(),
		"commands":new GridCmdContainerAjx(id+":cmd",{
			"cmdSearch":false,
			"cmdExport":false
		}),
		"head":new GridHead(id+":head",{
			"elements":[
				new GridRow(id+":head:row0",{
					"elements":cells
				})
			]
		}),
		"pagination":null,				
		"autoRefresh":false,
		"refreshInterval":0,
		"rowSelect":true
	};	
	DocumentTemplateFieldGrid.superclass.constructor.call(this,id,options);
}
extend(DocumentTemplateFieldGrid,GridAjx);

/* Constants */


/* private members */

/* protected*/


/* public methods */
