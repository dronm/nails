/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2022
 * @class
 * @classdesc Enumerator class. Created from template build/templates/enum.js.tmpl. !!!DO NOT MODIFY!!!
 
 * @extends EditSelect
 
 * @requires core/extend.js
 * @requires controls/EditSelect.js
 
 * @param string id 
 * @param {object} options
 */

function Enum_locales(id,options){
	options = options || {};
	options.addNotSelected = (options.addNotSelected!=undefined)? options.addNotSelected:true;
	options.options = [{"value":"ru",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"ru"],
checked:(options.defaultValue&&options.defaultValue=="ru")}
	];
	
	Enum_locales.superclass.constructor.call(this, id, options);
	
}
extend(Enum_locales,EditSelect);

Enum_locales.prototype.multyLangValues = {
	"ru_ru":"Русский"
};


