	
--Refrerece type
CREATE OR REPLACE FUNCTION salary_debets_ref(salary_debets)
  RETURNS json AS
$$
	SELECT json_build_object(
		'keys',json_build_object(
			'id',$1.id    
			),	
		'descr',$1.name,
		'dataType','salary_debets'
	);
$$
  LANGUAGE sql VOLATILE COST 100;
ALTER FUNCTION salary_debets_ref(salary_debets) OWNER TO nails;	
	
