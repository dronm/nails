/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2022
 * @class
 * @classdesc Grid column Enumerator class. Created from build/templates/enumGridColumn.js.tmpl !!!DO NOT MODIFY!!!
 
 * @extends GridColumnEnum
 
 * @requires core/extend.js
 * @requires controls/GridColumnEnum.js
 
 * @param {object} options
 */

function EnumGridColumn_specialist_status_types(options){
	options = options || {};
	
	options.multyLangValues = {};
	options.multyLangValues["ru"] = {};
	options.multyLangValues["ru"]["contract_signing"] = "Подписание контракта";
	
	options.multyLangValues["ru"]["contract_signed"] = "Заключен контракт";
	
	options.multyLangValues["ru"]["contract_terminated"] = "Расторгнут контракт";
	
	options.ctrlClass = options.ctrlClass || Enum_specialist_status_types;
	options.searchOptions = options.searchOptions || {};
	options.searchOptions.searchType = options.searchOptions.searchType || "on_match";
	options.searchOptions.typeChange = (options.searchOptions.typeChange!=undefined)? options.searchOptions.typeChange:false;
	
	EnumGridColumn_specialist_status_types.superclass.constructor.call(this,options);		
}
extend(EnumGridColumn_specialist_status_types, GridColumnEnum);

