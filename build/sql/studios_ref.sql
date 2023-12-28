--Refrerece type
CREATE OR REPLACE FUNCTION studios_ref(studios)
  RETURNS json AS
$$
	SELECT json_build_object(
		'keys',json_build_object(
			'id',$1.id    
			),	
		'descr',$1.name,
		'dataType','studios'
	);
$$
  LANGUAGE sql VOLATILE COST 100;
ALTER FUNCTION studios_ref(studios) OWNER TO ;	
	
