
/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>,2023

 * @class
 * @classdesc
 
 * @requires core/extend.js  
 * @requires controls/GridCmd.js

 * @param {string} id Object identifier
 * @param {namespace} options
*/
function SpecialistPeriodSalaryDetailForPayGridCmdMakeDocs(id,options){
	options = options || {};	

	options.showCmdControl = (options.showCmdControl!=undefined)? options.showCmdControl:true;
	
	options.glyph = "glyphicon-refresh";
	options.caption = " п/п в банк";
	options.title = "Сформировать платежные поручения";
	SpecialistPeriodSalaryDetailForPayGridCmdMakeDocs.superclass.constructor.call(this,id,options);
		
}
extend(SpecialistPeriodSalaryDetailForPayGridCmdMakeDocs, GridCmd);

/* Constants */


/* private members */

/* protected*/


/* public methods */
SpecialistPeriodSalaryDetailForPayGridCmdMakeDocs.prototype.onCommand = function(){
	var mark = this.m_grid.getHead().getElement("row0").getElement("mark");
	if(!mark){
		return;
	}	
	var keys = mark.getSelectedKeys();	
	if(!keys || !keys.length){
		return;
	}
	
	var self = this;
	WindowQuestion.show({
		"no":false,
		"callBackYes":function(){
			self.onCommandCont(keys);
		},
		"text":"Сформировать п/п на оплаты (" +  keys.length +")?"
	});
	
}

SpecialistPeriodSalaryDetailForPayGridCmdMakeDocs.prototype.onCommandCont = function(keys){
	var pm = (new SpecialistPeriodSalaryDetail_Controller()).getPublicMethod("documents_to_bank");
	pm.setFieldValue("ids", CommonHelper.serialize(keys));
	var self = this;
	pm.download("ViewXML", 0, function(){
		self.m_grid.onRefresh(function(){
			window.showTempOk("Сформирован файл выгрузки для банка", null, 1000);	
		});		
	});
}
