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

function DocumentTemplate_Controller(options){
	options = options || {};
	options.listModelClass = DocumentTemplateList_Model;
	options.objModelClass = DocumentTemplateDialog_Model;
	DocumentTemplate_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
		
}
extend(DocumentTemplate_Controller,ControllerObjServer);

			DocumentTemplate_Controller.prototype.addInsert = function(){
	DocumentTemplate_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.alias = "Код";options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Описание шаблона";
	var field = new FieldText("name",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("sql_query",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldJSONB("fields",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("need_signing",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("sign_image_name",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			DocumentTemplate_Controller.prototype.addUpdate = function(){
	DocumentTemplate_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.alias = "Код";options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	options.alias = "Описание шаблона";
	var field = new FieldText("name",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("sql_query",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldJSONB("fields",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("need_signing",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("sign_image_name",options);
	
	pm.addField(field);
	
	
}

			DocumentTemplate_Controller.prototype.addDelete = function(){
	DocumentTemplate_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
	options.alias = "Код";	
	pm.addField(new FieldInt("id",options));
}

			DocumentTemplate_Controller.prototype.addGetList = function(){
	DocumentTemplate_Controller.superclass.addGetList.call(this);
	
	
	
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
	
	pm.addField(new FieldString("name",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("file_preview",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldBool("need_signing",f_opts));
}

			DocumentTemplate_Controller.prototype.addGetObject = function(){
	DocumentTemplate_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
	f_opts.alias = "Код";	
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

		