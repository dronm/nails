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
function SpecialistRegRegister_View(id,options){
	options = options || {};	
	
	options.addElement = function(){
		this.addElement(new StudioEdit(id+":studio",{
			"required":true
		}));	
		this.addElement(new YClStaffEdit(id+":ycl_staff",{
			"required":true
		}));	
		
	}
	
	SpecialistRegRegister_View.superclass.constructor.call(this,id,options);
}
//ViewObjectAjx,ViewAjxList
extend(SpecialistRegRegister_View, View);

/* Constants */


/* private members */

/* protected*/


/* public methods */

