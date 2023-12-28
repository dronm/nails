package controllers

/**
 * Project: nails
 * Author: Andery Mikhalevich
 * Email: katrenplus@mail.ru
 * Date: 23/11/2023
 *
 * This is ConfirmationStatus controller implimentation file.
 *
 */

import (
	"reflect"	
	"io/ioutil"
	"strings"
	"net/http"
	"context"
	"fmt"
	"path/filepath"
	"os"
	"time"
	
	//"nails/models"
	
	"github.com/dronm/ds/pgds"
	
	"nails/nails_config"
	
	"osbe"
	//"osbe/model"
	"osbe/srv"
	"osbe/socket"
	"osbe/response"	
	"osbe/srv/httpSrv"
	"osbe/view"
	
	//"github.com/jackc/pgx/v4"
)

const (
	AUTH_PREF = "token "
	
	QR_CODE_FILE_NAME = "qr.png"
)

//Method implemenation get_qr
func (pm *Mgateway_Controller_get_qr) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	var cache_f *os.File
	var err error
	cache_fn := filepath.Join(app.GetBaseDir(), CACHE_DIR, QR_CODE_FILE_NAME)
	if view.FileExists(cache_fn) {
		//from cache		
		cache_f, err = os.Open(cache_fn)
		if err != nil {
			return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Mgateway_Controller_get_qr os.Open(): %v", err))
		}
		defer cache_f.Close()		
		
		if err := httpSrv.DownloadFile(resp, sock, cache_f, "qr.png", "", httpSrv.CONTENT_DISPOSITION_INLINE); err != nil {
			return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Mgateway_Controller_get_qr DownloadFile(): %v", err))
		}
		
	}else{
		//no cache, read from db && save
		d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
		pool_conn, conn_id, err_с := d_store.GetSecondary("")
		if err_с != nil {
			return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register Request.GetPrimary(): %v",err_с))
		}
		defer d_store.Release(pool_conn, conn_id)
		conn := pool_conn.Conn()
		
		var f_cont []byte//&fields.ValBytea{}
		if err := conn.QueryRow(context.Background(),
			`SELECT wa_qrcode FROM apps WHERE name = 'nails'`,
		).Scan(&f_cont); err != nil {
			return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Mgateway_Controller_get_qr conn.QueryRow(): %v",err))	
		}
		
		cache_f, err = os.Create(cache_fn)
		if err != nil {
			return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Mgateway_Controller_get_qr os.Create(): %v", err))
		}
		defer cache_f.Close()
		if _, err := cache_f.Write(f_cont); err != nil {
			return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Mgateway_Controller_get_qr Write(): %v", err))
		}
		sock_http, ok := sock.(*httpSrv.HTTPSocket)
		if !ok {
			return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, "Mgateway_Controller_get_qr no cache sock must be *HTTPSocket")
		}
		
		httpSrv.ServeContent(sock_http, &f_cont, "qr.png", "", time.Time{}, httpSrv.CONTENT_DISPOSITION_INLINE)		
	}
	
	return nil

}

//Method implemenation callback
func (pm *Mgateway_Controller_callback) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	http_sock, ok := sock.(*httpSrv.HTTPSocket)
	if !ok {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, "Mgateway_Controller_callback Not HTTPSocket type")
	}
	r := http_sock.Request
	
	auth := r.Header.Get("Authorization")

	if len(auth) < len(AUTH_PREF) || strings.ToLower(auth[0:len(AUTH_PREF)]) != AUTH_PREF {
		app.GetLogger().Errorf("Mgateway_Controller_callback: unauthorized access %s", r.RemoteAddr)
		http_sock.Response.WriteHeader(http.StatusUnauthorized)
		return nil
	}
	
	//b64.StdEncoding.DecodeString(
	token := auth[len(AUTH_PREF):]
	conf := app.GetConfig().(*nails_config.NailsConfig)
	if token != conf.NotifierParams.CallbackKey {
		app.GetLogger().Errorf("Mgateway_Controller_callback: unauthorized access, wrong token %s", r.RemoteAddr)
		http_sock.Response.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.GetLogger().Errorf("Mgateway_Controller_callback: ioutil.ReadAll() failed %v", err)
		http_sock.Response.WriteHeader(http.StatusBadGateway)
		return nil
		
	}
	defer r.Body.Close()
		
	cont := r.Header.Get("Content-Type")
	if cont == "image/png" {
		//qr from wa
		
		d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
		pool_conn, conn_id, err_с := d_store.GetPrimary()
		if err_с != nil {
			return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register Request.GetPrimary(): %v",err_с))
		}
		defer d_store.Release(pool_conn, conn_id)
		conn := pool_conn.Conn()
		
		if _, err := conn.Exec(context.Background(),
			`UPDATE apps SET wa_qrcode = $1 WHERE name = 'nails'`,
			buf,
		); err != nil {
			return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register conn.Exec() failed: %v",err))
		}
		
		f, err := os.Create(filepath.Join(app.GetBaseDir(), CACHE_DIR, QR_CODE_FILE_NAME))
		if err != nil {
			return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register os.Create() faled: %v",err))
		}
		defer f.Close()
		
		_, err = f.Write(buf)
		if err != nil {
			return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("Specialist_Controller_register f.Write() faled: %v",err))
		}
		
	}else{
fmt.Println("massahe update ", string(buf))		
	}
	
	return nil
}


