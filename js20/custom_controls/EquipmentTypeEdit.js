/** Copyright (c) 2019 
 *	Andrey Mikhalevich, Katren ltd.
 */
function EquipmentTypeEdit(id,options){

	options = options || {};	
	if (options.labelCaption!=""){
		options.labelCaption = options.labelCaption || "Оборудование:";
	}
	options.cmdInsert = (options.cmdInsert!=undefined)? options.cmdInsert:false;
	
	options.keyIds = options.keyIds || ["id"];
	
	//форма выбора из списка
	options.selectWinClass = EquipmentTypeList_Form;
	options.selectDescrIds = options.selectDescrIds || ["name"];
	
	//форма редактирования элемента
	//options.editWinClass = EquipmentType_Form;
	
	options.acMinLengthForQuery = (options.acMinLengthForQuery!=undefined)? options.acMinLengthForQuery:1;
	options.acController = new EquipmentType_Controller();
	options.acPublicMethod = options.acController.getPublicMethod("complete");
	options.acModel = new EquipmentType_Model();
	options.acPatternFieldId = options.acPatternFieldId || "name";
	options.acKeyFields = options.acKeyFields || [options.acModel.getField("id")];
	options.acDescrFields = options.acDescrFields || [options.acModel.getField("name")];
	options.acICase = options.acICase || "1";
	options.acMid = options.acMid || "1";
	
	EquipmentTypeEdit.superclass.constructor.call(this,id,options);
}
extend(EquipmentTypeEdit,EditRef);

/* Constants */


/* private members */

/* protected*/


/* public methods */

