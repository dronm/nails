/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023

 * @extends View
 * @requires core/extend.js
 * @requires controls/View.js     

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {object} options
 */
function Sign_View(id,options){
	options = options || {};	
	
	options.addElement = function(){
		var self = this;
		/*this.addElement(new ButtonCtrl(id+":clear",{
			"caption":" Очистить",
			"glyph":"glyphicon-remove",
			"title":"Очистить подпись",
			"onClick": function(){
				self.m_context.clearRect(0, 0, self.m_canvas.width, self.m_canvas.height);
				if(self.m_canvasToPreviewTimer){
					clearTimeout(self.m_canvasToPreviewTimer);
				}
				
			},
		}));
		
		this.addElement(new Control(id+":preview", "img", {
			"attrs":{
				"height":"50px","width":"50px"
			}
		}));
		*/
		this.addElement(new ControlContainer(id+":signCont", "div", {
			"className":"signCont",
			"elements":[
				new Control(id+":signCont:sign", "canvas", {})
			]
		}));
		
	}
	
	Sign_View.superclass.constructor.call(this,id,options);		
}
//ViewObjectAjx,ViewAjxList
extend(Sign_View, View);

/* Constants */


/* private members */
Sign_View.prototype.m_canvas;
Sign_View.prototype.m_context;
Sign_View.prototype.m_x;
Sign_View.prototype.m_y;
Sign_View.prototype.m_isDrawing;

/* protected*/


/* public methods */

Sign_View.prototype.setCanvasSize = function(v){
	var cont = this.getElement("signCont");
	cont.m_node.style = "width:" + v +"px;";
	this.m_canvas.width = v;
	this.m_canvas.height = v;		
}

Sign_View.prototype.toDOM = function(p){
	Sign_View.superclass.toDOM.call(this, p);
	
	var cont = this.getElement("signCont");
	this.m_canvas = cont.getElement("sign").m_node;
	this.m_context = this.m_canvas.getContext("2d");	

	//events
	var self = this;
	this.m_canvas.addEventListener("mousedown", (e) => {
		self.m_x = e.offsetX;
		self.m_y = e.offsetY;
		self.setDrawing(true);
	});
	this.m_canvas.addEventListener("touchstart", (e) => {
		e.preventDefault();
		var touches = e.changedTouches;
		if(touches && touches.length){
			self.m_x = touches[0].pageX;
			self.m_y = touches[0].pageY;
			self.setDrawing(true);
		}
	});
	this.m_canvas.addEventListener("mousemove", (e) => {
		if (self.m_isDrawing) {
			self.drawLine(self.m_x, self.m_y, e.offsetX, e.offsetY);
			self.m_x = e.offsetX;
			self.m_y = e.offsetY;		
		}	
	});
	this.m_canvas.addEventListener("touchmove", (e) => {
		e.preventDefault();
		if (self.m_isDrawing) {
			var touches = e.changedTouches;
			if(touches && touches.length){
				self.drawLine(self.m_x, self.m_y, touches[0].pageX, touches[0].pageY);
				self.m_x = touches[0].pageX;
				self.m_y = touches[0].pageY;		
			}
		}		
	});

	window.addEventListener("mouseup", (e) => {
		if (self.m_isDrawing) {
			self.drawLine(self.m_x, self.m_y, e.offsetX, e.offsetY);
			self.m_x = 0;
			self.m_y = 0;
			self.setDrawing(false);
		}
	});

	window.addEventListener("touchend", (e) => {			
		if (self.m_isDrawing) {
			e.preventDefault();
			var touches = e.changedTouches;
			if(touches && touches.length){			
				self.drawLine(self.m_x, self.m_y, touches[0].pageX, touches[0].pageY);
				self.m_x = 0;
				self.m_y = 0;
				self.setDrawing(false);
			}
		}
	});
	
}

Sign_View.prototype.drawLine = function(x1, y1, x2, y2) {
	this.m_context.beginPath();
	this.m_context.strokeStyle = "#222222";
	this.m_context.lineWidth = 4;
	this.m_context.moveTo(x1, y1);
	this.m_context.lineTo(x2, y2);
	this.m_context.stroke();
	this.m_context.closePath();
}

Sign_View.prototype.setDrawing = function(drawing) {
	this.m_isDrawing = drawing;
	if(drawing){		
//console.log("Startig drawing")
		/*if(this.m_canvasToPreviewTimer){
			clearTimeout(this.m_canvasToPreviewTimer);
			this.m_canvasToPreviewTimer = undefined;
			console.log("timeout cleared")
		}
		var n = this.getElement("preview").getNode();
		DOMHelper.hide(n);
		n.setAttribute("src", "#");			
		*/
		
	//}else if (!drawing && !this.m_canvasToPreviewTimer){
		/*
		 var self = this; 
		 this.m_canvasToPreviewTimer = setTimeout(function(){
			var n = self.getElement("preview").getNode();
			n.setAttribute("src", self.m_canvas.toDataURL());	
			DOMHelper.show(n);
			console.log("timeout done")
		}, 2500);
		*/
	}
}


Sign_View.prototype.getSignatureFile = function() {
	if(!this.m_canvas){
		throw Error("canvas not found")
	}
	var img_tp = "image/png";
	var url = this.m_canvas.toDataURL(img_tp, 1.0); //image/png
	var blob_bin = atob(url.split(',')[1]);
	var array = [];
	for(var i = 0; i < blob_bin.length; i++) {
	    array.push(blob_bin.charCodeAt(i));
	}
	var file = new Blob([new Uint8Array(array)], {type: img_tp});
	return file;
}

