package nails_config

import(
	"osbe/config"
)

type RedisSession struct {	
	Connect string `json:"connect"`
	Namespace string `json:"namespace"`
}

type ExtServiceParams struct {	
	AppName string `json:"appName"`
	Pwd string `json:"pwd"`
	Host string `json:"host"`
	CallbackKey string `json:"callbackKey"`
}

type YClients struct {	
	Login string `json:"login"`
	Pwd string `json:"pwd"`
	PartnerToken string `json:"partnerToken"`
	CompanyID int `json:"companyId"`
	RecordInfDoUpdate bool `json:"recordInfDoUpdate"`
	RecordInfUpdateInterval int `json:"recordInfUpdateInterval"`
}

type NailsConfig struct {
	config.AppConfig //base structure
	EventServerEnabled bool `json:"eventServerEnabled"`
	WsServerEnabled bool `json:"wsServerEnabled"`
	WsExtPort int `json:"wsExtPort"`
	WsExtPortTls int `json:"wsExtPortTls"`
	HTTPDir string `json:"httpDir"`
	JsDebug bool `json:"jsDebug"`
	HttpServer string `json:"httpServer"`
	JsHost string `json:"jsHost"`
	DadataKey string `json:"dadataKey"`
	Fop string `json:"fop"`
	FopConf string `json:"fopConf"`
	Redis RedisSession `json:"redis"`
	NotifierParams ExtServiceParams `json:"notifierParams"`
	QRExtractorParams ExtServiceParams `json:"qrExtractorParams"`
	PHP string `json:"php"`
	PHPDocConverter string `json:"phpDocConverter"`
	YClients YClients `json:"yClients"`		
}

