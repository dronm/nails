--Refrerece type
CREATE OR REPLACE FUNCTION template_batchs_ref(template_batches)
  RETURNS json AS
$$
	SELECT json_build_object(
		'keys',json_build_object(
			    
			),	
		'descr',$1.name,
		'dataType','template_batches'
	);
$$
  LANGUAGE sql VOLATILE COST 100;
ALTER FUNCTION template_batchs_ref(template_batches) OWNER TO nails;	
	
