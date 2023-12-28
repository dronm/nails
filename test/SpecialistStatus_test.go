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

const SpecialistStatus_contr = "SpecialistStatus_Controller"

func TestSpecialistStatus_insert(t *testing.T) {
	cl, params := GetClient()
	params["date_time"] = '2023-11-13PKT15:57:38'
	if err := cl.SendGet(SpecialistStatus_contr, "insert", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistStatus_get_list(t *testing.T) {
	cl, params := GetClient()
	if err := cl.SendGet(SpecialistStatus_contr, "get_list", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistStatus_get_object(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(SpecialistStatus_contr, "get_object", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistStatus_delete(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(SpecialistStatus_contr, "delete", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}



