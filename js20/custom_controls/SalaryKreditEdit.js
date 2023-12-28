/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023
 
 * @class
 * @classdesc
	
 * @param {string} id view identifier
 * @param {namespace} options
 */	
function SalaryKreditEdit(id,options){
	options = options || {};
	options.model = new SalaryKredit_Model();
	
	if (options.labelCaption!=""){
		options.labelCaption = options.labelCaption || "Удержание:";
	}
	
	options.keyIds = options.keyIds || ["id"];
	options.modelKeyFields = [options.model.getField("id")];
	options.modelDescrFields = [options.model.getField("name")];
	
	var contr = new SalaryKredit_Controller();
	options.readPublicMethod = contr.getPublicMethod("get_list");
	
	SalaryKreditEdit.superclass.constructor.call(this,id,options);
	
}
extend(SalaryKreditEdit,EditSelectRef);

