-- VIEW: public.specialist_period_salaries_list

DROP VIEW public.specialist_period_salaries_list;

CREATE OR REPLACE VIEW public.specialist_period_salaries_list AS
	SELECT
		t.id
		,t.date_time
		,t.period
		,month_period_rus(t.period) AS period_descr
		,t.studio_id
		,studios_ref(st) AS studios_ref
		,t.work_total
		,t.work_total_salary
		,t.hours
		,t.debet
		,t.kredit
		,t.rent_total
		,t.total
	FROM public.specialist_period_salaries AS t
	LEFT JOIN studios AS st ON st.id = t.studio_id
	ORDER BY period DESC
	;
	
ALTER VIEW public.specialist_period_salaries_list OWNER TO ;
