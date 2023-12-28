
/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>,2023

 * @class
 * @classdesc
 
 * @requires core/extend.js  
 * @requires controls/GridCmd.js

 * @param {string} id Object identifier
 * @param {namespace} options
*/
function SpecialistDocumentGridCmdSign(id,options){
	options = options || {};	

	options.showCmdControl = (options.showCmdControl!=undefined)? options.showCmdControl:true;
	
	var self = this;
	this.m_btn = new SpecialistDocumentBtnSign(id+"btn",{
		"cmd":(options.cmd!= undefined)? options.cmd:true
	});
	
	options.controls = [
		this.m_btn
	]
	
	SpecialistDocumentGridCmdSign.superclass.constructor.call(this,id,options);
		
}
extend(SpecialistDocumentGridCmdSign,GridCmd);

/* Constants */


/* private members */

/* protected*/


/* public methods */
SpecialistDocumentGridCmdSign.prototype.setGrid = function(v){
	SpecialistDocumentGridCmdSign.superclass.setGrid.call(this,v);
	
	this.m_btn.m_grid = v;
	this.m_grid = v;
}

