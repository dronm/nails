/*
 * Copyright (c) 2023
 * Andrey Mikhalevich, Katren ltd.
 */
function EquipmentTypeList_Form(options){
	options = options || {};	
	
	options.formName = "EquipmentTypeList";
	options.controller = "EquipmentType_Controller";
	options.method = "get_list";
	
	EquipmentTypeList_Form.superclass.constructor.call(this,options);
		
}
extend(EquipmentTypeList_Form,WindowFormObject);

/* Constants */


/* private members */

/* protected*/


/* public methods */

