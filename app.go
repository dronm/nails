package main

import(
	"fmt"
	"net"
	"time"
	"errors"
	"os"
	"encoding/json"
	"context"
	
	"nails/controllers"
	"nails/nails_config"
	
	"github.com/dronm/fnsnpd"	
	"github.com/dronm/ds"
	"github.com/dronm/ds/pgds"
	
	"osbe"
	"osbe/srv"
	"osbe/srv/httpSrv"
	"osbe/srv/wsSrv"
	"osbe/response"
	"osbe/evnt"	
	"osbe/socket"	
	"osbe/view"	
	"osbe/fields"
	rusTel "osbe/repo/validation/rusTel"
	clientSearch "osbe/repo/clientSearch"
	
	"osbe/permission"
	_ "osbe/permission/pg"	

	"github.com/dronm/session"	
	//_ "github.com/dronm/session/pg"
	_ "github.com/dronm/session/redis"	
	
	"github.com/labstack/gommon/log"
	
)

const (
	PAGE_TITLE = "Ногти"
	PAGE_AUTHOR = "Andrey Mikhalevich"
	PAGE_KEYWORS = ""
	PAGE_DESCRIPTION = ""
	DEF_COLOR_PALETTE = "pink-300"
	
	CONF_VIEW = "ConfirmationStatus"
	CONF_CONTR = "ConfirmationStatus"
	CONF_METH = "set_confirmed"	
)

type RegCaptcha struct {
	Id string `json:"id"`
	Width int `json:"width"`
	Height int `json:"height"`
	Count int `json:"count"`
}

type NailsVariables struct {
	httpSrv.ServerVars
	Role_id string `json:"role_id"`
	User_id int64 `json:"user_id"`
	User_name string `json:"user_name"`
	Photo string `json:"photo"`
	App_id string `json:"app_id"`
	Token string `json:"token"`
	TokenExpires string `json:"tokenExpires"`
	DefColorPalette string `json:"defColorPalette"`
	RegCaptcha string `json:"regCaptcha"`
	ValidTelPrefixes string `json:"validTelPrefixes"`
	WsPort int `json:"wsPort"`
	WsPortTls int `json:"wsPortTls"`
	UnsignedDocCount int `json:"unsignedDocCount"` //Количество не подписанных документов для мастера
	JoinContract string `json:"joinContract"`
}

type NailsApp struct {
	httpSrv.HTTPApplication	
	EvntServer *evnt.EvntSrv	
	//
	cancel context.CancelFunc		
	ctx context.Context
}

func (a *NailsApp) init(conf *nails_config.NailsConfig) {
	var err error
	
	a.Config = conf
	l := log.New("-")
	l.SetHeader("${time_rfc3339_nano} ${short_file}:${line} ${level} -${message}")
	l.SetLevel(a.GetLogLevel(conf.GetLogLevel()))
	a.Logger = l 
	
	//context
	a.ctx, a.cancel = context.WithCancel(context.Background())		
	
	//
	a.OnReloadConfig = func(){
		a.init(a.Config.(*nails_config.NailsConfig))
	}
	
	//metadata
	a.MD = NewMD(a.GetAppVersion())
	initMD(a.MD)
	
	//default per page count
	osbe.DocPerPageCountConstantID = "doc_per_page_count"
	
	//HTTPDir or current
	http_dir := conf.HTTPDir
	if http_dir == "" {
		http_dir, err = os.Getwd()//current working dir
		if err != nil {
			panic(fmt.Sprintf("os.Getwd(): %v", err))
		}
	}	
	a.UserTmplDir = http_dir+"/tmpl"
	a.UserTmplExtension = "html"

	a.setJavaScript(conf.JsDebug);	
	a.setCSS();
		
	//Event server
	a.EvntServer = evnt.NewEvntSrv(a.Logger, a.HandleEvent,
		[]string{"Permission.change",
			"Login.logout",
			"Login.destroy_session",
			"Attachment.clear_cache",
		},		
	)
			
	a.OnPublishEvent = a.EvntServer.PublishEvent
	a.MD.Controllers["Event"].(*evnt.Event_Controller).EvntServer = a.EvntServer

	//Db support
	db_conf := a.Config.GetDb()
	a.DataStorage, err = ds.NewProvider("pg", db_conf.Primary, a.EvntServer.OnNotification, db_conf.Secondaries)
	if err != nil {
		panic(err)
	}
	d_store,ok := a.DataStorage.(*pgds.PgProvider)
	if !ok {
		panic(errors.New("a.DataStorage is not of type *pgds.PgProvider"))
	}

	err = d_store.Primary.Connect()
	if err != nil {
		panic(err)
	}

	//peprmission support
	perm_conn_s := ""
	if db_conf.Secondaries != nil && len(db_conf.Secondaries) > 0 {
		for _,c := range db_conf.Secondaries {
			perm_conn_s = c
			break
		}		
	}else{
		perm_conn_s = db_conf.Primary
	}
	a.PermisManager, err = permission.NewManager("pg", perm_conn_s)
	if err != nil {
		panic(err)
	}
		
	//session support
	sess_conf := a.Config.GetSession()
	//d_store.Primary.Pool
	a.SessManager, err = session.NewManager("redis", sess_conf.MaxLifeTime, sess_conf.MaxIdleTime, sess_conf.DestroyAllTime, conf.Redis.Connect, conf.Redis.Namespace)
	//a.SessManager, err = session.NewManager("pg", sess_conf.MaxLifeTime, sess_conf.MaxIdleTime, db_conf.Primary, sess_conf.EncKey)
	if err != nil {
		panic(fmt.Sprintf("session.NewManager: %v", err))
	}
	lv := conf.GetLogLevel()
	var sess_lv session.LogLevel
	switch lv {
	case "debug": sess_lv = session.LOG_LEVEL_DEBUG
	case "warn": sess_lv = session.LOG_LEVEL_WARN
	default: sess_lv = session.LOG_LEVEL_ERROR
	}
	a.SessManager.StartGC(l.Output(), sess_lv)
	
	//Db connection to event server
	a.EvntServer.DbPool = d_store.Primary.Pool
	
	//http server
	server := &httpSrv.HTTPServer{}
	server.Address = conf.HttpServer
	server.AppID = conf.AppID
	server.Logger = a.Logger
	server.HTTPDir = http_dir
	server.AddURLShortcut("/email", CONF_CONTR, CONF_METH, CONF_VIEW, nil)//
	server.AddURLShortcut("/tel", CONF_CONTR, CONF_METH, CONF_VIEW, nil)//
	server.AddURLShortcut("/mgateway", "Mgateway", "callback", "ViewXML", nil)//
	server.AddURLShortcut("/qrextr", "QRExtractor", "callback", "ViewXML", nil)//
	server.AddURLShortcut("/Registration", "", "", "Registration", nil)//
	server.AddURLShortcut("/Gettoknow", "", "", "Gettoknow", nil)//
	server.AddURLShortcut("/Login", "", "", "Login", nil)//
	server.AddURLShortcut("/PasswordRecovery", "", "", "PasswordRecovery", nil)//
	
	//!!!Event custom socket
	server.OnConstructSocket = (func(srv *evnt.EvntSrv) srv.OnConstructSocketProto {
		//id string, id, 
		return func(conn net.Conn, token string, tokenExp time.Time) socket.ClientSocketer {
			sock := evnt.NewClientSocket(conn, token, tokenExp, srv)
			return sock
		}
	})(a.EvntServer)
	
	server.OnHandleRequest = a.HandleRequest	
	server.OnHandleSession = a.HandleSession
	server.OnHandleServerError = a.HandleServerError		
	
	//corresponding content types
	server.AddViewContentType("ViewXML", httpSrv.MIME_TYPE_xml, httpSrv.CHARSET_UTF8)
	server.AddViewContentType("ViewJSON", httpSrv.MIME_TYPE_json, httpSrv.CHARSET_UTF8)
	server.AddViewContentType("ViewHTML", httpSrv.MIME_TYPE_html, httpSrv.CHARSET_UTF8)
	server.AddViewContentType("ViewPDF", httpSrv.MIME_TYPE_pdf, "")
	server.AddViewContentType("ViewExcel", httpSrv.MIME_TYPE_xls, "")
	server.AddViewContentType("ViewText", httpSrv.MIME_TYPE_txt, httpSrv.CHARSET_UTF8)
	server.AddViewContentType("ViewCSV", httpSrv.MIME_TYPE_txt, httpSrv.CHARSET_UTF8)
	server.AllowedExtensions = []string{"ico","css","js","gif","png","mp3","woff","woff2","ttf","jpg"}
	server.Headers = map[string]string{
		"Pragma": "no-cache",
		"Cache-Control": "no-cache, no-store, max-age=0, must-revalidate",
		"Expires": "0",	
	}	
	a.AddServer("http", server)
	//views
	view.Init("ViewXML", map[string]interface{}{"BeforeRender": a.OnBeforeRenderXML,
				"DebugDir": http_dir + "/CACHE",
			})
	view.Init("ViewJSON", nil)
	view.Init("ViewHTML",
		map[string]interface{}{"SrvTemplateDir": conf.TemplateDir,
			"UserViewDir": http_dir + "/views",
			"TemplateExtension": "html.xsl",
			"TemplateTransform": a.XSLTransform,
			"BeforeRender": a.onBeforeRenderHTML,
			"DebugDir": http_dir + "/CACHE",
			"XMLDebug": false,
			"HTMLDebug": false,			
			})
	view.Init("ViewExcel",
		map[string]interface{}{"SrvTemplateDir": conf.TemplateDir,
			"UserViewDir": http_dir + "/views",
			"TemplateTransform": a.XSLTransform,			
			})
	view.Init("ViewPDF",
		map[string]interface{}{"SrvTemplateDir": conf.TemplateDir,
			"UserViewDir": http_dir + "/views",
			"TemplateTransform": a.XSLToPDFTransform,
			"Fop": conf.Fop,
			"ConfFile": conf.FopConf,
			"DebugDir": http_dir + "/CACHE",
			"Debug": false,
		})
	
	//web socket server
	ws_srv := &wsSrv.WSServer{srv.BaseServer{Address: conf.GetWSServer(),
		TlsAddress: conf.GetTLSWSServer(),
		TlsCert: conf.TLSCert,
		TlsKey: conf.TLSKey,
		Logger: a.Logger,
		AppID: conf.AppID,
		OnHandleJSONRequest: a.HandleJSONRequest,
		OnHandleSession: a.HandleSession,
		OnDestroySession: a.DestroySession,
		OnHandleServerError: a.HandleServerError,
		OnConstructSocket: server.OnConstructSocket,		
	},
	nil,
	nil}//a.OnCloseSocket
	a.AddServer("ws", ws_srv)
	
	//FNS person check
	fnsnpd.InitFNSPersonCheck(l.Output())
	
	//notifier
	controllers.InitNotifier(conf.NotifierParams.Host, conf.NotifierParams.AppName, conf.NotifierParams.Pwd)
	
	//client search
	clientSearch.DadataKey = conf.DadataKey
	
	//yClients
	controllers.YClientsAPIInit(conf.YClients.Login, conf.YClients.Pwd, conf.YClients.PartnerToken, conf.YClients.CompanyID)	
	//сервер обновления данных записей yclients
	//доп информация к коллбакам
	if conf.YClients.RecordInfDoUpdate {
		controllers.YClientsServe(d_store, a.GetLogger(), conf.YClients.RecordInfUpdateInterval, a.ctx)	
	}
}

func (a *NailsApp) onBeforeRenderHTML(sock *httpSrv.HTTPSocket, resp *response.Response) error {
	
	if err := a.BeforeRenderHTML(sock, resp); err != nil {
		return err
	}
	
	//+custom variables
	conf := a.GetConfig()
	cl_vars := &NailsVariables{ServerVars: httpSrv.ServerVars{Title: PAGE_TITLE,
		Author: PAGE_AUTHOR,
		Keywords: PAGE_KEYWORS,
		Description: PAGE_DESCRIPTION,
		Version: a.MD.Version.Value,		
	}}	
	if conf.(*nails_config.NailsConfig).JsDebug {
		cl_vars.Debug = 1
	}else{
		cl_vars.Debug = 0
	}
	cl_vars.App_id = conf.GetAppID()		
	cl_vars.DefColorPalette = DEF_COLOR_PALETTE	
	
	sess := sock.GetSession()	
	cl_vars.Role_id = sess.GetString(osbe.SESS_ROLE)		
	
	cl_vars.WsPort = conf.(*nails_config.NailsConfig).WsExtPort
	cl_vars.WsPortTls = conf.(*nails_config.NailsConfig).WsExtPortTls
	
	if sess.GetBool(controllers.SESS_VAR_LOGGED) {
		cl_vars.User_id = sess.GetInt(controllers.SESS_VAR_ID)
		cl_vars.Token = sock.Token
		cl_vars.TokenExpires = sock.TokenExpires.Format(fields.FORMAT_DATE_TIME_TZ1)				
		
		cl_vars.User_name = sess.GetString(controllers.SESS_VAR_NAME)
		cl_vars.Photo = sess.GetString(controllers.SESS_VAR_PHOTO)

		if cl_vars.Role_id == "specialist" {
			d_store,_ := a.GetDataStorage().(*pgds.PgProvider)
			pool_conn, conn_id, err_с := d_store.GetSecondary("")
			if err_с != nil {
				return err_с
			}
			defer d_store.Release(pool_conn, conn_id)
			conn := pool_conn.Conn()
			
			if err := conn.QueryRow(context.Background(),			
				`SELECT count(*) FROM specialist_documents_for_sign_list WHERE specialist_id = $1`,				
				sess.GetInt("SPECIALIST_ID"),
			).Scan(&cl_vars.UnsignedDocCount); err != nil {
				return err
			}
		}
		
	}else{
		reg_capt := RegCaptcha{Id: controllers.SPECIALIST_CHECK_CAPTCHA_ID,
			Width: controllers.SPECIALIST_CHECK_CAPTCHA_W,
			Height: controllers.SPECIALIST_CHECK_CAPTCHA_H,
			Count: controllers.SPECIALIST_CHECK_CAPTCHA_CNT,
		}
		reg_capt_b, err := json.Marshal(&reg_capt)
		if err != nil {
			return err
		}
		cl_vars.RegCaptcha = string(reg_capt_b)
		
		tel_pref := rusTel.GetValidPrefixes()
		tel_pref_b, err := json.Marshal(&tel_pref)
		if err != nil {
			return err
		}
		cl_vars.ValidTelPrefixes = string(tel_pref_b)
		
		d_store,_ := a.GetDataStorage().(*pgds.PgProvider)
		pool_conn, conn_id, err_с := d_store.GetSecondary("")
		if err_с != nil {
			return err_с
		}
		defer d_store.Release(pool_conn, conn_id)
		conn := pool_conn.Conn()
		
		if err := conn.QueryRow(context.Background(),			
			`SELECT const_join_contract_val()`,				
		).Scan(&cl_vars.JoinContract); err != nil {
			return err
		}		
	}	
	
	locale := sess.GetString(osbe.SESS_LOCALE)
	if locale == "" {
		locale = conf.GetDefaultLocale()
	}
	
	cl_vars.CurDate = time.Now().Unix()
	if cl_vars.Debug == 1 {
		cl_vars.ScriptId = osbe.GenUniqID(12)
	}else{
		cl_vars.ScriptId = cl_vars.Version
	}
	cl_vars.Locale_id = locale		
	
	resp.AddModel(httpSrv.NewServerVarsModel(cl_vars))

	return nil
}

func (a *NailsApp) GetLogLevel(logLevel string) log.Lvl {
	var lvl log.Lvl

	switch logLevel {
	case "debug":
		lvl = log.DEBUG
		break
	case "info":
		lvl = log.INFO
		break
	case "warn":
		lvl = log.WARN
		break
	case "error":
		lvl = log.ERROR
		break
	default:
		lvl = log.INFO
	}
	return lvl
}

//cleanup before socket close
/*
func (a *GLabApp) OnCloseSocket(sock socket.ClientSocketer){
	sess := sock.GetSession()
	if sess == nil {
		return
	}
	user_id := sess.GetInt(controllers.SESS_VAR_ID)
	if user_id == 0 {
		return
	}
	go func(app *GLabApp){
		d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
		var conn_id pgds.ServerID
		var pool_conn *pgxpool.Conn
		pool_conn, conn_id, err_с := d_store.GetPrimary()
		if err_с != nil {
			app.GetLogger().Errorf("OnCloseSocket GetPrimary(): %v", err_с)
			return
		}
		defer d_store.Release(pool_conn, conn_id)
		conn := pool_conn.Conn()
		
		if _, err := conn.Exec(context.Background(),
			`DELETE FROM user_operations WHERE user_id = $1 AND status='end'`,
			user_id,
		); err != nil {		
			app.GetLogger().Errorf("OnCloseSocket DELETE conn.Exec(): %v", err)	
		}
	}(a)
}
*/
