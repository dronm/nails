/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023

 * @extends ButtonCmd
 * @requires core/extend.js
 * @requires controls/ButtonCmd.js     

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {object} options
 */
function SpecialistRegBtnPrintApp(id,options){
	options = options || {};	
	
	options.caption = " Заявление";
	options.glyph = "glyphicon-print";
	options.title = "Печать заявления";
	
	var self = this;
	options.onClick = function(){
		self.print();
	}
	
	this.m_cmd = options.cmd;
	this.m_getId = options.getId;
	
	SpecialistRegBtnPrintApp.superclass.constructor.call(this,id,options);
}
//ViewObjectAjx,ViewAjxList
extend(SpecialistRegBtnPrintApp,ButtonCmd);

/* Constants */

/* private members */

/* protected*/


/* public methods */

SpecialistRegBtnPrintApp.prototype.print = function(){
	var id, name_full;
	if(this.m_cmd){
		this.m_grid.setModelToCurrentRow();
		id = this.m_grid.getModel().getFieldValue("id");
		name_full = this.m_grid.getModel().getFieldValue("name_full");
	}
	else{
		id = this.m_getId();
	}

	var offset = 0 * 20;
	var h = $( window ).width()/3*2;
	var left = $( window ).width()/2;
	var w = left - 20;	
	var pm = (new SpecialistReg_Controller()).getPublicMethod("print_app");
	pm.setFieldValue("id", id);
	window.getApp().openHrefDownload(window, pm, "ViewXML", "location=0,menubar=0,status=0,titlebar=0,top="+(50+offset)+",left="+(left+offset)+",width="+w+",height="+h);
}

