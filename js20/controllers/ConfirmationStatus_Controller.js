/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2017
 
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller_js20.xsl
 * ALL DIRECT MODIFICATIONS WILL BE LOST WITH THE NEXT BUILD PROCESS!!!
 
 * @class
 * @classdesc controller
 
 * @extends ControllerObjServer
  
 * @requires core/extend.js
 * @requires core/ControllerObjServer.js
  
 * @param {Object} options
 * @param {Model} options.listModelClass
 * @param {Model} options.objModelClass
 */ 

function ConfirmationStatus_Controller(options){
	options = options || {};
	ConfirmationStatus_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.add_set_confirmed();
		
}
extend(ConfirmationStatus_Controller,ControllerObjServer);

			ConfirmationStatus_Controller.prototype.add_set_confirmed = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('set_confirmed',opts);
	
				
	
	var options = {};
	
		options.required = true;
	
		options.maxlength = "36";
	
		pm.addField(new FieldString("secret",options));
	
			
	this.addPublicMethod(pm);
}

		