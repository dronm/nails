-- VIEW: public.studios_dialog

--DROP VIEW public.studios_dialog;

CREATE OR REPLACE VIEW public.studios_dialog AS
	SELECT
		t.id
		,t.name
		,t.firm_id
		,firms_ref(firms_ref_t) AS firms_ref
		,t.equipments
		,t.hour_rent_price
	FROM public.studios AS t
	LEFT JOIN firms AS firms_ref_t ON firms_ref_t.id = t.firm_id
	;
	
ALTER VIEW public.studios_dialog OWNER TO ;
