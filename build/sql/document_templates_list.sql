-- VIEW: public.document_templates_list

--DROP VIEW public.document_templates_list;

CREATE OR REPLACE VIEW public.document_templates_list AS
	SELECT
		t.id
		,t.name
		,att.content_info || jsonb_build_object('dataBase64',encode(att.content_preview, 'base64')) AS file_preview
		,t.need_signing
	FROM public.document_templates AS t
	LEFT JOIN attachments AS att ON att.ref->>'dataType' = 'document_templates' AND (att.ref->'keys'->>'id')::int = t.id
	ORDER BY t.name
	;
	
ALTER VIEW public.document_templates_list OWNER TO ;
