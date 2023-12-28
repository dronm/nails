-- VIEW: public.specialities_list

--DROP VIEW public.specialities_list;

CREATE OR REPLACE VIEW public.specialities_list AS
	SELECT
		t.id
		,t.name
	FROM public.specialities AS t
	ORDER BY name
	
	;
	
ALTER VIEW public.specialities_list OWNER TO ;
