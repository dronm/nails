	
--Refrerece type
CREATE OR REPLACE FUNCTION salary_kredits_ref(salary_kredits)
  RETURNS json AS
$$
	SELECT json_build_object(
		'keys',json_build_object(
			'id',$1.id    
			),	
		'descr',$1.name,
		'dataType','salary_kredits'
	);
$$
  LANGUAGE sql VOLATILE COST 100;
ALTER FUNCTION salary_kredits_ref(salary_kredits) OWNER TO nails;	
	
