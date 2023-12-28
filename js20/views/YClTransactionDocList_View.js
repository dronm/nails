/** Copyright (c) 2023
 *	Andrey Mikhalevich, Katren ltd.
 */
function YClTransactionDocList_View(id,options){	

	options.HEAD_TITLE = "Документы yclients";
	YClTransactionDocList_View.superclass.constructor.call(this,id,options);

	var model = (options.models && options.models.YClTransactionDocList_Model)? options.models.YClTransactionDocList_Model : new YClTransactionDocList_Model();
	var contr = new YClTransaction_Controller();
	
	var constants = {"doc_per_page_count":null,"grid_refresh_interval":null};
	window.getApp().getConstantManager().get(constants);

	this.m_onSelectDoc = options.onSelectDoc;

	var filters;
	var popup_menu = new PopUpMenu();
	var pagClass = window.getApp().getPaginationClass();
	var self = this;
	this.addElement(new GridAjx(id+":grid",{
		"model":model,
		"keyIds":["document_id"],
		"controller":contr,
		"readPublicMethod":contr.getPublicMethod("get_for_work_list"),
		"editInline":true,
		"editWinClass":null,
		"commands":new GridCmdContainerAjx(id+":grid:cmd",{
			"filters": filters,
			"exportFileName":"Транзакции_yclients",
			"cmdSearch": false,
			"cmdAllCommands": false,
			"cmdInsert":false,
			"cmdEdit":false
		}),		
		"popUpMenu":popup_menu,
		"filters":(options.detailFilters&&options.detailFilters.YClTransactionDoclList_Model)? options.detailFilters.YClTransactionDocList_Model:null,
		"head":new GridHead(id+"-grid:head",{
			"elements":[
				new GridRow(id+":grid:head:row0",{
					"elements":[
						(window.getApp().getServVar("role_id")=="specialist")? null : new GridCellHead(id+":grid:head:document_id",{
							"value":"ID документа",
							"columns":[
								new GridColumn({
									"field":model.getField("document_id")
								})
							],
							"sortable":true
						})
						,new GridCellHead(id+":grid:head:date",{
							"value":"Дата",
							"columns":[
								new GridColumnDateTime({
									"field":model.getField("date")
								})
							],
							"sortable":true,
							"sort":"desc"
						})
						,new GridCellHead(id+":grid:head:client",{
							"value":"Клиент",
							"columns":[
								new GridColumn({
									"field":model.getField("client")
								})
							]
						})
						,new GridCellHead(id+":grid:head:amount",{
							"value":"Сумма",
							"columns":[
								new GridColumnFloat({
									"field":model.getField("amount"),
									"precision":"2",
									"length":"15"
								})
							]
						})
						
					]
				})
			]
		}),
		"pagination":new pagClass(id+"_page",
			{"countPerPage":constants.doc_per_page_count.getValue()}),		
		
		"autoRefresh":options.detailFilters||options.onSelectDoc? true:false,
		"refreshInterval":constants.grid_refresh_interval.getValue()*1000,
		"rowSelect":false,
		"focus":true,
		/*"onSelect":!options.forSelect? null : function(row){
			alert(row.document_id.getValue());
		},*/
		"clickEvent":!options.onSelectDoc? null : function(e)	{
			var tr = DOMHelper.getParentByTagName(e.target, "tr");
			if(!tr){
				return;
			}
			var keys = CommonHelper.unserialize(tr.getAttribute("keys"));
			self.m_onSelectDoc(parseInt(keys.document_id, 10));
		}
	}));	
	
}
extend(YClTransactionDocList_View,ViewAjxList);

