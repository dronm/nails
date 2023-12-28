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
function SpecialityEquipmentGrid(id,options){
	var model = new SpecialityEquipment_Model({
		"sequences":{"id":0}
	});

	var self = this;
	var cells = [
		new GridCellHead(id+":head:equipment_types_ref",{
			"value":"Наименование",
			"columns":[
				new GridColumnRef({
					"field":model.getField("equipment_types_ref"),
					"ctrlClass":EquipmentTypeEdit,
					"ctrlBindFieldId": "equipment_types_ref",
					"ctrlOptions":{
						"labelCaption":""
					}					
				})
			]
		})
		,new GridCellHead(id+":head:measure_unit",{
			"value":"Ед-ца",
			"columns":[
				new GridColumn({
					"field":model.getField("measure_unit"),
					"ctrlClass":EditString,
					"ctrlOptions":{
					}					
				})
			]
		})
		,new GridCellHead(id+":head:quant",{
			"value":"Кол-во",
			"columns":[
				new GridColumn({
					"field":model.getField("quant"),
					"ctrlClass":EditNum,
					"ctrlOptions":{
					}					
				})
			]
		})
		
	];

	options = {
		"model":model,
		"keyIds":["id"],
		"controller":new SpecialityEquipment_Controller({"clientModel":model}),
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
	SpecialityEquipmentGrid.superclass.constructor.call(this,id,options);
}
extend(SpecialityEquipmentGrid,GridAjx);

/* Constants */


/* private members */

/* protected*/


/* public methods */
