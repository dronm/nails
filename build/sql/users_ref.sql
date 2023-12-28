-- Function: public.users_ref(users)

-- DROP FUNCTION public.users_ref(users);

CREATE OR REPLACE FUNCTION public.users_ref(users)
  RETURNS json AS
$BODY$
	SELECT json_build_object(
		'keys',json_build_object(
			'id',$1.id    
			),	
		'descr', coalesce($1.name_full, $1.name::text),
		'dataType','users'
	);
$BODY$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION public.users_ref(users) OWNER TO ;

