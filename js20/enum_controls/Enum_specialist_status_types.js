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

function Enum_specialist_status_types(id,options){
	options = options || {};
	options.addNotSelected = (options.addNotSelected!=undefined)? options.addNotSelected:true;
	options.options = [{"value":"contract_signing",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"contract_signing"],
checked:(options.defaultValue&&options.defaultValue=="contract_signing")}
	, {"value":"contract_signed",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"contract_signed"],
checked:(options.defaultValue&&options.defaultValue=="contract_signed")}
	, {"value":"contract_terminated",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"contract_terminated"],
checked:(options.defaultValue&&options.defaultValue=="contract_terminated")}
	];
	
	Enum_specialist_status_types.superclass.constructor.call(this, id, options);
	
}
extend(Enum_specialist_status_types,EditSelect);

Enum_specialist_status_types.prototype.multyLangValues = {
	"ru_contract_signing":"Подписание контракта", "ru_contract_signed":"Заключен контракт", "ru_contract_terminated":"Расторгнут контракт"
};


