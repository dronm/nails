-- VIEW: public.specialist_documents_for_sign_list

--DROP VIEW public.specialist_documents_for_sign_list;

CREATE OR REPLACE VIEW public.specialist_documents_for_sign_list AS
	SELECT
		t.id
		,t.specialist_id
		,sp.user_id AS user_id
		,t.specialists_ref
		,t.document_att_ref
		,t.date_time
		,t.opened
		,t.name
	FROM public.specialist_documents_list AS t
	LEFT JOIN specialists AS sp ON sp.id = t .id
	WHERE t.need_signing AND t.sign_date_time IS NULL
	ORDER BY t.date_time DESC
	;
	
ALTER VIEW public.specialist_documents_for_sign_list OWNER TO ;
