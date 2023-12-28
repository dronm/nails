-- VIEW: public.specialists_dialog

--DROP VIEW public.specialists_dialog;

CREATE OR REPLACE VIEW public.specialists_dialog AS
	SELECT
		t.id
		,t.name
		,t.inn
		,banks.banks_ref(bnk) AS banks_ref
		,t.bank_acc
		,t.studio_id
		,studios_ref(studios_ref_t) AS studios_ref
		
		,t.birthdate
		,t.address_reg
		,t.passport
		
		,(SELECT
			st.status_type
		FROM specialist_statuses AS st
		WHERE st.specialist_id = t.id
		ORDER BY st.date_time DESC
		LIMIT 1		
		) AS last_status_type
		
		,users_ref((SELECT u FROM users AS u WHERE u.id = t.user_id)) AS users_ref
		
		,t.equipments
		
		,(SELECT
			jsonb_agg(
				att.content_info || jsonb_build_object('dataBase64',encode(att.content_preview, 'base64'))
			)
		FROM attachments AS att
		WHERE att.ref->>'dataType' = 'specialists' AND (att.ref->'keys'->>'id')::int = t.id AND att.content_info->>'id' = 'passport'
		) AS passport_preview
		
		,ct.email
		,ct.tel
		,t.agent_percent
		,specialities_ref(spt) AS specialities_ref
		
	FROM public.specialists AS t
	LEFT JOIN studios AS studios_ref_t ON studios_ref_t.id = t.studio_id
	LEFT JOIN banks.banks AS bnk ON bnk.bik = t.bank_bik
	LEFT JOIN entity_contacts AS e_ct ON e_ct.entity_type = 'specialists' AND e_ct.entity_id = t.id
	LEFT JOIN contacts AS ct ON ct.id = e_ct.contact_id
	LEFT JOIN specialities AS spt ON spt.id = t.speciality_id	
	;
	
ALTER VIEW public.specialists_dialog OWNER TO ;
