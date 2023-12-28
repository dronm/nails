/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023
 
 * @extends ViewObjectAjx.js
 * @requires core/extend.js  
 * @requires controls/ViewObjectAjx.js 
 
 * @class
 * @classdesc
	
 * @param {string} id view identifier
 * @param {object} options
 */	
function SpecialityDialog_View(id,options){	

	options = options || {};
	
	options.controller = new Speciality_Controller();
	options.model = (options.models&&options.models.SpecialityDialog_Model)? options.models.SpecialityDialog_Model : new SpecialityDialog_Model();
	
	options.addElement = function(){
		this.addElement(new EditString(id+":name",{
			"required": true,
			"focus": true,
			"maxLength":500,
			"labelCaption": "Наименование:",
			"placeholder":"Наименование специальности",
			"title":"Наименование специальности"
		}));	

		this.addElement(new EditFloat(id+":agent_percent",{
			"required": true,
			"precision":"2",
			"labelCaption": "Агентское вознаграждение:",
			"title":"Процент агентского вознаграждения"
		}));	

		this.addElement(new SpecialityEquipmentGrid(id+":equipments",{
		}));

	}
	
	SpecialityDialog_View.superclass.constructor.call(this,id,options);
	
	//****************************************************
	//read
	this.setDataBindings([
		new DataBinding({"control":this.getElement("name")})
		,new DataBinding({"control":this.getElement("agent_percent")})
		,new DataBinding({"control":this.getElement("equipments")})
	]);
	
	//write
	this.setWriteBindings([
		new CommandBinding({"control":this.getElement("name")})
		,new CommandBinding({"control":this.getElement("agent_percent")})
		,new CommandBinding({"control":this.getElement("equipments"), "fieldId":"equipments"})
	]);
	
}
extend(SpecialityDialog_View, ViewObjectAjx);
