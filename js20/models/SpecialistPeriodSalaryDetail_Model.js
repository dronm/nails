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

function SpecialistPeriodSalaryDetail_Model(options){
	var id = 'SpecialistPeriodSalaryDetail_Model';
	options = options || {};
	
	options.fields = {};
	
				
	
	var filed_options = {};
	filed_options.primaryKey = true;	
	
	filed_options.autoInc = true;	
	
	options.fields.id = new FieldInt("id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.specialist_period_salary_id = new FieldInt("specialist_period_salary_id",filed_options);
	options.fields.specialist_period_salary_id.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Номер';
	filed_options.autoInc = false;	
	
	options.fields.line_num = new FieldInt("line_num",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.specialist_id = new FieldInt("specialist_id",filed_options);
	options.fields.specialist_id.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.studio_id = new FieldInt("studio_id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.period = new FieldDate("period",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Часов';
	filed_options.autoInc = false;	
	
	options.fields.hours = new FieldInt("hours",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.agent_percent = new FieldFloat("agent_percent",filed_options);
	options.fields.agent_percent.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Выручка';
	filed_options.autoInc = false;	
	
	options.fields.work_total = new FieldFloat("work_total",filed_options);
	options.fields.work_total.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Выручка для зп';
	filed_options.autoInc = false;	
	
	options.fields.work_total_salary = new FieldFloat("work_total_salary",filed_options);
	options.fields.work_total_salary.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Доп.начисления';
	filed_options.autoInc = false;	
	
	options.fields.debet = new FieldFloat("debet",filed_options);
	options.fields.debet.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Цена за аренду';
	filed_options.autoInc = false;	
	
	options.fields.kredit = new FieldFloat("kredit",filed_options);
	options.fields.kredit.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Сумма за аренду';
	filed_options.autoInc = false;	
	
	options.fields.rent_price = new FieldFloat("rent_price",filed_options);
	options.fields.rent_price.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Заработано';
	filed_options.autoInc = false;	
	
	options.fields.rent_total = new FieldFloat("rent_total",filed_options);
	options.fields.rent_total.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'К перечислению';
	filed_options.autoInc = false;	
	
	options.fields.total = new FieldFloat("total",filed_options);
	options.fields.total.getValidator().setMaxLength('15');
	
			
			
			
		SpecialistPeriodSalaryDetail_Model.superclass.constructor.call(this,id,options);
}
extend(SpecialistPeriodSalaryDetail_Model,ModelXML);

