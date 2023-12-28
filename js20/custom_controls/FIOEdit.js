/** Copyright (c) 2023
	Andrey Mikhalevich, Katren ltd.
 */
function FIOEdit(id,options){
	options = options || {};
	options.labelCaption = (options.labelCaption!=undefined)? options.labelCaption : "ФИО:";
	options.maxLength = (options.maxLength==undefined)? options.maxLength : 200;
	options.regExpression = "^([А-ЯA-Z]|[А-ЯA-Z][\x27а-яa-z]{1,}|[А-ЯA-Z][\x27а-яa-z]{1,}\-([А-ЯA-Z][\x27а-яa-z]{1,}|(оглы)|(кызы)))\040[А-ЯA-Z][\x27а-яa-z]{1,}(\040[А-ЯA-Z][\x27а-яa-z]{1,})?$";
	
	FIOEdit.superclass.constructor.call(this,id,options);	
}
extend(FIOEdit, EditString);
