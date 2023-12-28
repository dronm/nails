-- VIEW: public.users_list

--DROP VIEW public.users_list;

CREATE OR REPLACE VIEW public.users_list AS
	SELECT
		t.id
		,t.name
		,t.name_full
		,t.role_id
		,att.content_info || jsonb_build_object('dataBase64',encode(att.content_preview, 'base64')) AS photo
		
		,ct.email
		,ct.tel
		
	FROM public.users AS t
 	LEFT JOIN attachments AS att ON att.ref->>'dataType' = 'users' AND (att.ref->'keys'->>'id')::int = t.id AND att.content_info->>'id' = 'photo'
 	LEFT JOIN entity_contacts AS e_ct ON e_ct.entity_type = 'users' AND e_ct.entity_id = t.id
 	LEFT JOIN contacts AS ct ON ct.id = e_ct.contact_id
	
	ORDER BY t.name ASC
	;
	
ALTER VIEW public.users_list OWNER TO ;

