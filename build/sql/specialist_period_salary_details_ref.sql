-- FUNCTION: public.specialist_period_salary_details_ref(specialist_period_salary_details)

-- DROP FUNCTION IF EXISTS public.specialist_period_salary_details_ref(specialist_period_salary_details);

CREATE OR REPLACE FUNCTION public.specialist_period_salary_details_ref(
	specialist_period_salary_details)
    RETURNS json
    LANGUAGE 'sql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
	SELECT json_build_object(
		'keys',json_build_object(
			'id',$1.id    
			),	
		'descr','Зарплата ' || (SELECT st.name FROM studios AS st WHERE st.id=$1.studio_id)||' '||month_period_rus($1.period),
		'dataType','specialist_period_salary_details'
	);
$BODY$;

ALTER FUNCTION public.specialist_period_salary_details_ref(specialist_period_salary_details)
    OWNER TO ;

