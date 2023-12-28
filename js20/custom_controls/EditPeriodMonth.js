/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2019
 
 * @class
 * @classdesc Period Edit cotrol
 
 * @extends EditPeriodDate
 
 * @requires core/extend.js
 * @requires controls/ControlContainer.js
 * @requires controls/ButtonCmd.js               
 
 * @param string id 
 * @param {object} options
 * @param {object} options.filters
 * @param {string} options.dateFormat  
 * @param {function} options.onChange
 * @param {Date} [options.dateFrom=first date of current month]   
 */
 
function EditPeriodMonth(id,options){
	options = options || {};	
	options.template = options.template || window.getApp().getTemplate("EditPeriodMonth");

	options.cmdPeriodSelect = false;
	options.downTitle = options.downTitle || "Предыдущий месяц";
	options.upTitle = options.upTitle || "Следующий месяц";
	options.cmdControlTo = false;
	options.cmdControlFrom = false;
	
	options.cmdDownFast = false;
	options.cmdUpFast = false;
	this.setControlPeriodSelect(new PeriodSelect(this.getId()+":periodSelect",{}));
	
	//options.periodSelectClass = PeriodSelect;
	//options.periodSelectOptions = {"periodShift":true};
	
	this.m_dateFrom = options.dateFrom;	
	if(!this.m_dateFrom)this.m_dateFrom = this.getPeriodFrom();
	this.calcDateTo();
	
	this.m_filters = options.filters;	
	this.m_dateFormat = options.dateFormat;
	this.m_onChange = options.onChange;
	
	EditPeriodMonth.superclass.constructor.call(this,id,options);
}
extend(EditPeriodMonth,EditPeriodDate);

EditPeriodMonth.prototype.m_dateFrom;
EditPeriodMonth.prototype.m_dateTo;
EditPeriodMonth.prototype.m_timeFrom;

EditPeriodMonth.prototype.INCORRECT_VAL_CLASS="error";

EditPeriodMonth.prototype.getPeriodFrom = function(dt){
	return DateHelper.monthStart(dt);
}

EditPeriodMonth.prototype.addControls = function(){
	
	this.addElement(this.m_controlDownFast);
	this.addElement(this.m_controlDown);
	
	var self = this;	
	this.addElement(new Label(this.getId()+":inf",{
		"caption":this.getPeriodDescr(),
		"events":{
			"click":function(){
				self.picCustomDate();
			}
		}
	}));

	this.addElement(this.m_controlUp);
	this.addElement(this.m_controlUpFast);	
	
	this.m_errorControl = new ErrorControl(this.getId()+":error");
	this.addElement(this.m_errorControl);
}

EditPeriodMonth.prototype.picCustomDate = function(){
	var self = this;
	var p = $(this.getElement("inf").getNode());
	p.datepicker({
		format:{
			//called after date is selected
			toDisplay: function (date, format, language) {
				date = DateHelper.dateStart(date);
				//date = new Date(date.getTime() + DateHelper.timeToMS(constants.first_shift_start_time.getValue()));				
				self.setDateFrom(date);
			},
			//called in ctrl edit?
			toValue: function (date, format, language) {
			}																	
		},
		language:"ru",
		daysOfWeekHighlighted:"0,6",
		autoclose:true,
		todayHighlight:true,
		orientation: "bottom right",
		//container:form,
		showOnFocus:false,
		clearBtn:true
	});
	
	p.on("hide", function(ev){
		//self.getEditControl().applyMask();
	});					
	
	p.datepicker("show");
}

EditPeriodMonth.prototype.go = function(sign){
	var t = (sign>0)?  this.m_dateTo.getTime() : this.m_dateFrom.getTime();
	this.setDateFrom(DateHelper.monthStart( new Date(t + sign*24*60*60*1000)) );
}

EditPeriodMonth.prototype.setDateFrom = function(dt){
	this.m_dateFrom = this.getPeriodFrom(dt);
	this.calcDateTo();
	this.updateDateInf();
		
	if(this.m_grid){
		this.applyFilter();
		this.m_grid.onRefresh();
	}
	else if(this.m_onChange){
		this.m_onChange(this.m_dateFrom,this.m_dateTo);
	}
}
EditPeriodMonth.prototype.getDateFrom = function(){
	return this.m_dateFrom;
}

EditPeriodMonth.prototype.getDateTo = function(){
	return this.m_dateTo;
}

EditPeriodMonth.prototype.calcDateTo = function(){	
	this.m_dateTo = DateHelper.monthEnd(this.m_dateFrom);
}

EditPeriodMonth.prototype.updateDateInf = function(){	
	this.getElement("inf").setValue(this.getPeriodDescr());
}

EditPeriodMonth.prototype.getPeriodDescr = function(){	
	return (DateHelper.format(this.m_dateFrom,"FF Y"));
}

EditPeriodMonth.prototype.applyFilter = function(v){
	if(this.m_filters&&this.m_filters.length){
		this.m_filters[0].val = DateHelper.format(this.m_dateFrom,this.m_dateFormat);
		if(this.m_filters.length>1){
			this.m_filters[1].val = DateHelper.format(this.m_dateTo,this.m_dateFormat);
		}
	}
}

EditPeriodMonth.prototype.setGrid = function(v){
	this.m_grid = v;
	if(this.m_filters&&this.m_filters.length){
		this.applyFilter();
		
		for (var i=0;i<this.m_filters.length;i++){
			this.m_grid.setFilter(this.m_filters[i]);
		}
		
	}
}
EditPeriodMonth.prototype.isNull = function(){
	return isNaN(this.m_dateFrom.getTime());
}

EditPeriodMonth.prototype.focus = function(){
}

EditPeriodMonth.prototype.reset = function(){
	this.m_dateFrom = undefined;
	this.m_dateTo = undefined;
}

EditPeriodMonth.prototype.setValid = function(){
	DOMHelper.delClass(this.m_node,this.INCORRECT_VAL_CLASS);
	this.m_errorControl.clear();
}

EditPeriodMonth.prototype.setNotValid = function(erStr){
	DOMHelper.addClass(this.m_node,this.INCORRECT_VAL_CLASS);
	this.m_errorControl.setValue(erStr);
}

