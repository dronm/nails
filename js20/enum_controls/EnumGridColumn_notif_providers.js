/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2022
 * @class
 * @classdesc Grid column Enumerator class. Created from build/templates/enumGridColumn.js.tmpl !!!DO NOT MODIFY!!!
 
 * @extends GridColumnEnum
 
 * @requires core/extend.js
 * @requires controls/GridColumnEnum.js
 
 * @param {object} options
 */

function EnumGridColumn_notif_providers(options){
	options = options || {};
	
	options.multyLangValues = {};
	options.multyLangValues["ru"] = {};
	options.multyLangValues["ru"]["email"] = "Электронная почта";
	
	options.multyLangValues["ru"]["sms"] = "СМС";
	
	options.multyLangValues["ru"]["wa"] = "WhatsUp";
	
	options.multyLangValues["ru"]["tm"] = "Telegram";
	
	options.multyLangValues["ru"]["vb"] = "Viber";
	
	options.ctrlClass = options.ctrlClass || Enum_notif_providers;
	options.searchOptions = options.searchOptions || {};
	options.searchOptions.searchType = options.searchOptions.searchType || "on_match";
	options.searchOptions.typeChange = (options.searchOptions.typeChange!=undefined)? options.searchOptions.typeChange:false;
	
	EnumGridColumn_notif_providers.superclass.constructor.call(this,options);		
}
extend(EnumGridColumn_notif_providers, GridColumnEnum);

