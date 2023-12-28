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

function SpecialistPeriodSalaryDetailList_Model(options){
	var id = 'SpecialistPeriodSalaryDetailList_Model';
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
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Номер';
	filed_options.autoInc = false;	
	
	options.fields.line_num = new FieldInt("line_num",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.specialist_id = new FieldInt("specialist_id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.specialists_ref = new FieldJSON("specialists_ref",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.period = new FieldDate("period",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.period_descr = new FieldString("period_descr",filed_options);
	
				
	
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
	filed_options.alias = 'Часов';
	filed_options.autoInc = false;	
	
	options.fields.hours = new FieldInt("hours",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.agent_percent = new FieldFloat("agent_percent",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Выручка';
	filed_options.autoInc = false;	
	
	options.fields.work_total = new FieldFloat("work_total",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Выручка для зп';
	filed_options.autoInc = false;	
	
	options.fields.work_total_salary = new FieldFloat("work_total_salary",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Доп.начисления';
	filed_options.autoInc = false;	
	
	options.fields.debet = new FieldFloat("debet",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Цена за аренду';
	filed_options.autoInc = false;	
	
	options.fields.kredit = new FieldFloat("kredit",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Сумма за аренду';
	filed_options.autoInc = false;	
	
	options.fields.rent_price = new FieldFloat("rent_price",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Заработано';
	filed_options.autoInc = false;	
	
	options.fields.rent_total = new FieldFloat("rent_total",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'К перечислению';
	filed_options.autoInc = false;	
	
	options.fields.total = new FieldFloat("total",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Сумма по чекам по данной строке';
	filed_options.autoInc = false;	
	
	options.fields.receipt_total = new FieldFloat("receipt_total",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Привью чеков';
	filed_options.autoInc = false;	
	
	options.fields.receipt_photos = new FieldJSONB("receipt_photos",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Чек проверен ФНС';
	filed_options.autoInc = false;	
	
	options.fields.receipt_checked = new FieldBool("receipt_checked",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Текст ошибки проверки чека';
	filed_options.autoInc = false;	
	
	options.fields.receipt_error = new FieldText("receipt_error",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Ссылка на чек';
	filed_options.autoInc = false;	
	
	options.fields.receipt_href = new FieldText("receipt_href",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Платежное поручение';
	filed_options.autoInc = false;	
	
	options.fields.bank_payments_ref = new FieldJSON("bank_payments_ref",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Представление';
	filed_options.autoInc = false;	
	
	options.fields.descr = new FieldString("descr",filed_options);
	
		SpecialistPeriodSalaryDetailList_Model.superclass.constructor.call(this,id,options);
}
extend(SpecialistPeriodSalaryDetailList_Model,ModelXML);

