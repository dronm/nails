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

const SpecialistDocument_contr = "SpecialistDocument_Controller"

func TestSpecialistDocument_insert(t *testing.T) {
	cl, params := GetClient()
	params["date_time"] = '2023-11-28PKT07:55:55'
	if err := cl.SendGet(SpecialistDocument_contr, "insert", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistDocument_get_list(t *testing.T) {
	cl, params := GetClient()
	if err := cl.SendGet(SpecialistDocument_contr, "get_list", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistDocument_get_object(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(SpecialistDocument_contr, "get_object", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistDocument_delete(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(SpecialistDocument_contr, "delete", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}



