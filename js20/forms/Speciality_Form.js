/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function Speciality_Form(options){
	options = options || {};	
	
	options.formName = "SpecialityDialog";
	options.controller = "Speciality_Controller";
	options.method = "get_object";
	
	Speciality_Form.superclass.constructor.call(this,options);
	
}
extend(Speciality_Form,WindowFormObject);

/* Constants */


/* private members */

/* protected*/


/* public methods */

