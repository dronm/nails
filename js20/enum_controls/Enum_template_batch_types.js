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

function Enum_template_batch_types(id,options){
	options = options || {};
	options.addNotSelected = (options.addNotSelected!=undefined)? options.addNotSelected:true;
	options.options = [{"value":"specialist_registration",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"specialist_registration"],
checked:(options.defaultValue&&options.defaultValue=="specialist_registration")}
	, {"value":"specialist_salary",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"specialist_salary"],
checked:(options.defaultValue&&options.defaultValue=="specialist_salary")}
	];
	
	Enum_template_batch_types.superclass.constructor.call(this, id, options);
	
}
extend(Enum_template_batch_types,EditSelect);

Enum_template_batch_types.prototype.multyLangValues = {
	"ru_specialist_registration":"Регистрация мастера", "ru_specialist_salary":"Расчет зарплаты"
};


