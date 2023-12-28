-- VIEW: public.template_batch_items_list

--DROP VIEW public.template_batch_items_list;

CREATE OR REPLACE VIEW public.template_batch_items_list AS
	SELECT
		t.id
		,t.template_batch_id
		,document_templates_ref(templates_ref_t) AS templates_ref
		,studios_ref(st) AS studios_ref
	FROM public.template_batch_items AS t
	LEFT JOIN document_templates AS templates_ref_t ON templates_ref_t.id = t.template_id
	LEFT JOIN studios AS st ON st.id = t.studio_id
	ORDER BY
		st.name,
		templates_ref_t.name	
	;
	
ALTER VIEW public.template_batch_items_list OWNER TO ;
