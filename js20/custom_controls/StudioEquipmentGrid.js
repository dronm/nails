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
function StudioEquipmentGrid(id,options){
	var model = new StudioEquipment_Model({
		"sequences":{"id":0}
	});

	var self = this;
	var cells = [
		new GridCellHead(id+":head:name",{
			"value":"Наименование",
			"columns":[
				new GridColumn({
					"field":model.getField("name"),
					"ctrlClass":EditString,
					"ctrlOptions":{
						"maxLength":1000
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
						"events":{
							"input":function(e){
								var view = self.getEditViewObj();
								if(!view){
									return;
								}
								view.getElement("total").setValue(view.getElement("price").getValue() * view.getElement("quant").getValue());
							}
						}
					}					
				})
			]
		})
		,new GridCellHead(id+":head:price",{
			"value":"Цена",
			"columns":[
				new GridColumnFloat({
					"field":model.getField("price"),
					"precision":"2",
					"ctrlClass":EditMoney,
					"ctrlOptions":{
						"events":{
							"input":function(e){
								var view = self.getEditViewObj();
								if(!view){
									return;
								}
								view.getElement("total").setValue(view.getElement("price").getValue() * view.getElement("quant").getValue());
							}
						}
					}					
				})
			]
		})
		,new GridCellHead(id+":head:total",{
			"value":"Сумма",
			"columns":[
				new GridColumnFloat({
					"field":model.getField("total"),
					"precision":"2",
					"ctrlClass":EditMoney,
					"ctrlOptions":{
						"events":{
							"input":function(e){
								var view = self.getEditViewObj();
								if(!view){
									return;
								}
								view.getElement("price").setValue(Math.round(view.getElement("total").getValue() / view.getElement("quant").getValue() / 100) * 100);
							}
						}
					}					
				})
			]
		})
		
	];

	options = {
		"model":model,
		"keyIds":["id"],
		"controller":new StudioEquipment_Controller({"clientModel":model}),
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
		"foot":new GridFoot(id+"grid:foot",{
			"autoCalc":true,			
			"elements":[
				new GridRow(id+":grid:foot:row0",{
					"elements":[
						new GridCell(id+":grid:foot:total_sp1",{
							"colSpan":"3"
						})											
						,new GridCellFoot(id+":grid:foot:tot_total",{
							"attrs":{"align":"right"},
							"calcOper":"sum",
							"calcFieldId":"total",
							"gridColumn":new GridColumnFloat({"id":"tot_total"})
						})
					]
				})		
			]
		}),
		
		"pagination":null,				
		"autoRefresh":false,
		"refreshInterval":0,
		"rowSelect":true
	};	
	StudioEquipmentGrid.superclass.constructor.call(this,id,options);
}
extend(StudioEquipmentGrid,GridAjx);

/* Constants */


/* private members */

/* protected*/


/* public methods */
