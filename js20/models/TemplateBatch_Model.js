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

function TemplateBatch_Model(options){
	var id = 'TemplateBatch_Model';
	options = options || {};
	
	options.fields = {};
	
			
				
			
				
	
	var filed_options = {};
	filed_options.primaryKey = true;	
	
	filed_options.autoInc = true;	
	
	options.fields.id = new FieldInt("id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Наименование пакета';
	filed_options.autoInc = false;	
	
	options.fields.name = new FieldText("name",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.template_batch_type = new FieldEnum("template_batch_type",filed_options);
	filed_options.enumValues = 'specialist_registration,specialist_salary';
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Студия или пусто, если все';
	filed_options.autoInc = false;	
	
	options.fields.studio_id = new FieldInt("studio_id",filed_options);
	
		TemplateBatch_Model.superclass.constructor.call(this,id,options);
}
extend(TemplateBatch_Model,ModelXML);

