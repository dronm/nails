package controllers

import(
	"math/rand"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"context"
	"time"
	"os/exec"
	"strings"
	
	"github.com/dronm/ds/pgds"
	"osbe"
	"osbe/response"
	"osbe/socket"
	
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4"
	//"github.com/jackc/pgconn"
)

const (
	RESP_ER_PREDEFINED_DEL = 1100
	RESP_ER_PREDEFINED_DESCR_DEL = "Удалять предопределенный элемент запрещено."
	
	RESP_ER_PREDEFINED_UPD = 1101
	RESP_ER_PREDEFINED_DESCR_UPD = "Редактировать предопределенный элемент запрещено."

	RESP_ER_NOT_ALLOWED_CODE = 1200
	RESP_ER_NOT_ALLOWED_DESCR = "Запрещено менять чужой докуент"

	
	MOD_HEAD_Q = `INSERT INTO object_mod_log
		(user_descr, object_type, object_id, object_descr, date_time, action)
		VALUES `
	
	CACHE_DIR = "CACHE"
	CACHE_PATH = "CACHE/"
	
	PWD_LEN = 6
	DEF_PASSWORD = "123456" //временно	
)


func GetMd5(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func genUniqID(maxLen int) string{
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, maxLen)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)	
}

func genUniqNumberID(maxLen int) string{
	var letterRunes = []rune("1234567890")
	b := make([]rune, maxLen)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)	
}

func GetUUID() (string, error) {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func checkPredefined(app osbe.Applicationer, itemID int64, table string) error {
	d_store,_ := app.GetDataStorage().(*pgds.PgProvider)
	var conn_id pgds.ServerID
	var pool_conn *pgxpool.Conn
	defer d_store.Release(pool_conn, conn_id)
	pool_conn, conn_id, err := d_store.GetSecondary("")
	if err != nil {
		return err
	}
	conn := pool_conn.Conn()
	
	predefined := false
	if err := conn.QueryRow(context.Background(), fmt.Sprintf("SELECT predefined FROM %s WHERE id=$1",table), itemID).Scan(&predefined);err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("SELECT predefined Scan(): %v",err))
	}
	if predefined {
		return osbe.NewPublicMethodError(RESP_ER_PREDEFINED_UPD, RESP_ER_PREDEFINED_DESCR_UPD)
	}
	return nil
}
	
func currentUserID(sock socket.ClientSocketer) int64{
	sess := sock.GetSession()	
	user_id := sess.GetInt("USER_ID")
	if user_id == 0 {
		user_id = 1
	}
	
	return user_id
}


//emits event
func EmitEvent(conn *pgx.Conn, evntID string, objID int64) {
	conn.Exec(context.Background(), fmt.Sprintf(`SELECT pg_notify('%s', json_build_object('params', json_build_object('keys', json_build_object('id', %d)))::text)`, evntID, objID))
}

func GenUserPwd() (string, string) {
	pwd := DEF_PASSWORD //genUniqID(PWD_LEN)	
	return pwd, GetMd5(pwd)
}	
	
func IsDigit(s string) bool {
	is_not_digit := func(c rune) bool { return c < '0' || c > '9' }
	return strings.IndexFunc(s, is_not_digit) == -1
}
	
	
