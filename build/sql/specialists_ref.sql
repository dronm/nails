--Refrerece type
CREATE OR REPLACE FUNCTION specialists_ref(specialists)
  RETURNS json AS
$$
	SELECT json_build_object(
		'keys',json_build_object(
			'id',$1.id    
			),	
		'descr',$1.name,
		'dataType','specialists'
	);
$$
  LANGUAGE sql VOLATILE COST 100;
ALTER FUNCTION specialists_ref(specialists) OWNER TO ;	
	
