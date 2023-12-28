/* Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function SpecialistPeriodSalaryDialog_View(id,options){	

	options = options || {};
	options.HEAD_TITLE = "Расчет зарплаты";
	options.cmdSave = true;
	options.controller = new SpecialistPeriodSalary_Controller();
	options.model = (options.models&&options.models.SpecialistPeriodSalaryList_Model)? options.models.SpecialistPeriodSalaryList_Model : new SpecialistPeriodSalaryList_Model();

	var self = this;
	options.addElement = function(){
		
		this.addElement(new EditDateTime(id+":date_time", {
			"inline":true,
			"enabled":false
		}));
/*
		this.addElement(new EditPeriodMonth(id+":period", {
			"labelCaption":"Период:",
			"title":"Период расчета зарплаты"
		}));
*/
		this.addElement(new EditDate(id+":period", {
			"labelCaption":"Период:",
			"title":"Период расчета зарплаты"
		}));

		this.addElement(new StudioEdit(id+":studios_ref", {
			"labelCaption":"Салон:"
		}));

		//grid
		this.addElement(new SpecialistPeriodSalaryDetailList_View(id+":details",{
			"detail":true,
			"view": this
		}));

		this.addElement(new ButtonCmd(id+":fill_doc",{
			"caption":"Заполнить",
			"title": "Заполнить табличную часть по студии за период",
			"onClick":function(){
				self.fill_doc();
			}
		}));
	}	
	
	SpecialistPeriodSalaryDialog_View.superclass.constructor.call(this,id,options);
	
	//read
	var r_bd = [
		new DataBinding({"control":this.getElement("date_time")})
		,new DataBinding({"control":this.getElement("period")})
		,new DataBinding({"control":this.getElement("studios_ref")})
	];
	this.setDataBindings(r_bd);
	
	//write
	var wr_b = [
		new CommandBinding({"control":this.getElement("date_time")})
		,new CommandBinding({"control":this.getElement("period")})
		,new CommandBinding({"control":this.getElement("studios_ref"), "fieldId": "studio_id"})
	];
	this.setWriteBindings(wr_b);
	
	this.addDetailDataSet({
		"control":this.getElement("details").getElement("grid"),
		"controlFieldId":"specialist_period_salary_id",
		"field":this.m_model.getField("id")
	});
	
}
extend(SpecialistPeriodSalaryDialog_View,ViewObjectAjx);


SpecialistPeriodSalaryDialog_View.prototype.fill_doc_cont = function(){
	var id = this.getElement("id").getValue();
	if (!id){
		throw new Error("Документ не записан!");
	}
	
	var self = this;
	WindowQuestion.show({
		"text":"Заполнить?",
		"no": false,
		"callBackYes": function(){
			self.fill_doc_call(id)
		}
	});
}

SpecialistPeriodSalaryDialog_View.prototype.fill_doc = function(){
	//if (this.getCmd() != "insert" || (!this.m_view.getModified() && !this.getModified()) ){
	if (this.getCmd() != "insert"  && !this.getModified() ){
		this.fill_doc_cont();
	}
	else{	
		var self = this;
		this.getControlOK().setEnabled(false);
		this.getControlSave().setEnabled(false);				
		this.onSave(
			function(){
				self.fill_doc_cont();
			},
			null,
			function(){
				self.getControlOK().setEnabled(true);
				self.getControlSave().setEnabled(true);				
			}
		);			
	}	
}

SpecialistPeriodSalaryDialog_View.prototype.fill_doc_call = function(id){
	var self = this;
	var pm = (new SpecialistPeriodSalary_Controller()).getPublicMethod("fill_doc");
	pm.setFieldValue("id", id);
	window.setGlobalWait(true);
	pm.run({
		"ok":function(resp){
			//Update fieds
			self.getElement("details").getElement("grid").onRefresh(function(){
				window.showTempNote("Документ заполнен",null,5000);
			});			
		}
		,"all":function(){
			window.setGlobalWait(false);
		}
	})
}

