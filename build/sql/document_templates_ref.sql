--Refrerece type
CREATE OR REPLACE FUNCTION document_templates_ref(document_templates)
  RETURNS json AS
$$
	SELECT json_build_object(
		'keys',json_build_object(
			'id',$1.id    
			),	
		'descr',$1.name,
		'dataType','document_templates'
	);
$$
  LANGUAGE sql VOLATILE COST 100;
ALTER FUNCTION document_templates_ref(document_templates) OWNER TO ;	
	
