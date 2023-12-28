-- VIEW: public.specialist_salary_debets_list

--DROP VIEW public.specialist_salary_debets_list;

CREATE OR REPLACE VIEW public.specialist_salary_debets_list AS
	SELECT
		t.id
		,t.date_time
		,t.specialist_id
		,specialists_ref(specialists_ref_t) AS specialists_ref
		,salary_debets_ref(salary_debets_ref_t) AS salary_debets_ref
		,t.total
	FROM public.specialist_salary_debets AS t
	LEFT JOIN specialists AS specialists_ref_t ON specialists_ref_t.id = t.specialist_id
	LEFT JOIN salary_debets AS salary_debets_ref_t ON salary_debets_ref_t.id = t.salary_debet_id
	ORDER BY date_time DESC
	;
	
ALTER VIEW public.specialist_salary_debets_list OWNER TO ;
