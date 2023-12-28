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

function SpecialistDocument_Controller(options){
	options = options || {};
	options.listModelClass = SpecialistDocumentList_Model;
	options.objModelClass = SpecialistDocumentList_Model;
	SpecialistDocument_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
	this.add_get_for_sign_list();
	this.add_sign();
	this.add_get_file();
		
}
extend(SpecialistDocument_Controller,ControllerObjServer);

			SpecialistDocument_Controller.prototype.addInsert = function(){
	SpecialistDocument_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("specialist_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("template_att_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("document_att_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Изображение подписи";
	var field = new FieldBytea("sign_img",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldDateTimeTZ("date_time",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDateTimeTZ("sign_date_time",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDateTimeTZ("open_date_time",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("need_signing",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Описание, представление";
	var field = new FieldText("name",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldInt("document_template_id",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			SpecialistDocument_Controller.prototype.addUpdate = function(){
	SpecialistDocument_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("specialist_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("template_att_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("document_att_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Изображение подписи";
	var field = new FieldBytea("sign_img",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDateTimeTZ("date_time",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDateTimeTZ("sign_date_time",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldDateTimeTZ("open_date_time",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("need_signing",options);
	
	pm.addField(field);
	
	var options = {};
	options.alias = "Описание, представление";
	var field = new FieldText("name",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("document_template_id",options);
	
	pm.addField(field);
	
	
}

			SpecialistDocument_Controller.prototype.addDelete = function(){
	SpecialistDocument_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			SpecialistDocument_Controller.prototype.addGetList = function(){
	SpecialistDocument_Controller.superclass.addGetList.call(this);
	
	
	
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
	
	pm.addField(new FieldInt("specialist_id",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("specialists_ref",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldJSON("document_att_ref",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldDateTimeTZ("date_time",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldDateTimeTZ("sign_date_time",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldBool("signed",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldDateTimeTZ("open_date_time",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldBool("opened",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldBool("need_signing",f_opts));
	var f_opts = {};
	f_opts.alias = "Описание, представление";
	pm.addField(new FieldText("name",f_opts));
}

			SpecialistDocument_Controller.prototype.addGetObject = function(){
	SpecialistDocument_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

			SpecialistDocument_Controller.prototype.add_get_for_sign_list = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('get_for_sign_list',opts);
	
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

			SpecialistDocument_Controller.prototype.add_sign = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('sign',opts);
	
	pm.setRequestType('post');
	
	pm.setEncType(ServConnector.prototype.ENCTYPES.MULTIPART);
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldInt("id",options));
	
				
	
	var options = {};
	
		pm.addField(new FieldBytea("signature",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "36";
	
		pm.addField(new FieldString("operation_id",options));
	
			
	this.addPublicMethod(pm);
}

			SpecialistDocument_Controller.prototype.add_get_file = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('get_file',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("ref",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "36";
	
		pm.addField(new FieldString("content_id",options));
	
				
	
	var options = {};
	
		pm.addField(new FieldInt("inline",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldInt("doc_id",options));
	
			
	this.addPublicMethod(pm);
}

		