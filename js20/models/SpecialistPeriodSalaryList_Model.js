/**	
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/models/Model_js.xsl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2017
 * @class
 * @classdesc Model class. Created from template build/templates/models/Model_js.xsl. !!!DO NOT MODEFY!!!
 
 * @extends ModelXML
 
 * @requires core/extend.js
 * @requires core/ModelXML.js
 
 * @param {string} id 
 * @param {Object} options
 */

function SpecialistPeriodSalaryList_Model(options){
	var id = 'SpecialistPeriodSalaryList_Model';
	options = options || {};
	
	options.fields = {};
	
			
				
			
			
				
				
				
				
				
				
				
			
				
	
	var filed_options = {};
	filed_options.primaryKey = true;	
	
	filed_options.autoInc = true;	
	
	options.fields.id = new FieldInt("id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.date_time = new FieldDateTimeTZ("date_time",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.period = new FieldDate("period",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.studio_id = new FieldInt("studio_id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.studios_ref = new FieldJSON("studios_ref",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.period_descr = new FieldString("period_descr",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Вырчка';
	filed_options.autoInc = false;	
	
	options.fields.work_total = new FieldFloat("work_total",filed_options);
	options.fields.work_total.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Вырчка для з/п';
	filed_options.autoInc = false;	
	
	options.fields.work_total_salary = new FieldFloat("work_total_salary",filed_options);
	options.fields.work_total_salary.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Всего часов';
	filed_options.autoInc = false;	
	
	options.fields.hours = new FieldInt("hours",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Всего доп начислений';
	filed_options.autoInc = false;	
	
	options.fields.debet = new FieldFloat("debet",filed_options);
	options.fields.debet.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Всего доп удержаний';
	filed_options.autoInc = false;	
	
	options.fields.kredit = new FieldFloat("kredit",filed_options);
	options.fields.kredit.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Всего доп удержаний';
	filed_options.autoInc = false;	
	
	options.fields.rent_total = new FieldFloat("rent_total",filed_options);
	options.fields.rent_total.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'К перечислению';
	filed_options.autoInc = false;	
	
	options.fields.total = new FieldFloat("total",filed_options);
	options.fields.total.getValidator().setMaxLength('15');
	
		SpecialistPeriodSalaryList_Model.superclass.constructor.call(this,id,options);
}
extend(SpecialistPeriodSalaryList_Model,ModelXML);

