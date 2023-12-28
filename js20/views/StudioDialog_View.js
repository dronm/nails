/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023

 * @extends ViewObjectAjx
 * @requires core/extend.js  

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {Object} options
 */
function StudioDialog_View(id,options){	

	options = options || {};
	
	options.model = (options.models&&options.models.StudioDialog_Model)? options.models.StudioDialog_Model : new StudioDialog_Model();
	options.controller = options.controller || new Studio_Controller();
	
	var self = this;
	options.addElement = function(){
		this.addElement(new EditString(id+":name",{
			"labelCaption":"Наименование:",
			"required":true,
			"focus":true,
			"maxLength":"500",
			"placeholder":"Наименование салона",
			"title":"Наименование салона"
		}));	

		this.addElement(new EditMoney(id+":hour_rent_price",{
			"labelCaption":"Цена часа аренды:",
			"required":true,
			"title":"Цена часа аренды рабочего места"
		}));	

		this.addElement(new FirmEdit(id+":firms_ref",{
		}));
	}
	
	StudioDialog_View.superclass.constructor.call(this,id,options);
	
	
	//****************************************************	
	
	//read
	var read_b = [
		new DataBinding({"control":this.getElement("name")})
		,new DataBinding({"control":this.getElement("hour_rent_price")})
		,new DataBinding({"control":this.getElement("firms_ref"), "fieldId":"firm_id"})				
	];
	this.setDataBindings(read_b);
	
	//write
	var write_b = [
		new CommandBinding({"control":this.getElement("name")})
		,new CommandBinding({"control":this.getElement("hour_rent_price")})
		,new CommandBinding({"control":this.getElement("firms_ref"), "fieldId":"firm_id"})		
	];
	this.setWriteBindings(write_b);
	
}
extend(StudioDialog_View,ViewObjectAjx);

