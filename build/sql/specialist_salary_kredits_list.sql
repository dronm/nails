-- VIEW: public.specialist_salary_kredits_list

--DROP VIEW public.specialist_salary_kredits_list;

CREATE OR REPLACE VIEW public.specialist_salary_kredits_list AS
	SELECT
		t.id
		,t.date_time
		,t.specialist_id
		,specialists_ref(specialists_ref_t) AS specialists_ref
		,salary_kredits_ref(salary_kredits_ref_t) AS salary_kredits_ref
		,t.total
	FROM public.specialist_salary_kredits AS t
	LEFT JOIN specialists AS specialists_ref_t ON specialists_ref_t.id = t.specialist_id
	LEFT JOIN salary_kredits AS salary_kredits_ref_t ON salary_kredits_ref_t.id = t.salary_kredit_id
	ORDER BY date_time DESC
	;
	
ALTER VIEW public.specialist_salary_kredits_list OWNER TO ;
