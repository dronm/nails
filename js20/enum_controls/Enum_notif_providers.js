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

function Enum_notif_providers(id,options){
	options = options || {};
	options.addNotSelected = (options.addNotSelected!=undefined)? options.addNotSelected:true;
	options.options = [{"value":"email",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"email"],
checked:(options.defaultValue&&options.defaultValue=="email")}
	, {"value":"sms",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"sms"],
checked:(options.defaultValue&&options.defaultValue=="sms")}
	, {"value":"wa",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"wa"],
checked:(options.defaultValue&&options.defaultValue=="wa")}
	, {"value":"tm",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"tm"],
checked:(options.defaultValue&&options.defaultValue=="tm")}
	, {"value":"vb",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"vb"],
checked:(options.defaultValue&&options.defaultValue=="vb")}
	];
	
	Enum_notif_providers.superclass.constructor.call(this, id, options);
	
}
extend(Enum_notif_providers,EditSelect);

Enum_notif_providers.prototype.multyLangValues = {
	"ru_email":"Электронная почта", "ru_sms":"СМС", "ru_wa":"WhatsUp", "ru_tm":"Telegram", "ru_vb":"Viber"
};


