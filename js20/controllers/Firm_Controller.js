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

function Firm_Controller(options){
	options = options || {};
	options.listModelClass = FirmList_Model;
	options.objModelClass = FirmDialog_Model;
	Firm_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
		
}
extend(Firm_Controller,ControllerObjServer);

			Firm_Controller.prototype.addInsert = function(){
	Firm_Controller.superclass.addInsert.call(this);
	
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
	
	var field = new FieldString("inn",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Полное наименование";
	var field = new FieldText("name_full",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Адрес юридический";
	var field = new FieldText("legal_address",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Адрес почтовый";
	var field = new FieldText("post_address",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "КПП";
	var field = new FieldString("kpp",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldString("ogrn",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldString("okpo",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("okved",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Расчетный счет";
	var field = new FieldString("bank_acc",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Банк";
	var field = new FieldString("bank_bik",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			Firm_Controller.prototype.addUpdate = function(){
	Firm_Controller.superclass.addUpdate.call(this);
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
	options.alias = "Полное наименование";
	var field = new FieldText("name_full",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Адрес юридический";
	var field = new FieldText("legal_address",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Адрес почтовый";
	var field = new FieldText("post_address",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "КПП";
	var field = new FieldString("kpp",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldString("ogrn",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldString("okpo",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("okved",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Расчетный счет";
	var field = new FieldString("bank_acc",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Банк";
	var field = new FieldString("bank_bik",options);
	
	pm.addField(field);
	
	
}

			Firm_Controller.prototype.addDelete = function(){
	Firm_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			Firm_Controller.prototype.addGetList = function(){
	Firm_Controller.superclass.addGetList.call(this);
	
	
	
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
	
	pm.addField(new FieldString("ogrn",f_opts));
	pm.getField(this.PARAM_ORD_FIELDS).setValue("name");
	
}

			Firm_Controller.prototype.addGetObject = function(){
	Firm_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

		