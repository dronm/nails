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

const SpecialistPeriodSalary_contr = "SpecialistPeriodSalary_Controller"

func TestSpecialistPeriodSalary_insert(t *testing.T) {
	cl, params := GetClient()
	params["period"] = '2023-12-12'
	if err := cl.SendGet(SpecialistPeriodSalary_contr, "insert", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistPeriodSalary_get_list(t *testing.T) {
	cl, params := GetClient()
	if err := cl.SendGet(SpecialistPeriodSalary_contr, "get_list", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistPeriodSalary_get_object(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(SpecialistPeriodSalary_contr, "get_object", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSpecialistPeriodSalary_delete(t *testing.T) {
	cl, params := GetClient()
	params["id"] = 1
	if err := cl.SendGet(SpecialistPeriodSalary_contr, "delete", VIEW, "", params); err != nil {
		t.Fatalf("%v", err)
	}
}



