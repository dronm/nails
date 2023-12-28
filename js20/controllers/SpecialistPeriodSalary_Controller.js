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

function SpecialistPeriodSalary_Controller(options){
	options = options || {};
	options.listModelClass = SpecialistPeriodSalaryList_Model;
	options.objModelClass = SpecialistPeriodSalaryList_Model;
	SpecialistPeriodSalary_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
	this.add_fill_doc();
		
}
extend(SpecialistPeriodSalary_Controller,ControllerObjServer);

			SpecialistPeriodSalary_Controller.prototype.addInsert = function(){
	SpecialistPeriodSalary_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDateTimeTZ("date_time",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldDate("period",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldFloat("total",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Вырчка";
	var field = new FieldFloat("work_total",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Вырчка для з/п";
	var field = new FieldFloat("work_total_salary",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Всего часов";
	var field = new FieldInt("hours",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Всего доп начислений";
	var field = new FieldFloat("debet",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Всего доп удержаний";
	var field = new FieldFloat("kredit",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Всего доп удержаний";
	var field = new FieldFloat("rent_total",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("studio_id",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			SpecialistPeriodSalary_Controller.prototype.addUpdate = function(){
	SpecialistPeriodSalary_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDateTimeTZ("date_time",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDate("period",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldFloat("total",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Вырчка";
	var field = new FieldFloat("work_total",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Вырчка для з/п";
	var field = new FieldFloat("work_total_salary",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Всего часов";
	var field = new FieldInt("hours",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Всего доп начислений";
	var field = new FieldFloat("debet",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Всего доп удержаний";
	var field = new FieldFloat("kredit",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Всего доп удержаний";
	var field = new FieldFloat("rent_total",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("studio_id",options);
	
	pm.addField(field);
	
	
}

			SpecialistPeriodSalary_Controller.prototype.addDelete = function(){
	SpecialistPeriodSalary_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			SpecialistPeriodSalary_Controller.prototype.addGetList = function(){
	SpecialistPeriodSalary_Controller.superclass.addGetList.call(this);
	
	
	
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
	
	pm.addField(new FieldDateTimeTZ("date_time",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("period_descr",f_opts));
	var f_opts = {};
	f_opts.alias = "Вырчка";
	pm.addField(new FieldFloat("work_total",f_opts));
	var f_opts = {};
	f_opts.alias = "Вырчка для з/п";
	pm.addField(new FieldFloat("work_total_salary",f_opts));
	var f_opts = {};
	f_opts.alias = "Всего часов";
	pm.addField(new FieldInt("hours",f_opts));
	var f_opts = {};
	f_opts.alias = "Всего доп начислений";
	pm.addField(new FieldFloat("debet",f_opts));
	var f_opts = {};
	f_opts.alias = "Всего доп удержаний";
	pm.addField(new FieldFloat("kredit",f_opts));
	var f_opts = {};
	f_opts.alias = "Всего доп удержаний";
	pm.addField(new FieldFloat("rent_total",f_opts));
	var f_opts = {};
	f_opts.alias = "К перечислению";
	pm.addField(new FieldFloat("total",f_opts));
	pm.getField(this.PARAM_ORD_FIELDS).setValue("period");
	
}

			SpecialistPeriodSalary_Controller.prototype.addGetObject = function(){
	SpecialistPeriodSalary_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

			SpecialistPeriodSalary_Controller.prototype.add_fill_doc = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('fill_doc',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldInt("id",options));
	
			
	this.addPublicMethod(pm);
}

		