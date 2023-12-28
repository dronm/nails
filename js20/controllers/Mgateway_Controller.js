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

function Mgateway_Controller(options){
	options = options || {};
	Mgateway_Controller.superclass.constructor.call(this,options);	
	
	//methods
	this.add_callback();
	this.add_get_qr();
		
}
extend(Mgateway_Controller,ControllerObjServer);

			Mgateway_Controller.prototype.add_callback = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('callback',opts);
	
	this.addPublicMethod(pm);
}

			Mgateway_Controller.prototype.add_get_qr = function(){
	var opts = {"controller":this};	
	var pm = new PublicMethodServer('get_qr',opts);
	
	this.addPublicMethod(pm);
}

		