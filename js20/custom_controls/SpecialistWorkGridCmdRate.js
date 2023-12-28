
/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>,2023

 * @class
 * @classdesc
 
 * @requires core/extend.js  
 * @requires controls/GridCmd.js

 * @param {string} id Object identifier
 * @param {namespace} options
*/
function SpecialistWorkGridCmdRate(id,options){
	options = options || {};	

	options.showCmdControl = (options.showCmdControl!=undefined)? options.showCmdControl:true;
	
	options.glyph = "glyphicon-star-empty";
	options.caption = " Оценить";
	SpecialistWorkGridCmdRate.superclass.constructor.call(this,id,options);
		
}
extend(SpecialistWorkGridCmdRate, GridCmd);

/* Constants */


/* private members */

/* protected*/


/* public methods */
SpecialistWorkGridCmdRate.prototype.onCommand = function(){
	//this.m_grid.setModelToCurrentRow();
	//id = this.m_grid.getModel().getFieldValue("id");
	
	var self = this;
	(new WindowFormModalBS("adminRate",{
		"dialogWidth":"60%",
		"cmdOk":true,
		"cmdCancel":false,		
		"onClickOk":function(){
			if(!this.getContent().m_updates.length){
				this.close();
			}
			var dlg = this;
			var pm = (new SpecialistWork_Controller()).getPublicMethod("set_admin_rates");
			pm.setFieldValue("rates", CommonHelper.serialize(this.getContent().m_updates));
			pm.run({
				"ok":function(){
					dlg.close();
					self.m_grid.onRefresh();
				}
			});
		},
		"cmdClose":true,
		"content":new AdminRate_View("adminRate:view")
	})).open();	
}

