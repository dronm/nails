/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>,2017
 
 * @class
 * @classdesc Visual error control
 
 * @param {string} id
 * @param {Object} options
 * @param {string} [options.errorClassName=this.DEF_ERROR_CLASS]
 */
function ErrorControl(id,options){
	options = options || {};
	
	options.templateOptions = {
		"TEXT_CLASS":this.TEXT_CLASS,
		"PICT_CLASS":this.PICT_CLASS
	};
	options.visible = false;
	
	ErrorControl.superclass.constructor.call(this, id, options);
	
}
extend(ErrorControl,EditInfo);

/* constants */
ErrorControl.prototype.PICT_CLASS = "icon-cancel-circle2";
ErrorControl.prototype.TEXT_CLASS = "text-danger";

