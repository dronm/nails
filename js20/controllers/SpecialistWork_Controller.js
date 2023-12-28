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

function SpecialistWork_Controller(options){
	options = options || {};
	options.listModelClass = SpecialistWorkList_Model;
	options.objModelClass = SpecialistWorkList_Model;
	SpecialistWork_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
	this.add_get_for_rate_list();
	this.add_set_admin_rates();
		
}
extend(SpecialistWork_Controller,ControllerObjServer);

			SpecialistWork_Controller.prototype.addInsert = function(){
	SpecialistWork_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("specialist_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("studio_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDateTimeTZ("date_time",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("admin_rate",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("ycl_document_id",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			SpecialistWork_Controller.prototype.addUpdate = function(){
	SpecialistWork_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("specialist_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("studio_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDateTimeTZ("date_time",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("admin_rate",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("ycl_document_id",options);
	
	pm.addField(field);
	
	
}

			SpecialistWork_Controller.prototype.addDelete = function(){
	SpecialistWork_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			SpecialistWork_Controller.prototype.addGetList = function(){
	SpecialistWork_Controller.superclass.addGetList.call(this);
	
	
	
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
	
	pm.addField(new FieldInt("specialist_id",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("specialists_ref",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldInt("studio_id",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("studios_ref",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldDateTimeTZ("date_time",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("photo",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldInt("admin_rate",f_opts));
}

			SpecialistWork_Controller.prototype.addGetObject = function(){
	SpecialistWork_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

			SpecialistWork_Controller.prototype.add_get_for_rate_list = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('get_for_rate_list',opts);
	
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

	this.addPublicMethod(pm);
}

			SpecialistWork_Controller.prototype.add_set_admin_rates = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('set_admin_rates',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldJSON("rates",options));
	
			
	this.addPublicMethod(pm);
}

		