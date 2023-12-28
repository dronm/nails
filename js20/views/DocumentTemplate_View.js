/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023
 
 * @extends ViewObjectAjx.js
 * @requires core/extend.js  
 * @requires controls/ViewObjectAjx.js 
 
 * @class
 * @classdesc
	
 * @param {string} id view identifier
 * @param {object} options
 * @param {object} options.models All data models
 * @param {object} options.variantStorage {name,model}
 */	
function DocumentTemplate_View(id,options){	

	options = options || {};

	options.controller = new DocumentTemplate_Controller();
	options.model = (options && options.models && options.models.DocumentTemplateDialog_Model)? options.models.DocumentTemplateDialog_Model : new DocumentTemplateDialog_Model();
	
	options.addElement = function(){
		this.addElement(new EditString(id+":name",{
			"labelCaption":"Наименование:",
			"required":true,
			"maxLength":500
		}));	
			
		this.addElement(new EditText(id+":sql_query",{
			"labelCaption":"Запрос к базе данных:"
		}));	

		this.addElement(new EditCheckBox(id+":need_signing",{
			"labelCaption":"Требовать подпись:",
			"title":"Документ должен быть подписан"
		}));	

		this.addElement(new EditString(id+":sign_image_name",{
			"labelCaption":"Картинка для замены:",
			"title":"Наименование файла картинки, которая будет заменена на подпись"
		}));	
		
		var self = this;
		this.addElement(new EditFile(id+":file_preview",{
			"maxWidth":"100",
			"maxHeight":"100",
			"multipleFiles":true
			,"showHref":false
			,"showPic":true
			,"onDeleteFile":function(fileId,callBack){
				self.m_attachManager.deleteAttachment(fileId,callBack);
			}
			,"onFileAdded":function(fileId){
				//self.addAttachment(fileId);
				self.m_attachManager.addAttachment(fileId);
			}
			,"onDownload":function(fileId){
				self.m_attachManager.downloadAttachment(fileId);
			}
			,"allowedFileExtList":"xlsx,docx"
		}));	

		//********* field grid ***********************
		this.addElement(new DocumentTemplateFieldGrid(id+":fields",{
		}));		

	}
	
	DocumentTemplate_View.superclass.constructor.call(this,id,options);
	
	this.m_attachManager = new AttachmentManager({
		"view": this,
		"dataType": "document_templates",
		"attachmentViewName": "file_preview"
	});
	
	//****************************************************
	//read
	this.setDataBindings([
		new DataBinding({"control":this.getElement("name")})
		,new DataBinding({"control":this.getElement("sql_query")})
		,new DataBinding({"control":this.getElement("need_signing")})
		,new DataBinding({"control":this.getElement("sign_image_name")})		
		,new DataBinding({"control":this.getElement("fields"), "fieldId":"fields"})
		,new DataBinding({"control":this.getElement("file_preview")})
	]);
	
	//write
	this.setWriteBindings([
		new CommandBinding({"control":this.getElement("name")})
		,new CommandBinding({"control":this.getElement("sql_query")})
		,new CommandBinding({"control":this.getElement("need_signing")})
		,new CommandBinding({"control":this.getElement("sign_image_name")})
		,new CommandBinding({"control":this.getElement("fields"),"fieldId":"fields"})
	]);
			
}
extend(DocumentTemplate_View,ViewObjectAjx);

