/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2022

 * @extends View
 * @requires core/extend.js
 * @requires controls/View.js     

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {object} options
 */
function Mgateway_View(id,options){
	options = options || {};	
	
	var self = this;
	options.addElement = function(){
		this.addElement(new ButtonCmd(id+":showQR", {
			"caption": "Показать QR код",
			"title": "Открыть QR для сканирования",
			"onClick":function(){
				self.showQR()
			}
		}));
	}
	
	Mgateway_View.superclass.constructor.call(this,id,options);
}
//ViewObjectAjx,ViewAjxList
extend(Mgateway_View,View);

/* Constants */


/* private members */

/* protected*/


/* public methods */
Mgateway_View.prototype.showQR = function(){
	var pm = (new Mgateway_Controller()).getPublicMethod("get_qr");
	var offset = 0;
	var h = $( window ).width()/3*2;
	var left = $( window ).width()/2;
	var w = left - 20;	
	window.getApp().openHrefDownload(window, pm, "ViewXML", "noopener=1,location=0,menubar=0,status=0,titlebar=0,top="+(50+offset)+",left="+(left+offset)+",width="+w+",height="+h);	
}

