package models

/*
type ConfirmationStatus_set_confirmed struct {
	Secret fields.ValText `json:"secret" required:"true"`
}
*/
type ConfirmationStatus_set_confirmed map[string]string

type ConfirmationStatus_set_confirmed_argv struct {
	Argv *ConfirmationStatus_set_confirmed `json:"argv"`	
}

