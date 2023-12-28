/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2022
 * @class
 * @classdesc Grid column Enumerator class. Created from build/templates/enumGridColumn.js.tmpl !!!DO NOT MODIFY!!!
 
 * @extends GridColumnEnum
 
 * @requires core/extend.js
 * @requires controls/GridColumnEnum.js
 
 * @param {object} options
 */

function EnumGridColumn_notif_types(options){
	options = options || {};
	
	options.multyLangValues = {};
	options.multyLangValues["ru"] = {};
	options.multyLangValues["ru"]["new_specialist"] = "Добавлен самозанятый";
	
	options.multyLangValues["ru"]["tel_check"] = "Проверка телефона";
	
	options.multyLangValues["ru"]["email_check"] = "Проверка электронной почты";
	
	options.multyLangValues["ru"]["docs_for_sign"] = "Документы для подписания";
	
	options.multyLangValues["ru"]["signed_docs"] = "Подписанные документы";
	
	options.multyLangValues["ru"]["new_account"] = "Новый аккаунт";
	
	options.ctrlClass = options.ctrlClass || Enum_notif_types;
	options.searchOptions = options.searchOptions || {};
	options.searchOptions.searchType = options.searchOptions.searchType || "on_match";
	options.searchOptions.typeChange = (options.searchOptions.typeChange!=undefined)? options.searchOptions.typeChange:false;
	
	EnumGridColumn_notif_types.superclass.constructor.call(this,options);		
}
extend(EnumGridColumn_notif_types, GridColumnEnum);

