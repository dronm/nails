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

function ConfirmationStatus_Model(options){
	var id = 'ConfirmationStatus_Model';
	options = options || {};
	
	options.fields = {};
	
				
	
	var filed_options = {};
	filed_options.primaryKey = true;	
	
	filed_options.autoInc = true;	
	
	options.fields.id = new FieldInt("id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Объект проверки,';
	filed_options.autoInc = false;	
	
	options.fields.ref = new FieldJSONB("ref",filed_options);
	options.fields.ref.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.field = new FieldString("field",filed_options);
	options.fields.field.getValidator().setRequired(true);
	options.fields.field.getValidator().setMaxLength('50');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.secret = new FieldText("secret",filed_options);
	options.fields.secret.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.defValue = true;
	
	filed_options.autoInc = false;	
	
	options.fields.confirmed = new FieldBool("confirmed",filed_options);
	options.fields.confirmed.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.defValue = true;
	
	filed_options.autoInc = false;	
	
	options.fields.tries = new FieldInt("tries",filed_options);
	options.fields.tries.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.try_date_time = new FieldDateTimeTZ("try_date_time",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.defValue = true;
	
	filed_options.autoInc = false;	
	
	options.fields.date_time = new FieldDateTimeTZ("date_time",filed_options);
	options.fields.date_time.getValidator().setRequired(true);
	
			
		ConfirmationStatus_Model.superclass.constructor.call(this,id,options);
}
extend(ConfirmationStatus_Model,ModelXML);

