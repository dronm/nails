-- VIEW: public.specialist_works_list

--DROP VIEW public.specialist_works_list;

CREATE OR REPLACE VIEW public.specialist_works_list AS
	SELECT
		t.id
		,t.specialist_id
		,specialists_ref(sp) AS specialists_ref
		,t.studio_id
		,studios_ref(st) AS studios_ref
		,t.date_time
		,tr.admin_rate
		,att.content_info || jsonb_build_object('dataBase64',encode(att.content_preview, 'base64')) AS photo
		,coalesce(tr.len_hour, 0.0) AS hours
		,coalesce(tr.amount, 0.0) AS amount
		
	FROM public.specialist_works AS t
	LEFT JOIN attachments AS att ON att.ref->>'dataType' = 'specialist_works' AND (att.ref->'keys'->>'id')::int = t.id
	LEFT JOIN specialists AS sp ON sp.id = t.specialist_id
	LEFT JOIN studios AS st ON st.id = t.studio_id
	LEFT JOIN ycl_transactions_doc_all_list AS tr ON tr.document_id = t.ycl_document_id
	ORDER BY t.date_time DESC	
	;
	
ALTER VIEW public.specialist_works_list OWNER TO ;
