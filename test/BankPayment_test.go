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

const BankPayment_contr = "BankPayment_Controller"

func TestBankPayment_insert(t *testing.T) {
	cl, params := GetClient()
	params["specialist_period_salary_detail_id"] = 1
	if err := cl.SendGet(BankPayment_contr, "insert", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBankPayment_get_list(t *testing.T) {
	cl, params := GetClient()
	if err := cl.SendGet(BankPayment_contr, "get_list", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBankPayment_get_object(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(BankPayment_contr, "get_object", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBankPayment_delete(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(BankPayment_contr, "delete", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}



