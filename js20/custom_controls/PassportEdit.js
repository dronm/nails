/** Copyright (c) 2023
	Andrey Mikhalevich, Katren ltd.
 *
 * @Param picture
 */
function PassportEdit(id,options){
	options = options || {};
	options.viewClass = Passport_View,
	options.headTitle = "Данные паспорта";
	options.viewTemplate = "Passport_View";
	options.strictValidation = true;
	options.formatValue = function(v){
		if(!v || !v.series){
			return "Введите данные паспорта"
		}
		return "Паспорт: " + v.series+" "+v.num+", "+v.issue_body + ", " + DateHelper.format(DateHelper.strtotime(v.issue_date), "d/m/y");
	}
	var self = this;
	options.title = "Данные паспорта";	
	options.value = "Введите данные паспорта";
	if(options.picture !== false){
		options.buttonSelect = new ButtonCtrl(this.getId() + ":btnFromStrore", {
			"glyph":"glyphicon-folder-open",
			"title":"Выбрать файл",
			"onClick":function(){
				self.m_docFileNode.removeAttribute("capture");
				$(self.m_docFileNode).click();
			}
		});
		options.buttonOpen = new ButtonCtrl(this.getId() + ":btnFromCam", {
			"glyph": "glyphicon-camera", //"glyphicon-phone",
			"title":"Открыть камеру",
			"onClick":function(){
				self.m_docFileNode.setAttribute("capture", "camera");
				$(self.m_docFileNode).click();
			}
		});
	}	
	options.errorControl = new ErrorControl(id+":error");
	var info_ctrl = new EditInfo(id+":info", {
		"templateOptions":{
			"TEXT_CLASS":"text-info",
			"PICT_CLASS":"icon-checkmark"
		}
	});
	options.infoControlElements = [options.errorControl, info_ctrl];
	
	PassportEdit.superclass.constructor.call(this,id,options);	
}
extend(PassportEdit, EditModalDialog);

PassportEdit.prototype.m_docFileNode;

PassportEdit.prototype.validate = function(){
	if(!PassportEdit.superclass.validate.call(this)){
		return false;
	}	
	if(!this.m_docFileNode || !this.m_docFileNode.files || !this.m_docFileNode.files.length){
		this.setNotValid("Файл не прикреплен");
		return false;
	}
	return true;
}

PassportEdit.prototype.toDOM = function(p){
	PassportEdit.superclass.toDOM.call(this, p);	
	
	this.m_docFileNode = document.createElement("input");
	this.m_docFileNode["type"] = "file";
	this.m_docFileNode["id"] = this.getId()+ ":docFile";
	this.m_docFileNode["name"] = "focFile";
	this.m_docFileNode["accept"] = "image/*";
	this.m_docFileNode.setAttribute("class", "hidden");
	var self = this;
	this.m_docFileNode.addEventListener("change", function(e) {
		var tp = "image/";
		if(e.target.files[0]["type"].substring(0, tp.length) != tp){
			self.setNotValid("Данный тип файлов не поддерживается.");
			return;
		}
		self.getInfoControls().getElement("info").setValue("Файл: " + e.target.files[0]["name"]);
		self.setValid();
	});
	p.appendChild(this.m_docFileNode);
}


