-- VIEW: public.document_templates_dialog

--DROP VIEW public.document_templates_dialog;

CREATE OR REPLACE VIEW public.document_templates_dialog AS
	SELECT
		t.id
		,t.name
		,t.fields
		,t.sql_query
		,att.content_info || jsonb_build_object('dataBase64',encode(att.content_preview, 'base64')) AS file_preview
		,t.need_signing
		,t.sign_image_name
	FROM public.document_templates AS t
	LEFT JOIN attachments AS att ON att.ref->>'dataType' = 'document_templates' AND (att.ref->'keys'->>'id')::int = t.id
	
	;
	
ALTER VIEW public.document_templates_dialog OWNER TO ;
