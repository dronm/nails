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

function MainMenuConstructor_Model(options){
	var id = 'MainMenuConstructor_Model';
	options = options || {};
	
	options.fields = {};
	
			
				
			
				
	
	var filed_options = {};
	filed_options.primaryKey = true;	
	filed_options.alias = 'Код';
	filed_options.autoInc = true;	
	
	options.fields.id = new FieldInt("id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.role_id = new FieldEnum("role_id",filed_options);
	filed_options.enumValues = 'admin,boss,accountant,chief_eng,supplier,constr_planner,constr_manager';
	options.fields.role_id.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.user_id = new FieldInt("user_id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Содержание';
	filed_options.autoInc = false;	
	
	options.fields.content = new FieldText("content",filed_options);
	options.fields.content.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Содержание для модели,заполняется при записи из контроллера!';
	filed_options.autoInc = false;	
	
	options.fields.model_content = new FieldText("model_content",filed_options);
	
									
		MainMenuConstructor_Model.superclass.constructor.call(this,id,options);
}
extend(MainMenuConstructor_Model,ModelXML);

