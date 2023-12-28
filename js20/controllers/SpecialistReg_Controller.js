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

function SpecialistReg_Controller(options){
	options = options || {};
	options.listModelClass = SpecialistRegList_Model;
	options.objModelClass = SpecialistRegDialog_Model;
	SpecialistReg_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
	this.add_get_by_operation_id();
	this.add_register();
	this.add_print_app();
		
}
extend(SpecialistReg_Controller,ControllerObjServer);

			SpecialistReg_Controller.prototype.addInsert = function(){
	SpecialistReg_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldString("user_operation_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldText("name_full",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldText("name",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldString("inn",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldString("email",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("email_sent",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldString("tel",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("tel_sent",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldJSONB("banks_ref",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("studio_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Дата рождения";
	var field = new FieldDate("birthdate",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Адрес регистрации";
	var field = new FieldText("address_reg",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Паспорт";
	var field = new FieldJSON("passport",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("passport_uploaded",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldDateTimeTZ("date_time",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldBool("inn_checked",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldBool("inn_fns_ok",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			SpecialistReg_Controller.prototype.addUpdate = function(){
	SpecialistReg_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldString("user_operation_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("name_full",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("name",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldString("inn",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldString("email",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("email_sent",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldString("tel",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("tel_sent",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldJSONB("banks_ref",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("studio_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Дата рождения";
	var field = new FieldDate("birthdate",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Адрес регистрации";
	var field = new FieldText("address_reg",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Паспорт";
	var field = new FieldJSON("passport",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("passport_uploaded",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDateTimeTZ("date_time",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("inn_checked",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("inn_fns_ok",options);
	
	pm.addField(field);
	
	
}

			SpecialistReg_Controller.prototype.addDelete = function(){
	SpecialistReg_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			SpecialistReg_Controller.prototype.addGetList = function(){
	SpecialistReg_Controller.superclass.addGetList.call(this);
	
	
	
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
	
	pm.addField(new FieldString("inn",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("name_full",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("tel",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldBool("tel_confirmed",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldBool("tel_sent",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("email",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldBool("email_confirmed",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldBool("email_sent",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldDateTimeTZ("date_time",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldDate("birthdate",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldBool("inn_checked",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldBool("inn_fns_ok",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("banks_ref",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("bank_acc",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSONB("passport_preview",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldBool("operation_result",f_opts));
}

			SpecialistReg_Controller.prototype.addGetObject = function(){
	SpecialistReg_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

			SpecialistReg_Controller.prototype.add_get_by_operation_id = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('get_by_operation_id',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "36";
	
		pm.addField(new FieldString("operation_id",options));
	
			
	this.addPublicMethod(pm);
}

			SpecialistReg_Controller.prototype.add_register = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('register',opts);
	
	pm.setRequestType('post');
	
	pm.setEncType(ServConnector.prototype.ENCTYPES.MULTIPART);
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldInt("id",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldInt("studio_id",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldInt("ycl_staff_id",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "36";
	
		pm.addField(new FieldString("operation_id",options));
	
			
	this.addPublicMethod(pm);
}

			SpecialistReg_Controller.prototype.add_print_app = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('print_app',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldInt("id",options));
	
			
	this.addPublicMethod(pm);
}

		