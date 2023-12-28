/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function SpecialistPeriodSalaryDetailEdit(id,options){

	options = options || {};	
	if (options.labelCaption!=""){
		options.labelCaption = options.labelCaption || "Начисление з/п:";
	}
	options.cmdInsert = (options.cmdInsert!=undefined)? options.cmdInsert:true;
	
	options.keyIds = options.keyIds || ["id"];
	
	//форма выбора из списка
	options.selectWinClass = SpecialistPeriodSalaryDetailList_Form;
	options.selectDescrIds = options.selectDescrIds || ["descr"];
	
	//форма редактирования элемента
	options.editWinClass = null;
	
	options.acMinLengthForQuery = (options.acMinLengthForQuery!=undefined)? options.acMinLengthForQuery:1;
	options.acController = new SpecialistPeriodSalaryDetail_Controller();
	options.acPublicMethod = options.acController.getPublicMethod("complete");
	options.acModel = new SpecialistPeriodSalaryDetailList_Model();
	options.acPatternFieldId = options.acPatternFieldId || "descr";
	options.acKeyFields = options.acKeyFields || [options.acModel.getField("id")];
	options.acDescrFields = options.acDescrFields || [options.acModel.getField("descr")];
	options.acICase = options.acICase || "1";
	options.acMid = options.acMid || "1";
	
	SpecialistPeriodSalaryDetailEdit.superclass.constructor.call(this,id,options);
}
extend(SpecialistPeriodSalaryDetailEdit,EditRef);

/* Constants */


/* private members */

/* protected*/


/* public methods */

