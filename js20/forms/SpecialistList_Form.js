/*
 * Copyright (c) 2023
 * Andrey Mikhalevich, Katren ltd.
 */
function SpecialistList_Form(options){
	options = options || {};	
	
	options.formName = "SpecialistList";
	options.controller = "Specialist_Controller";
	options.method = "get_list";
	
	SpecialistList_Form.superclass.constructor.call(this,options);
		
}
extend(SpecialistList_Form, WindowFormObject);

/* Constants */


/* private members */

/* protected*/


/* public methods */


