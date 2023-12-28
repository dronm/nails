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

const SpecialistReg_contr = "SpecialistReg_Controller"

func TestSpecialistReg_insert(t *testing.T) {
	cl, params := GetClient()
	params["inn_fns_ok"] = true
	if err := cl.SendGet(SpecialistReg_contr, "insert", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistReg_get_list(t *testing.T) {
	cl, params := GetClient()
	if err := cl.SendGet(SpecialistReg_contr, "get_list", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistReg_get_object(t *testing.T) {
	cl, params := GetClient()
	params["user_operation_id"] = 1
	if err := cl.SendGet(SpecialistReg_contr, "get_object", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistReg_delete(t *testing.T) {
	cl, params := GetClient()
	params["user_operation_id"] = `Some text string`
	if err := cl.SendGet(SpecialistReg_contr, "delete", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}



