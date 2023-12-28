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

const SpecialistDocumentOnRegister_contr = "SpecialistDocumentOnRegister_Controller"

func TestSpecialistDocumentOnRegister_insert(t *testing.T) {
	cl, params := GetClient()
	params["date_time"] = '2023-11-26PKT08:27:24'
	if err := cl.SendGet(SpecialistDocumentOnRegister_contr, "insert", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistDocumentOnRegister_get_list(t *testing.T) {
	cl, params := GetClient()
	if err := cl.SendGet(SpecialistDocumentOnRegister_contr, "get_list", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistDocumentOnRegister_get_object(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(SpecialistDocumentOnRegister_contr, "get_object", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistDocumentOnRegister_delete(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(SpecialistDocumentOnRegister_contr, "delete", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}



