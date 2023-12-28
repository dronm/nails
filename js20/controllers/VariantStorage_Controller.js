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

function VariantStorage_Controller(options){
	options = options || {};
	options.objModelClass = VariantStorage_Model;
	VariantStorage_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.addInsert();
	this.add_upsert_filter_data();
	this.add_upsert_col_visib_data();
	this.add_upsert_col_order_data();
	this.addUpdate();
	this.addDelete();
	this.addGetList();
	this.addGetObject();
	this.add_get_filter_data();
	this.add_get_col_visib_data();
	this.add_get_col_order_data();
		
}
extend(VariantStorage_Controller,ControllerObjServer);

			VariantStorage_Controller.prototype.addInsert = function(){
	VariantStorage_Controller.superclass.addInsert.call(this);
	
	var pm = this.getInsert();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("user_id",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldText("storage_name",options);
	
	pm.addField(field);
	
	var options = {};
	options.required = true;
	var field = new FieldText("variant_name",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("default_variant",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldJSONB("filter_data",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("col_visib_data",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("col_order_data",options);
	
	pm.addField(field);
	
	pm.addField(new FieldInt("ret_id",{}));
	
	
}

			VariantStorage_Controller.prototype.add_upsert_filter_data = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('upsert_filter_data',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("storage_name",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("variant_name",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("filter_data",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldBool("default_variant",options));
	
			
	this.addPublicMethod(pm);
}
			
			VariantStorage_Controller.prototype.add_upsert_col_visib_data = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('upsert_col_visib_data',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("storage_name",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("variant_name",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("col_visib",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldBool("default_variant",options));
	
			
	this.addPublicMethod(pm);
}
			
			VariantStorage_Controller.prototype.add_upsert_col_order_data = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('upsert_col_order_data',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("storage_name",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("variant_name",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("col_order",options));
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldBool("default_variant",options));
	
			
	this.addPublicMethod(pm);
}
						
			VariantStorage_Controller.prototype.addUpdate = function(){
	VariantStorage_Controller.superclass.addUpdate.call(this);
	var pm = this.getUpdate();
	
	var options = {};
	options.primaryKey = true;options.autoInc = true;
	var field = new FieldInt("id",options);
	
	pm.addField(field);
	
	field = new FieldInt("old_id",{});
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldInt("user_id",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("storage_name",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("variant_name",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldBool("default_variant",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldJSONB("filter_data",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("col_visib_data",options);
	
	pm.addField(field);
	
	var options = {};
	
	var field = new FieldText("col_order_data",options);
	
	pm.addField(field);
	
	
}

			VariantStorage_Controller.prototype.addDelete = function(){
	VariantStorage_Controller.superclass.addDelete.call(this);
	var pm = this.getDelete();
	var options = {"required":true};
		
	pm.addField(new FieldInt("id",options));
}

			VariantStorage_Controller.prototype.addGetList = function(){
	VariantStorage_Controller.superclass.addGetList.call(this);
	
	
	
	var pm = this.getGetList();
	
	pm.addField(new FieldInt(this.PARAM_COUNT));
	pm.addField(new FieldInt(this.PARAM_FROM));
	pm.addField(new FieldString(this.PARAM_COND_FIELDS));
	pm.addField(new FieldString(this.PARAM_COND_SGNS));
	pm.addField(new FieldString(this.PARAM_COND_VALS));
	pm.addField(new FieldString(this.PARAM_COND_ICASE));
	pm.addField(new FieldString(this.PARAM_ORD_FIELDS));
	pm.addField(new FieldString(this.PARAM_ORD_DIRECTS));
	pm.addField(new FieldString(this.PARAM_FIELD_SEP));

		var options = {};
		
			options.required = true;
						
		pm.addField(new FieldString("storage_name",options));
	
		var options = {};
						
		pm.addField(new FieldString("variant_name",options));
	
}

			VariantStorage_Controller.prototype.addGetObject = function(){
	VariantStorage_Controller.superclass.addGetObject.call(this);
	
	var pm = this.getGetObject();
	var f_opts = {};
		
	pm.addField(new FieldInt("id",f_opts));
	
	pm.addField(new FieldString("mode"));
}

			VariantStorage_Controller.prototype.add_get_filter_data = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('get_filter_data',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("storage_name",options));
	
				
	
	var options = {};
	
		pm.addField(new FieldString("variant_name",options));
	
			
	this.addPublicMethod(pm);
}

			VariantStorage_Controller.prototype.add_get_col_visib_data = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('get_col_visib_data',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("storage_name",options));
	
				
	
	var options = {};
	
		pm.addField(new FieldString("variant_name",options));
	
			
	this.addPublicMethod(pm);
}
			
			VariantStorage_Controller.prototype.add_get_col_order_data = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('get_col_order_data',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		pm.addField(new FieldString("storage_name",options));
	
				
	
	var options = {};
	
		pm.addField(new FieldString("variant_name",options));
	
			
	this.addPublicMethod(pm);
}
			
		