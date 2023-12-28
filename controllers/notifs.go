package controllers

import(
	"errors"
	"fmt"
	"strings"
	
	"osbe/repo/notif"
	
	"github.com/jackc/pgx/v4"	
)

//Здесь собраны функции отправки различных сообщений
//специфичные для проекта

var NailsNotifier *notif.Notifier

func InitNotifier(host, appName, pwd string){
	NailsNotifier = notif.NewNotifier(host, appName, pwd)	
}


func NotifSend(batch []*notif.NotifMessage) error {
	resp, err := NailsNotifier.Send(batch)		
	if err != nil {
		return err
	}
	
	var e strings.Builder
	for _, r := range resp {
		if r.Error != "" {
			if e.Len() > 0 {
				e.WriteString(", ")
			}
			e.WriteString(fmt.Sprintf("Error sending message ID: %d, text: %s", r.ID, r.Error))
		}
	}
	if e.Len() > 0 {
		return errors.New(e.String())
	}
	return nil
}

//Отправляем wa + email + attachments
func NotifNewAccount(conn *pgx.Conn, specialistID int64, pwd string, attachments []string, attachmentAlias []string) error {
	mag_type := "new_account"
	
	//wa meassage
	wa_msg, err := notif.NewWAMessageFromSQL(conn, "wa_" + mag_type, []interface{}{specialistID, pwd})	
	if err != nil {
		return err
	}	
	
	//email message
	e_msg, err := notif.NewEmailMessageFromSQL(conn, "email_" + mag_type, attachments, attachmentAlias, []interface{}{specialistID, pwd})	
	if err != nil {
		return err
	}
		
	//send all
	return NotifSend([]*notif.NotifMessage{wa_msg.NewNotif(mag_type), e_msg.NewNotif(mag_type)})
}

//wa message to admin
func NotifNewSpecialist(conn *pgx.Conn) error {
	msg_wa_list, err := notif.NewWAMessagesFromSQL(conn, "wa_new_specialist", nil)
	if err != nil {
		return err
	}
	notif_list := make([]*notif.NotifMessage, len(msg_wa_list))
	for i, msg := range msg_wa_list {
		notif_list[i] = msg.NewNotif("new_specialist")
	}
	return NotifSend(notif_list)
}

//wa + sms
func NotifTelCheck(conn *pgx.Conn, tel, telKey string) error {
	masg_type := "tel_check"
	
	if len(tel) == 10 {
		tel = "7" + tel
	}
	//+wa
	msg_wa, err := notif.NewWAMessageFromSQL(conn, "wa_" + masg_type, []interface{}{tel, telKey})
	if err != nil {
		return err
	}
	//+sms
	msg_sms, err := notif.NewSMSMessageFromSQL(conn, "sms_" + masg_type, []interface{}{tel, telKey})
	if err != nil {
		return err
		
	}
	//both providers to one message
	notif_msg := notif.NewNotifMessage(masg_type, notif.PROV_WA, notif.PROV_SMS)		
	msg_wa.SetParams(notif_msg) //wa		
	msg_sms.SetParams(notif_msg) //sms		
			
	return NotifSend([]*notif.NotifMessage{notif_msg})	
}

