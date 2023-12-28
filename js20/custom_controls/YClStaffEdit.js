/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023
 
 * @class
 * @classdesc
	
 * @param {string} id view identifier
 * @param {namespace} options
 */	
function YClStaffEdit(id,options){
	options = options || {};
	options.model = new YClStaffList_Model();
	
	if (options.labelCaption!=""){
		options.labelCaption = options.labelCaption || "Мастер yclients:";
	}
	
	options.keyIds = options.keyIds || ["id"];
	options.modelKeyFields = [options.model.getField("id")];
	options.modelDescrFields = [options.model.getField("name")];
	
	var contr = new YClStaff_Controller();
	options.readPublicMethod = contr.getPublicMethod("get_list");
	
	YClStaffEdit.superclass.constructor.call(this,id,options);
	
}
extend(YClStaffEdit,EditSelectRef);

