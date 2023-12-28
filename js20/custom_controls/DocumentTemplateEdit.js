/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023
 
 * @class
 * @classdesc
	
 * @param {string} id view identifier
 * @param {namespace} options
 */	
function DocumentTemplateEdit(id,options){
	options = options || {};
	options.model = new DocumentTemplateList_Model();
	
	if (options.labelCaption!=""){
		options.labelCaption = options.labelCaption || "Шаблон:";
	}
	
	options.keyIds = options.keyIds || ["id"];
	options.modelKeyFields = [options.model.getField("id")];
	options.modelDescrFields = [options.model.getField("name")];
	
	var contr = new DocumentTemplate_Controller();
	options.readPublicMethod = contr.getPublicMethod("get_list");
	
	DocumentTemplateEdit.superclass.constructor.call(this,id,options);
	
}
extend(DocumentTemplateEdit,EditSelectRef);

