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

function Enum_notif_types(id,options){
	options = options || {};
	options.addNotSelected = (options.addNotSelected!=undefined)? options.addNotSelected:true;
	options.options = [{"value":"new_specialist",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"new_specialist"],
checked:(options.defaultValue&&options.defaultValue=="new_specialist")}
	, {"value":"tel_check",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"tel_check"],
checked:(options.defaultValue&&options.defaultValue=="tel_check")}
	, {"value":"email_check",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"email_check"],
checked:(options.defaultValue&&options.defaultValue=="email_check")}
	, {"value":"docs_for_sign",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"docs_for_sign"],
checked:(options.defaultValue&&options.defaultValue=="docs_for_sign")}
	, {"value":"signed_docs",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"signed_docs"],
checked:(options.defaultValue&&options.defaultValue=="signed_docs")}
	, {"value":"new_account",
"descr":this.multyLangValues[window.getApp().getLocale()+"_"+"new_account"],
checked:(options.defaultValue&&options.defaultValue=="new_account")}
	];
	
	Enum_notif_types.superclass.constructor.call(this, id, options);
	
}
extend(Enum_notif_types,EditSelect);

Enum_notif_types.prototype.multyLangValues = {
	"ru_new_specialist":"Добавлен самозанятый", "ru_tel_check":"Проверка телефона", "ru_email_check":"Проверка электронной почты", "ru_docs_for_sign":"Документы для подписания", "ru_signed_docs":"Подписанные документы", "ru_new_account":"Новый аккаунт"
};


