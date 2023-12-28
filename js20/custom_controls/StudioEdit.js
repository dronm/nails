/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023
 
 * @class
 * @classdesc
	
 * @param {string} id view identifier
 * @param {namespace} options
 */	
function StudioEdit(id,options){
	options = options || {};
	options.model = new StudioList_Model();
	
	if (options.labelCaption!=""){
		options.labelCaption = options.labelCaption || "Салон:";
	}
	options.title = "Салон";
	
	options.keyIds = options.keyIds || ["id"];
	options.modelKeyFields = [options.model.getField("id")];
	options.modelDescrFields = [options.model.getField("name")];
	
	var contr = new Studio_Controller();
	options.readPublicMethod = contr.getPublicMethod("get_list");
	
	StudioEdit.superclass.constructor.call(this,id,options);
	
}
extend(StudioEdit,EditSelectRef);

