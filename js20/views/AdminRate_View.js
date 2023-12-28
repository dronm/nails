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
function AdminRate_View(id,options){
	options = options || {};	
	//options.template = window.getApp().getTemplate("AdminRate");
	
	var self = this;
	options.addElement = function(){
		var pic_size = "500px";
		this.addElement(new Control(id+":photo", "img", {
			"attrs":{
				"alt":"Не выбрано",
				"width": pic_size,
				"height": pic_size,
				"max-width": pic_size,
				"max-height": pic_size
			},
			"alt":"Не выбрано",
			"src":"#"
		}));
		
		this.addElement(new ButtonCmd(id+":prev",{
			"glyph":"glyphicon-step-backward",
			"title":"Предыдущее фото",
			"enabled":false,
			"onClick":function(){
				self.setIndex(self.m_index - 1);
			}
		}));
/*		
		this.addElement(new ButtonCmd(id+":first",{
			"glyph":"glyphicon-fast-backward",
			"title":"Первое фото",
			"enabled":false,
			"onClick":function(){
				self.setIndex(0);
			}
		}));
*/		
		this.addElement(new ButtonCmd(id+":next",{
			"glyph":"glyphicon-step-forward",
			"title":"Следующее фото",
			"enabled":false,
			"onClick":function(){
				self.setIndex(self.m_index + 1);
			}
		}));
/*		
		this.addElement(new ButtonCmd(id+":last",{
			"glyph":"glyphicon-fast-forward",
			"title":"Последнее фото",
			"enabled":false,
			"onClick":function(){
				self.setIndex(self.m_model.getTotalCount() - 1);
			}
		}));
*/		
		for(var i=0; i <= 10; i ++){
			this.addElement(new Control(id+":r" + i, "SPAN", {
				"title":"Выбрать оценку",
				"events":{
					"click":function(e){
						var rt = parseInt(DOMHelper.getText(e.target), 10);	
						self.m_model.setFieldValue("admin_rate", rt);
						self.m_model.recUpdate();
						
						for(var i=0; i <= 10; i ++){
							DOMHelper.delClass(document.getElementById(self.getId() + ":r" + i), "bg-success");		
						}
						
						DOMHelper.addClass(e.target, "bg-success");		
						setTimeout(function(){
							self.setIndex(self.m_index + 1);
						}, 500);								
						self.m_updates.push({"id":self.m_model.getFieldValue("id"), "admin_rate": rt});
						/*
						var pm = (new SpecialistWork_Controller()).getPublicMethod("update");
						pm.setFieldValue("old_id", self.m_model.getFieldValue("id"));
						pm.setFieldValue("admin_rate", rt);
						pm.run({
							"ok":function(){
								setTimeout(function(){
									self.setIndex(self.m_index + 1);
								}, 500);								
							}
							,"fail":function(resp, err){
								DOMHelper.delClass(e.target, "bg-success");
								throw new Error(err);
							}
						});
						*/
					}
				}
			}));
			
		}
		
		this.addElement(new ButtonCtrl(id+":rotate", {
			"glyph":"glyphicon-repeat",
			"title":"Повернуть фото по часовой стрелке",
			"events":{
				"click":function(){
					self.m_rotation = (self.m_rotation + 90) % 360;
					self.getElement("photo").m_node.style.transform = `rotate(${self.m_rotation}deg)`;
				}
			}
		}));
		
	}
	
	this.m_index = -1;
	this.m_updates = [];
	
	AdminRate_View.superclass.constructor.call(this,id,options);
	
}
//ViewObjectAjx,ViewAjxList
extend(AdminRate_View, View);

/* Constants */


/* private members */

/* protected*/
AdminRate_View.prototype.m_model;
AdminRate_View.prototype.m_index;
AdminRate_View.prototype.m_rotation;
AdminRate_View.prototype.FETCH_COUNT = 1;
AdminRate_View.prototype.REST_FOR_FETCH = 2;

AdminRate_View.prototype.setIndex = function(ind){		
	var cnt = this.m_model.getRowCount();
console.log("setIndex ind=",ind, " getRowCount=", cnt)			
	if(ind < cnt){
		this.m_index = ind;
		
		var tot_cnt = this.m_model.getTotCount()
		if(cnt - ind <= this.REST_FOR_FETCH && cnt < tot_cnt){
console.log("tot_cnt=",tot_cnt, "fetching new data")						
			this.fetch();
		}
		this.m_rotation = 0;
		
		this.getElement("photo").setAttr("src", "//:0");
		this.getElement("prev").setEnabled(false);
		this.getElement("next").setEnabled(false);
		DOMHelper.hide(this.getId() + ":rating");
		for(var i=0; i <= 10; i ++){
			DOMHelper.delClass(document.getElementById(this.getId() + ":r" + i), "bg-success");		
		}
		
		if(ind > 0){
			this.getElement("prev").setEnabled(true);			
		}
		if(ind < tot_cnt - 1){
			this.getElement("next").setEnabled(true);
		}		
		
		this.m_model.getRow(ind);
console.log("getRow ID=",this.m_model.getFieldValue("id"))
console.log("getRow admin_rate=",this.m_model.getFieldValue("admin_rate"))
console.log("getRow isNull=",this.m_model.getField("admin_rate").isNull())
console.log("getRow isSet=",this.m_model.getField("admin_rate").isSet())
		this.getElement("photo").setAttr("src", "data:image/png;base64, " + this.m_model.getFieldValue("photo").dataBase64);
		this.getElement("photo").m_node.style = undefined;
		
		var rt = this.m_model.getFieldValue("admin_rate");
		if(rt!=undefined){
			DOMHelper.addClass(document.getElementById(this.getId() + ":r" + rt), "bg-success");		
		}
		
		DOMHelper.setText(this.getId() + ":inf", (ind+1) + " из " + tot_cnt);
		
		DOMHelper.show(this.getId() + ":rating");
	}
}

AdminRate_View.prototype.setModel = function(m){
	var new_ind;
	if(!this.m_model){
		this.m_model = m;	
	}else{
		//append
console.log("appending new model, current index:", this.m_index, " getRowCount=",this.m_model.getRowCount())
		var cur = this.m_model.m_currentRow;
		while(m.getNextRow()){
			this.m_model.resetFields();
			this.m_model.setFields(m.getFields());
			this.m_model.recInsert();
		}
		this.m_model.m_currentRow = cur;
console.log("getRowCount=",this.m_model.getRowCount())		
	}
}

/* public methods */

AdminRate_View.prototype.toDOM = function(p){
	AdminRate_View.superclass.toDOM.call(this, p);
	var self = this;
	this.fetch(function(){
		self.setIndex(0);		
	});
}

AdminRate_View.prototype.fetch = function(callback){
	var contr = new SpecialistWork_Controller();
	var pm = contr.getPublicMethod("get_for_rate_list");
	pm.setFieldValue("count", this.FETCH_COUNT);
	pm.setFieldValue("from", this.m_index + 1);
	var self = this;
	pm.run({
		"ok":function(resp){
			self.setModel(resp.getModel("SpecialistWorkForRateList_Model"));
			if(callback){
				callback();
			}
		}
	})
}
