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

function NotifTemplate_Controller(options){
	options = options || {};
	options.listModelClass = NotifTemplateList_Model;
	options.objModelClass = NotifTemplate_Model;
	NotifTemplate_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
		
}
extend(NotifTemplate_Controller,ControllerObjServer);

			NotifTemplate_Controller.prototype.addInsert = function(){
	NotifTemplate_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Способ доставки";options.required = true;	
	options.enumValues = 'email,sms,wa,tm,vb';
	var field = new FieldEnum("notif_provider",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Тип сообщения";options.required = true;	
	options.enumValues = 'new_specialist,tel_check,email_check,docs_for_sign,signed_docs,new_account';
	var field = new FieldEnum("notif_type",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Шаблон";options.required = true;
	var field = new FieldText("template",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Комментарий";options.required = true;
	var field = new FieldText("comment_text",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Поля";options.required = true;
	var field = new FieldJSON("fields",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Специфичные для провайдера поля, например subject для писем, фичи оформления для месенджеров";
	var field = new FieldJSON("provider_values",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			NotifTemplate_Controller.prototype.addUpdate = function(){
	NotifTemplate_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	options.alias = "Способ доставки";	
	options.enumValues = 'email,sms,wa,tm,vb';
	options.enumValues+= (options.enumValues=='')? '':',';
	options.enumValues+= 'null';
	
	var field = new FieldEnum("notif_provider",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Тип сообщения";	
	options.enumValues = 'new_specialist,tel_check,email_check,docs_for_sign,signed_docs,new_account';
	options.enumValues+= (options.enumValues=='')? '':',';
	options.enumValues+= 'null';
	
	var field = new FieldEnum("notif_type",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Шаблон";
	var field = new FieldText("template",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Комментарий";
	var field = new FieldText("comment_text",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Поля";
	var field = new FieldJSON("fields",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Специфичные для провайдера поля, например subject для писем, фичи оформления для месенджеров";
	var field = new FieldJSON("provider_values",options);
	
	pm.addField(field);
	
	
}

			NotifTemplate_Controller.prototype.addDelete = function(){
	NotifTemplate_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			NotifTemplate_Controller.prototype.addGetList = function(){
	NotifTemplate_Controller.superclass.addGetList.call(this);
	
	
	
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
	
	pm.addField(new FieldString("notif_provider",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldString("notif_type",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldText("template",f_opts));
}

			NotifTemplate_Controller.prototype.addGetObject = function(){
	NotifTemplate_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

		