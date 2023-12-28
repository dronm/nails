-- VIEW: public.ycl_transactions_doc_all_list

DROP VIEW specialist_works_list;
DROP VIEW public.ycl_transactions_doc_all_list;

CREATE OR REPLACE VIEW public.ycl_transactions_doc_all_list AS
	SELECT
		t.document_id
		,t.date
		,sum(t.amount) AS amount
		,t.data->'client'->>'name' || ' (' || (t.data->'client'->>'phone')::text || ')' AS client
		,t.specialist_id
		,specialists_ref(sp) AS specialists_ref
		,round(t.seance_length / 3600, 2) AS len_hour
		
		,att.content_info || jsonb_build_object('dataBase64',encode(att.content_data, 'base64')) AS photo
		,sp_w.admin_rate
		
	FROM public.ycl_transactions AS t	
	LEFT JOIN specialists AS sp ON sp.id = t.specialist_id
	LEFT JOIN specialist_works AS sp_w ON sp_w.ycl_document_id = (t.data->>'document_id')::int
	LEFT JOIN attachments AS att ON att.ref->>'dataType' = 'specialist_works' AND (att.ref->'keys'->>'id')::int = sp_w.id
	GROUP BY		
		t.document_id,
		t.seance_length,
		t.date,
		t.data->'client',
		t.specialist_id,
		sp.*,
		att.content_info,
		att.content_data,
		sp_w.admin_rate
	ORDER BY
		t.document_id,		
		t.date
	;
	
ALTER VIEW public.ycl_transactions_doc_all_list OWNER TO ;
