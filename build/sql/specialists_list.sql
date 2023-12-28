-- VIEW: public.specialists_list

--DROP VIEW public.specialists_list;

CREATE OR REPLACE VIEW public.specialists_list AS
	SELECT
		t.id
		,t.name
		,t.inn
		,banks.banks_ref(bn) AS banks_ref
		,t.bank_acc
		,t.studio_id
		,studios_ref(studios_ref_t) AS studios_ref
		,ct.email
		,ct.tel
		,att.content_info || jsonb_build_object('dataBase64',encode(att.content_preview, 'base64')) AS photo
		,specialities_ref(spt) AS specialities_ref
		
	FROM public.specialists AS t
	LEFT JOIN studios AS studios_ref_t ON studios_ref_t.id = t.studio_id
	LEFT JOIN banks.banks AS bn ON bn.bik = t.bank_bik
	LEFT JOIN entity_contacts AS e_ct ON e_ct.entity_id = t.id AND e_ct.entity_type = 'specialists'
	LEFT JOIN contacts AS ct ON ct.id = e_ct.contact_id
	LEFT JOIN attachments AS att ON att.ref->>'dataType' = 'users' AND (att.ref->'keys'->>'id')::int = t.user_id AND att.content_info->>'id' = 'photo'
	LEFT JOIN specialities AS spt ON spt.id = t.speciality_id	
	ORDER BY t.name
	;
	
ALTER VIEW public.specialists_list OWNER TO ;
