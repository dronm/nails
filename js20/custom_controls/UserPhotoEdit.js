/**	
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2022

 * @extends EditFile
 * @requires core/extend.js
 * @requires controls/EditFile.js     

 * @class
 * @classdesc
 
 * @param {string} id - Object identifier
 * @param {object} options
 */
function UserPhotoEdit(id,options){
	options = options || {};	
	
	options.labelCaption = "Фотография:";
	options.showHref = false;
	options.showPic = true;
	options.fileInfTemplateOptions = {
		"IMG_CLASS": "userPhoto"
	}
	
	this.m_view = options.view;
	this.m_photoAdded = options.photoAdded;
	this.m_photoDeleted = options.photoDeleted;
	
	var self = this;
	
	options.onDeleteFile = function(fileId){
		self.deletePhoto(fileId);
		if(self.m_photoDeleted){
			self.m_photoDeleted();
		}
	};
	options.onFileAdded = function(fileId){
		self.photoAdded(fileId);
		if(self.m_photoAdded){
			self.m_photoAdded();
		}
	}
	options.allowedFileExtList = "jpg,jpeg,png";
	options.maxWidth = 250;
	options.maxHeight = 250;
	
	UserPhotoEdit.superclass.constructor.call(this, id, options);
}
//ViewObjectAjx,ViewAjxList
extend(UserPhotoEdit, EditFile);

/* Constants */


/* private members */

/* protected*/


/* public methods */

UserPhotoEdit.prototype.deletePhotoCont = function(id){
	var self = this;
	var pm = (new User_Controller()).getPublicMethod("delete_photo");
	pm.setFieldValue("user_id", id);
	pm.run({
		"ok":function(resp){
			var ctrl = self.m_view.getElement("photo");
			if(ctrl){
				ctrl.reset();
			}
		}
	});
}

UserPhotoEdit.prototype.deletePhoto = function(fileId){
	var ctrl_id = this.m_view.getElement("id");
	if(ctrl_id){
		var id = ctrl_id.getValue();
		if(!id){
			return;
		}
		var self = this;
		WindowQuestion.show({
			"cancel":false,
			"text":"Удалить фотографию?",
			"callBack": function(r)	{
				if(r == WindowQuestion.RES_YES){
					self.deletePhotoCont(id);
				}
			}
		});
	}
}

UserPhotoEdit.prototype.photoAdded = function(fileId){
	if(!window["FileReader"]){
		return;
	}
	var fl = this.m_view.getElement("photo").getValue();
	if(!fl || !fl.length){
		return;
	}
	DOMHelper.hide(document.getElementById(fileId+"-preview"));
	var self = this;
	var reader = new FileReader();
	reader.readAsDataURL(fl[0]);
	reader.onload = function () {
		var n = document.getElementById(fileId+"-preview");
		if(n){
			n.src = reader.result;
			n.onload = function(){
				var dim = CommonHelper.calculateImgRatioFit(this.width, this.height, self.m_maxWidth, self.m_maxWidth);
				this.width = dim.width;
				this.height = dim.height;
				DOMHelper.show(this);			
			}
		}		
		//status
		var n = document.getElementById(fileId+"-pic");
		if(n){
			n.setAttribute("class", "glyphicon glyphicon-ok");
		}
	};
	reader.onerror = function (error) {
		throw new Error(error);
	};
}

