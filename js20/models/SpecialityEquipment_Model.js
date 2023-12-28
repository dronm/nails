/**	
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/models/Model_js.xsl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 *
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2017
 * @class
 * @classdesc Model class. Created from template build/templates/models/Model_js.xsl. !!!DO NOT MODEFY!!!
 
 * @extends ModelJSON
 
 * @requires core/extend.js
 * @requires core/ModelJSON.js
 
 * @param {string} id 
 * @param {Object} options
 */

function SpecialityEquipment_Model(options){
	var id = 'SpecialityEquipment_Model';
	options = options || {};
	
	options.fields = {};
	
				
	
	var filed_options = {};
	filed_options.primaryKey = true;	
	
	filed_options.autoInc = true;	
	
	options.fields.id = new FieldInt("id",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Вид оборудования';
	filed_options.autoInc = false;	
	
	options.fields.equipment_types_ref = new FieldJSONB("equipment_types_ref",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Количество';
	filed_options.autoInc = false;	
	
	options.fields.quant = new FieldInt("quant",filed_options);
	
				
	
	var filed_options = {};
	filed_options.primaryKey = false;	
	filed_options.alias = 'Едииница';
	filed_options.autoInc = false;	
	
	options.fields.measure_unit = new FieldString("measure_unit",filed_options);
	
		SpecialityEquipment_Model.superclass.constructor.call(this,id,options);
}
extend(SpecialityEquipment_Model,ModelJSON);

