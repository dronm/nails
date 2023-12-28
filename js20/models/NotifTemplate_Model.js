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

function NotifTemplate_Model(options){
	var id = 'NotifTemplate_Model';
	options = options || {};
	
	options.fields = {};
	
			
				
				
			
				
	
	var filed_options = {};
	filed_options.primaryKey = true;	
	
	filed_options.autoInc = true;	
	
	options.fields.id = new FieldInt("id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Способ доставки';
	filed_options.autoInc = false;	
	
	options.fields.notif_provider = new FieldEnum("notif_provider",filed_options);
	filed_options.enumValues = 'email,sms,wa,tm,vb';
	options.fields.notif_provider.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Тип сообщения';
	filed_options.autoInc = false;	
	
	options.fields.notif_type = new FieldEnum("notif_type",filed_options);
	filed_options.enumValues = 'new_specialist';
	options.fields.notif_type.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Шаблон';
	filed_options.autoInc = false;	
	
	options.fields.template = new FieldText("template",filed_options);
	options.fields.template.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Комментарий';
	filed_options.autoInc = false;	
	
	options.fields.comment_text = new FieldText("comment_text",filed_options);
	options.fields.comment_text.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Поля';
	filed_options.autoInc = false;	
	
	options.fields.fields = new FieldJSON("fields",filed_options);
	options.fields.fields.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Специфичные для провайдера поля, например subject для писем, фичи оформления для месенджеров';
	filed_options.autoInc = false;	
	
	options.fields.provider_values = new FieldJSON("provider_values",filed_options);
	
			
		NotifTemplate_Model.superclass.constructor.call(this,id,options);
}
extend(NotifTemplate_Model,ModelXML);

