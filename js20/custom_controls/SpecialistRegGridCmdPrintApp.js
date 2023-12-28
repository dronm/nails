
/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>,2023

 * @class
 * @classdesc
 
 * @requires core/extend.js  
 * @requires controls/GridCmd.js

 * @param {string} id Object identifier
 * @param {namespace} options
*/
function SpecialistRegGridCmdPrintApp(id,options){
	options = options || {};	

	options.showCmdControl = (options.showCmdControl!=undefined)? options.showCmdControl:true;
	
	var self = this;
	this.m_btn = new SpecialistRegBtnPrintApp(id+"btn",{
		"cmd":(options.cmd!= undefined)? options.cmd:true
	});
	
	options.controls = [
		this.m_btn
	]
	
	SpecialistRegGridCmdPrintApp.superclass.constructor.call(this,id,options);
		
}
extend(SpecialistRegGridCmdPrintApp,GridCmd);

/* Constants */


/* private members */

/* protected*/


/* public methods */
SpecialistRegGridCmdPrintApp.prototype.setGrid = function(v){
	SpecialistRegGridCmdPrintApp.superclass.setGrid.call(this,v);
	
	this.m_btn.m_grid = v;
	this.m_grid = v;
}

