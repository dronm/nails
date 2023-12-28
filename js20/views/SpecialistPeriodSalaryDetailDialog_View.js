/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023
 
 * @extends ViewObjectAjx.js
 * @requires core/extend.js  
 * @requires controls/ViewObjectAjx.js 
 
 * @class
 * @classdesc
	
 * @param {string} id view identifier
 * @param {object} options
 */	
function SpecialistPeriodSalaryDetailDialog_View(id,options){	

	options = options || {};
	
	options.template = window.getApp().getTemplate("SpecialistPeriodSalaryDetailDialog");
	options.controller = new SpecialistPeriodSalaryDetail_Controller();
	options.model = (options.models&&options.models.SpecialistPeriodSalaryDetailList_Model)? options.models.SpecialistPeriodSalaryDetailList_Model : new SpecialistPeriodSalaryDetailList_Model();
	
	options.addElement = function(){
		this.addElement(new SpecialistEdit(id+":specialists_ref",{
		}));	

		var self = this;
		this.addElement(new EditFloat(id+":hours",{
			"required": true,
			"precision":"2",
			"labelCaption": "Часы:",
			"title":"Отработанное количество часов",
			"events":{
				"input":function(){
					self.recalc();
				}
			}
		}));	

		this.addElement(new EditMoney(id+":work_total",{
			"required": true,
			"labelCaption": "Общая выручка:",
			"title":"Сумма общей выручки по сотруднику",
			"events":{
				"input":function(){
					self.recalc();
				}
			}
		}));	

		this.addElement(new EditMoney(id+":debet",{
			"required": false,
			"labelCaption": "Доп.начисления:",
			"title":"Сумма дополнительных начислений по сотруднику",
			"events":{
				"input":function(){
					self.recalc();
				}
			}
		}));	

		this.addElement(new EditMoney(id+":kredit",{
			"required": false,
			"labelCaption": "Доп.удержания:",
			"title":"Сумма дополнительных удержаний по сотруднику",
			"events":{
				"input":function(){
					self.recalc();
				}
			}
		}));	

		this.addElement(new EditMoney(id+":rent_price",{
			"required": true,
			"labelCaption": "Цена аренды:",
			"title":"Цена аренды за 1 час",
			"events":{
				"input":function(){
					self.recalc();
				}
			}
		}));	

		this.addElement(new EditMoney(id+":rent_total",{
			"required": true,
			"labelCaption": "Сумма за аренду:",
			"title":"Сумма удержания за аренду",
			"events":{
				"input":function(){
					self.recalc();
				}
			}			
		}));	

		this.addElement(new EditMoney(id+":work_total_salary",{
			"required": true,
			"labelCaption": "Для начисления:",
			"title":"Сумма для начисления",
			"events":{
				"input":function(){
					self.recalc();
				}
			}
		}));	

		this.addElement(new EditMoney(id+":total",{
			"required": true,
			"labelCaption": "Итого:",
			"title":"Итоговая сумма начисления",
			"events":{
				"input":function(){
					self.recalc();
				}
			}
		}));	

		this.addElement(new EditFloat(id+":agent_percent",{
			"required": true,
			"precision":"2",
			"labelCaption": "Агентское вознаграждение:",
			"title":"Процент агентского вознаграждения",
			"events":{
				"input":function(){
					self.recalc();
				}
			}
		}));	

	}
	
	SpecialistPeriodSalaryDetailDialog_View.superclass.constructor.call(this,id,options);
	
	//****************************************************
	//read
	this.setDataBindings([
		new DataBinding({"control":this.getElement("specialists_ref")})
		,new DataBinding({"control":this.getElement("hours")})
		,new DataBinding({"control":this.getElement("debet")})
		,new DataBinding({"control":this.getElement("kredit")})
		,new DataBinding({"control":this.getElement("rent_price")})
		,new DataBinding({"control":this.getElement("rent_total")})
		,new DataBinding({"control":this.getElement("work_total")})
		,new DataBinding({"control":this.getElement("work_total_salary")})
		,new DataBinding({"control":this.getElement("total")})
		,new DataBinding({"control":this.getElement("agent_percent")})
	]);
	
	//write
	this.setWriteBindings([
		new CommandBinding({"control":this.getElement("specialists_ref"), "fieldId":"specialist_id"})
		,new CommandBinding({"control":this.getElement("hours")})
		,new CommandBinding({"control":this.getElement("debet")})
		,new CommandBinding({"control":this.getElement("kredit")})
		,new CommandBinding({"control":this.getElement("rent_price")})
		,new CommandBinding({"control":this.getElement("rent_total")})
		,new CommandBinding({"control":this.getElement("work_total")})
		,new CommandBinding({"control":this.getElement("work_total_salary")})
		,new CommandBinding({"control":this.getElement("total")})
		,new CommandBinding({"control":this.getElement("agent_percent")})
		
	]);
	
}
extend(SpecialistPeriodSalaryDetailDialog_View, ViewObjectAjx);

SpecialistPeriodSalaryDetailDialog_View.prototype.recalc = function(){
	var w = this.getElement("work_total").getValue();
	
	var rent = this.getElement("rent_price").getValue() * this.getElement("hours").getValue();
	this.getElement("rent_total").setValue(rent);
	
	var ws = w + this.getElement("debet").getValue() - this.getElement("kredit").getValue() - rent;
	this.getElement("work_total_salary").setValue(ws);
	
	var v = ws - Math.round(ws * this.getElement("agent_percent").getValue() / 100 * 100) / 100;
	this.getElement("total").setValue(v);	
}


