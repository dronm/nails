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

const SalaryKredit_contr = "SalaryKredit_Controller"

func TestSalaryKredit_insert(t *testing.T) {
	cl, params := GetClient()
	params["name"] = ``
	if err := cl.SendGet(SalaryKredit_contr, "insert", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSalaryKredit_get_list(t *testing.T) {
	cl, params := GetClient()
	if err := cl.SendGet(SalaryKredit_contr, "get_list", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSalaryKredit_get_object(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(SalaryKredit_contr, "get_object", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSalaryKredit_delete(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(SalaryKredit_contr, "delete", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}



