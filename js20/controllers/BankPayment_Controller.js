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

function BankPayment_Controller(options){
	options = options || {};
	options.listModelClass = BankPaymentList_Model;
	options.objModelClass = BankPaymentList_Model;
	BankPayment_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
		
}
extend(BankPayment_Controller,ControllerObjServer);

			BankPayment_Controller.prototype.addInsert = function(){
	BankPayment_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Дата создания";
	var field = new FieldDateTimeTZ("date_time",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Дата документа";
	var field = new FieldDateTimeTZ("document_date",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Номер документа";
	var field = new FieldInt("document_num",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Сумма документа";
	var field = new FieldFloat("document_total",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Назначение платежа";
	var field = new FieldText("document_comment",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Расчетный счет плательщика";
	var field = new FieldString("payer_acc",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Кор. счет банка плательщика";
	var field = new FieldString("payer_bank_acc",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Бик банка плательщика";
	var field = new FieldString("payer_bank_bik",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Банк плательщика";
	var field = new FieldText("payer_bank",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Город банка плательщика";
	var field = new FieldText("payer_bank_place",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Расчетный счет получателя";
	var field = new FieldString("rec_acc",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Кор. счет банка получателя";
	var field = new FieldString("rec_bank_acc",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Бик банка получателя";
	var field = new FieldString("rec_bank_bik",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Банк получателя";
	var field = new FieldText("rec_bank",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Город банка получателя";
	var field = new FieldText("rec_bank_place",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("specialist_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("specialist_period_salary_detail_id",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			BankPayment_Controller.prototype.addUpdate = function(){
	BankPayment_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	options.alias = "Дата создания";
	var field = new FieldDateTimeTZ("date_time",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Дата документа";
	var field = new FieldDateTimeTZ("document_date",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Номер документа";
	var field = new FieldInt("document_num",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Сумма документа";
	var field = new FieldFloat("document_total",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Назначение платежа";
	var field = new FieldText("document_comment",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Расчетный счет плательщика";
	var field = new FieldString("payer_acc",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Кор. счет банка плательщика";
	var field = new FieldString("payer_bank_acc",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Бик банка плательщика";
	var field = new FieldString("payer_bank_bik",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Банк плательщика";
	var field = new FieldText("payer_bank",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Город банка плательщика";
	var field = new FieldText("payer_bank_place",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Расчетный счет получателя";
	var field = new FieldString("rec_acc",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Кор. счет банка получателя";
	var field = new FieldString("rec_bank_acc",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Бик банка получателя";
	var field = new FieldString("rec_bank_bik",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Банк получателя";
	var field = new FieldText("rec_bank",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Город банка получателя";
	var field = new FieldText("rec_bank_place",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("specialist_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("specialist_period_salary_detail_id",options);
	
	pm.addField(field);
	
	
}

			BankPayment_Controller.prototype.addDelete = function(){
	BankPayment_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			BankPayment_Controller.prototype.addGetList = function(){
	BankPayment_Controller.superclass.addGetList.call(this);
	
	
	
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
	f_opts.alias = "Дата создания";
	pm.addField(new FieldDateTimeTZ("date_time",f_opts));
	var f_opts = {};
	f_opts.alias = "Дата документа";
	pm.addField(new FieldDateTimeTZ("document_date",f_opts));
	var f_opts = {};
	f_opts.alias = "Номер документа";
	pm.addField(new FieldInt("document_num",f_opts));
	var f_opts = {};
	f_opts.alias = "Сумма документа";
	pm.addField(new FieldFloat("document_total",f_opts));
	var f_opts = {};
	f_opts.alias = "Назначение платежа";
	pm.addField(new FieldText("document_comment",f_opts));
	var f_opts = {};
	f_opts.alias = "Расчетный счет плательщика";
	pm.addField(new FieldString("payer_acc",f_opts));
	var f_opts = {};
	f_opts.alias = "Кор. счет банка плательщика";
	pm.addField(new FieldString("payer_bank_acc",f_opts));
	var f_opts = {};
	f_opts.alias = "Бик банка плательщика";
	pm.addField(new FieldString("payer_bank_bik",f_opts));
	var f_opts = {};
	f_opts.alias = "Банк плательщика";
	pm.addField(new FieldText("payer_bank",f_opts));
	var f_opts = {};
	f_opts.alias = "Город банка плательщика";
	pm.addField(new FieldText("payer_bank_place",f_opts));
	var f_opts = {};
	f_opts.alias = "Расчетный счет получателя";
	pm.addField(new FieldString("rec_acc",f_opts));
	var f_opts = {};
	f_opts.alias = "Кор. счет банка получателя";
	pm.addField(new FieldString("rec_bank_acc",f_opts));
	var f_opts = {};
	f_opts.alias = "Бик банка получателя";
	pm.addField(new FieldString("rec_bank_bik",f_opts));
	var f_opts = {};
	f_opts.alias = "Банк получателя";
	pm.addField(new FieldText("rec_bank",f_opts));
	var f_opts = {};
	f_opts.alias = "Город банка получателя";
	pm.addField(new FieldText("rec_bank_place",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldInt("specialist_id",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("specialists_ref",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldInt("specialist_period_salary_detail_id",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("specialist_period_salary_details_ref",f_opts));
	pm.getField(this.PARAM_ORD_FIELDS).setValue("document_num");
	
}

			BankPayment_Controller.prototype.addGetObject = function(){
	BankPayment_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

		