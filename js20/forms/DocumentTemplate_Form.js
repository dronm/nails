/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function DocumentTemplate_Form(options){
	options = options || {};	
	
	options.formName = "DocumentTemplate";
	options.controller = "DocumentTemplate_Controller";
	options.method = "get_object";
	
	DocumentTemplate_Form.superclass.constructor.call(this,options);
	
}
extend(DocumentTemplate_Form,WindowFormObject);

/* Constants */


/* private members */

/* protected*/


/* public methods */

