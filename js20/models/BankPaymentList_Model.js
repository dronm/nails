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

function BankPaymentList_Model(options){
	var id = 'BankPaymentList_Model';
	options = options || {};
	
	options.fields = {};
	
			
				
			
			
				
					
				
	
	var filed_options = {};
	filed_options.primaryKey = true;	
	
	filed_options.autoInc = true;	
	
	options.fields.id = new FieldInt("id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Дата создания';
	filed_options.autoInc = false;	
	
	options.fields.date_time = new FieldDateTimeTZ("date_time",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Дата документа';
	filed_options.autoInc = false;	
	
	options.fields.document_date = new FieldDateTimeTZ("document_date",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Номер документа';
	filed_options.autoInc = false;	
	
	options.fields.document_num = new FieldInt("document_num",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Сумма документа';
	filed_options.autoInc = false;	
	
	options.fields.document_total = new FieldFloat("document_total",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Назначение платежа';
	filed_options.autoInc = false;	
	
	options.fields.document_comment = new FieldText("document_comment",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Расчетный счет плательщика';
	filed_options.autoInc = false;	
	
	options.fields.payer_acc = new FieldString("payer_acc",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Кор. счет банка плательщика';
	filed_options.autoInc = false;	
	
	options.fields.payer_bank_acc = new FieldString("payer_bank_acc",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Бик банка плательщика';
	filed_options.autoInc = false;	
	
	options.fields.payer_bank_bik = new FieldString("payer_bank_bik",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Банк плательщика';
	filed_options.autoInc = false;	
	
	options.fields.payer_bank = new FieldText("payer_bank",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Город банка плательщика';
	filed_options.autoInc = false;	
	
	options.fields.payer_bank_place = new FieldText("payer_bank_place",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Расчетный счет получателя';
	filed_options.autoInc = false;	
	
	options.fields.rec_acc = new FieldString("rec_acc",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Кор. счет банка получателя';
	filed_options.autoInc = false;	
	
	options.fields.rec_bank_acc = new FieldString("rec_bank_acc",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Бик банка получателя';
	filed_options.autoInc = false;	
	
	options.fields.rec_bank_bik = new FieldString("rec_bank_bik",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Банк получателя';
	filed_options.autoInc = false;	
	
	options.fields.rec_bank = new FieldText("rec_bank",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Город банка получателя';
	filed_options.autoInc = false;	
	
	options.fields.rec_bank_place = new FieldText("rec_bank_place",filed_options);
	
				
	
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
	
	options.fields.specialist_period_salary_detail_id = new FieldInt("specialist_period_salary_detail_id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.specialist_period_salary_details_ref = new FieldJSON("specialist_period_salary_details_ref",filed_options);
	
		BankPaymentList_Model.superclass.constructor.call(this,id,options);
}
extend(BankPaymentList_Model,ModelXML);

