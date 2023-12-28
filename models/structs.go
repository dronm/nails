package models

import(
	"osbe/fields"	
)

type Doc_id struct {
	Id fields.ValInt `json:"id" required:"true"`
}
type Doc_id_argv struct {
	Argv *Doc_id `json:"argv"`	
}

type DocAsync_id struct {
	Id fields.ValInt `json:"id" required:"true"`
	Operation_id fields.ValText `json:"operation_id" required:"true"`
}
type DocAsync_id_argv struct {
	Argv *DocAsync_id `json:"argv"`	
}

type Mgateway_callback map[string]string

type Mgateway_callback_argv struct {
	Argv *Mgateway_callback `json:"argv"`	
}

//
type QRExtractor_callback map[string]string

type QRExtractor_callback_argv struct {
	Argv *QRExtractor_callback `json:"argv"`	
}

//
type SpecialistWork_set_admin_rates struct {
	Rates [] struct {
		Id int `json:"id"`
		Admin_rate int `json:"admin_rate"`
	} `json:"rates"`	
}
type SpecialistWork_set_admin_rates_argv struct {
	Argv *SpecialistWork_set_admin_rates `json:"argv"`	
}

//
type SpecialistReceipt_add struct {
	Operation_id fields.ValText `json:"operation_id" required:"true"`
	Specialist_period_salary_detail_id fields.ValInt `json:"specialist_period_salary_detail_id" required:"true"`
}
type SpecialistReceipt_add_argv struct {
	Argv *SpecialistReceipt_add `json:"argv"`	
}

//All object key array
type Object_key_ar struct {
	Ids []Doc_id `json:"ids" required:"true"`
}
type Object_key_ar_argv struct {
	Argv *Object_key_ar `json:"argv"`	
}

