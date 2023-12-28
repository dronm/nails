/**
 * @author Andrey Mikhalevich <katrenplus@mail.ru>, 2023
 
 * @class
 * @classdesc
	
 * @param {object} options
 */	
function App_nails(options){
	options = options || {};
	options.lang = "ru";
	options.paginationClass = Pagination;
	options.soundEnabled = false;
	
	App_nails.superclass.constructor.call(this,"dpassport",options);	
}
extend(App_nails,App);

/* Constants */
App_nails.prototype.SERV_RESPONSE_MODEL_ID = "Response";
App_nails.prototype.EVENT_MODEL_ID = "Event";
App_nails.prototype.INSERTED_KEY_MODEL_ID = "InsertedKey";

/* private members */

/* protected*/


/* public methods */
App_nails.prototype.getSidebarId = function(){
	return this.getServVar("user_name")+"_"+"sidebar-xs";
}
App_nails.prototype.toggleSidebar = function(){
	var id = this.getSidebarId();
	this.storageSet(id,(this.storageGet(id)=="xs")? "":"xs");
}


App_nails.prototype.makeItemCurrent = function(elem){
	if (elem){		
		var l = DOMHelper.getElementsByAttr("active", document.body, "class", false, "LI");
		for(var i=0; i < l.length; i++){
			//DOMHelper.delClass(l[i], "active");
			//groups!!!						
			var ch_l = DOMHelper.getElementsByAttr("has-ul", l[i], "class", true, "A");
			if(!ch_l || !ch_l.length){
				//not a group
				DOMHelper.delClass(l[i], "active");
			}
		}
		var it = (elem.tagName.toUpperCase()=="LI")? elem:elem.parentNode;
		DOMHelper.addClass(it, "active");
		/*if (elem.nextSibling){
			elem.nextSibling.style="display: block;";
		}*/
		var p = DOMHelper.getParentByTagName(it, "LI");
		if(p && !DOMHelper.hasClass(p, "active")){
			DOMHelper.addClass(p, "active");
			var ch_l = DOMHelper.getElementsByAttr("hidden-ul", p, "class", true, "UL");
			if(ch_l && ch_l.length){
				ch_l[0].style="display: block;";
			}
		}
	}
}

App_nails.prototype.showMenuItem = function(item,c,f,t,extra,title){
	App_nails.superclass.showMenuItem.call(this,item,c,f,t,extra,title);
	this.makeItemCurrent(item);
}

/**
 * cont is for context, for massage being shown in context window not in parent
 */
App_nails.prototype.openHrefDownload = function(cont, pm, viewId, href){
	var wnd = pm.openHref(viewId, href);
	if(pm.fieldExists("inline") && (!wnd || !wnd.top)){
		pm.setFieldValue("inline", 0);
		pm.download();	
		cont.showTempWarn("Всплывающие окна заблокированы, файл открыт в режиме скачивания",null,10000);
	}

}

App_nails.prototype.formatContactList = function(f, cell){
	var contact_list = f.contact_list.getValue();
	if(!contact_list || !contact_list.length){
		return;
	}
	var cell_n = cell.getNode();
	var ul_tag = document.createElement("ul");
	
	for(var i = 0; i< contact_list.length; i++){
		var tel = contact_list[i].tel;
		var tel_m = tel;
		if(tel_m && tel_m.length==10){
			tel_m = "+7"+tel;
		}
		else if(tel_m && tel_m.length==11){
			tel_m = "+7"+tel.substr(1);
		}
		var li_tag = document.createElement("li");
		li_tag.textContent = contact_list[i]["name"];
		if(contact_list[i]["email"]){
			li_tag.textContent+= ", "+contact_list[i]["email"];
		}
		if(contact_list[i]["post"]){
			li_tag.textContent+= ", "+contact_list[i]["post"];
		}		
		
		if(tel_m && tel_m.length){
			li_tag.textContent+= ", ";
			var t_tag = document.createElement("A");
			t_tag.setAttribute("href","tel:"+tel_m);
			t_tag.textContent = CommonHelper.maskFormat(tel, window.getApp().getPhoneEditMask());
			t_tag.setAttribute("title","Позвонить на номер "+t_tag.textContent);
			if(window.getWidthType()!="sm"){
				t_tag.setAttribute("tel",tel_m);
				EventHelper.add(t_tag,"click",function(e){
					e.preventDefault();
					window.getApp().makeCall(this.getAttribute("tel"));
				});
			}
			
			li_tag.appendChild(t_tag);
		}
		
		ul_tag.appendChild(li_tag);
	}	
	
	cell_n.appendChild(ul_tag);
}

App_nails.prototype.TMInviteContact = function(contactRef, callBack){
	var pm = (new TmOutMessage_Controller()).getPublicMethod("tm_invite_contact");
	pm.setFieldValue("contact_id", contactRef.getKey("id"));
	window.setGlobalWait(true);
	pm.run({
		"all":function(){
			window.setGlobalWait(false);
			if(callBack){
				callBack();//disable controls
			}
		}
		,"ok":function(){
			window.showTempNote("Отправлено приглашение в Telegram",null,5000);
		}
	});
}

App_nails.prototype.makeCallContinue = function(tel){
	/*f(this.getServVar("debug")==1){
		window.showTempNote("ТЕСТ Набор номера: "+tel,null,10000);
		return;
	}*/
	
	var pm = (new Caller_Controller()).getPublicMethod("call");
	pm.setFieldValue("tel",tel);
	var self = this;
	pm.run({
		"ok":function(resp){
			if(self.m_appSrv){
				var m = resp.getModel("Call_Model");
				if(m && m.getNextRow()){
					//subscribe
					self.callActions(m.getFieldValue("call_id"),tel);
				}
			}
			window.showTempNote("Набор номера: "+tel,null,10000);			
		}
	})
}

App_nails.prototype.makeCall = function(tel){
	if(!window.Caller_Controller){
		throw new Error("Контроллер Caller_Controlle не определен!");
	}
	var self = this;
	WindowQuestion.show({
		"cancel":false,
		"text":"Набрать номер "+tel+"?",
		"callBack":function(res){
			if(res==WindowQuestion.RES_YES){
				self.makeCallContinue(tel);
			}
		}
	});
}

//tr - table row
//
App_nails.prototype.openGridAttachment = function(gridModel, tr, previewFieldId, attDataType, inline, ind){
	if(!tr){
		return;
	}
	var keys = CommonHelper.unserialize(tr.getAttribute("keys"));
	var ind = gridModel.locate(keys,true);
	if(ind == undefined || ind < 0 || !gridModel.getRow(ind)){
		return;
	}
	var content_id;
	var preview = gridModel.getFieldValue(previewFieldId);
	if(preview && preview.length && preview[0].id){
		content_id = preview[0].id;
		
	}else if(preview && preview.id){
		content_id = preview.id;
	}
	var pm = (new Attachment_Controller()).getPublicMethod("get_file");
	pm.setFieldValue("ref", CommonHelper.serialize(new RefType({"keys":{"id": gridModel.getFieldValue("id")}, "dataType": attDataType})));
	pm.setFieldValue("content_id", content_id);
	pm.setFieldValue("inline", inline? "1":"0");
	if(inline){
		if(ind==undefined){
			ind = 0;
		}
		var offset = ind * 20;
		var h = $( window ).width()/3*2;
		var left = $( window ).width()/2;
		var w = left - 20;	
		this.openHrefDownload(window, pm, "ViewXML", "location=0,menubar=0,status=0,titlebar=0,top="+(50+offset)+",left="+(left+offset)+",width="+w+",height="+h);
	}else{
		//download
		pm.download();	
		window.showTempNote("Загрузка файла",null,3000);
	}
}

App_nails.prototype.addWorkSrvCont = function(id, fl){
	var self = this;
	var pm = (new Attachment_Controller()).getPublicMethod("add_file");
	pm.setFieldValue("ref", CommonHelper.serialize({"keys": {"id": id}, "dataType": "specialist_works"}));
	pm.setFieldValue("content_info", CommonHelper.serialize({"id": CommonHelper.uniqid(), "name": "Фото", "size": 0}));
	pm.setFieldValue("content_data", [fl]);
	pm.run({
		"ok":function(){
			window.showTempOk("Фото загружено.",null,2000);
			var audio = new Audio('sounds/mixkit-long-pop-2358.mp3');
			audio.play();			
		}
	})
}

App_nails.prototype.addWorkSrv = function(documentId, fl){
	var pm = (new SpecialistWork_Controller()).getPublicMethod("insert");
	pm.setFieldValue("specialist_id", 0);
	pm.setFieldValue("studio_id", 0);
	pm.setFieldValue("ycl_document_id", documentId);
	var self = this;
	pm.run({
		"ok":function(resp){
			var m = resp.getModel("InsertedKey");
			if(m && m.getNextRow()){
				self.addWorkSrvCont(m.getFieldValue("id"), fl);
			}
		}
	});
}

App_nails.prototype.addPhoto = function(callback){
	var self = this;
	var input = document.createElement('input');
	input.type = 'file';
	input.capture = 'camera';
	input.accept = 'image/*';
	input.addEventListener("change", function(e) {
		callback(e.target.files[0]);
	});	
	input.click();		
}

App_nails.prototype.addWork = function(){
	//1) select transaction
	var self = this;
	var trans_view = new YClTransactionDocList_View("transactions:view", {
			"onSelectDoc": function(documentId){
				window.getApp().m_form.close();
				self.addPhoto(function(fl){
					self.addWorkSrv(documentId, fl);
				});
			}
	});
	this.m_form = new WindowFormModalBS("transactions",{
		"dialogWidth":"80%",
		"cmdOk":false,
		"cmdCancel":true,		
		"onClickCancel":function(){
			this.close();
		},
		"cmdClose":false,
		"content":trans_view
	});
	this.m_form.open();
	
	
	return false;
}



