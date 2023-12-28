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

const Bank_contr = "Bank_Controller"


func TestBank_get_list(t *testing.T) {
	cl, params := GetClient()
	if err := cl.SendGet(Bank_contr, "get_list", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBank_get_object(t *testing.T) {
	cl, params := GetClient()
	params["bik"] = `Some text`
	if err := cl.SendGet(Bank_contr, "get_object", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}


func TestBank_complete(t *testing.T) {
	cl, params := GetClient()
	params["count"] = 10
	params["ic"] = 1
	params["mid"] = 1
	params["bik"] = `Some text`
	if err := cl.SendGet(Bank_contr, "complete", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}


