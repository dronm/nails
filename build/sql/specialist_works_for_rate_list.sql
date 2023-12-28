-- VIEW: public.specialist_works_for_rate_list

--DROP VIEW public.specialist_works_for_rate_list;

CREATE OR REPLACE VIEW public.specialist_works_for_rate_list AS
	SELECT
		t.id
		,t.specialist_id
		,specialists_ref(sp) AS specialists_ref
		,t.studio_id
		,studios_ref(st) AS studios_ref
		,t.date_time
		,admin_rate
		,att.content_info || jsonb_build_object('dataBase64',encode(att.content_data, 'base64')) AS photo
	FROM public.specialist_works AS t
	LEFT JOIN attachments AS att ON att.ref->>'dataType' = 'specialist_works' AND (att.ref->'keys'->>'id')::int = t.id
	LEFT JOIN specialists AS sp ON sp.id = t.specialist_id
	LEFT JOIN studios AS st ON st.id = t.studio_id
	WHERE admin_rate IS NULL
	ORDER BY t.date_time DESC	
	;
	
ALTER VIEW public.specialist_works_for_rate_list OWNER TO ;
