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

function Specialist_Model(options){
	var id = 'Specialist_Model';
	options = options || {};
	
	options.fields = {};
	
				
	
	var filed_options = {};
	filed_options.primaryKey = true;	
	
	filed_options.autoInc = true;	
	
	options.fields.id = new FieldInt("id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.name = new FieldText("name",filed_options);
	options.fields.name.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.inn = new FieldString("inn",filed_options);
	options.fields.inn.getValidator().setRequired(true);
	options.fields.inn.getValidator().setMaxLength('12');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'БИК банка';
	filed_options.autoInc = false;	
	
	options.fields.bank_bik = new FieldString("bank_bik",filed_options);
	options.fields.bank_bik.getValidator().setMaxLength('9');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Банковский счет';
	filed_options.autoInc = false;	
	
	options.fields.bank_acc = new FieldString("bank_acc",filed_options);
	options.fields.bank_acc.getValidator().setMaxLength('20');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.studio_id = new FieldInt("studio_id",filed_options);
	options.fields.studio_id.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Дата рождения';
	filed_options.autoInc = false;	
	
	options.fields.birthdate = new FieldDate("birthdate",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Адрес регистрации';
	filed_options.autoInc = false;	
	
	options.fields.address_reg = new FieldText("address_reg",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Паспорт';
	filed_options.autoInc = false;	
	
	options.fields.passport = new FieldJSON("passport",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Список оборудования для одного рабочего места';
	filed_options.autoInc = false;	
	
	options.fields.equipments = new FieldJSONB("equipments",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.user_id = new FieldInt("user_id",filed_options);
	options.fields.user_id.getValidator().setRequired(true);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.ycl_staff_id = new FieldInt("ycl_staff_id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.agent_percent = new FieldFloat("agent_percent",filed_options);
	options.fields.agent_percent.getValidator().setMaxLength('15');
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	
	filed_options.autoInc = false;	
	
	options.fields.speciality_id = new FieldInt("speciality_id",filed_options);
	
			
			
		Specialist_Model.superclass.constructor.call(this,id,options);
}
extend(Specialist_Model,ModelXML);

