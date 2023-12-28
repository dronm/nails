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

function YClStaff_Controller(options){
	options = options || {};
	options.listModelClass = YClStaffList_Model;
	options.objModelClass = YClStaffList_Model;
	YClStaff_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
	this.add_api_get_all();
		
}
extend(YClStaff_Controller,ControllerObjServer);

			YClStaff_Controller.prototype.addInsert = function(){
	YClStaff_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("name",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldJSONB("data",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			YClStaff_Controller.prototype.addUpdate = function(){
	YClStaff_Controller.superclass.addUpdate.call(this);
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
	
	var field = new FieldJSONB("data",options);
	
	pm.addField(field);
	
	
}

			YClStaff_Controller.prototype.addDelete = function(){
	YClStaff_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			YClStaff_Controller.prototype.addGetList = function(){
	YClStaff_Controller.superclass.addGetList.call(this);
	
	
	
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
	
	pm.addField(new FieldText("specialization",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldText("avatar",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldText("avatar_big",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldFloat("rating",f_opts));
	var f_opts = {};
	
	pm.addField(new FieldInt("votes_count",f_opts));
}

			YClStaff_Controller.prototype.addGetObject = function(){
	YClStaff_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
	pm.addField(new FieldString("lsn"));
}

			YClStaff_Controller.prototype.add_api_get_all = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('api_get_all',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "36";
	
		pm.addField(new FieldString("operation_id",options));
	
			
	this.addPublicMethod(pm);
}

		