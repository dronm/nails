/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function Specialist_Form(options){
	options = options || {};	
	
	options.formName = "SpecialistDialog";
	options.controller = "Specialist_Controller";
	options.method = "get_object";
	
	Specialist_Form.superclass.constructor.call(this,options);
	
}
extend(Specialist_Form,WindowFormObject);

/* Constants */


/* private members */

/* protected*/


/* public methods */

