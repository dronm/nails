/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2017
 
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller_js20.xsl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 
 * @class
 * @classdesc controller
 
 * @extends ControllerObjClient
  
 * @requires core/extend.js
 * @requires core/ControllerObjClient.js
  
 * @param {Object} options
 * @param {Model} options.listModelClass
 * @param {Model} options.objModelClass
 */ 

function SpecialityEquipment_Controller(options){
	options = options || {};
	options.listModelClass = SpecialityEquipment_Model;
	options.objModelClass = SpecialityEquipment_Model;
	SpecialityEquipment_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
		
}
extend(SpecialityEquipment_Controller,ControllerObjClient);

			SpecialityEquipment_Controller.prototype.addInsert = function(){
	SpecialityEquipment_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Вид оборудования";
	var field = new FieldJSONB("equipment_types_ref",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Количество";
	var field = new FieldInt("quant",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Едииница";
	var field = new FieldString("measure_unit",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			SpecialityEquipment_Controller.prototype.addUpdate = function(){
	SpecialityEquipment_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	options.alias = "Вид оборудования";
	var field = new FieldJSONB("equipment_types_ref",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Количество";
	var field = new FieldInt("quant",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Едииница";
	var field = new FieldString("measure_unit",options);
	
	pm.addField(field);
	
	
}

			SpecialityEquipment_Controller.prototype.addDelete = function(){
	SpecialityEquipment_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			SpecialityEquipment_Controller.prototype.addGetList = function(){
	SpecialityEquipment_Controller.superclass.addGetList.call(this);
	
	
	
	var pm = this.getGetList();
	
	pm.addField(new FieldInt(this.PARAM_COUNT));
	pm.addField(new FieldInt(this.PARAM_FROM));
	pm.addField(new FieldString(this.PARAM_COND_FIELDS));
	pm.addField(new FieldString(this.PARAM_COND_SGNS));
	pm.addField(new FieldString(this.PARAM_COND_VALS));
	pm.addField(new FieldString(this.PARAM_COND_JOINS));
	pm.addField(new FieldString(this.PARAM_COND_ICASE));
	pm.addField(new FieldString(this.PARAM_ORD_FIELDS));
	pm.addField(new FieldString(this.PARAM_ORD_DIRECTS));
	pm.addField(new FieldString(this.PARAM_FIELD_SEP));
	pm.addField(new FieldString(this.PARAM_FIELD_LSN));
	pm.addField(new FieldString(this.PARAM_EXP_FNAME));

	var f_opts = {};
	
	pm.addField(new FieldInt("id",f_opts));
	var f_opts = {};
	f_opts.alias = "Вид оборудования";
	pm.addField(new FieldJSONB("equipment_types_ref",f_opts));
	var f_opts = {};
	f_opts.alias = "Количество";
	pm.addField(new FieldInt("quant",f_opts));
	var f_opts = {};
	f_opts.alias = "Едииница";
	pm.addField(new FieldString("measure_unit",f_opts));
}

			SpecialityEquipment_Controller.prototype.addGetObject = function(){
	SpecialityEquipment_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

		