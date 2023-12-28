package tests

/**
 * Andrey Mikhalevich 18/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller_test.go.tmpl
 */

import(
	"testing"
	
	"nails/enums"
)

const NotifTemplate_contr = "NotifTemplate_Controller"

func TestNotifTemplate_insert(t *testing.T) {
	cl, params := GetClient()
	params["fields"] = `{&quot;a&quot;:&quot;1&quot;,&quot;b&quot;:&quot;abc&quot;}`
	if err := cl.SendGet(NotifTemplate_contr, "insert", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestNotifTemplate_get_list(t *testing.T) {
	cl, params := GetClient()
	if err := cl.SendGet(NotifTemplate_contr, "get_list", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestNotifTemplate_get_object(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(NotifTemplate_contr, "get_object", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestNotifTemplate_delete(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(NotifTemplate_contr, "delete", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}



