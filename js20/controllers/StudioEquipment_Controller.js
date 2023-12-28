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

function StudioEquipment_Controller(options){
	options = options || {};
	options.listModelClass = StudioEquipment_Model;
	options.objModelClass = StudioEquipment_Model;
	StudioEquipment_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
		
}
extend(StudioEquipment_Controller,ControllerObjClient);

			StudioEquipment_Controller.prototype.addInsert = function(){
	StudioEquipment_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Наименование";
	var field = new FieldText("name",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Количество";
	var field = new FieldInt("quant",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Единица";
	var field = new FieldString("measure_unit",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Цена";
	var field = new FieldFloat("price",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Сумма";
	var field = new FieldFloat("total",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			StudioEquipment_Controller.prototype.addUpdate = function(){
	StudioEquipment_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	options.alias = "Наименование";
	var field = new FieldText("name",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Количество";
	var field = new FieldInt("quant",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Единица";
	var field = new FieldString("measure_unit",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Цена";
	var field = new FieldFloat("price",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Сумма";
	var field = new FieldFloat("total",options);
	
	pm.addField(field);
	
	
}

			StudioEquipment_Controller.prototype.addDelete = function(){
	StudioEquipment_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			StudioEquipment_Controller.prototype.addGetList = function(){
	StudioEquipment_Controller.superclass.addGetList.call(this);
	
	
	
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
	f_opts.alias = "Наименование";
	pm.addField(new FieldText("name",f_opts));
	var f_opts = {};
	f_opts.alias = "Количество";
	pm.addField(new FieldInt("quant",f_opts));
	var f_opts = {};
	f_opts.alias = "Единица";
	pm.addField(new FieldString("measure_unit",f_opts));
	var f_opts = {};
	f_opts.alias = "Цена";
	pm.addField(new FieldFloat("price",f_opts));
	var f_opts = {};
	f_opts.alias = "Сумма";
	pm.addField(new FieldFloat("total",f_opts));
}

			StudioEquipment_Controller.prototype.addGetObject = function(){
	StudioEquipment_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

		