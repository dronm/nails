-- VIEW: public.studios_list

--DROP VIEW public.studios_list;

CREATE OR REPLACE VIEW public.studios_list AS
	SELECT
		t.id
		,t.name
	FROM public.studios AS t
	
	ORDER BY name ASC
	;
	
ALTER VIEW public.studios_list OWNER TO ;
