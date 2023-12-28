-- VIEW: public.bank_payments_list

--DROP VIEW public.bank_payments_list;

CREATE OR REPLACE VIEW public.bank_payments_list AS
	SELECT
		t.id
		,t.date_time
		,t.document_date
		,t.document_num
		,t.document_total
		,t.document_comment
		,t.payer_acc
		,t.payer_bank_acc
		,t.payer_bank_bik
		,t.payer_bank
		,t.payer_bank_place
		,t.rec_acc
		,t.rec_bank_acc
		,t.rec_bank_bik
		,t.rec_bank
		,t.rec_bank_place
		,t.specialist_id
		,specialists_ref(specialists_ref_t) AS specialists_ref
		,t.specialist_period_salary_detail_id
		,specialist_period_salary_details_ref(specialist_period_salary_details_ref_t) AS specialist_period_salary_details_ref
	FROM public.bank_payments AS t
	LEFT JOIN specialists AS specialists_ref_t ON specialists_ref_t.id = t.specialist_id
	LEFT JOIN specialist_period_salary_details AS specialist_period_salary_details_ref_t ON specialist_period_salary_details_ref_t.id = t.specialist_period_salary_detail_id
	ORDER BY t.document_num DESC	
	;
	
ALTER VIEW public.bank_payments_list OWNER TO ;
