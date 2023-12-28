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

function Login_Controller(options){
	options = options || {};
	options.listModelClass = LoginList_Model;
	options.objModelClass = LoginList_Model;
	Login_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addGetList();
	this.addGetObject();
		
}
extend(Login_Controller,ControllerObjServer);

			Login_Controller.prototype.addGetList = function(){
	Login_Controller.superclass.addGetList.call(this);
	
	
	
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
	f_opts.alias = "Дата входа";
	pm.addField(new FieldDateTimeTZ("date_time_in",f_opts));
	var f_opts = {};
	f_opts.alias = "Дата выхода";
	pm.addField(new FieldDateTimeTZ("date_time_out",f_opts));
	var f_opts = {};
	f_opts.alias = "IP адрес";
	pm.addField(new FieldString("ip",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldInt("user_id",f_opts));
	var f_opts = {};
	f_opts.alias = "Пользователь";
	pm.addField(new FieldJSON("users_ref",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("pub_key",f_opts));
	var f_opts = {};
	f_opts.alias = "Время последнего обращения";
	pm.addField(new FieldDateTimeTZ("set_date_time",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("user_agent_descr",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("user_agent",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("headers",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldDateTimeTZ("sess_create_time",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldDateTimeTZ("accessed_time",f_opts));
}

			Login_Controller.prototype.addGetObject = function(){
	Login_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

		