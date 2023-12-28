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

function Specialist_Controller(options){
	options = options || {};
	options.listModelClass = SpecialistList_Model;
	options.objModelClass = SpecialistDialog_Model;
	Specialist_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
	this.add_register();
	this.addComplete();
	this.add_send_reg_documents();
	this.add_send_salary_documents();
		
}
extend(Specialist_Controller,ControllerObjServer);

			Specialist_Controller.prototype.addInsert = function(){
	Specialist_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
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
	options.alias = "БИК банка";
	var field = new FieldString("bank_bik",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Банковский счет";
	var field = new FieldString("bank_acc",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
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
	options.alias = "Список оборудования для одного рабочего места";
	var field = new FieldJSONB("equipments",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("user_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("ycl_staff_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldFloat("agent_percent",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("speciality_id",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			Specialist_Controller.prototype.addUpdate = function(){
	Specialist_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("name",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldString("inn",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "БИК банка";
	var field = new FieldString("bank_bik",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Банковский счет";
	var field = new FieldString("bank_acc",options);
	
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
	options.alias = "Список оборудования для одного рабочего места";
	var field = new FieldJSONB("equipments",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("user_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("ycl_staff_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldFloat("agent_percent",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("speciality_id",options);
	
	pm.addField(field);
	
	
}

			Specialist_Controller.prototype.addDelete = function(){
	Specialist_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			Specialist_Controller.prototype.addGetList = function(){
	Specialist_Controller.superclass.addGetList.call(this);
	
	
	
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
	
	pm.addField(new FieldText("name",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("inn",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("tel",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("email",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("banks_ref",f_opts));
	var f_opts = {};
	f_opts.alias = "Банковский счет";
	pm.addField(new FieldString("bank_acc",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldInt("studio_id",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("studios_ref",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("photo",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("specialities_ref",f_opts));
}

			Specialist_Controller.prototype.addGetObject = function(){
	Specialist_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

			Specialist_Controller.prototype.add_register = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('register',opts);
	
	pm.setRequestType('post');
	
	pm.setEncType(ServConnector.prototype.ENCTYPES.MULTIPART);
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "36";
	
		pm.addField(new FieldString("operation_id",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "12";
	
		pm.addField(new FieldString("inn",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldJSON("banks_ref",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "20";
	
		pm.addField(new FieldString("bank_acc",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "200";
	
		pm.addField(new FieldString("name_full",options));
	
				
	
	var options = {};
	
		options.maxlength = "100";
	
		pm.addField(new FieldString("name",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "50";
	
		pm.addField(new FieldString("email",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "11";
	
		pm.addField(new FieldString("tel",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldDate("birthdate",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "500";
	
		pm.addField(new FieldString("address_reg",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "10";
	
		pm.addField(new FieldString("captcha",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldJSON("passport",options));
	
				
	
	var options = {};
	
		pm.addField(new FieldText("passport_file_content",options));
	
			
	this.addPublicMethod(pm);
}

			Specialist_Controller.prototype.addComplete = function(){
	Specialist_Controller.superclass.addComplete.call(this);
	
	var f_opts = {};
	
	var pm = this.getComplete();
	pm.addField(new FieldText("name",f_opts));
	pm.addField(new FieldInt("count", {}));
	pm.getField(this.PARAM_ORD_FIELDS).setValue("name");	
}

			Specialist_Controller.prototype.add_send_reg_documents = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('send_reg_documents',opts);
	
				
	
	var options = {};
	
		options.alias = "specialist ID";
	
		options.required = true;
	
		pm.addField(new FieldInt("id",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("operation_id",options));
	
			
	this.addPublicMethod(pm);
}

			Specialist_Controller.prototype.add_send_salary_documents = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('send_salary_documents',opts);
	
				
	
	var options = {};
	
		options.alias = "period salary detail ID";
	
		options.required = true;
	
		pm.addField(new FieldInt("id",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("operation_id",options));
	
			
	this.addPublicMethod(pm);
}

		