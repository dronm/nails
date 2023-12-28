/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023

 * @extends View
 * @requires core/extend.js
 * @requires controls/View.js     

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {object} options
 */
function Passport_View(id,options){
	options = options || {};	
	
	options.addElement = function(){
		this.addElement(new EditNum(id+":series", {
			"labelCaption": "Серия:",
			"cmdSmClear":true,
			"placeholder":"Серия",
			"title":"Серия паспорта",
			"maxLength":"4",
			"fixLength":true,
			"required":true,
			"focus":true
		}));
		
		this.addElement(new EditNum(id+":num", {
			"labelCaption": "Номер:",
			"cmdSmClear":true,
			"placeholder":"Номер",
			"title":"Номер паспорта",
			"maxLength":"6",
			"fixLength":true,
			"required":true
		}));
		this.addElement(new EditString(id+":issue_body", {
			"labelCaption": "Выдан:",
			"cmdSmClear":true,
			"maxLength":"500",
			"placeholder":"Кто выдал документ",
			"title":"Кто выдал документ",
			"required":true
		}));
		this.addElement(new EditDateInlineValidation(id+":issue_date", {
			"labelCaption": "Дата:",
			"title":"Дата выдачи документа",
			"placeholder":"Дата выдачи документа",
			"required":true
		}));
		
		this.addElement(new EditNum(id+":dep_code", {
			"labelCaption": "Код:",
			"cmdSmClear":true,
			"placeholder":"Код",
			"title":"Код подразделения",
			"maxLength":"6",
			"required":true,
			"fixLength":true
		}));
		
	}
	
	Passport_View.superclass.constructor.call(this,id,options);
}
//ViewObjectAjx,ViewAjxList
extend(Passport_View, EditJSON);

/* Constants */


/* private members */

/* protected*/


/* public methods */

