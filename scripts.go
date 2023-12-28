package main

/**
 * Andrey Mikhalevich 15/12/21
 * This file is part of the OSBE framework
 * 
 * This file is generated from the template build/templates/scripts.go.tmpl
 * All direct modification will be lost with the next build.
 * Edit template instead.
*/

import(
	"path/filepath"
	
	"osbe/srv/httpSrv"
)

func (a *NailsApp) setCSS() {
	bs := a.GetBaseDir()
	a.CSSModel = httpSrv.NewLinkModel(9)
	a.CSSModel.Rows[0] = &httpSrv.TagLink{Href: "js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/icons/icomoon/styles.css",
		Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/icons/icomoon/styles.css")),
	}
	a.CSSModel.Rows[1] = &httpSrv.TagLink{Href: "js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/bootstrap.css",
		Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/bootstrap.css")),
	}
	a.CSSModel.Rows[2] = &httpSrv.TagLink{Href: "js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/core.min.css",
		Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/core.min.css")),
	}
	a.CSSModel.Rows[3] = &httpSrv.TagLink{Href: "js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/components.min.css",
		Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/components.min.css")),
	}
	a.CSSModel.Rows[4] = &httpSrv.TagLink{Href: "js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/colors.min.css",
		Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/colors.min.css")),
	}
	a.CSSModel.Rows[5] = &httpSrv.TagLink{Href: "js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/icons/fontawesome/styles.min.css",
		Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/icons/fontawesome/styles.min.css")),
	}
	a.CSSModel.Rows[6] = &httpSrv.TagLink{Href: "js20/ext/bootstrap-datepicker/bootstrap-datepicker.standalone.min.css",
		Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/bootstrap-datepicker/bootstrap-datepicker.standalone.min.css")),
	}
	a.CSSModel.Rows[7] = &httpSrv.TagLink{Href: "js20/custom-css/style.css",
		Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom-css/style.css")),
	}
	a.CSSModel.Rows[8] = &httpSrv.TagLink{Href: "js20/custom-css/print.css",
		Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom-css/print.css")),
	}
}

func (a *NailsApp) setJavaScript(debugMode bool) {
	bs := a.GetBaseDir()
	if debugMode {
		a.JavaScriptModel = httpSrv.NewScriptModel(608)
		a.JavaScriptModel.Rows[0] = &httpSrv.TagScript{Src: "js20/assets/js/core/libraries/jquery.min.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/assets/js/core/libraries/jquery.min.js")),
			}
		a.JavaScriptModel.Rows[1] = &httpSrv.TagScript{Src: "js20/assets/js/core/libraries/bootstrap.min.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/assets/js/core/libraries/bootstrap.min.js")),
			}
		a.JavaScriptModel.Rows[2] = &httpSrv.TagScript{Src: "js20/assets/js/plugins/loaders/blockui.min.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/assets/js/plugins/loaders/blockui.min.js")),
			}
		a.JavaScriptModel.Rows[3] = &httpSrv.TagScript{Src: "js20/assets/js/core/app.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/assets/js/core/app.js")),
			}
		a.JavaScriptModel.Rows[4] = &httpSrv.TagScript{Src: "js20/ext/bootstrap-datepicker/bootstrap-datepicker.min.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/bootstrap-datepicker/bootstrap-datepicker.min.js")),
			}
		a.JavaScriptModel.Rows[5] = &httpSrv.TagScript{Src: "js20/ext/bootstrap-datepicker/bootstrap-datepicker.ru.min.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/bootstrap-datepicker/bootstrap-datepicker.ru.min.js")),
			}
		a.JavaScriptModel.Rows[6] = &httpSrv.TagScript{Src: "js20/jquery.maskedinput.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/jquery.maskedinput.js")),
			}
		a.JavaScriptModel.Rows[7] = &httpSrv.TagScript{Src: "js20/ext/cleave/cleave.min.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/cleave/cleave.min.js")),
			}
		a.JavaScriptModel.Rows[8] = &httpSrv.TagScript{Src: "js20/ext/cleave/cleave-phone.ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/cleave/cleave-phone.ru.js")),
			}
		a.JavaScriptModel.Rows[9] = &httpSrv.TagScript{Src: "js20/ext/jshash-2.2/md5-min.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/jshash-2.2/md5-min.js")),
			}
		a.JavaScriptModel.Rows[10] = &httpSrv.TagScript{Src: "js20/ext/mustache/mustache.min.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/mustache/mustache.min.js")),
			}
		a.JavaScriptModel.Rows[11] = &httpSrv.TagScript{Src: "js20/ext/DragnDrop/DragMaster.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/DragnDrop/DragMaster.js")),
			}
		a.JavaScriptModel.Rows[12] = &httpSrv.TagScript{Src: "js20/ext/DragnDrop/DragObject.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/DragnDrop/DragObject.js")),
			}
		a.JavaScriptModel.Rows[13] = &httpSrv.TagScript{Src: "js20/ext/DragnDrop/DropTarget.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/DragnDrop/DropTarget.js")),
			}
		a.JavaScriptModel.Rows[14] = &httpSrv.TagScript{Src: "js20/assets/js/plugins/forms/styling/switchery.min.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/assets/js/plugins/forms/styling/switchery.min.js")),
			}
		a.JavaScriptModel.Rows[15] = &httpSrv.TagScript{Src: "js20/assets/js/plugins/forms/styling/uniform.min.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/assets/js/plugins/forms/styling/uniform.min.js")),
			}
		a.JavaScriptModel.Rows[16] = &httpSrv.TagScript{Src: "js20/core/extend.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/extend.js")),
			}
		a.JavaScriptModel.Rows[17] = &httpSrv.TagScript{Src: "js20/core/App.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/App.js")),
			}
		a.JavaScriptModel.Rows[18] = &httpSrv.TagScript{Src: "js20/core/AppWin.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/AppWin.js")),
			}
		a.JavaScriptModel.Rows[19] = &httpSrv.TagScript{Src: "js20/core/CommonHelper.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/CommonHelper.js")),
			}
		a.JavaScriptModel.Rows[20] = &httpSrv.TagScript{Src: "js20/core/DOMHelper.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/DOMHelper.js")),
			}
		a.JavaScriptModel.Rows[21] = &httpSrv.TagScript{Src: "js20/core/DateHelper.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/DateHelper.js")),
			}
		a.JavaScriptModel.Rows[22] = &httpSrv.TagScript{Src: "js20/core/EventHelper.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/EventHelper.js")),
			}
		a.JavaScriptModel.Rows[23] = &httpSrv.TagScript{Src: "js20/core/FatalException.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FatalException.js")),
			}
		a.JavaScriptModel.Rows[24] = &httpSrv.TagScript{Src: "js20/core/DbException.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/DbException.js")),
			}
		a.JavaScriptModel.Rows[25] = &httpSrv.TagScript{Src: "js20/core/VersException.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/VersException.js")),
			}
		a.JavaScriptModel.Rows[26] = &httpSrv.TagScript{Src: "js20/core/ConstantManager.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ConstantManager.js")),
			}
		a.JavaScriptModel.Rows[27] = &httpSrv.TagScript{Src: "js20/core/AppSrv.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/AppSrv.js")),
			}
		a.JavaScriptModel.Rows[28] = &httpSrv.TagScript{Src: "js20/core/AppSrv.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/AppSrv.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[29] = &httpSrv.TagScript{Src: "js20/core/CookieManager.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/CookieManager.js")),
			}
		a.JavaScriptModel.Rows[30] = &httpSrv.TagScript{Src: "js20/core/ServConnector.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ServConnector.js")),
			}
		a.JavaScriptModel.Rows[31] = &httpSrv.TagScript{Src: "js20/core/Response.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/Response.js")),
			}
		a.JavaScriptModel.Rows[32] = &httpSrv.TagScript{Src: "js20/core/ResponseXML.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ResponseXML.js")),
			}
		a.JavaScriptModel.Rows[33] = &httpSrv.TagScript{Src: "js20/core/ResponseJSON.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ResponseJSON.js")),
			}
		a.JavaScriptModel.Rows[34] = &httpSrv.TagScript{Src: "js20/core/PublicMethod.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/PublicMethod.js")),
			}
		a.JavaScriptModel.Rows[35] = &httpSrv.TagScript{Src: "js20/core/PublicMethodServer.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/PublicMethodServer.js")),
			}
		a.JavaScriptModel.Rows[36] = &httpSrv.TagScript{Src: "js20/core/Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/Controller.js")),
			}
		a.JavaScriptModel.Rows[37] = &httpSrv.TagScript{Src: "js20/core/ControllerObj.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ControllerObj.js")),
			}
		a.JavaScriptModel.Rows[38] = &httpSrv.TagScript{Src: "js20/core/ControllerObjServer.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ControllerObjServer.js")),
			}
		a.JavaScriptModel.Rows[39] = &httpSrv.TagScript{Src: "js20/core/ControllerObjClient.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ControllerObjClient.js")),
			}
		a.JavaScriptModel.Rows[40] = &httpSrv.TagScript{Src: "js20/core/Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/Model.js")),
			}
		a.JavaScriptModel.Rows[41] = &httpSrv.TagScript{Src: "js20/core/ModelXML.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ModelXML.js")),
			}
		a.JavaScriptModel.Rows[42] = &httpSrv.TagScript{Src: "js20/core/ModelObjectXML.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ModelObjectXML.js")),
			}
		a.JavaScriptModel.Rows[43] = &httpSrv.TagScript{Src: "js20/core/ModelServRespXML.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ModelServRespXML.js")),
			}
		a.JavaScriptModel.Rows[44] = &httpSrv.TagScript{Src: "js20/core/ModelJSON.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ModelJSON.js")),
			}
		a.JavaScriptModel.Rows[45] = &httpSrv.TagScript{Src: "js20/core/ModelObjectJSON.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ModelObjectJSON.js")),
			}
		a.JavaScriptModel.Rows[46] = &httpSrv.TagScript{Src: "js20/core/ModelServRespJSON.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ModelServRespJSON.js")),
			}
		a.JavaScriptModel.Rows[47] = &httpSrv.TagScript{Src: "js20/core/ModelXMLTree.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ModelXMLTree.js")),
			}
		a.JavaScriptModel.Rows[48] = &httpSrv.TagScript{Src: "js20/core/Validator.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/Validator.js")),
			}
		a.JavaScriptModel.Rows[49] = &httpSrv.TagScript{Src: "js20/core/ValidatorString.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ValidatorString.js")),
			}
		a.JavaScriptModel.Rows[50] = &httpSrv.TagScript{Src: "js20/core/ValidatorBool.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ValidatorBool.js")),
			}
		a.JavaScriptModel.Rows[51] = &httpSrv.TagScript{Src: "js20/core/ValidatorDate.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ValidatorDate.js")),
			}
		a.JavaScriptModel.Rows[52] = &httpSrv.TagScript{Src: "js20/core/ValidatorDateTime.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ValidatorDateTime.js")),
			}
		a.JavaScriptModel.Rows[53] = &httpSrv.TagScript{Src: "js20/core/ValidatorTime.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ValidatorTime.js")),
			}
		a.JavaScriptModel.Rows[54] = &httpSrv.TagScript{Src: "js20/core/ValidatorInt.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ValidatorInt.js")),
			}
		a.JavaScriptModel.Rows[55] = &httpSrv.TagScript{Src: "js20/core/ValidatorFloat.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ValidatorFloat.js")),
			}
		a.JavaScriptModel.Rows[56] = &httpSrv.TagScript{Src: "js20/core/ValidatorEnum.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ValidatorEnum.js")),
			}
		a.JavaScriptModel.Rows[57] = &httpSrv.TagScript{Src: "js20/core/ValidatorJSON.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ValidatorJSON.js")),
			}
		a.JavaScriptModel.Rows[58] = &httpSrv.TagScript{Src: "js20/core/ValidatorArray.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ValidatorArray.js")),
			}
		a.JavaScriptModel.Rows[59] = &httpSrv.TagScript{Src: "js20/core/ValidatorEmail.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ValidatorEmail.js")),
			}
		a.JavaScriptModel.Rows[60] = &httpSrv.TagScript{Src: "js20/core/Field.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/Field.js")),
			}
		a.JavaScriptModel.Rows[61] = &httpSrv.TagScript{Src: "js20/core/FieldString.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldString.js")),
			}
		a.JavaScriptModel.Rows[62] = &httpSrv.TagScript{Src: "js20/core/FieldEnum.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldEnum.js")),
			}
		a.JavaScriptModel.Rows[63] = &httpSrv.TagScript{Src: "js20/core/FieldBool.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldBool.js")),
			}
		a.JavaScriptModel.Rows[64] = &httpSrv.TagScript{Src: "js20/core/FieldDate.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldDate.js")),
			}
		a.JavaScriptModel.Rows[65] = &httpSrv.TagScript{Src: "js20/core/FieldDateTime.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldDateTime.js")),
			}
		a.JavaScriptModel.Rows[66] = &httpSrv.TagScript{Src: "js20/core/FieldDateTimeTZ.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldDateTimeTZ.js")),
			}
		a.JavaScriptModel.Rows[67] = &httpSrv.TagScript{Src: "js20/core/FieldTime.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldTime.js")),
			}
		a.JavaScriptModel.Rows[68] = &httpSrv.TagScript{Src: "js20/core/FieldInt.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldInt.js")),
			}
		a.JavaScriptModel.Rows[69] = &httpSrv.TagScript{Src: "js20/core/FieldBigInt.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldBigInt.js")),
			}
		a.JavaScriptModel.Rows[70] = &httpSrv.TagScript{Src: "js20/core/FieldSmallInt.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldSmallInt.js")),
			}
		a.JavaScriptModel.Rows[71] = &httpSrv.TagScript{Src: "js20/core/FieldFloat.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldFloat.js")),
			}
		a.JavaScriptModel.Rows[72] = &httpSrv.TagScript{Src: "js20/core/FieldPassword.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldPassword.js")),
			}
		a.JavaScriptModel.Rows[73] = &httpSrv.TagScript{Src: "js20/core/FieldText.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldText.js")),
			}
		a.JavaScriptModel.Rows[74] = &httpSrv.TagScript{Src: "js20/core/FieldInterval.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldInterval.js")),
			}
		a.JavaScriptModel.Rows[75] = &httpSrv.TagScript{Src: "js20/core/FieldJSON.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldJSON.js")),
			}
		a.JavaScriptModel.Rows[76] = &httpSrv.TagScript{Src: "js20/core/FieldJSONB.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldJSONB.js")),
			}
		a.JavaScriptModel.Rows[77] = &httpSrv.TagScript{Src: "js20/core/FieldArray.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldArray.js")),
			}
		a.JavaScriptModel.Rows[78] = &httpSrv.TagScript{Src: "js20/core/FieldBytea.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/FieldBytea.js")),
			}
		a.JavaScriptModel.Rows[79] = &httpSrv.TagScript{Src: "js20/core/ModelFilter.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/ModelFilter.js")),
			}
		a.JavaScriptModel.Rows[80] = &httpSrv.TagScript{Src: "js20/core/RefType.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/RefType.js")),
			}
		a.JavaScriptModel.Rows[81] = &httpSrv.TagScript{Src: "js20/core/rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/core/rs_ru.js")),
			}
		a.JavaScriptModel.Rows[82] = &httpSrv.TagScript{Src: "js20/controls/DataBinding.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/DataBinding.js")),
			}
		a.JavaScriptModel.Rows[83] = &httpSrv.TagScript{Src: "js20/controls/Command.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/Command.js")),
			}
		a.JavaScriptModel.Rows[84] = &httpSrv.TagScript{Src: "js20/controls/CommandBinding.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/CommandBinding.js")),
			}
		a.JavaScriptModel.Rows[85] = &httpSrv.TagScript{Src: "js20/controls/Control.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/Control.js")),
			}
		a.JavaScriptModel.Rows[86] = &httpSrv.TagScript{Src: "js20/controls/rs/Control.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/Control.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[87] = &httpSrv.TagScript{Src: "js20/controls/ControlContainer.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ControlContainer.js")),
			}
		a.JavaScriptModel.Rows[88] = &httpSrv.TagScript{Src: "js20/controls/rs/ControlContainer.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ControlContainer.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[89] = &httpSrv.TagScript{Src: "js20/controls/View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/View.js")),
			}
		a.JavaScriptModel.Rows[90] = &httpSrv.TagScript{Src: "js20/controls/ViewAjx.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ViewAjx.js")),
			}
		a.JavaScriptModel.Rows[91] = &httpSrv.TagScript{Src: "js20/controls/rs/ViewAjx.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ViewAjx.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[92] = &httpSrv.TagScript{Src: "js20/controls/ViewAjxList.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ViewAjxList.js")),
			}
		a.JavaScriptModel.Rows[93] = &httpSrv.TagScript{Src: "js20/controls/Calculator.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/Calculator.js")),
			}
		a.JavaScriptModel.Rows[94] = &httpSrv.TagScript{Src: "js20/controls/rs/Calculator.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/Calculator.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[95] = &httpSrv.TagScript{Src: "js20/controls/Button.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/Button.js")),
			}
		a.JavaScriptModel.Rows[96] = &httpSrv.TagScript{Src: "js20/controls/ButtonCtrl.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonCtrl.js")),
			}
		a.JavaScriptModel.Rows[97] = &httpSrv.TagScript{Src: "js20/controls/ButtonEditCtrl.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonEditCtrl.js")),
			}
		a.JavaScriptModel.Rows[98] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonEditCtrl.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonEditCtrl.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[99] = &httpSrv.TagScript{Src: "js20/controls/ButtonCalc.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonCalc.js")),
			}
		a.JavaScriptModel.Rows[100] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonCalc.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonCalc.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[101] = &httpSrv.TagScript{Src: "js20/controls/ButtonCalendar.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonCalendar.js")),
			}
		a.JavaScriptModel.Rows[102] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonCalendar.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonCalendar.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[103] = &httpSrv.TagScript{Src: "js20/controls/ButtonClear.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonClear.js")),
			}
		a.JavaScriptModel.Rows[104] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonClear.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonClear.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[105] = &httpSrv.TagScript{Src: "js20/custom_controls/ButtonCmd.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/ButtonCmd.js")),
			}
		a.JavaScriptModel.Rows[106] = &httpSrv.TagScript{Src: "js20/controls/ButtonExpToExcel.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonExpToExcel.js")),
			}
		a.JavaScriptModel.Rows[107] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonExpToExcel.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonExpToExcel.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[108] = &httpSrv.TagScript{Src: "js20/controls/ButtonExpToPDF.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonExpToPDF.js")),
			}
		a.JavaScriptModel.Rows[109] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonExpToPDF.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonExpToPDF.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[110] = &httpSrv.TagScript{Src: "js20/controls/ButtonOpen.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonOpen.js")),
			}
		a.JavaScriptModel.Rows[111] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonOpen.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonOpen.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[112] = &httpSrv.TagScript{Src: "js20/controls/ButtonInsert.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonInsert.js")),
			}
		a.JavaScriptModel.Rows[113] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonInsert.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonInsert.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[114] = &httpSrv.TagScript{Src: "js20/controls/ButtonPrint.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonPrint.js")),
			}
		a.JavaScriptModel.Rows[115] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonPrint.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonPrint.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[116] = &httpSrv.TagScript{Src: "js20/controls/ButtonPrintList.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonPrintList.js")),
			}
		a.JavaScriptModel.Rows[117] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonPrintList.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonPrintList.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[118] = &httpSrv.TagScript{Src: "js20/controls/ButtonSelectRef.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonSelectRef.js")),
			}
		a.JavaScriptModel.Rows[119] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonSelectRef.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonSelectRef.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[120] = &httpSrv.TagScript{Src: "js20/controls/ButtonSelectDataType.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonSelectDataType.js")),
			}
		a.JavaScriptModel.Rows[121] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonSelectDataType.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonSelectDataType.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[122] = &httpSrv.TagScript{Src: "js20/controls/ButtonMakeSelection.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonMakeSelection.js")),
			}
		a.JavaScriptModel.Rows[123] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonMakeSelection.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonMakeSelection.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[124] = &httpSrv.TagScript{Src: "js20/controls/ButtonToggle.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonToggle.js")),
			}
		a.JavaScriptModel.Rows[125] = &httpSrv.TagScript{Src: "js20/controls/Label.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/Label.js")),
			}
		a.JavaScriptModel.Rows[126] = &httpSrv.TagScript{Src: "js20/controls/Edit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/Edit.js")),
			}
		a.JavaScriptModel.Rows[127] = &httpSrv.TagScript{Src: "js20/controls/rs/Edit.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/Edit.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[128] = &httpSrv.TagScript{Src: "js20/controls/EditRef.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditRef.js")),
			}
		a.JavaScriptModel.Rows[129] = &httpSrv.TagScript{Src: "js20/controls/EditComplete.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditComplete.js")),
			}
		a.JavaScriptModel.Rows[130] = &httpSrv.TagScript{Src: "js20/controls/EditString.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditString.js")),
			}
		a.JavaScriptModel.Rows[131] = &httpSrv.TagScript{Src: "js20/controls/EditText.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditText.js")),
			}
		a.JavaScriptModel.Rows[132] = &httpSrv.TagScript{Src: "js20/controls/EditNum.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditNum.js")),
			}
		a.JavaScriptModel.Rows[133] = &httpSrv.TagScript{Src: "js20/controls/EditInt.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditInt.js")),
			}
		a.JavaScriptModel.Rows[134] = &httpSrv.TagScript{Src: "js20/controls/EditFloat.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditFloat.js")),
			}
		a.JavaScriptModel.Rows[135] = &httpSrv.TagScript{Src: "js20/controls/EditMoney.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditMoney.js")),
			}
		a.JavaScriptModel.Rows[136] = &httpSrv.TagScript{Src: "js20/controls/EditPhone.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditPhone.js")),
			}
		a.JavaScriptModel.Rows[137] = &httpSrv.TagScript{Src: "js20/controls/EditEmail.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditEmail.js")),
			}
		a.JavaScriptModel.Rows[138] = &httpSrv.TagScript{Src: "js20/controls/EditPercent.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditPercent.js")),
			}
		a.JavaScriptModel.Rows[139] = &httpSrv.TagScript{Src: "js20/controls/EditDate.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditDate.js")),
			}
		a.JavaScriptModel.Rows[140] = &httpSrv.TagScript{Src: "js20/controls/EditDateTime.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditDateTime.js")),
			}
		a.JavaScriptModel.Rows[141] = &httpSrv.TagScript{Src: "js20/controls/EditTime.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditTime.js")),
			}
		a.JavaScriptModel.Rows[142] = &httpSrv.TagScript{Src: "js20/controls/EditPassword.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditPassword.js")),
			}
		a.JavaScriptModel.Rows[143] = &httpSrv.TagScript{Src: "js20/controls/EditCheckBox.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditCheckBox.js")),
			}
		a.JavaScriptModel.Rows[144] = &httpSrv.TagScript{Src: "js20/controls/EditContainer.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditContainer.js")),
			}
		a.JavaScriptModel.Rows[145] = &httpSrv.TagScript{Src: "js20/controls/rs/EditContainer.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/EditContainer.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[146] = &httpSrv.TagScript{Src: "js20/controls/EditRadioGroup.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditRadioGroup.js")),
			}
		a.JavaScriptModel.Rows[147] = &httpSrv.TagScript{Src: "js20/controls/EditRadio.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditRadio.js")),
			}
		a.JavaScriptModel.Rows[148] = &httpSrv.TagScript{Src: "js20/controls/EditSelect.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditSelect.js")),
			}
		a.JavaScriptModel.Rows[149] = &httpSrv.TagScript{Src: "js20/controls/EditSelectRef.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditSelectRef.js")),
			}
		a.JavaScriptModel.Rows[150] = &httpSrv.TagScript{Src: "js20/controls/rs/EditSelectRef.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/EditSelectRef.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[151] = &httpSrv.TagScript{Src: "js20/controls/EditSelectOption.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditSelectOption.js")),
			}
		a.JavaScriptModel.Rows[152] = &httpSrv.TagScript{Src: "js20/controls/EditSelectOptionRef.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditSelectOptionRef.js")),
			}
		a.JavaScriptModel.Rows[153] = &httpSrv.TagScript{Src: "js20/controls/EditRadioGroupRef.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditRadioGroupRef.js")),
			}
		a.JavaScriptModel.Rows[154] = &httpSrv.TagScript{Src: "js20/controls/rs/EditRadioGroupRef.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/EditRadioGroupRef.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[155] = &httpSrv.TagScript{Src: "js20/controls/PrintObj.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/PrintObj.js")),
			}
		a.JavaScriptModel.Rows[156] = &httpSrv.TagScript{Src: "js20/controls/ControlForm.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ControlForm.js")),
			}
		a.JavaScriptModel.Rows[157] = &httpSrv.TagScript{Src: "js20/controls/HiddenKey.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/HiddenKey.js")),
			}
		a.JavaScriptModel.Rows[158] = &httpSrv.TagScript{Src: "js20/controls/EditJSON.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditJSON.js")),
			}
		a.JavaScriptModel.Rows[159] = &httpSrv.TagScript{Src: "js20/controls/EditModalDialog.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditModalDialog.js")),
			}
		a.JavaScriptModel.Rows[160] = &httpSrv.TagScript{Src: "js20/controls/rs/EditModalDialog.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/EditModalDialog.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[161] = &httpSrv.TagScript{Src: "js20/controls/GridColumn.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridColumn.js")),
			}
		a.JavaScriptModel.Rows[162] = &httpSrv.TagScript{Src: "js20/controls/GridColumnBool.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridColumnBool.js")),
			}
		a.JavaScriptModel.Rows[163] = &httpSrv.TagScript{Src: "js20/controls/GridColumnPhone.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridColumnPhone.js")),
			}
		a.JavaScriptModel.Rows[164] = &httpSrv.TagScript{Src: "js20/controls/GridColumnEmail.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridColumnEmail.js")),
			}
		a.JavaScriptModel.Rows[165] = &httpSrv.TagScript{Src: "js20/controls/GridColumnFloat.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridColumnFloat.js")),
			}
		a.JavaScriptModel.Rows[166] = &httpSrv.TagScript{Src: "js20/controls/GridColumnByte.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridColumnByte.js")),
			}
		a.JavaScriptModel.Rows[167] = &httpSrv.TagScript{Src: "js20/controls/GridColumnDate.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridColumnDate.js")),
			}
		a.JavaScriptModel.Rows[168] = &httpSrv.TagScript{Src: "js20/controls/GridColumnDateTime.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridColumnDateTime.js")),
			}
		a.JavaScriptModel.Rows[169] = &httpSrv.TagScript{Src: "js20/controls/GridColumnEnum.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridColumnEnum.js")),
			}
		a.JavaScriptModel.Rows[170] = &httpSrv.TagScript{Src: "js20/controls/GridColumnRef.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridColumnRef.js")),
			}
		a.JavaScriptModel.Rows[171] = &httpSrv.TagScript{Src: "js20/controls/GridColumnPicture.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridColumnPicture.js")),
			}
		a.JavaScriptModel.Rows[172] = &httpSrv.TagScript{Src: "js20/controls/GridCell.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCell.js")),
			}
		a.JavaScriptModel.Rows[173] = &httpSrv.TagScript{Src: "js20/controls/GridCellHead.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCellHead.js")),
			}
		a.JavaScriptModel.Rows[174] = &httpSrv.TagScript{Src: "js20/controls/GridCellHeadMark.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCellHeadMark.js")),
			}
		a.JavaScriptModel.Rows[175] = &httpSrv.TagScript{Src: "js20/controls/GridCellFoot.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCellFoot.js")),
			}
		a.JavaScriptModel.Rows[176] = &httpSrv.TagScript{Src: "js20/controls/GridCellPhone.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCellPhone.js")),
			}
		a.JavaScriptModel.Rows[177] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCellPhone.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCellPhone.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[178] = &httpSrv.TagScript{Src: "js20/controls/GridCellDetailToggle.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCellDetailToggle.js")),
			}
		a.JavaScriptModel.Rows[179] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCellDetailToggle.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCellDetailToggle.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[180] = &httpSrv.TagScript{Src: "js20/controls/GridHead.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridHead.js")),
			}
		a.JavaScriptModel.Rows[181] = &httpSrv.TagScript{Src: "js20/controls/GridRow.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridRow.js")),
			}
		a.JavaScriptModel.Rows[182] = &httpSrv.TagScript{Src: "js20/controls/GridFoot.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridFoot.js")),
			}
		a.JavaScriptModel.Rows[183] = &httpSrv.TagScript{Src: "js20/controls/GridBody.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridBody.js")),
			}
		a.JavaScriptModel.Rows[184] = &httpSrv.TagScript{Src: "js20/controls/Grid.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/Grid.js")),
			}
		a.JavaScriptModel.Rows[185] = &httpSrv.TagScript{Src: "js20/controls/rs/Grid.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/Grid.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[186] = &httpSrv.TagScript{Src: "js20/controls/GridSearchInf.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridSearchInf.js")),
			}
		a.JavaScriptModel.Rows[187] = &httpSrv.TagScript{Src: "js20/controls/rs/GridSearchInf.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridSearchInf.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[188] = &httpSrv.TagScript{Src: "js20/controls/VariantStorage.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/VariantStorage.js")),
			}
		a.JavaScriptModel.Rows[189] = &httpSrv.TagScript{Src: "js20/controls/GridCommands.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCommands.js")),
			}
		a.JavaScriptModel.Rows[190] = &httpSrv.TagScript{Src: "js20/controls/GridCmd.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmd.js")),
			}
		a.JavaScriptModel.Rows[191] = &httpSrv.TagScript{Src: "js20/controls/GridCmdContainer.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdContainer.js")),
			}
		a.JavaScriptModel.Rows[192] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdContainer.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdContainer.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[193] = &httpSrv.TagScript{Src: "js20/controls/GridCmdContainerAjx.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdContainerAjx.js")),
			}
		a.JavaScriptModel.Rows[194] = &httpSrv.TagScript{Src: "js20/controls/GridCmdContainerDOC.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdContainerDOC.js")),
			}
		a.JavaScriptModel.Rows[195] = &httpSrv.TagScript{Src: "js20/controls/GridCmdInsert.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdInsert.js")),
			}
		a.JavaScriptModel.Rows[196] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdInsert.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdInsert.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[197] = &httpSrv.TagScript{Src: "js20/controls/GridCmdEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdEdit.js")),
			}
		a.JavaScriptModel.Rows[198] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdEdit.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdEdit.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[199] = &httpSrv.TagScript{Src: "js20/controls/GridCmdCopy.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdCopy.js")),
			}
		a.JavaScriptModel.Rows[200] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdCopy.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdCopy.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[201] = &httpSrv.TagScript{Src: "js20/controls/GridCmdDelete.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdDelete.js")),
			}
		a.JavaScriptModel.Rows[202] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdDelete.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdDelete.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[203] = &httpSrv.TagScript{Src: "js20/controls/GridCmdColManager.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdColManager.js")),
			}
		a.JavaScriptModel.Rows[204] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdColManager.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdColManager.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[205] = &httpSrv.TagScript{Src: "js20/controls/GridCmdPrint.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdPrint.js")),
			}
		a.JavaScriptModel.Rows[206] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdPrint.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdPrint.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[207] = &httpSrv.TagScript{Src: "js20/controls/GridCmdRefresh.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdRefresh.js")),
			}
		a.JavaScriptModel.Rows[208] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdRefresh.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdRefresh.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[209] = &httpSrv.TagScript{Src: "js20/controls/GridCmdPrintObj.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdPrintObj.js")),
			}
		a.JavaScriptModel.Rows[210] = &httpSrv.TagScript{Src: "js20/controls/GridCmdSearch.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdSearch.js")),
			}
		a.JavaScriptModel.Rows[211] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdSearch.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdSearch.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[212] = &httpSrv.TagScript{Src: "js20/controls/GridCmdExport.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdExport.js")),
			}
		a.JavaScriptModel.Rows[213] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdExport.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdExport.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[214] = &httpSrv.TagScript{Src: "js20/controls/GridCmdAllCommands.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdAllCommands.js")),
			}
		a.JavaScriptModel.Rows[215] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdAllCommands.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdAllCommands.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[216] = &httpSrv.TagScript{Src: "js20/controls/GridCmdDOCUnprocess.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdDOCUnprocess.js")),
			}
		a.JavaScriptModel.Rows[217] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdDOCUnprocess.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdDOCUnprocess.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[218] = &httpSrv.TagScript{Src: "js20/controls/GridCmdDOCShowActs.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdDOCShowActs.js")),
			}
		a.JavaScriptModel.Rows[219] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdDOCShowActs.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdDOCShowActs.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[220] = &httpSrv.TagScript{Src: "js20/controls/GridCmdRowUp.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdRowUp.js")),
			}
		a.JavaScriptModel.Rows[221] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdRowUp.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdRowUp.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[222] = &httpSrv.TagScript{Src: "js20/controls/GridCmdRowDown.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdRowDown.js")),
			}
		a.JavaScriptModel.Rows[223] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdRowDown.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdRowDown.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[224] = &httpSrv.TagScript{Src: "js20/controls/GridCmdFilter.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdFilter.js")),
			}
		a.JavaScriptModel.Rows[225] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdFilter.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdFilter.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[226] = &httpSrv.TagScript{Src: "js20/controls/GridCmdFilterView.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdFilterView.js")),
			}
		a.JavaScriptModel.Rows[227] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdFilterView.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdFilterView.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[228] = &httpSrv.TagScript{Src: "js20/controls/GridCmdFilterSave.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdFilterSave.js")),
			}
		a.JavaScriptModel.Rows[229] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdFilterSave.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdFilterSave.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[230] = &httpSrv.TagScript{Src: "js20/controls/GridCmdFilterOpen.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCmdFilterOpen.js")),
			}
		a.JavaScriptModel.Rows[231] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCmdFilterOpen.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCmdFilterOpen.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[232] = &httpSrv.TagScript{Src: "js20/controls/ViewGridColManager.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ViewGridColManager.js")),
			}
		a.JavaScriptModel.Rows[233] = &httpSrv.TagScript{Src: "js20/controls/rs/ViewGridColManager.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ViewGridColManager.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[234] = &httpSrv.TagScript{Src: "js20/controls/ViewGridColParam.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ViewGridColParam.js")),
			}
		a.JavaScriptModel.Rows[235] = &httpSrv.TagScript{Src: "js20/controls/rs/ViewGridColParam.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ViewGridColParam.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[236] = &httpSrv.TagScript{Src: "js20/controls/ViewGridColVisibility.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ViewGridColVisibility.js")),
			}
		a.JavaScriptModel.Rows[237] = &httpSrv.TagScript{Src: "js20/controls/rs/ViewGridColVisibility.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ViewGridColVisibility.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[238] = &httpSrv.TagScript{Src: "js20/controls/ViewGridColOrder.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ViewGridColOrder.js")),
			}
		a.JavaScriptModel.Rows[239] = &httpSrv.TagScript{Src: "js20/controls/rs/ViewGridColOrder.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ViewGridColOrder.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[240] = &httpSrv.TagScript{Src: "js20/controls/VariantStorageSaveView.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/VariantStorageSaveView.js")),
			}
		a.JavaScriptModel.Rows[241] = &httpSrv.TagScript{Src: "js20/controls/rs/VariantStorageSaveView.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/VariantStorageSaveView.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[242] = &httpSrv.TagScript{Src: "js20/controls/VariantStorageOpenView.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/VariantStorageOpenView.js")),
			}
		a.JavaScriptModel.Rows[243] = &httpSrv.TagScript{Src: "js20/controls/rs/VariantStorageOpenView.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/VariantStorageOpenView.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[244] = &httpSrv.TagScript{Src: "js20/controls/GridAjx.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridAjx.js")),
			}
		a.JavaScriptModel.Rows[245] = &httpSrv.TagScript{Src: "js20/controls/rs/GridAjx.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridAjx.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[246] = &httpSrv.TagScript{Src: "js20/controls/TreeAjx.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/TreeAjx.js")),
			}
		a.JavaScriptModel.Rows[247] = &httpSrv.TagScript{Src: "js20/controls/GridAjxDOCT.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridAjxDOCT.js")),
			}
		a.JavaScriptModel.Rows[248] = &httpSrv.TagScript{Src: "js20/controls/GridAjxMaster.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridAjxMaster.js")),
			}
		a.JavaScriptModel.Rows[249] = &httpSrv.TagScript{Src: "js20/controls/GridCommandsAjx.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCommandsAjx.js")),
			}
		a.JavaScriptModel.Rows[250] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCommandsAjx.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCommandsAjx.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[251] = &httpSrv.TagScript{Src: "js20/controls/GridCommandsDOC.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCommandsDOC.js")),
			}
		a.JavaScriptModel.Rows[252] = &httpSrv.TagScript{Src: "js20/controls/rs/GridCommandsDOC.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridCommandsDOC.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[253] = &httpSrv.TagScript{Src: "js20/controls/GridPagination.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridPagination.js")),
			}
		a.JavaScriptModel.Rows[254] = &httpSrv.TagScript{Src: "js20/controls/rs/GridPagination.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridPagination.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[255] = &httpSrv.TagScript{Src: "js20/controls/GridFilter.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridFilter.js")),
			}
		a.JavaScriptModel.Rows[256] = &httpSrv.TagScript{Src: "js20/controls/rs/GridFilter.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/GridFilter.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[257] = &httpSrv.TagScript{Src: "js20/controls/EditPeriodDate.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditPeriodDate.js")),
			}
		a.JavaScriptModel.Rows[258] = &httpSrv.TagScript{Src: "js20/controls/rs/EditPeriodDate.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/EditPeriodDate.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[259] = &httpSrv.TagScript{Src: "js20/controls/EditPeriodDateTime.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditPeriodDateTime.js")),
			}
		a.JavaScriptModel.Rows[260] = &httpSrv.TagScript{Src: "js20/controls/ConstantGrid.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ConstantGrid.js")),
			}
		a.JavaScriptModel.Rows[261] = &httpSrv.TagScript{Src: "js20/controls/rs/ConstantGrid.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ConstantGrid.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[262] = &httpSrv.TagScript{Src: "js20/controls/ButtonOK.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonOK.js")),
			}
		a.JavaScriptModel.Rows[263] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonOK.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonOK.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[264] = &httpSrv.TagScript{Src: "js20/controls/ButtonSave.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonSave.js")),
			}
		a.JavaScriptModel.Rows[265] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonSave.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonSave.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[266] = &httpSrv.TagScript{Src: "js20/controls/ButtonCancel.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonCancel.js")),
			}
		a.JavaScriptModel.Rows[267] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonCancel.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonCancel.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[268] = &httpSrv.TagScript{Src: "js20/controls/ViewObjectAjx.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ViewObjectAjx.js")),
			}
		a.JavaScriptModel.Rows[269] = &httpSrv.TagScript{Src: "js20/controls/rs/ViewObjectAjx.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ViewObjectAjx.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[270] = &httpSrv.TagScript{Src: "js20/controls/ViewGridEditInlineAjx.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ViewGridEditInlineAjx.js")),
			}
		a.JavaScriptModel.Rows[271] = &httpSrv.TagScript{Src: "js20/controls/rs/ViewGridEditInlineAjx.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ViewGridEditInlineAjx.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[272] = &httpSrv.TagScript{Src: "js20/controls/ViewDOC.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ViewDOC.js")),
			}
		a.JavaScriptModel.Rows[273] = &httpSrv.TagScript{Src: "js20/controls/rs/ViewDOC.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ViewDOC.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[274] = &httpSrv.TagScript{Src: "js20/controls/WindowPrint.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/WindowPrint.js")),
			}
		a.JavaScriptModel.Rows[275] = &httpSrv.TagScript{Src: "js20/controls/rs/WindowPrint.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/WindowPrint.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[276] = &httpSrv.TagScript{Src: "js20/controls/WindowQuestion.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/WindowQuestion.js")),
			}
		a.JavaScriptModel.Rows[277] = &httpSrv.TagScript{Src: "js20/controls/rs/WindowQuestion.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/WindowQuestion.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[278] = &httpSrv.TagScript{Src: "js20/controls/WindowSearch.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/WindowSearch.js")),
			}
		a.JavaScriptModel.Rows[279] = &httpSrv.TagScript{Src: "js20/controls/rs/WindowSearch.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/WindowSearch.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[280] = &httpSrv.TagScript{Src: "js20/controls/WindowForm.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/WindowForm.js")),
			}
		a.JavaScriptModel.Rows[281] = &httpSrv.TagScript{Src: "js20/controls/rs/WindowForm.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/WindowForm.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[282] = &httpSrv.TagScript{Src: "js20/controls/WindowFormObject.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/WindowFormObject.js")),
			}
		a.JavaScriptModel.Rows[283] = &httpSrv.TagScript{Src: "js20/controls/rs/WindowFormObject.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/WindowFormObject.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[284] = &httpSrv.TagScript{Src: "js20/controls/WindowFormModalBS.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/WindowFormModalBS.js")),
			}
		a.JavaScriptModel.Rows[285] = &httpSrv.TagScript{Src: "js20/controls/rs/WindowFormModalBS.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/WindowFormModalBS.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[286] = &httpSrv.TagScript{Src: "js20/controls/WindowMessage.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/WindowMessage.js")),
			}
		a.JavaScriptModel.Rows[287] = &httpSrv.TagScript{Src: "js20/controls/WindowTempMessage.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/WindowTempMessage.js")),
			}
		a.JavaScriptModel.Rows[288] = &httpSrv.TagScript{Src: "js20/controls/GridCellHeadDOCProcessed.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCellHeadDOCProcessed.js")),
			}
		a.JavaScriptModel.Rows[289] = &httpSrv.TagScript{Src: "js20/controls/GridCellHeadDOCDate.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCellHeadDOCDate.js")),
			}
		a.JavaScriptModel.Rows[290] = &httpSrv.TagScript{Src: "js20/controls/GridCellHeadDOCNumber.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/GridCellHeadDOCNumber.js")),
			}
		a.JavaScriptModel.Rows[291] = &httpSrv.TagScript{Src: "js20/controls/actb.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/actb.js")),
			}
		a.JavaScriptModel.Rows[292] = &httpSrv.TagScript{Src: "js20/controls/rs/actb.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/actb.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[293] = &httpSrv.TagScript{Src: "js20/controls/RepCommands.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/RepCommands.js")),
			}
		a.JavaScriptModel.Rows[294] = &httpSrv.TagScript{Src: "js20/controls/rs/RepCommands.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/RepCommands.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[295] = &httpSrv.TagScript{Src: "js20/controls/ViewReport.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ViewReport.js")),
			}
		a.JavaScriptModel.Rows[296] = &httpSrv.TagScript{Src: "js20/controls/rs/ViewReport.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ViewReport.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[297] = &httpSrv.TagScript{Src: "js20/controls/PopUpMenu.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/PopUpMenu.js")),
			}
		a.JavaScriptModel.Rows[298] = &httpSrv.TagScript{Src: "js20/controls/PopOver.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/PopOver.js")),
			}
		a.JavaScriptModel.Rows[299] = &httpSrv.TagScript{Src: "js20/controls/PeriodSelect.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/PeriodSelect.js")),
			}
		a.JavaScriptModel.Rows[300] = &httpSrv.TagScript{Src: "js20/controls/rs/PeriodSelect.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/PeriodSelect.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[301] = &httpSrv.TagScript{Src: "js20/controls/WindowAbout.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/WindowAbout.js")),
			}
		a.JavaScriptModel.Rows[302] = &httpSrv.TagScript{Src: "js20/controls/rs/WindowAbout.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/WindowAbout.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[303] = &httpSrv.TagScript{Src: "js20/controls/MainMenuTree.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/MainMenuTree.js")),
			}
		a.JavaScriptModel.Rows[304] = &httpSrv.TagScript{Src: "js20/controls/rs/MainMenuTree.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/MainMenuTree.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[305] = &httpSrv.TagScript{Src: "js20/controls/BigFileUploader.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/BigFileUploader.js")),
			}
		a.JavaScriptModel.Rows[306] = &httpSrv.TagScript{Src: "js20/controls/rs/BigFileUploader.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/BigFileUploader.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[307] = &httpSrv.TagScript{Src: "js20/controls/EditInfo.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/EditInfo.js")),
			}
		a.JavaScriptModel.Rows[308] = &httpSrv.TagScript{Src: "js20/controls/ButtonEditPeriodValues.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonEditPeriodValues.js")),
			}
		a.JavaScriptModel.Rows[309] = &httpSrv.TagScript{Src: "js20/controls/rs/ButtonEditPeriodValues.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/ButtonEditPeriodValues.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[310] = &httpSrv.TagScript{Src: "js20/controls/ButtonOrgSearch.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/ButtonOrgSearch.js")),
			}
		a.JavaScriptModel.Rows[311] = &httpSrv.TagScript{Src: "js20/controls/Captcha.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/Captcha.js")),
			}
		a.JavaScriptModel.Rows[312] = &httpSrv.TagScript{Src: "js20/controls/rs/Captcha.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/Captcha.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[313] = &httpSrv.TagScript{Src: "js20/forms/ViewList_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/ViewList_Form.js")),
			}
		a.JavaScriptModel.Rows[314] = &httpSrv.TagScript{Src: "js20/forms/MainMenuConstructor_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/MainMenuConstructor_Form.js")),
			}
		a.JavaScriptModel.Rows[315] = &httpSrv.TagScript{Src: "js20/forms/User_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/User_Form.js")),
			}
		a.JavaScriptModel.Rows[316] = &httpSrv.TagScript{Src: "js20/forms/UserList_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/UserList_Form.js")),
			}
		a.JavaScriptModel.Rows[317] = &httpSrv.TagScript{Src: "js20/forms/PostList_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/PostList_Form.js")),
			}
		a.JavaScriptModel.Rows[318] = &httpSrv.TagScript{Src: "js20/forms/Post_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/Post_Form.js")),
			}
		a.JavaScriptModel.Rows[319] = &httpSrv.TagScript{Src: "js20/forms/ContactList_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/ContactList_Form.js")),
			}
		a.JavaScriptModel.Rows[320] = &httpSrv.TagScript{Src: "js20/forms/Contact_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/Contact_Form.js")),
			}
		a.JavaScriptModel.Rows[321] = &httpSrv.TagScript{Src: "js20/forms/Firm_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/Firm_Form.js")),
			}
		a.JavaScriptModel.Rows[322] = &httpSrv.TagScript{Src: "js20/forms/Bank_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/Bank_Form.js")),
			}
		a.JavaScriptModel.Rows[323] = &httpSrv.TagScript{Src: "js20/forms/BankList_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/BankList_Form.js")),
			}
		a.JavaScriptModel.Rows[324] = &httpSrv.TagScript{Src: "js20/forms/Firm_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/Firm_Form.js")),
			}
		a.JavaScriptModel.Rows[325] = &httpSrv.TagScript{Src: "js20/forms/Studio_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/Studio_Form.js")),
			}
		a.JavaScriptModel.Rows[326] = &httpSrv.TagScript{Src: "js20/forms/NotifTemplate_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/NotifTemplate_Form.js")),
			}
		a.JavaScriptModel.Rows[327] = &httpSrv.TagScript{Src: "js20/forms/DocumentTemplate_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/DocumentTemplate_Form.js")),
			}
		a.JavaScriptModel.Rows[328] = &httpSrv.TagScript{Src: "js20/forms/SpecialistReg_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/SpecialistReg_Form.js")),
			}
		a.JavaScriptModel.Rows[329] = &httpSrv.TagScript{Src: "js20/forms/Specialist_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/Specialist_Form.js")),
			}
		a.JavaScriptModel.Rows[330] = &httpSrv.TagScript{Src: "js20/forms/Speciality_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/Speciality_Form.js")),
			}
		a.JavaScriptModel.Rows[331] = &httpSrv.TagScript{Src: "js20/forms/EquipmentTypeList_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/EquipmentTypeList_Form.js")),
			}
		a.JavaScriptModel.Rows[332] = &httpSrv.TagScript{Src: "js20/forms/SpecialistList_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/SpecialistList_Form.js")),
			}
		a.JavaScriptModel.Rows[333] = &httpSrv.TagScript{Src: "js20/forms/SpecialistPeriodSalary_Form.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/forms/SpecialistPeriodSalary_Form.js")),
			}
		a.JavaScriptModel.Rows[334] = &httpSrv.TagScript{Src: "js20/views/SpecialistPeriodSalaryDetailList_From.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistPeriodSalaryDetailList_From.js")),
			}
		a.JavaScriptModel.Rows[335] = &httpSrv.TagScript{Src: "js20/controllers/User_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/User_Controller.js")),
			}
		a.JavaScriptModel.Rows[336] = &httpSrv.TagScript{Src: "js20/views/Login_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/Login_View.js")),
			}
		a.JavaScriptModel.Rows[337] = &httpSrv.TagScript{Src: "js20/views/MainMenuConstructorList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/MainMenuConstructorList_View.js")),
			}
		a.JavaScriptModel.Rows[338] = &httpSrv.TagScript{Src: "js20/views/MainMenuConstructor_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/MainMenuConstructor_View.js")),
			}
		a.JavaScriptModel.Rows[339] = &httpSrv.TagScript{Src: "js20/views/ViewList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/ViewList_View.js")),
			}
		a.JavaScriptModel.Rows[340] = &httpSrv.TagScript{Src: "js20/views/About_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/About_View.js")),
			}
		a.JavaScriptModel.Rows[341] = &httpSrv.TagScript{Src: "js20/views/ConstantList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/ConstantList_View.js")),
			}
		a.JavaScriptModel.Rows[342] = &httpSrv.TagScript{Src: "js20/views/LoginList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/LoginList_View.js")),
			}
		a.JavaScriptModel.Rows[343] = &httpSrv.TagScript{Src: "js20/views/UserList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/UserList_View.js")),
			}
		a.JavaScriptModel.Rows[344] = &httpSrv.TagScript{Src: "js20/views/UserDialog_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/UserDialog_View.js")),
			}
		a.JavaScriptModel.Rows[345] = &httpSrv.TagScript{Src: "js20/views/UserProfile_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/UserProfile_View.js")),
			}
		a.JavaScriptModel.Rows[346] = &httpSrv.TagScript{Src: "js20/views/TimeZoneLocale_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/TimeZoneLocale_View.js")),
			}
		a.JavaScriptModel.Rows[347] = &httpSrv.TagScript{Src: "js20/views/LoginDeviceList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/LoginDeviceList_View.js")),
			}
		a.JavaScriptModel.Rows[348] = &httpSrv.TagScript{Src: "js20/views/LoginDeviceBanList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/LoginDeviceBanList_View.js")),
			}
		a.JavaScriptModel.Rows[349] = &httpSrv.TagScript{Src: "js20/views/Registration_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/Registration_View.js")),
			}
		a.JavaScriptModel.Rows[350] = &httpSrv.TagScript{Src: "js20/views/rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/rs_ru.js")),
			}
		a.JavaScriptModel.Rows[351] = &httpSrv.TagScript{Src: "js20/views/PostList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/PostList_View.js")),
			}
		a.JavaScriptModel.Rows[352] = &httpSrv.TagScript{Src: "js20/views/PostDialog_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/PostDialog_View.js")),
			}
		a.JavaScriptModel.Rows[353] = &httpSrv.TagScript{Src: "js20/views/ContactDialog_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/ContactDialog_View.js")),
			}
		a.JavaScriptModel.Rows[354] = &httpSrv.TagScript{Src: "js20/views/ContactList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/ContactList_View.js")),
			}
		a.JavaScriptModel.Rows[355] = &httpSrv.TagScript{Src: "js20/views/EntityContactList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/EntityContactList_View.js")),
			}
		a.JavaScriptModel.Rows[356] = &httpSrv.TagScript{Src: "js20/views/BankList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/BankList_View.js")),
			}
		a.JavaScriptModel.Rows[357] = &httpSrv.TagScript{Src: "js20/views/Bank_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/Bank_View.js")),
			}
		a.JavaScriptModel.Rows[358] = &httpSrv.TagScript{Src: "js20/views/FirmDialog_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/FirmDialog_View.js")),
			}
		a.JavaScriptModel.Rows[359] = &httpSrv.TagScript{Src: "js20/views/FirmList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/FirmList_View.js")),
			}
		a.JavaScriptModel.Rows[360] = &httpSrv.TagScript{Src: "js20/views/StudioDialog_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/StudioDialog_View.js")),
			}
		a.JavaScriptModel.Rows[361] = &httpSrv.TagScript{Src: "js20/views/StudioList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/StudioList_View.js")),
			}
		a.JavaScriptModel.Rows[362] = &httpSrv.TagScript{Src: "js20/views/Passport_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/Passport_View.js")),
			}
		a.JavaScriptModel.Rows[363] = &httpSrv.TagScript{Src: "js20/views/NotifTemplate_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/NotifTemplate_View.js")),
			}
		a.JavaScriptModel.Rows[364] = &httpSrv.TagScript{Src: "js20/views/NotifTemplateList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/NotifTemplateList_View.js")),
			}
		a.JavaScriptModel.Rows[365] = &httpSrv.TagScript{Src: "js20/views/SpecialistRegList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistRegList_View.js")),
			}
		a.JavaScriptModel.Rows[366] = &httpSrv.TagScript{Src: "js20/views/SpecialistRegDialog_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistRegDialog_View.js")),
			}
		a.JavaScriptModel.Rows[367] = &httpSrv.TagScript{Src: "js20/views/DocumentTemplateList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/DocumentTemplateList_View.js")),
			}
		a.JavaScriptModel.Rows[368] = &httpSrv.TagScript{Src: "js20/views/DocumentTemplate_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/DocumentTemplate_View.js")),
			}
		a.JavaScriptModel.Rows[369] = &httpSrv.TagScript{Src: "js20/views/SpecialistDocumentOnRegisterList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistDocumentOnRegisterList_View.js")),
			}
		a.JavaScriptModel.Rows[370] = &httpSrv.TagScript{Src: "js20/views/Mgateway_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/Mgateway_View.js")),
			}
		a.JavaScriptModel.Rows[371] = &httpSrv.TagScript{Src: "js20/views/SpecialistRegRegister_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistRegRegister_View.js")),
			}
		a.JavaScriptModel.Rows[372] = &httpSrv.TagScript{Src: "js20/views/SpecialistDialog_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistDialog_View.js")),
			}
		a.JavaScriptModel.Rows[373] = &httpSrv.TagScript{Src: "js20/views/SpecialistList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistList_View.js")),
			}
		a.JavaScriptModel.Rows[374] = &httpSrv.TagScript{Src: "js20/views/SpecialistWorkList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistWorkList_View.js")),
			}
		a.JavaScriptModel.Rows[375] = &httpSrv.TagScript{Src: "js20/views/SpecialistDocumentList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistDocumentList_View.js")),
			}
		a.JavaScriptModel.Rows[376] = &httpSrv.TagScript{Src: "js20/views/SpecialistDocumentForSignList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistDocumentForSignList_View.js")),
			}
		a.JavaScriptModel.Rows[377] = &httpSrv.TagScript{Src: "js20/views/EquipmentTypeList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/EquipmentTypeList_View.js")),
			}
		a.JavaScriptModel.Rows[378] = &httpSrv.TagScript{Src: "js20/views/ItemList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/ItemList_View.js")),
			}
		a.JavaScriptModel.Rows[379] = &httpSrv.TagScript{Src: "js20/views/YClTransactionList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/YClTransactionList_View.js")),
			}
		a.JavaScriptModel.Rows[380] = &httpSrv.TagScript{Src: "js20/views/YClTransactionDocList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/YClTransactionDocList_View.js")),
			}
		a.JavaScriptModel.Rows[381] = &httpSrv.TagScript{Src: "js20/views/YClStaffList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/YClStaffList_View.js")),
			}
		a.JavaScriptModel.Rows[382] = &httpSrv.TagScript{Src: "js20/views/SpecialityList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialityList_View.js")),
			}
		a.JavaScriptModel.Rows[383] = &httpSrv.TagScript{Src: "js20/views/SpecialityDialog_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialityDialog_View.js")),
			}
		a.JavaScriptModel.Rows[384] = &httpSrv.TagScript{Src: "js20/views/AdminRate_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/AdminRate_View.js")),
			}
		a.JavaScriptModel.Rows[385] = &httpSrv.TagScript{Src: "js20/views/SalaryDebetList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SalaryDebetList_View.js")),
			}
		a.JavaScriptModel.Rows[386] = &httpSrv.TagScript{Src: "js20/views/SalaryKreditList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SalaryKreditList_View.js")),
			}
		a.JavaScriptModel.Rows[387] = &httpSrv.TagScript{Src: "js20/views/SpecialistSalaryDebetList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistSalaryDebetList_View.js")),
			}
		a.JavaScriptModel.Rows[388] = &httpSrv.TagScript{Src: "js20/views/SpecialistSalaryKreditList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistSalaryKreditList_View.js")),
			}
		a.JavaScriptModel.Rows[389] = &httpSrv.TagScript{Src: "js20/views/SpecialistPeriodSalaryList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistPeriodSalaryList_View.js")),
			}
		a.JavaScriptModel.Rows[390] = &httpSrv.TagScript{Src: "js20/views/SpecialistPeriodSalaryDialog_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistPeriodSalaryDialog_View.js")),
			}
		a.JavaScriptModel.Rows[391] = &httpSrv.TagScript{Src: "js20/views/SpecialistPeriodSalaryDetailList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistPeriodSalaryDetailList_View.js")),
			}
		a.JavaScriptModel.Rows[392] = &httpSrv.TagScript{Src: "js20/views/SpecialistPeriodSalaryDetailForPayList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistPeriodSalaryDetailForPayList_View.js")),
			}
		a.JavaScriptModel.Rows[393] = &httpSrv.TagScript{Src: "js20/views/SpecialistPeriodSalaryDetailDialog_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistPeriodSalaryDetailDialog_View.js")),
			}
		a.JavaScriptModel.Rows[394] = &httpSrv.TagScript{Src: "js20/views/YClTransactionDocAllList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/YClTransactionDocAllList_View.js")),
			}
		a.JavaScriptModel.Rows[395] = &httpSrv.TagScript{Src: "js20/views/TemplateBatchList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/TemplateBatchList_View.js")),
			}
		a.JavaScriptModel.Rows[396] = &httpSrv.TagScript{Src: "js20/views/TemplateBatchItemList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/TemplateBatchItemList_View.js")),
			}
		a.JavaScriptModel.Rows[397] = &httpSrv.TagScript{Src: "js20/views/BankPaymentList_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/BankPaymentList_View.js")),
			}
		a.JavaScriptModel.Rows[398] = &httpSrv.TagScript{Src: "js20/views/SpecialistRegTemplateBatch_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistRegTemplateBatch_View.js")),
			}
		a.JavaScriptModel.Rows[399] = &httpSrv.TagScript{Src: "js20/views/SpecialistSalaryTemplateBatch_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/SpecialistSalaryTemplateBatch_View.js")),
			}
		a.JavaScriptModel.Rows[400] = &httpSrv.TagScript{Src: "js20/tmpl/App.templates.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/tmpl/App.templates.js")),
			}
		a.JavaScriptModel.Rows[401] = &httpSrv.TagScript{Src: "js20/custom_controls/App_nails.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/App_nails.js")),
			}
		a.JavaScriptModel.Rows[402] = &httpSrv.TagScript{Src: "js20/custom_controls/ViewSectionSelect.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/ViewSectionSelect.js")),
			}
		a.JavaScriptModel.Rows[403] = &httpSrv.TagScript{Src: "js20/custom_controls/ViewEditRef.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/ViewEditRef.js")),
			}
		a.JavaScriptModel.Rows[404] = &httpSrv.TagScript{Src: "js20/custom_controls/ErrorControl.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/ErrorControl.js")),
			}
		a.JavaScriptModel.Rows[405] = &httpSrv.TagScript{Src: "js20/custom_controls/Pagination.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/Pagination.js")),
			}
		a.JavaScriptModel.Rows[406] = &httpSrv.TagScript{Src: "js20/custom_controls/EditFile.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/EditFile.js")),
			}
		a.JavaScriptModel.Rows[407] = &httpSrv.TagScript{Src: "js20/controls/rs/EditFile.rs_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controls/rs/EditFile.rs_ru.js")),
			}
		a.JavaScriptModel.Rows[408] = &httpSrv.TagScript{Src: "js20/custom_controls/UserEditRef.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/UserEditRef.js")),
			}
		a.JavaScriptModel.Rows[409] = &httpSrv.TagScript{Src: "js20/custom_controls/TimeZoneLocaleEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/TimeZoneLocaleEdit.js")),
			}
		a.JavaScriptModel.Rows[410] = &httpSrv.TagScript{Src: "js20/custom_controls/UserPhotoEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/UserPhotoEdit.js")),
			}
		a.JavaScriptModel.Rows[411] = &httpSrv.TagScript{Src: "js20/custom_controls/PostEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/PostEdit.js")),
			}
		a.JavaScriptModel.Rows[412] = &httpSrv.TagScript{Src: "js20/custom_controls/ContactEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/ContactEdit.js")),
			}
		a.JavaScriptModel.Rows[413] = &httpSrv.TagScript{Src: "js20/custom_controls/BankEditRef.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/BankEditRef.js")),
			}
		a.JavaScriptModel.Rows[414] = &httpSrv.TagScript{Src: "js20/custom_controls/EditINNSelfEmployed.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/EditINNSelfEmployed.js")),
			}
		a.JavaScriptModel.Rows[415] = &httpSrv.TagScript{Src: "js20/custom_controls/FirmEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/FirmEdit.js")),
			}
		a.JavaScriptModel.Rows[416] = &httpSrv.TagScript{Src: "js20/custom_controls/StudioEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/StudioEdit.js")),
			}
		a.JavaScriptModel.Rows[417] = &httpSrv.TagScript{Src: "js20/custom_controls/BankAccEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/BankAccEdit.js")),
			}
		a.JavaScriptModel.Rows[418] = &httpSrv.TagScript{Src: "js20/custom_controls/EditEmailInlineValidation.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/EditEmailInlineValidation.js")),
			}
		a.JavaScriptModel.Rows[419] = &httpSrv.TagScript{Src: "js20/custom_controls/EditDateInlineValidation.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/EditDateInlineValidation.js")),
			}
		a.JavaScriptModel.Rows[420] = &httpSrv.TagScript{Src: "js20/custom_controls/PassportEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/PassportEdit.js")),
			}
		a.JavaScriptModel.Rows[421] = &httpSrv.TagScript{Src: "js20/custom_controls/FIOEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/FIOEdit.js")),
			}
		a.JavaScriptModel.Rows[422] = &httpSrv.TagScript{Src: "js20/custom_controls/SignEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SignEdit.js")),
			}
		a.JavaScriptModel.Rows[423] = &httpSrv.TagScript{Src: "js20/custom_controls/Sign_View.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/Sign_View.js")),
			}
		a.JavaScriptModel.Rows[424] = &httpSrv.TagScript{Src: "js20/custom_controls/ReportTemplateFieldGrid.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/ReportTemplateFieldGrid.js")),
			}
		a.JavaScriptModel.Rows[425] = &httpSrv.TagScript{Src: "js20/custom_controls/ReportTemplateEditRef.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/ReportTemplateEditRef.js")),
			}
		a.JavaScriptModel.Rows[426] = &httpSrv.TagScript{Src: "js20/custom_controls/DocumentTemplateFieldGrid.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/DocumentTemplateFieldGrid.js")),
			}
		a.JavaScriptModel.Rows[427] = &httpSrv.TagScript{Src: "js20/custom_controls/StudioEquipmentGrid.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/StudioEquipmentGrid.js")),
			}
		a.JavaScriptModel.Rows[428] = &httpSrv.TagScript{Src: "js20/custom_controls/SpecialityEquipmentGrid.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SpecialityEquipmentGrid.js")),
			}
		a.JavaScriptModel.Rows[429] = &httpSrv.TagScript{Src: "js20/custom_controls/AttachmentManager.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/AttachmentManager.js")),
			}
		a.JavaScriptModel.Rows[430] = &httpSrv.TagScript{Src: "js20/custom_controls/SpecialistRegGridCmdRegister.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SpecialistRegGridCmdRegister.js")),
			}
		a.JavaScriptModel.Rows[431] = &httpSrv.TagScript{Src: "js20/custom_controls/SpecialistRegBtnRegister.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SpecialistRegBtnRegister.js")),
			}
		a.JavaScriptModel.Rows[432] = &httpSrv.TagScript{Src: "js20/custom_controls/SpecialistRegGridCmdPrintApp.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SpecialistRegGridCmdPrintApp.js")),
			}
		a.JavaScriptModel.Rows[433] = &httpSrv.TagScript{Src: "js20/custom_controls/SpecialistRegBtnPrintApp.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SpecialistRegBtnPrintApp.js")),
			}
		a.JavaScriptModel.Rows[434] = &httpSrv.TagScript{Src: "js20/custom_controls/DocumentTemplateEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/DocumentTemplateEdit.js")),
			}
		a.JavaScriptModel.Rows[435] = &httpSrv.TagScript{Src: "js20/custom_controls/SpecialistEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SpecialistEdit.js")),
			}
		a.JavaScriptModel.Rows[436] = &httpSrv.TagScript{Src: "js20/custom_controls/EquipmentTypeEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/EquipmentTypeEdit.js")),
			}
		a.JavaScriptModel.Rows[437] = &httpSrv.TagScript{Src: "js20/custom_controls/SpecialistDocumentGridCmdSign.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SpecialistDocumentGridCmdSign.js")),
			}
		a.JavaScriptModel.Rows[438] = &httpSrv.TagScript{Src: "js20/custom_controls/SpecialistDocumentBtnSign.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SpecialistDocumentBtnSign.js")),
			}
		a.JavaScriptModel.Rows[439] = &httpSrv.TagScript{Src: "js20/custom_controls/SpecialistWorkGridCmdRate.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SpecialistWorkGridCmdRate.js")),
			}
		a.JavaScriptModel.Rows[440] = &httpSrv.TagScript{Src: "js20/custom_controls/YClStaffEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/YClStaffEdit.js")),
			}
		a.JavaScriptModel.Rows[441] = &httpSrv.TagScript{Src: "js20/custom_controls/SpecialityEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SpecialityEdit.js")),
			}
		a.JavaScriptModel.Rows[442] = &httpSrv.TagScript{Src: "js20/custom_controls/SalaryDebetEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SalaryDebetEdit.js")),
			}
		a.JavaScriptModel.Rows[443] = &httpSrv.TagScript{Src: "js20/custom_controls/SalaryKreditEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SalaryKreditEdit.js")),
			}
		a.JavaScriptModel.Rows[444] = &httpSrv.TagScript{Src: "js20/custom_controls/SpecialistPeriodSalaryDetailEdit.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SpecialistPeriodSalaryDetailEdit.js")),
			}
		a.JavaScriptModel.Rows[445] = &httpSrv.TagScript{Src: "js20/custom_controls/YClStaffGridCmdApiGetAll.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/YClStaffGridCmdApiGetAll.js")),
			}
		a.JavaScriptModel.Rows[446] = &httpSrv.TagScript{Src: "js20/custom_controls/YClTransactionGridCmdApiGet.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/YClTransactionGridCmdApiGet.js")),
			}
		a.JavaScriptModel.Rows[447] = &httpSrv.TagScript{Src: "js20/custom_controls/SpecialistPeriodSalaryDetailForPayGridCmdMakeDocs.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/SpecialistPeriodSalaryDetailForPayGridCmdMakeDocs.js")),
			}
		a.JavaScriptModel.Rows[448] = &httpSrv.TagScript{Src: "js20/custom_controls/EditPeriodMonth.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/EditPeriodMonth.js")),
			}
		a.JavaScriptModel.Rows[449] = &httpSrv.TagScript{Src: "js20/enum_controls/Enum_data_types.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/Enum_data_types.js")),
			}
		a.JavaScriptModel.Rows[450] = &httpSrv.TagScript{Src: "js20/enum_controls/EnumGridColumn_data_types.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/EnumGridColumn_data_types.js")),
			}
		a.JavaScriptModel.Rows[451] = &httpSrv.TagScript{Src: "js20/enum_controls/EnumGridColumn_locales.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/EnumGridColumn_locales.js")),
			}
		a.JavaScriptModel.Rows[452] = &httpSrv.TagScript{Src: "js20/enum_controls/EnumGridColumn_role_types.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/EnumGridColumn_role_types.js")),
			}
		a.JavaScriptModel.Rows[453] = &httpSrv.TagScript{Src: "js20/enum_controls/Enum_locales.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/Enum_locales.js")),
			}
		a.JavaScriptModel.Rows[454] = &httpSrv.TagScript{Src: "js20/enum_controls/Enum_role_types.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/Enum_role_types.js")),
			}
		a.JavaScriptModel.Rows[455] = &httpSrv.TagScript{Src: "js20/models/TimeZoneLocale_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/TimeZoneLocale_Model.js")),
			}
		a.JavaScriptModel.Rows[456] = &httpSrv.TagScript{Src: "js20/models/User_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/User_Model.js")),
			}
		a.JavaScriptModel.Rows[457] = &httpSrv.TagScript{Src: "js20/models/UserList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/UserList_Model.js")),
			}
		a.JavaScriptModel.Rows[458] = &httpSrv.TagScript{Src: "js20/models/UserDialog_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/UserDialog_Model.js")),
			}
		a.JavaScriptModel.Rows[459] = &httpSrv.TagScript{Src: "js20/models/Login_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/Login_Model.js")),
			}
		a.JavaScriptModel.Rows[460] = &httpSrv.TagScript{Src: "js20/models/LoginList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/LoginList_Model.js")),
			}
		a.JavaScriptModel.Rows[461] = &httpSrv.TagScript{Src: "js20/models/ConstantList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/ConstantList_Model.js")),
			}
		a.JavaScriptModel.Rows[462] = &httpSrv.TagScript{Src: "js20/models/View_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/View_Model.js")),
			}
		a.JavaScriptModel.Rows[463] = &httpSrv.TagScript{Src: "js20/models/ViewList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/ViewList_Model.js")),
			}
		a.JavaScriptModel.Rows[464] = &httpSrv.TagScript{Src: "js20/models/ViewSectionList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/ViewSectionList_Model.js")),
			}
		a.JavaScriptModel.Rows[465] = &httpSrv.TagScript{Src: "js20/models/MainMenuConstructor_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/MainMenuConstructor_Model.js")),
			}
		a.JavaScriptModel.Rows[466] = &httpSrv.TagScript{Src: "js20/models/MainMenuConstructorList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/MainMenuConstructorList_Model.js")),
			}
		a.JavaScriptModel.Rows[467] = &httpSrv.TagScript{Src: "js20/models/MainMenuConstructorDialog_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/MainMenuConstructorDialog_Model.js")),
			}
		a.JavaScriptModel.Rows[468] = &httpSrv.TagScript{Src: "js20/models/MainMenuContent_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/MainMenuContent_Model.js")),
			}
		a.JavaScriptModel.Rows[469] = &httpSrv.TagScript{Src: "js20/models/VariantStorage_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/VariantStorage_Model.js")),
			}
		a.JavaScriptModel.Rows[470] = &httpSrv.TagScript{Src: "js20/models/VariantStorageList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/VariantStorageList_Model.js")),
			}
		a.JavaScriptModel.Rows[471] = &httpSrv.TagScript{Src: "js20/models/About_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/About_Model.js")),
			}
		a.JavaScriptModel.Rows[472] = &httpSrv.TagScript{Src: "js20/views/rs_common_ru.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/views/rs_common_ru.js")),
			}
		a.JavaScriptModel.Rows[473] = &httpSrv.TagScript{Src: "js20/models/UserLogin_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/UserLogin_Model.js")),
			}
		a.JavaScriptModel.Rows[474] = &httpSrv.TagScript{Src: "js20/models/UserProfile_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/UserProfile_Model.js")),
			}
		a.JavaScriptModel.Rows[475] = &httpSrv.TagScript{Src: "js20/models/Post_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/Post_Model.js")),
			}
		a.JavaScriptModel.Rows[476] = &httpSrv.TagScript{Src: "js20/controllers/Post_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/Post_Controller.js")),
			}
		a.JavaScriptModel.Rows[477] = &httpSrv.TagScript{Src: "js20/controllers/TimeZoneLocale_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/TimeZoneLocale_Controller.js")),
			}
		a.JavaScriptModel.Rows[478] = &httpSrv.TagScript{Src: "js20/controllers/Constant_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/Constant_Controller.js")),
			}
		a.JavaScriptModel.Rows[479] = &httpSrv.TagScript{Src: "js20/controllers/Enum_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/Enum_Controller.js")),
			}
		a.JavaScriptModel.Rows[480] = &httpSrv.TagScript{Src: "js20/controllers/MainMenuConstructor_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/MainMenuConstructor_Controller.js")),
			}
		a.JavaScriptModel.Rows[481] = &httpSrv.TagScript{Src: "js20/controllers/MainMenuContent_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/MainMenuContent_Controller.js")),
			}
		a.JavaScriptModel.Rows[482] = &httpSrv.TagScript{Src: "js20/controllers/SrvStatistics_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SrvStatistics_Controller.js")),
			}
		a.JavaScriptModel.Rows[483] = &httpSrv.TagScript{Src: "js20/controllers/VariantStorage_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/VariantStorage_Controller.js")),
			}
		a.JavaScriptModel.Rows[484] = &httpSrv.TagScript{Src: "js20/controllers/View_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/View_Controller.js")),
			}
		a.JavaScriptModel.Rows[485] = &httpSrv.TagScript{Src: "js20/controllers/About_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/About_Controller.js")),
			}
		a.JavaScriptModel.Rows[486] = &httpSrv.TagScript{Src: "js20/custom_controls/App.enums.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/custom_controls/App.enums.js")),
			}
		a.JavaScriptModel.Rows[487] = &httpSrv.TagScript{Src: "js20/enum_controls/Enum_mail_types.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/Enum_mail_types.js")),
			}
		a.JavaScriptModel.Rows[488] = &httpSrv.TagScript{Src: "js20/enum_controls/EnumGridColumn_mail_types.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/EnumGridColumn_mail_types.js")),
			}
		a.JavaScriptModel.Rows[489] = &httpSrv.TagScript{Src: "js20/enum_controls/Enum_notif_types.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/Enum_notif_types.js")),
			}
		a.JavaScriptModel.Rows[490] = &httpSrv.TagScript{Src: "js20/enum_controls/EnumGridColumn_notif_types.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/EnumGridColumn_notif_types.js")),
			}
		a.JavaScriptModel.Rows[491] = &httpSrv.TagScript{Src: "js20/models/LoginDeviceList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/LoginDeviceList_Model.js")),
			}
		a.JavaScriptModel.Rows[492] = &httpSrv.TagScript{Src: "js20/models/UserOperation_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/UserOperation_Model.js")),
			}
		a.JavaScriptModel.Rows[493] = &httpSrv.TagScript{Src: "js20/models/UserOperationDialog_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/UserOperationDialog_Model.js")),
			}
		a.JavaScriptModel.Rows[494] = &httpSrv.TagScript{Src: "js20/controllers/UserOperation_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/UserOperation_Controller.js")),
			}
		a.JavaScriptModel.Rows[495] = &httpSrv.TagScript{Src: "js20/models/Contact_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/Contact_Model.js")),
			}
		a.JavaScriptModel.Rows[496] = &httpSrv.TagScript{Src: "js20/models/ContactList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/ContactList_Model.js")),
			}
		a.JavaScriptModel.Rows[497] = &httpSrv.TagScript{Src: "js20/models/ContactDialog_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/ContactDialog_Model.js")),
			}
		a.JavaScriptModel.Rows[498] = &httpSrv.TagScript{Src: "js20/models/EntityContact_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/EntityContact_Model.js")),
			}
		a.JavaScriptModel.Rows[499] = &httpSrv.TagScript{Src: "js20/models/EntityContactList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/EntityContactList_Model.js")),
			}
		a.JavaScriptModel.Rows[500] = &httpSrv.TagScript{Src: "js20/controllers/Contact_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/Contact_Controller.js")),
			}
		a.JavaScriptModel.Rows[501] = &httpSrv.TagScript{Src: "js20/models/Bank_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/Bank_Model.js")),
			}
		a.JavaScriptModel.Rows[502] = &httpSrv.TagScript{Src: "js20/models/BankList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/BankList_Model.js")),
			}
		a.JavaScriptModel.Rows[503] = &httpSrv.TagScript{Src: "js20/controllers/EntityContact_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/EntityContact_Controller.js")),
			}
		a.JavaScriptModel.Rows[504] = &httpSrv.TagScript{Src: "js20/controllers/Bank_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/Bank_Controller.js")),
			}
		a.JavaScriptModel.Rows[505] = &httpSrv.TagScript{Src: "js20/controllers/Captcha_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/Captcha_Controller.js")),
			}
		a.JavaScriptModel.Rows[506] = &httpSrv.TagScript{Src: "js20/models/FirmList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/FirmList_Model.js")),
			}
		a.JavaScriptModel.Rows[507] = &httpSrv.TagScript{Src: "js20/models/FirmDialog_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/FirmDialog_Model.js")),
			}
		a.JavaScriptModel.Rows[508] = &httpSrv.TagScript{Src: "js20/controllers/Firm_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/Firm_Controller.js")),
			}
		a.JavaScriptModel.Rows[509] = &httpSrv.TagScript{Src: "js20/models/Studio_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/Studio_Model.js")),
			}
		a.JavaScriptModel.Rows[510] = &httpSrv.TagScript{Src: "js20/controllers/Studio_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/Studio_Controller.js")),
			}
		a.JavaScriptModel.Rows[511] = &httpSrv.TagScript{Src: "js20/controllers/StudioEquipment_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/StudioEquipment_Controller.js")),
			}
		a.JavaScriptModel.Rows[512] = &httpSrv.TagScript{Src: "js20/models/StudioEquipment_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/StudioEquipment_Model.js")),
			}
		a.JavaScriptModel.Rows[513] = &httpSrv.TagScript{Src: "js20/models/StudioDialog_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/StudioDialog_Model.js")),
			}
		a.JavaScriptModel.Rows[514] = &httpSrv.TagScript{Src: "js20/models/Specialist_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/Specialist_Model.js")),
			}
		a.JavaScriptModel.Rows[515] = &httpSrv.TagScript{Src: "js20/models/SpecialistList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistList_Model.js")),
			}
		a.JavaScriptModel.Rows[516] = &httpSrv.TagScript{Src: "js20/models/SpecialistDialog_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistDialog_Model.js")),
			}
		a.JavaScriptModel.Rows[517] = &httpSrv.TagScript{Src: "js20/controllers/Specialist_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/Specialist_Controller.js")),
			}
		a.JavaScriptModel.Rows[518] = &httpSrv.TagScript{Src: "js20/enum_controls/Enum_specialist_status_types.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/Enum_specialist_status_types.js")),
			}
		a.JavaScriptModel.Rows[519] = &httpSrv.TagScript{Src: "js20/enum_controls/EnumGridColumn_specialist_status_types.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/EnumGridColumn_specialist_status_types.js")),
			}
		a.JavaScriptModel.Rows[520] = &httpSrv.TagScript{Src: "js20/models/SpecialistStatus_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistStatus_Model.js")),
			}
		a.JavaScriptModel.Rows[521] = &httpSrv.TagScript{Src: "js20/controllers/SpecialistStatus_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SpecialistStatus_Controller.js")),
			}
		a.JavaScriptModel.Rows[522] = &httpSrv.TagScript{Src: "js20/enum_controls/Enum_notif_providers.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/Enum_notif_providers.js")),
			}
		a.JavaScriptModel.Rows[523] = &httpSrv.TagScript{Src: "js20/enum_controls/EnumGridColumn_notif_providers.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/EnumGridColumn_notif_providers.js")),
			}
		a.JavaScriptModel.Rows[524] = &httpSrv.TagScript{Src: "js20/models/NotifTemplate_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/NotifTemplate_Model.js")),
			}
		a.JavaScriptModel.Rows[525] = &httpSrv.TagScript{Src: "js20/models/NotifTemplateList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/NotifTemplateList_Model.js")),
			}
		a.JavaScriptModel.Rows[526] = &httpSrv.TagScript{Src: "js20/controllers/NotifTemplate_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/NotifTemplate_Controller.js")),
			}
		a.JavaScriptModel.Rows[527] = &httpSrv.TagScript{Src: "js20/models/ConfirmationStatus_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/ConfirmationStatus_Model.js")),
			}
		a.JavaScriptModel.Rows[528] = &httpSrv.TagScript{Src: "js20/models/SpecialistRegDialog_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistRegDialog_Model.js")),
			}
		a.JavaScriptModel.Rows[529] = &httpSrv.TagScript{Src: "js20/controllers/SpecialistReg_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SpecialistReg_Controller.js")),
			}
		a.JavaScriptModel.Rows[530] = &httpSrv.TagScript{Src: "js20/models/SpecialistReg_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistReg_Model.js")),
			}
		a.JavaScriptModel.Rows[531] = &httpSrv.TagScript{Src: "js20/controllers/ConfirmationStatus_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/ConfirmationStatus_Controller.js")),
			}
		a.JavaScriptModel.Rows[532] = &httpSrv.TagScript{Src: "js20/controllers/ClientSearch_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/ClientSearch_Controller.js")),
			}
		a.JavaScriptModel.Rows[533] = &httpSrv.TagScript{Src: "js20/controllers/Login_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/Login_Controller.js")),
			}
		a.JavaScriptModel.Rows[534] = &httpSrv.TagScript{Src: "js20/controllers/LoginDevice_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/LoginDevice_Controller.js")),
			}
		a.JavaScriptModel.Rows[535] = &httpSrv.TagScript{Src: "js20/controllers/ReportTemplateField_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/ReportTemplateField_Controller.js")),
			}
		a.JavaScriptModel.Rows[536] = &httpSrv.TagScript{Src: "js20/models/ReportTemplateField_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/ReportTemplateField_Model.js")),
			}
		a.JavaScriptModel.Rows[537] = &httpSrv.TagScript{Src: "js20/models/DocumentTemplate_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/DocumentTemplate_Model.js")),
			}
		a.JavaScriptModel.Rows[538] = &httpSrv.TagScript{Src: "js20/models/DocumentTemplateList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/DocumentTemplateList_Model.js")),
			}
		a.JavaScriptModel.Rows[539] = &httpSrv.TagScript{Src: "js20/controllers/DocumentTemplate_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/DocumentTemplate_Controller.js")),
			}
		a.JavaScriptModel.Rows[540] = &httpSrv.TagScript{Src: "js20/controllers/DocumentTemplateField_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/DocumentTemplateField_Controller.js")),
			}
		a.JavaScriptModel.Rows[541] = &httpSrv.TagScript{Src: "js20/models/DocumentTemplateField_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/DocumentTemplateField_Model.js")),
			}
		a.JavaScriptModel.Rows[542] = &httpSrv.TagScript{Src: "js20/models/DocumentTemplateDialog_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/DocumentTemplateDialog_Model.js")),
			}
		a.JavaScriptModel.Rows[543] = &httpSrv.TagScript{Src: "js20/models/Attachment_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/Attachment_Model.js")),
			}
		a.JavaScriptModel.Rows[544] = &httpSrv.TagScript{Src: "js20/models/AttachmentList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/AttachmentList_Model.js")),
			}
		a.JavaScriptModel.Rows[545] = &httpSrv.TagScript{Src: "js20/controllers/Attachment_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/Attachment_Controller.js")),
			}
		a.JavaScriptModel.Rows[546] = &httpSrv.TagScript{Src: "js20/models/SpecialistRegList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistRegList_Model.js")),
			}
		a.JavaScriptModel.Rows[547] = &httpSrv.TagScript{Src: "js20/models/StudioList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/StudioList_Model.js")),
			}
		a.JavaScriptModel.Rows[548] = &httpSrv.TagScript{Src: "js20/models/SpecialistDocument_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistDocument_Model.js")),
			}
		a.JavaScriptModel.Rows[549] = &httpSrv.TagScript{Src: "js20/models/SpecialistDocumentList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistDocumentList_Model.js")),
			}
		a.JavaScriptModel.Rows[550] = &httpSrv.TagScript{Src: "js20/models/SpecialistDocumentOnRegister_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistDocumentOnRegister_Model.js")),
			}
		a.JavaScriptModel.Rows[551] = &httpSrv.TagScript{Src: "js20/models/SpecialistDocumentOnRegisterList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistDocumentOnRegisterList_Model.js")),
			}
		a.JavaScriptModel.Rows[552] = &httpSrv.TagScript{Src: "js20/controllers/SpecialistDocumentOnRegister_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SpecialistDocumentOnRegister_Controller.js")),
			}
		a.JavaScriptModel.Rows[553] = &httpSrv.TagScript{Src: "js20/models/SpecialistWork_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistWork_Model.js")),
			}
		a.JavaScriptModel.Rows[554] = &httpSrv.TagScript{Src: "js20/models/SpecialistWorkList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistWorkList_Model.js")),
			}
		a.JavaScriptModel.Rows[555] = &httpSrv.TagScript{Src: "js20/controllers/SpecialistWork_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SpecialistWork_Controller.js")),
			}
		a.JavaScriptModel.Rows[556] = &httpSrv.TagScript{Src: "js20/controllers/SpecialistDocument_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SpecialistDocument_Controller.js")),
			}
		a.JavaScriptModel.Rows[557] = &httpSrv.TagScript{Src: "js20/models/SpecialistDocumentForSignList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistDocumentForSignList_Model.js")),
			}
		a.JavaScriptModel.Rows[558] = &httpSrv.TagScript{Src: "js20/models/EquipmentType_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/EquipmentType_Model.js")),
			}
		a.JavaScriptModel.Rows[559] = &httpSrv.TagScript{Src: "js20/controllers/EquipmentType_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/EquipmentType_Controller.js")),
			}
		a.JavaScriptModel.Rows[560] = &httpSrv.TagScript{Src: "js20/controllers/Item_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/Item_Controller.js")),
			}
		a.JavaScriptModel.Rows[561] = &httpSrv.TagScript{Src: "js20/models/Speciality_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/Speciality_Model.js")),
			}
		a.JavaScriptModel.Rows[562] = &httpSrv.TagScript{Src: "js20/models/Item_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/Item_Model.js")),
			}
		a.JavaScriptModel.Rows[563] = &httpSrv.TagScript{Src: "js20/controllers/Speciality_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/Speciality_Controller.js")),
			}
		a.JavaScriptModel.Rows[564] = &httpSrv.TagScript{Src: "js20/models/SpecialityList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialityList_Model.js")),
			}
		a.JavaScriptModel.Rows[565] = &httpSrv.TagScript{Src: "js20/models/SpecialityDialog_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialityDialog_Model.js")),
			}
		a.JavaScriptModel.Rows[566] = &httpSrv.TagScript{Src: "js20/controllers/SpecialityEquipment_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SpecialityEquipment_Controller.js")),
			}
		a.JavaScriptModel.Rows[567] = &httpSrv.TagScript{Src: "js20/models/SpecialityEquipment_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialityEquipment_Model.js")),
			}
		a.JavaScriptModel.Rows[568] = &httpSrv.TagScript{Src: "js20/models/YClStaff_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/YClStaff_Model.js")),
			}
		a.JavaScriptModel.Rows[569] = &httpSrv.TagScript{Src: "js20/models/YClStaffList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/YClStaffList_Model.js")),
			}
		a.JavaScriptModel.Rows[570] = &httpSrv.TagScript{Src: "js20/models/YClTransaction_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/YClTransaction_Model.js")),
			}
		a.JavaScriptModel.Rows[571] = &httpSrv.TagScript{Src: "js20/controllers/YClStaff_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/YClStaff_Controller.js")),
			}
		a.JavaScriptModel.Rows[572] = &httpSrv.TagScript{Src: "js20/models/YClVisit_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/YClVisit_Model.js")),
			}
		a.JavaScriptModel.Rows[573] = &httpSrv.TagScript{Src: "js20/models/YClTransactionDocList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/YClTransactionDocList_Model.js")),
			}
		a.JavaScriptModel.Rows[574] = &httpSrv.TagScript{Src: "js20/models/YClTransactionList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/YClTransactionList_Model.js")),
			}
		a.JavaScriptModel.Rows[575] = &httpSrv.TagScript{Src: "js20/controllers/YClTransaction_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/YClTransaction_Controller.js")),
			}
		a.JavaScriptModel.Rows[576] = &httpSrv.TagScript{Src: "js20/models/SpecialistWorkForRateList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistWorkForRateList_Model.js")),
			}
		a.JavaScriptModel.Rows[577] = &httpSrv.TagScript{Src: "js20/controllers/SalaryDebet_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SalaryDebet_Controller.js")),
			}
		a.JavaScriptModel.Rows[578] = &httpSrv.TagScript{Src: "js20/controllers/SalaryKredit_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SalaryKredit_Controller.js")),
			}
		a.JavaScriptModel.Rows[579] = &httpSrv.TagScript{Src: "js20/models/SalaryDebet_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SalaryDebet_Model.js")),
			}
		a.JavaScriptModel.Rows[580] = &httpSrv.TagScript{Src: "js20/models/SalaryKredit_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SalaryKredit_Model.js")),
			}
		a.JavaScriptModel.Rows[581] = &httpSrv.TagScript{Src: "js20/models/SpecialistSalaryDebet_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistSalaryDebet_Model.js")),
			}
		a.JavaScriptModel.Rows[582] = &httpSrv.TagScript{Src: "js20/models/SpecialistSalaryKredit_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistSalaryKredit_Model.js")),
			}
		a.JavaScriptModel.Rows[583] = &httpSrv.TagScript{Src: "js20/models/SpecialistSalaryKreditList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistSalaryKreditList_Model.js")),
			}
		a.JavaScriptModel.Rows[584] = &httpSrv.TagScript{Src: "js20/controllers/SpecialistSalaryKredit_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SpecialistSalaryKredit_Controller.js")),
			}
		a.JavaScriptModel.Rows[585] = &httpSrv.TagScript{Src: "js20/controllers/SpecialistSalaryDebet_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SpecialistSalaryDebet_Controller.js")),
			}
		a.JavaScriptModel.Rows[586] = &httpSrv.TagScript{Src: "js20/models/SpecialistSalaryDebetList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistSalaryDebetList_Model.js")),
			}
		a.JavaScriptModel.Rows[587] = &httpSrv.TagScript{Src: "js20/models/SpecialistPeriodSalary_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistPeriodSalary_Model.js")),
			}
		a.JavaScriptModel.Rows[588] = &httpSrv.TagScript{Src: "js20/controllers/SpecialistPeriodSalary_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SpecialistPeriodSalary_Controller.js")),
			}
		a.JavaScriptModel.Rows[589] = &httpSrv.TagScript{Src: "js20/models/SpecialistPeriodSalaryList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistPeriodSalaryList_Model.js")),
			}
		a.JavaScriptModel.Rows[590] = &httpSrv.TagScript{Src: "js20/models/YClTransactionDocAllList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/YClTransactionDocAllList_Model.js")),
			}
		a.JavaScriptModel.Rows[591] = &httpSrv.TagScript{Src: "js20/models/SpecialistPeriodSalaryDetail_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistPeriodSalaryDetail_Model.js")),
			}
		a.JavaScriptModel.Rows[592] = &httpSrv.TagScript{Src: "js20/models/SpecialistPeriodSalaryDetailList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistPeriodSalaryDetailList_Model.js")),
			}
		a.JavaScriptModel.Rows[593] = &httpSrv.TagScript{Src: "js20/controllers/SpecialistPeriodSalaryDetail_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SpecialistPeriodSalaryDetail_Controller.js")),
			}
		a.JavaScriptModel.Rows[594] = &httpSrv.TagScript{Src: "js20/models/SpecialistReceipt_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/SpecialistReceipt_Model.js")),
			}
		a.JavaScriptModel.Rows[595] = &httpSrv.TagScript{Src: "js20/controllers/SpecialistReceipt_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/SpecialistReceipt_Controller.js")),
			}
		a.JavaScriptModel.Rows[596] = &httpSrv.TagScript{Src: "js20/controllers/QRExtractor_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/QRExtractor_Controller.js")),
			}
		a.JavaScriptModel.Rows[597] = &httpSrv.TagScript{Src: "js20/controllers/TemplateBatch_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/TemplateBatch_Controller.js")),
			}
		a.JavaScriptModel.Rows[598] = &httpSrv.TagScript{Src: "js20/models/TemplateBatch_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/TemplateBatch_Model.js")),
			}
		a.JavaScriptModel.Rows[599] = &httpSrv.TagScript{Src: "js20/controllers/TemplateBatchItem_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/TemplateBatchItem_Controller.js")),
			}
		a.JavaScriptModel.Rows[600] = &httpSrv.TagScript{Src: "js20/models/TemplateBatchItem_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/TemplateBatchItem_Model.js")),
			}
		a.JavaScriptModel.Rows[601] = &httpSrv.TagScript{Src: "js20/models/TemplateBatchItemList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/TemplateBatchItemList_Model.js")),
			}
		a.JavaScriptModel.Rows[602] = &httpSrv.TagScript{Src: "js20/enum_controls/Enum_template_batch_types.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/Enum_template_batch_types.js")),
			}
		a.JavaScriptModel.Rows[603] = &httpSrv.TagScript{Src: "js20/enum_controls/EnumGridColumn_template_batch_types.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/enum_controls/EnumGridColumn_template_batch_types.js")),
			}
		a.JavaScriptModel.Rows[604] = &httpSrv.TagScript{Src: "js20/models/BankPayment_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/BankPayment_Model.js")),
			}
		a.JavaScriptModel.Rows[605] = &httpSrv.TagScript{Src: "js20/models/BankPaymentList_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/BankPaymentList_Model.js")),
			}
		a.JavaScriptModel.Rows[606] = &httpSrv.TagScript{Src: "js20/controllers/BankPayment_Controller.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/controllers/BankPayment_Controller.js")),
			}
		a.JavaScriptModel.Rows[607] = &httpSrv.TagScript{Src: "js20/models/Firm_Model.js",
				Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/models/Firm_Model.js")),
			}
	}else{
		a.JavaScriptModel = httpSrv.NewScriptModel(13)
		a.JavaScriptModel.Rows[0] = &httpSrv.TagScript{Src: "js20/assets/js/core/libraries/jquery.min.js",
			Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/assets/js/core/libraries/jquery.min.js")),
		}
		a.JavaScriptModel.Rows[1] = &httpSrv.TagScript{Src: "js20/assets/js/core/libraries/bootstrap.min.js",
			Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/assets/js/core/libraries/bootstrap.min.js")),
		}
		a.JavaScriptModel.Rows[2] = &httpSrv.TagScript{Src: "js20/assets/js/plugins/loaders/blockui.min.js",
			Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/assets/js/plugins/loaders/blockui.min.js")),
		}
		a.JavaScriptModel.Rows[3] = &httpSrv.TagScript{Src: "js20/assets/js/core/app.js",
			Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/assets/js/core/app.js")),
		}
		a.JavaScriptModel.Rows[4] = &httpSrv.TagScript{Src: "js20/ext/bootstrap-datepicker/bootstrap-datepicker.min.js",
			Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/bootstrap-datepicker/bootstrap-datepicker.min.js")),
		}
		a.JavaScriptModel.Rows[5] = &httpSrv.TagScript{Src: "js20/ext/bootstrap-datepicker/bootstrap-datepicker.ru.min.js",
			Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/bootstrap-datepicker/bootstrap-datepicker.ru.min.js")),
		}
		a.JavaScriptModel.Rows[6] = &httpSrv.TagScript{Src: "js20/ext/cleave/cleave.min.js",
			Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/cleave/cleave.min.js")),
		}
		a.JavaScriptModel.Rows[7] = &httpSrv.TagScript{Src: "js20/ext/cleave/cleave-phone.ru.js",
			Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/cleave/cleave-phone.ru.js")),
		}
		a.JavaScriptModel.Rows[8] = &httpSrv.TagScript{Src: "js20/ext/jshash-2.2/md5-min.js",
			Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/jshash-2.2/md5-min.js")),
		}
		a.JavaScriptModel.Rows[9] = &httpSrv.TagScript{Src: "js20/ext/mustache/mustache.min.js",
			Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/ext/mustache/mustache.min.js")),
		}
		a.JavaScriptModel.Rows[10] = &httpSrv.TagScript{Src: "js20/assets/js/plugins/forms/styling/switchery.min.js",
			Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/assets/js/plugins/forms/styling/switchery.min.js")),
		}
		a.JavaScriptModel.Rows[11] = &httpSrv.TagScript{Src: "js20/assets/js/plugins/forms/styling/uniform.min.js",
			Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/assets/js/plugins/forms/styling/uniform.min.js")),
		}
		a.JavaScriptModel.Rows[12] = &httpSrv.TagScript{Src: "js20/lib.js",
			Modified: httpSrv.ScriptModifiedTime(filepath.Join(bs, "js20/lib.js")),
		}
	}
}
