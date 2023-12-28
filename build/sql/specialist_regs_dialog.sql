-- VIEW: public.specialist_regs_dialog

--DROP VIEW public.specialist_regs_dialog;

CREATE OR REPLACE VIEW public.specialist_regs_dialog AS
	SELECT
		t.id
		,t.user_operation_id
		,t.inn
		,t.banks_ref
		,t.bank_acc
		,t.name_full
		,t.tel
		,t.email
		,t.address_reg
		,birthdate
		,t.name
		,t.passport
		,t.date_time
		,t.inn_checked
		,t.inn_fns_ok
		,(SELECT
			st.confirmed
		FROM confirmation_status AS st
		WHERE (st.ref->'keys'->>'id')::int = t.id AND st.ref->>'dataType'='specialist_regs' AND st.field='tel'
		LIMIT 1
		) AS tel_confirmed
		,tel_sent
		
		,(SELECT
			st.confirmed
		FROM confirmation_status AS st
		WHERE (st.ref->'keys'->>'id')::int = t.id AND st.ref->>'dataType'='specialist_regs' AND st.field='email'
		LIMIT 1
		) AS email_confirmed
		,email_sent
		
		,(SELECT
			jsonb_agg(
				att.content_info || jsonb_build_object('dataBase64',encode(att.content_preview, 'base64'))
			)
		FROM attachments AS att
		WHERE att.ref->>'dataType' = 'specialist_regs' AND (att.ref->'keys'->>'id')::int = t.id
		) AS passport_preview
		,passport_uploaded
		
		,(SELECT
			CASE
				WHEN op.status = 'end' AND coalesce(op.error_text,'')='' THEN TRUE
				WHEN op.status = 'end' AND coalesce(op.error_text,'')<>'' THEN FALSE
				ELSE NULL
			END
		FROM user_operations AS op
		WHERE op.operation_id = t.user_operation_id) AS operation_result
		
	FROM public.specialist_regs AS t
	;
	
ALTER VIEW public.specialist_regs_dialog OWNER TO ;

