-- FUNCTION: public.specialities_ref(specialities)

-- DROP FUNCTION IF EXISTS public.specialities_ref(specialities);

CREATE OR REPLACE FUNCTION public.specialities_ref(
	specialities)
    RETURNS json
    LANGUAGE 'sql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
	SELECT json_build_object(
		'keys',json_build_object(
			'id',$1.id    
			),	
		'descr',$1.name,
		'dataType','specialities'
	);
$BODY$;

ALTER FUNCTION public.specialities_ref(specialities)
    OWNER TO ;

