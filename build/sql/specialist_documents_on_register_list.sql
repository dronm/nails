-- VIEW: public.specialist_documents_on_register_list

--DROP VIEW public.specialist_documents_on_register_list;

CREATE OR REPLACE VIEW public.specialist_documents_on_register_list AS
	SELECT
		t.id
		,document_templates_ref(document_templates_ref_t) AS document_templates_ref
		,t.date_time
		,t.need_signing
	FROM public.specialist_documents_on_register AS t
	LEFT JOIN document_templates AS document_templates_ref_t ON document_templates_ref_t.id = t.document_template_id
	ORDER BY document_templates_ref_t.name
	;
	
ALTER VIEW public.specialist_documents_on_register_list OWNER TO ;
