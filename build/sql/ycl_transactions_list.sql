-- VIEW: public.ycl_transactions_list

--DROP VIEW public.ycl_transactions_list;

CREATE OR REPLACE VIEW public.ycl_transactions_list AS
	SELECT
		t.id
		,t.document_id
		,t.date
		,t.amount
		,t.data->'client'->>'name' AS client
		,t.data->'client'->>'phone' AS client_phone
		,specialists_ref(specialists_ref_t) AS specialists_ref
		,json_build_object(
			'keys', json_build_object('id', ycl_staff.id),
			'descr', ycl_staff.name
		) AS ycl_staff_ref
		,t.specialist_id
		,t.seance_length
		
	FROM public.ycl_transactions AS t
	LEFT JOIN specialists AS specialists_ref_t ON specialists_ref_t.id = t.specialist_id
	LEFT JOIN ycl_staff ON ycl_staff.id = t.staff_id
	ORDER BY t.date DESC
	;
	
ALTER VIEW public.ycl_transactions_list OWNER TO ;
