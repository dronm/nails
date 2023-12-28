/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function Studio_Form(options){
	options = options || {};	
	
	options.formName = "StudioDialog";
	options.controller = "Studio_Controller";
	options.method = "get_object";
	
	Studio_Form.superclass.constructor.call(this,options);
	
}
extend(Studio_Form,WindowFormObject);

/* Constants */


/* private members */

/* protected*/


/* public methods */

