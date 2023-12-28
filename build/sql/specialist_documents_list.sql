-- View: public.specialist_documents_list

-- DROP VIEW public.specialist_documents_list;

CREATE OR REPLACE VIEW public.specialist_documents_list
 AS
	SELECT t.id,
		t.specialist_id,
		specialists_ref(specialists_ref_t.*) AS specialists_ref,
		att.content_info || jsonb_build_object('dataBase64', encode(att.content_preview, 'base64'::text)) AS document_att_ref,
		t.date_time,
		t.sign_date_time,
		t.sign_date_time IS NOT NULL AS signed,
		t.open_date_time,
		t.open_date_time IS NOT NULL AS opened,
		t.need_signing,
		t.name
	FROM specialist_documents t
	LEFT JOIN specialists specialists_ref_t ON specialists_ref_t.id = t.specialist_id
	LEFT JOIN attachments att ON att.id = t.document_att_id
	ORDER BY
		t.specialist_id,
		t.date_time DESC;

ALTER TABLE public.specialist_documents_list
    OWNER TO nails;


