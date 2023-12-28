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

Sign_View.prototype.m_mousePos;
Sign_View.prototype.m_lastPos;

Sign_View.prototype.m_drawing;

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

	this.m_canvas.strokeStyle = "#222222";
	this.m_canvas.lineWidth = 4;
	
	this.m_mousePos = {x: 0, y: 0};
	this.m_lastPos = this.m_mousePos;
	
	window.requestAnimFrame = (function(callback) {
		return window.requestAnimationFrame ||
			window.webkitRequestAnimationFrame ||
			window.mozRequestAnimationFrame ||
			window.oRequestAnimationFrame ||
			window.msRequestAnimaitonFrame ||
			function(callback) {
				window.setTimeout(callback, 1000 / 60);
			};
      	});
      	
	//events
	var self = this;
	this.m_canvas.addEventListener("mousedown", function(e) {
		self.m_drawing = true;
		self.m_lastPos = getMousePos(self.m_canvas, e);
	}, false);	

	this.m_canvas.addEventListener("mouseup", function(e) {
		self.m_drawing = false;
	}, false);
	
	this.m_canvas.addEventListener("mousemove", function(e) {
		self.m_mousePos = getMousePos(self.m_canvas, e);
	}, false);
	
	// Add touch event support for mobile
	this.m_canvas.addEventListener("touchstart", function(e) {

	}, false);
	
	this.m_canvas.addEventListener("touchmove", function(e) {
		var touch = e.touches[0];
		var me = new MouseEvent("mousemove", {
			clientX: touch.clientX,
			clientY: touch.clientY
		});
		self.m_canvas.dispatchEvent(me);
	}, false);
	
	this.m_canvas.addEventListener("touchstart", function(e) {
		self.m_mousePos = getTouchPos(self.m_canvas, e);
		var touch = e.touches[0];
		var me = new MouseEvent("mousedown", {
			clientX: touch.clientX,
			clientY: touch.clientY
		});
		self.m_canvas.dispatchEvent(me);
	}, false);
	
	this.m_canvas.addEventListener("touchend", function(e) {
		var me = new MouseEvent("mouseup", {});
		self.m_canvas.dispatchEvent(me);
	}, false);
	
	// Prevent scrolling when touching the canvas
	document.body.addEventListener("touchstart", function(e) {
		if (e.target == canvas) {
			e.preventDefault();
		}
	}, false);
	
	document.body.addEventListener("touchend", function(e) {
		if (e.target == canvas) {
			e.preventDefault();
		}
	}, false);
	
	document.body.addEventListener("touchmove", function(e) {
		if (e.target == canvas) {
			e.preventDefault();
		}
	}, false);
	
	(function drawLoop() {
		requestAnimFrame(drawLoop);
		renderCanvas();
	})();								
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

//******************
function getMousePos(canvasDom, mouseEvent) {
	var rect = canvasDom.getBoundingClientRect();
	return {
		x: mouseEvent.clientX - rect.left,
		y: mouseEvent.clientY - rect.top
	}
}

function renderCanvas() {
	if (this.m_drawing) {
		this.m_context.moveTo(this.m_lastPos.x, this.m_lastPos.y);
		this.m_context.lineTo(this.m_mousePos.x, this.m_mousePos.y);
		this.m_context.stroke();
		this.m_lastPos = this.m_mousePos;
	}
}

function getTouchPos(canvasDom, touchEvent) {
	var rect = canvasDom.getBoundingClientRect();
	return {
		x: touchEvent.touches[0].clientX - rect.left,
		y: touchEvent.touches[0].clientY - rect.top
	}
}

function clearCanvas() {
	canvas.width = canvas.width;
}

