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

function SpecialistReceipt_Controller(options){
	options = options || {};
	options.listModelClass = SpecialistReceipt_Model;
	options.objModelClass = SpecialistReceipt_Model;
	SpecialistReceipt_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
	this.add_add();
		
}
extend(SpecialistReceipt_Controller,ControllerObjServer);

			SpecialistReceipt_Controller.prototype.addInsert = function(){
	SpecialistReceipt_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Дата загрузки";
	var field = new FieldDateTimeTZ("date_time",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Дата документа";
	var field = new FieldDateTimeTZ("document_date_time",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Сумма документа";
	var field = new FieldFloat("document_total",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Сумма проверен ФНС";
	var field = new FieldBool("document_parsed",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("specialist_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("specialist_period_salary_detail_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Ошибка проверки сека ФНС";
	var field = new FieldText("document_error",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "ИД сервиса qrextr";
	var field = new FieldText("qrextr_request_id",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			SpecialistReceipt_Controller.prototype.addUpdate = function(){
	SpecialistReceipt_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	options.alias = "Дата загрузки";
	var field = new FieldDateTimeTZ("date_time",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Дата документа";
	var field = new FieldDateTimeTZ("document_date_time",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Сумма документа";
	var field = new FieldFloat("document_total",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Сумма проверен ФНС";
	var field = new FieldBool("document_parsed",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("specialist_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("specialist_period_salary_detail_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Ошибка проверки сека ФНС";
	var field = new FieldText("document_error",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "ИД сервиса qrextr";
	var field = new FieldText("qrextr_request_id",options);
	
	pm.addField(field);
	
	
}

			SpecialistReceipt_Controller.prototype.addDelete = function(){
	SpecialistReceipt_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			SpecialistReceipt_Controller.prototype.addGetList = function(){
	SpecialistReceipt_Controller.superclass.addGetList.call(this);
	
	
	
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
	f_opts.alias = "Дата загрузки";
	pm.addField(new FieldDateTimeTZ("date_time",f_opts));
	var f_opts = {};
	f_opts.alias = "Дата документа";
	pm.addField(new FieldDateTimeTZ("document_date_time",f_opts));
	var f_opts = {};
	f_opts.alias = "Сумма документа";
	pm.addField(new FieldFloat("document_total",f_opts));
	var f_opts = {};
	f_opts.alias = "Сумма проверен ФНС";
	pm.addField(new FieldBool("document_parsed",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldInt("specialist_id",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldInt("specialist_period_salary_detail_id",f_opts));
	var f_opts = {};
	f_opts.alias = "Ошибка проверки сека ФНС";
	pm.addField(new FieldText("document_error",f_opts));
	var f_opts = {};
	f_opts.alias = "ИД сервиса qrextr";
	pm.addField(new FieldText("qrextr_request_id",f_opts));
}

			SpecialistReceipt_Controller.prototype.addGetObject = function(){
	SpecialistReceipt_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

			SpecialistReceipt_Controller.prototype.add_add = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('add',opts);
	
	pm.setRequestType('post');
	
	pm.setEncType(ServConnector.prototype.ENCTYPES.MULTIPART);
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "36";
	
		pm.addField(new FieldString("operation_id",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldInt("specialist_period_salary_detail_id",options));
	
				
	
	var options = {};
	
		pm.addField(new FieldText("file_content",options));
	
			
	this.addPublicMethod(pm);
}

		