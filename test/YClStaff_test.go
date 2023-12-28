package tests

/**
 * Andrey Mikhalevich 18/12/21
 * This file is part of the OSBE framework
 *
 * THIS FILE IS GENERATED FROM TEMPLATE build/templates/controllers/Controller_test.go.tmpl
 */

import(
	"testing"
	
)

const YClStaff_contr = "YClStaff_Controller"

func TestYClStaff_insert(t *testing.T) {
	cl, params := GetClient()
	if err := cl.SendGet(YClStaff_contr, "insert", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestYClStaff_get_list(t *testing.T) {
	cl, params := GetClient()
	if err := cl.SendGet(YClStaff_contr, "get_list", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestYClStaff_get_object(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(YClStaff_contr, "get_object", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestYClStaff_delete(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(YClStaff_contr, "delete", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}



