/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023
 
 * @class
 * @classdesc
	
 * @param {string} id view identifier
 * @param {namespace} options
 */	
function SpecialityEdit(id,options){
	options = options || {};
	options.model = new SpecialityList_Model();
	
	if (options.labelCaption!=""){
		options.labelCaption = options.labelCaption || "Специальность:";
	}
	
	options.keyIds = options.keyIds || ["id"];
	options.modelKeyFields = [options.model.getField("id")];
	options.modelDescrFields = [options.model.getField("name")];
	
	var contr = new Speciality_Controller();
	options.readPublicMethod = contr.getPublicMethod("get_list");
	
	SpecialityEdit.superclass.constructor.call(this,id,options);
	
}
extend(SpecialityEdit,EditSelectRef);

