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

const SpecialistWork_contr = "SpecialistWork_Controller"

func TestSpecialistWork_insert(t *testing.T) {
	cl, params := GetClient()
	params["studio_id"] = 1
	if err := cl.SendGet(SpecialistWork_contr, "insert", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistWork_get_list(t *testing.T) {
	cl, params := GetClient()
	if err := cl.SendGet(SpecialistWork_contr, "get_list", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistWork_get_object(t *testing.T) {
	cl, params := GetClient()
	if err := cl.SendGet(SpecialistWork_contr, "get_object", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistWork_delete(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(SpecialistWork_contr, "delete", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}



