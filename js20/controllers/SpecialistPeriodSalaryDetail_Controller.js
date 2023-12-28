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

function SpecialistPeriodSalaryDetail_Controller(options){
	options = options || {};
	options.listModelClass = SpecialistPeriodSalaryDetailList_Model;
	options.objModelClass = SpecialistPeriodSalaryDetailList_Model;
	SpecialistPeriodSalaryDetail_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
	this.add_get_for_pay_list();
	this.addComplete();
	this.add_documents_to_bank();
		
}
extend(SpecialistPeriodSalaryDetail_Controller,ControllerObjServer);

			SpecialistPeriodSalaryDetail_Controller.prototype.addInsert = function(){
	SpecialistPeriodSalaryDetail_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("specialist_period_salary_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Номер";
	var field = new FieldInt("line_num",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("specialist_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("studio_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDate("period",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Часов";
	var field = new FieldInt("hours",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldFloat("agent_percent",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Выручка";
	var field = new FieldFloat("work_total",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Выручка для зп";
	var field = new FieldFloat("work_total_salary",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Доп.начисления";
	var field = new FieldFloat("debet",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Цена за аренду";
	var field = new FieldFloat("kredit",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Сумма за аренду";
	var field = new FieldFloat("rent_price",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Заработано";
	var field = new FieldFloat("rent_total",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "К перечислению";
	var field = new FieldFloat("total",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			SpecialistPeriodSalaryDetail_Controller.prototype.addUpdate = function(){
	SpecialistPeriodSalaryDetail_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("specialist_period_salary_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Номер";
	var field = new FieldInt("line_num",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("specialist_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("studio_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDate("period",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Часов";
	var field = new FieldInt("hours",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldFloat("agent_percent",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Выручка";
	var field = new FieldFloat("work_total",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Выручка для зп";
	var field = new FieldFloat("work_total_salary",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Доп.начисления";
	var field = new FieldFloat("debet",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Цена за аренду";
	var field = new FieldFloat("kredit",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Сумма за аренду";
	var field = new FieldFloat("rent_price",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Заработано";
	var field = new FieldFloat("rent_total",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "К перечислению";
	var field = new FieldFloat("total",options);
	
	pm.addField(field);
	
	
}

			SpecialistPeriodSalaryDetail_Controller.prototype.addDelete = function(){
	SpecialistPeriodSalaryDetail_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			SpecialistPeriodSalaryDetail_Controller.prototype.addGetList = function(){
	SpecialistPeriodSalaryDetail_Controller.superclass.addGetList.call(this);
	
	
	
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
	
	pm.addField(new FieldInt("specialist_period_salary_id",f_opts));
	var f_opts = {};
	f_opts.alias = "Номер";
	pm.addField(new FieldInt("line_num",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldInt("specialist_id",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("specialists_ref",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldDate("period",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("period_descr",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldInt("studio_id",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("studios_ref",f_opts));
	var f_opts = {};
	f_opts.alias = "Часов";
	pm.addField(new FieldInt("hours",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldFloat("agent_percent",f_opts));
	var f_opts = {};
	f_opts.alias = "Выручка";
	pm.addField(new FieldFloat("work_total",f_opts));
	var f_opts = {};
	f_opts.alias = "Выручка для зп";
	pm.addField(new FieldFloat("work_total_salary",f_opts));
	var f_opts = {};
	f_opts.alias = "Доп.начисления";
	pm.addField(new FieldFloat("debet",f_opts));
	var f_opts = {};
	f_opts.alias = "Цена за аренду";
	pm.addField(new FieldFloat("kredit",f_opts));
	var f_opts = {};
	f_opts.alias = "Сумма за аренду";
	pm.addField(new FieldFloat("rent_price",f_opts));
	var f_opts = {};
	f_opts.alias = "Заработано";
	pm.addField(new FieldFloat("rent_total",f_opts));
	var f_opts = {};
	f_opts.alias = "К перечислению";
	pm.addField(new FieldFloat("total",f_opts));
	var f_opts = {};
	f_opts.alias = "Сумма по чекам по данной строке";
	pm.addField(new FieldFloat("receipt_total",f_opts));
	var f_opts = {};
	f_opts.alias = "Привью чеков";
	pm.addField(new FieldJSONB("receipt_photos",f_opts));
	var f_opts = {};
	f_opts.alias = "Чек проверен ФНС";
	pm.addField(new FieldBool("receipt_checked",f_opts));
	var f_opts = {};
	f_opts.alias = "Текст ошибки проверки чека";
	pm.addField(new FieldText("receipt_error",f_opts));
	var f_opts = {};
	f_opts.alias = "Ссылка на чек";
	pm.addField(new FieldText("receipt_href",f_opts));
	var f_opts = {};
	f_opts.alias = "Платежное поручение";
	pm.addField(new FieldJSON("bank_payments_ref",f_opts));
	var f_opts = {};
	f_opts.alias = "Представление";
	pm.addField(new FieldString("descr",f_opts));
	pm.getField(this.PARAM_ORD_FIELDS).setValue("line_num");
	
}

			SpecialistPeriodSalaryDetail_Controller.prototype.addGetObject = function(){
	SpecialistPeriodSalaryDetail_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

			SpecialistPeriodSalaryDetail_Controller.prototype.add_get_for_pay_list = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('get_for_pay_list',opts);
	
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

			SpecialistPeriodSalaryDetail_Controller.prototype.addComplete = function(){
	SpecialistPeriodSalaryDetail_Controller.superclass.addComplete.call(this);
	
	var f_opts = {};
	f_opts.alias = "";
	var pm = this.getComplete();
	pm.addField(new FieldString("descr",f_opts));
	pm.addField(new FieldInt("count", {}));
	pm.getField(this.PARAM_ORD_FIELDS).setValue("descr");	
}

			SpecialistPeriodSalaryDetail_Controller.prototype.add_documents_to_bank = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('documents_to_bank',opts);
	
				
	
	var options = {};
	
		options.alias = "Массив идентификаторов документов";
	
		options.required = true;
	
		pm.addField(new FieldString("ids",options));
	
			
	this.addPublicMethod(pm);
}

		