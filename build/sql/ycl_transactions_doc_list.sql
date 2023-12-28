-- VIEW: public.ycl_transactions_doc_list

--DROP VIEW public.ycl_transactions_doc_list;

CREATE OR REPLACE VIEW public.ycl_transactions_doc_list AS
	SELECT
		t.document_id AS document_id
		,t.date AS date
		,sum(t.amount) AS amount
		,t.data->'client'->>'name' || ' (' || (t.data->'client'->>'phone')::text || ')' AS client
		,t.specialist_id
		,t.seance_length AS seance_length
		
	FROM public.ycl_transactions AS t	
	WHERE NOT EXISTS (SELECT id FROM specialist_works WHERE ycl_document_id=(t.data->>'document_id')::int)
	GROUP BY		
		t.document_id,
		t.seance_length,
		t.date,
		t.data->'client',
		t.specialist_id
	ORDER BY
		t.document_id,		
		t.date
		
	;
	
ALTER VIEW public.ycl_transactions_doc_list OWNER TO ;
