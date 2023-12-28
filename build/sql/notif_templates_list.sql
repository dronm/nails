-- VIEW: public.notif_templates_list

--DROP VIEW public.notif_templates_list;

CREATE OR REPLACE VIEW public.notif_templates_list AS
	SELECT
		t.id
		,t.notif_provider
		,t.notif_type
		,t.template
	FROM public.notif_templates AS t
	ORDER BY t.notif_provider, t.notif_type
	
	;
	
ALTER VIEW public.notif_templates_list OWNER TO ;
