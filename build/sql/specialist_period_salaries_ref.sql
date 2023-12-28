--Refrerece type
CREATE OR REPLACE FUNCTION specialist_period_salaries_ref(specialist_period_salaries)
  RETURNS json AS
$$
	SELECT json_build_object(
		'keys',json_build_object(
			'id',$1.id    
			),	
		'descr',$1.id || '(' || to_char($1.date_time, 'DD/MM/YY') || ', ' || month_period_rus($1.date_time::DATE)||')',
		'dataType','specialist_period_salaries'
	);
$$
  LANGUAGE sql VOLATILE COST 100;
ALTER FUNCTION specialist_period_salaries_ref(specialist_period_salaries) OWNER TO nails;	
	

