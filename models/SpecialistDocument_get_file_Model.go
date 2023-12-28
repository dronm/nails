package models

//Controller method model
import (
		
	"osbe/fields"
	
	"osbe/repo/docAttachment"
)

type SpecialistDocument_get_file struct {
	Ref docAttachment.Ref_Type `json:"ref" required:"true"`
	Content_id fields.ValText `json:"content_id" required:"true"`
	Inline fields.ValInt `json:"inline"`
	Doc_id fields.ValInt `json:"doc_id" required:"true"`
}

type SpecialistDocument_get_file_argv struct {
	Argv *SpecialistDocument_get_file `json:"argv"`	
}

