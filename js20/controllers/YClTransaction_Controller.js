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

function YClTransaction_Controller(options){
	options = options || {};
	options.listModelClass = YClTransactionList_Model;
	options.objModelClass = YClTransactionList_Model;
	YClTransaction_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
	this.add_get_for_work_list();
	this.add_get_doc_all_list();
	this.add_api_get();
		
}
extend(YClTransaction_Controller,ControllerObjServer);

			YClTransaction_Controller.prototype.addInsert = function(){
	YClTransaction_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldDateTimeTZ("created_at",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldJSONB("data",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("specialist_id",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			YClTransaction_Controller.prototype.addUpdate = function(){
	YClTransaction_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDateTimeTZ("created_at",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldJSONB("data",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("specialist_id",options);
	
	pm.addField(field);
	
	
}

			YClTransaction_Controller.prototype.addDelete = function(){
	YClTransaction_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			YClTransaction_Controller.prototype.addGetList = function(){
	YClTransaction_Controller.superclass.addGetList.call(this);
	
	
	
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
	
	pm.addField(new FieldInt("document_id",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldDateTimeTZ("date",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldFloat("amount",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("client",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("client_phone",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("specialists_ref",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("ycl_staff_ref",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldInt("specialist_id",f_opts));
}

			YClTransaction_Controller.prototype.addGetObject = function(){
	YClTransaction_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

			YClTransaction_Controller.prototype.add_get_for_work_list = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('get_for_work_list',opts);
	
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

			YClTransaction_Controller.prototype.add_get_doc_all_list = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('get_doc_all_list',opts);
	
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

			YClTransaction_Controller.prototype.add_api_get = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('api_get',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "36";
	
		pm.addField(new FieldString("operation_id",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldDate("date_from",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldDate("date_to",options));
	
			
	this.addPublicMethod(pm);
}

		