/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function NotifTemplate_Form(options){
	options = options || {};	
	
	options.formName = "NotifTemplate";
	options.controller = "NotifTemplate_Controller";
	options.method = "get_object";
	
	NotifTemplate_Form.superclass.constructor.call(this,options);
	
}
extend(NotifTemplate_Form,WindowFormObject);

/* Constants */


/* private members */

/* protected*/


/* public methods */

