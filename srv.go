package main

import (
	"os"	
	"fmt"
	"path/filepath"
	"io/ioutil"
	
	"nails/nails_config"
	
	"osbe/config"
	_ "osbe/view/json"
	_ "osbe/view/xml"
	_ "osbe/view/html"
	_ "osbe/view/excel"
	_ "osbe/view/pdf"
)

var APP_VERSION string

func main() {	
	
	//Configuration file: first argument or PROG_FILE_NAME.json
	var ini_file string
	if len(os.Args) >= 2 {
		ini_file = os.Args[1]
	}else{
		ini_file = os.Args[0]+".json"
	}

	ex, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("os.Executable() faled: %v",err))
	}
	base_dir := filepath.Dir(ex)
	if err:= ReadVersion(base_dir); err != nil {
		panic(fmt.Sprintf("ReadVersion: %v",err))
	}

	conf := &nails_config.NailsConfig{}
	err = config.ReadConf(ini_file, conf)
	if err != nil {
		panic(fmt.Sprintf("ReadConf: %v, file: %s", err, ini_file))
	}
	app := &NailsApp{}
	app.BaseDir = base_dir
	app.ConfigFileName = ini_file
	app.init(conf)
	
	//other servers should be placed here
	
	//event server
	if conf.EventServerEnabled {
		go app.EvntServer.Run()
	}
	//ws server
	if conf.WsServerEnabled {
		go app.GetServer("ws").Run()
	}
	
	//http server in main gorouting
	app.GetServer("http").Run()
}

func ReadVersion(fileDir string) error{	
	ver, err := ioutil.ReadFile(filepath.Join(fileDir, "version.txt"))
	if err != nil {
		return err
	}
	if []rune(string(ver))[len(ver)-1] == 10 {
		ver = ver[0:len(ver)-1]
	}
	APP_VERSION = string(ver)
	return err
}

