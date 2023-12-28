/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2017
 
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller_js20.xsl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 
 * @class
 * @classdesc controller
 
 * @extends ControllerObjServer
  
 * @requires core/extend.js
 * @requires core/ControllerObjServer.js
  
 * @param {Object} options
 * @param {Model} options.listModelClass
 * @param {Model} options.objModelClass
 */ 

function TemplateBatchItem_Controller(options){
	options = options || {};
	options.listModelClass = TemplateBatchItemList_Model;
	options.objModelClass = TemplateBatchItemList_Model;
	TemplateBatchItem_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
		
}
extend(TemplateBatchItem_Controller,ControllerObjServer);

			TemplateBatchItem_Controller.prototype.addInsert = function(){
	TemplateBatchItem_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("template_batch_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("template_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Студия или пусто, если все";
	var field = new FieldInt("studio_id",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			TemplateBatchItem_Controller.prototype.addUpdate = function(){
	TemplateBatchItem_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("template_batch_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("template_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Студия или пусто, если все";
	var field = new FieldInt("studio_id",options);
	
	pm.addField(field);
	
	
}

			TemplateBatchItem_Controller.prototype.addDelete = function(){
	TemplateBatchItem_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			TemplateBatchItem_Controller.prototype.addGetList = function(){
	TemplateBatchItem_Controller.superclass.addGetList.call(this);
	
	
	
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
	
	pm.addField(new FieldInt("template_batch_id",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("templates_ref",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("studios_ref",f_opts));
}

			TemplateBatchItem_Controller.prototype.addGetObject = function(){
	TemplateBatchItem_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

		