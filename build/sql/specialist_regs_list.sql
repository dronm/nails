-- VIEW: public.specialist_regs_list

DROP VIEW public.specialist_regs_list;

CREATE OR REPLACE VIEW public.specialist_regs_list AS
	SELECT
		t.id
		,t.inn
		,t.name_full
		
		,t.tel
		,t.tel_sent
		,(SELECT
			st.confirmed
		FROM confirmation_status AS st
		WHERE (st.ref->'keys'->>'id')::int = t.id AND st.ref->>'dataType'='specialist_regs' AND st.field='tel'
		LIMIT 1
		) AS tel_confirmed
		
		,t.email
		,t.email_sent
		,(SELECT
			st.confirmed
		FROM confirmation_status AS st
		WHERE (st.ref->'keys'->>'id')::int = t.id AND st.ref->>'dataType'='specialist_regs' AND st.field='email'
		LIMIT 1
		) AS email_confirmed
		
		,t.date_time
		,t.inn_fns_ok
		,t.inn_checked
		,t.birthdate
		,t.banks_ref
		,t.bank_acc
		
		,att.content_info || jsonb_build_object('dataBase64',encode(att.content_preview, 'base64')) AS passport_preview
		,(SELECT
			CASE
				WHEN op.status = 'end' AND coalesce(op.error_text,'')='' THEN TRUE
				WHEN op.status = 'end' AND coalesce(op.error_text,'')<>'' THEN FALSE
				ELSE NULL
			END
		FROM user_operations AS op
		WHERE op.operation_id = t.user_operation_id) AS operation_result

		
	FROM public.specialist_regs AS t
	LEFT JOIN attachments AS att ON att.ref->>'dataType' = 'specialist_regs' AND (att.ref->'keys'->>'id')::int = t.id
	ORDER BY operation_result, t.date_time DESC
	
	;
	
ALTER VIEW public.specialist_regs_list OWNER TO ;
