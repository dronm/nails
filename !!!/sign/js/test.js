let isDrawing = false;
let x = 0;
let y = 0;

const canvas = document.getElementById("sign");
canvas.width = document.getElementById('sign_cont').clientWidth;
canvas.height = document.getElementById('sign_cont').clientHeight;

const context = canvas.getContext("2d");

canvas.addEventListener("mousedown", (e) => {
	x = e.offsetX;
	y = e.offsetY;
	isDrawing = true;
});
canvas.addEventListener("touchstart", (e) => {
	e.preventDefault();
	var touches = e.changedTouches;
	if(touches && touches.length){
		x = touches[0].pageX;
		y = touches[0].pageY;
		isDrawing = true;
	}
});
canvas.addEventListener("mousemove", (e) => {
	if (isDrawing) {
		drawLine(context, x, y, e.offsetX, e.offsetY);
		x = e.offsetX;
		y = e.offsetY;		
	}	
});
canvas.addEventListener("touchmove", (e) => {
	e.preventDefault();
	if (isDrawing) {
		var touches = e.changedTouches;
		if(touches && touches.length){
			drawLine(context, x, y, touches[0].pageX, touches[0].pageY);
			x = touches[0].pageX;
			y = touches[0].pageY;		
		}
	}		
});

window.addEventListener("mouseup", (e) => {
	if (isDrawing) {
		drawLine(context, x, y, e.offsetX, e.offsetY);
		x = 0;
		y = 0;
		isDrawing = false;
	}
});

window.addEventListener("touchend", (e) => {	
	if (isDrawing) {
		e.preventDefault();
		var touches = e.changedTouches;
		if(touches && touches.length){			
			drawLine(context, x, y, touches[0].pageX, touches[0].pageY);
			x = 0;
			y = 0;
			isDrawing = false;
		}
	}
});


var btn = document.getElementById("btn_clear");
btn.addEventListener("click", (e) => {
	context.clearRect(0, 0, canvas.width, canvas.height);
});	

function drawLine(context, x1, y1, x2, y2) {
	context.beginPath();
	context.strokeStyle = "black";
	context.lineWidth = 1;
	context.moveTo(x1, y1);
	context.lineTo(x2, y2);
	context.stroke();
	context.closePath();
}
