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

function SpecialistDocument_Model(options){
	var id = 'SpecialistDocument_Model';
	options = options || {};
	
	options.fields = {};
	
				
	
	var filed_options = {};
	filed_options.primaryKey = true;	
	
	filed_options.autoInc = true;	
	
	options.fields.id = new FieldInt("id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.specialist_id = new FieldInt("specialist_id",filed_options);
	options.fields.specialist_id.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.template_att_id = new FieldInt("template_att_id",filed_options);
	options.fields.template_att_id.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.document_att_id = new FieldInt("document_att_id",filed_options);
	options.fields.document_att_id.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Изображение подписи';
	filed_options.autoInc = false;	
	
	options.fields.sign_img = new FieldBytea("sign_img",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.defValue = true;
	
	filed_options.autoInc = false;	
	
	options.fields.date_time = new FieldDateTimeTZ("date_time",filed_options);
	options.fields.date_time.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.sign_date_time = new FieldDateTimeTZ("sign_date_time",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.open_date_time = new FieldDateTimeTZ("open_date_time",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.need_signing = new FieldBool("need_signing",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Описание, представление';
	filed_options.autoInc = false;	
	
	options.fields.name = new FieldText("name",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.document_template_id = new FieldInt("document_template_id",filed_options);
	options.fields.document_template_id.getValidator().setRequired(true);
	
			
			
		SpecialistDocument_Model.superclass.constructor.call(this,id,options);
}
extend(SpecialistDocument_Model,ModelXML);

