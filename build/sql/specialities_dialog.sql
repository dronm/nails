-- VIEW: public.specialities_dialog

--DROP VIEW public.specialities_dialog;

CREATE OR REPLACE VIEW public.specialities_dialog AS
	SELECT
		t.id
		,t.name
		,t.equipments
		,t.agent_percent
	FROM public.specialities AS t
	;
	
ALTER VIEW public.specialities_dialog OWNER TO ;
