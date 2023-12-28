/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023

 * @extends EditDate
 * @requires core/extend.js
 * @requires controls/EditDate.js     

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {object} options
 */
function EditDateInlineValidation(id,options){
	options = options || {};	
	
	options.required = (options.required!=undefined)? options.required : true;
	options.labelCaption = (options.labelCaption!=undefined)? options.labelCaption : "Дата:";
	options.placeholder = options.placeholder || "Дата 01/01/2023";
	
	options.editMask = "";
	options.formatterOptions = {"date": true, "datePattern":['d', 'm', 'Y']};
	
	options.regExpression = new RegExp(/^\d{2}(\/)\d{2}\1\d{4}$/);
	options.regExpressionInvalidMessage = "Неверный формат";
	
	options.events = options.events || {
		"focus":function(){
			this.setValid();
		}
		,"blur": function(){
			this.validate();
		}
	};
	
	EditDateInlineValidation.superclass.constructor.call(this, id, options);
}
//ViewObjectAjx,ViewAjxList
extend(EditDateInlineValidation, EditDate);

/* Constants */


/* private members */

/* protected*/


/* public methods */

