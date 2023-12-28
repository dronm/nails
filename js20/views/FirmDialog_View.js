/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023

 * @extends ViewObjectAjx
 * @requires core/extend.js  

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {Object} options
 */
function FirmDialog_View(id,options){	

	options = options || {};
	
	options.model = (options.models&&options.models.FirmDialog_Model)? options.models.FirmDialog_Model : new FirmDialog_Model();
	options.controller = options.controller || new Firm_Controller();
	
	var self = this;
	options.addElement = function(){
		this.addElement(new EditString(id+":name",{
			"labelCaption":"Наименование:",
			"required":true,
			"focus":true,
			"maxLength":"250",
			"placeholder":"Краткое наименование контрагента",
			"title":"Краткое наименование контрагента для поиска"
		}));	

		this.addElement(new EditString(id+":name_full",{
			"labelCaption":"Полное наименование:",
			"required":false,
			"focus":true,
			"maxLength":"1000",
			"placeholder":"Официальное наименование контрагента",
			"title":"Полное наименование контрагента в соответствии с учредительными документами",
			"events":{
				"focus": function(){
					var v = self.getElement("name_full").getValue();
					if(!v || !v.length){
						self.getElement("name_full").setValue(self.getElement("name").getValue());
					}
				}
			}
		}));	

		this.addElement(new EditNum(id+":inn",{
			"required":false,
			"maxLength":12,
			"labelCaption":"ИНН:",
			"buttonSelect": new ButtonOrgSearch(id+":btnOrgSearch",{
				"viewContext":this,
				"onGetAttrs": function(){
					var attrs = this.DEF_ATTRS;
					attrs["Адрес"] = "legal_address"
					return attrs;
				}
			}),
		}));	

		this.addElement(new EditNum(id+":kpp",{
			"required":false,
			"maxLength":10,
			"labelCaption":"КПП:"
		}));	

		this.addElement(new EditString(id+":post_address",{		
			"maxLength":1000,
			"labelCaption":"Почтовый адрес:",
			"placeholder":"Почтовый адрес",
			"title":"Почтовый адрес организации"
		}));	

		this.addElement(new EditString(id+":legal_address",{		
			"maxLength":1000,
			"labelCaption":"Юридический адрес:",
			"placeholder":"Юридический адрес",
			"title":"Адрес организации в соответствии с учредительыми документами"
		}));	

		this.addElement(new EditNum(id+":ogrn",{		
			"maxLength":15,
			"labelCaption":"ОГРН:",
			"title":"Регистрационный номер организации",
			"required":false
		}));	
		
		this.addElement(new EditNum(id+":okpo",{		
			"maxLength":20,
			"labelCaption":"ОКПО:",
			"title":"ОКПО организации"
		}));	
		
		this.addElement(new EditString(id+":okved",{
			"maxLength":50,
			"labelCaption":"ОКВЭД:",
			"title":"Код ОКВЭД организации"
		}));

		this.addElement(new BankAccEdit(id+":bank_acc",{
			"cmdSmClear":true,
			"labelCaption":"Расчетный счет:",
			"placeholder": "Расчетный счет",
			"title": "Расчетный счет организации",
			"events":{
				"input": function(){
					self.checkBankAcc();					
				}
			},
			"onClear":function(){
				self.getElement("bank_acc").setValid();
			}
		}));
		
		var er_ctrl = new ErrorControl(id+":banks_ref:error");
		var info_ctrl = new EditInfo(id+":banks_ref:info", {
			"templateOptions":{
				"TEXT_CLASS":"text-info",
				"PICT_CLASS":"icon-checkmark"
			}
		});
		this.addElement(new BankEditRef(id+":banks_ref",{
			"errorControl": er_ctrl,
			"infoControlElements": [er_ctrl, info_ctrl],
			"onSelect":function(f){
				self.getElement("banks_ref").setValid();
				var bik = f.bik.getValue();
				self.getElement("bank_acc").setBik(bik);
				var descr = f["name"].getValue();
				var kor = f["korshet"].getValue();
				if(kor){
					descr+= " "+kor;
				}
				var gor = f["gor"].getValue();
				if(gor){
					descr+= " "+gor;
				}
				self.getElement("banks_ref").getInfoControls().getElement("info").setValue(descr);
				self.checkBankAcc();
			},
			"onClear":function(){
				self.getElement("bank_acc").setBik(undefined);
				//self.getElement("banks_ref").getErrorControl().setValue("");
				self.getElement("banks_ref").getInfoControls().getElement("info").clear();
				self.checkBankAcc();
			}
		}));

	}
	
	FirmDialog_View.superclass.constructor.call(this,id,options);
	
	
	//****************************************************	
	
	//read
	var read_b = [
		new DataBinding({"control":this.getElement("name")}),
		new DataBinding({"control":this.getElement("name_full")}),
		new DataBinding({"control":this.getElement("inn")}),
		new DataBinding({"control":this.getElement("kpp")}),
		new DataBinding({"control":this.getElement("legal_address")}),
		new DataBinding({"control":this.getElement("post_address")}),
		new DataBinding({"control":this.getElement("ogrn")}),
		new DataBinding({"control":this.getElement("okved")}),		
		new DataBinding({"control":this.getElement("okpo")}),
		new DataBinding({"control":this.getElement("bank_acc")}),
		new DataBinding({"control":this.getElement("banks_ref")})
	];
	this.setDataBindings(read_b);
	
	//write
	var write_b = [
		new CommandBinding({"control":this.getElement("name")}),
		new CommandBinding({"control":this.getElement("name_full")}),
		new CommandBinding({"control":this.getElement("inn")}),
		new CommandBinding({"control":this.getElement("kpp")}),
		new CommandBinding({"control":this.getElement("legal_address")}),
		new CommandBinding({"control":this.getElement("post_address")}),
		new CommandBinding({"control":this.getElement("ogrn")}),
		new CommandBinding({"control":this.getElement("okved")}),		
		new CommandBinding({"control":this.getElement("okpo")}),
		new CommandBinding({"control":this.getElement("bank_acc")}),
		new CommandBinding({"control":this.getElement("banks_ref"), "fieldId": "bank_bik"})
	];
	this.setWriteBindings(write_b);
	
}
extend(FirmDialog_View,ViewObjectAjx);

FirmDialog_View.prototype.checkBankAcc = function(){	
	var v = document.getElementById(this.getId() + ":bank_acc").value;
	var bank_ctrl = this.getElement("bank_acc");
	if (v.length == 20 && bank_ctrl.m_bik && bank_ctrl.m_bik.length){
		if(!bank_ctrl.validate()){
			return;
		}
	}else{
		bank_ctrl.setValid();
	}
}

