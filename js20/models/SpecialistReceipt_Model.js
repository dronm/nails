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

function SpecialistReceipt_Model(options){
	var id = 'SpecialistReceipt_Model';
	options = options || {};
	
	options.fields = {};
	
				
	
	var filed_options = {};
	filed_options.primaryKey = true;	
	
	filed_options.autoInc = true;	
	
	options.fields.id = new FieldInt("id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.defValue = true;
	filed_options.alias = 'Дата загрузки';
	filed_options.autoInc = false;	
	
	options.fields.date_time = new FieldDateTimeTZ("date_time",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Дата документа';
	filed_options.autoInc = false;	
	
	options.fields.document_date_time = new FieldDateTimeTZ("document_date_time",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.defValue = true;
	filed_options.alias = 'Сумма документа';
	filed_options.autoInc = false;	
	
	options.fields.document_total = new FieldFloat("document_total",filed_options);
	options.fields.document_total.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.defValue = true;
	filed_options.alias = 'Сумма проверен ФНС';
	filed_options.autoInc = false;	
	
	options.fields.document_parsed = new FieldBool("document_parsed",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.specialist_id = new FieldInt("specialist_id",filed_options);
	options.fields.specialist_id.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.specialist_period_salary_detail_id = new FieldInt("specialist_period_salary_detail_id",filed_options);
	options.fields.specialist_period_salary_detail_id.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Ошибка проверки сека ФНС';
	filed_options.autoInc = false;	
	
	options.fields.document_error = new FieldText("document_error",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'ИД сервиса qrextr';
	filed_options.autoInc = false;	
	
	options.fields.qrextr_request_id = new FieldText("qrextr_request_id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.operation_id = new FieldString("operation_id",filed_options);
	options.fields.operation_id.getValidator().setMaxLength('36');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.document_href = new FieldText("document_href",filed_options);
	
			
			
			
		SpecialistReceipt_Model.superclass.constructor.call(this,id,options);
}
extend(SpecialistReceipt_Model,ModelXML);

