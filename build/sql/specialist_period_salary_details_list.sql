-- VIEW: public.specialist_period_salary_details_list

--DROP VIEW public.specialist_period_salary_details_list;

CREATE OR REPLACE VIEW public.specialist_period_salary_details_list AS
	SELECT
		t.id
		,t.specialist_period_salary_id
		,t.line_num
		,t.specialist_id
		,specialists_ref(specialists_ref_t) AS specialists_ref
		,t.studio_id
		,studios_ref(studios_ref_t) AS studios_ref
		,t.period		
		,month_period_rus(t.period) AS period_descr
		,t.hours
		,t.agent_percent
		,t.work_total
		,t.work_total_salary		
		,t.debet
		,t.kredit
		,t.rent_price
		,t.rent_total
		,t.total
		,coalesce(rct.receipt_total, 0) AS receipt_total
		,rct.receipt_error
		,rct.receipt_photos
		,coalesce(rct.receipt_checked, FALSE) AS receipt_checked
		,rct.receipt_href
		,bank_payments_ref(pp) AS bank_payments_ref
		
		,person_init(specialists_ref_t.name) || ' '|| month_period_rus(t.period) AS descr
		
	FROM public.specialist_period_salary_details AS t
	LEFT JOIN specialists AS specialists_ref_t ON specialists_ref_t.id = t.specialist_id
	LEFT JOIN studios AS studios_ref_t ON studios_ref_t.id = t.studio_id
	LEFT JOIN bank_payments AS pp ON pp.specialist_period_salary_detail_id = t.id
	LEFT JOIN (
		SELECT
			t.specialist_period_salary_detail_id,
			sum(t.document_total) AS receipt_total,
			string_agg(t.document_error, ', ') AS receipt_error,
			string_agg(t.document_href, ', ') AS receipt_href,
			jsonb_agg(
				att.content_info || jsonb_build_object(
					'dataBase64', encode(att.content_preview, 'base64'),
					'receipt_id', t.id
				)
			) AS receipt_photos,
			bool_and(t.document_parsed) AS receipt_checked
		FROM specialist_receipts AS t
		LEFT JOIN attachments AS att ON att.ref->>'dataType' = 'specialist_receipts' AND (att.ref->'keys'->>'id')::int = t.id
		GROUP BY t.specialist_period_salary_detail_id
	) AS rct ON rct.specialist_period_salary_detail_id = t.id
	
	ORDER BY
		t.specialist_period_salary_id,
		t.line_num
	;
	
ALTER VIEW public.specialist_period_salary_details_list OWNER TO ;
