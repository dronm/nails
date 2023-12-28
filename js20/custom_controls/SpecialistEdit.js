/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function SpecialistEdit(id,options){

	options = options || {};	
	if (options.labelCaption!=""){
		options.labelCaption = options.labelCaption || "Сотрудник:";
	}
	options.cmdInsert = (options.cmdInsert!=undefined)? options.cmdInsert:false;
	
	options.keyIds = options.keyIds || ["id"];
	
	//форма выбора из списка
	options.selectWinClass = SpecialistList_Form;
	options.selectDescrIds = options.selectDescrIds || ["name"];
	
	//форма редактирования элемента
	options.editWinClass = Specialist_Form;
	
	options.acMinLengthForQuery = (options.acMinLengthForQuery!=undefined)? options.acMinLengthForQuery:1;
	options.acController = new Specialist_Controller();
	options.acPublicMethod = options.acController.getPublicMethod("complete");
	options.acModel = new SpecialistList_Model();
	options.acPatternFieldId = options.acPatternFieldId || "name";
	options.acKeyFields = options.acKeyFields || [options.acModel.getField("id")];
	options.acDescrFields = options.acDescrFields || [options.acModel.getField("name")];
	options.acICase = options.acICase || "1";
	options.acMid = options.acMid || "1";
	
	SpecialistEdit.superclass.constructor.call(this,id,options);
}
extend(SpecialistEdit,EditRef);

/* Constants */


/* private members */

/* protected*/


/* public methods */

