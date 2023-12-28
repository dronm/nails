-- Function: specialist_documents_process()

-- DROP FUNCTION specialist_documents_process();

CREATE OR REPLACE FUNCTION specialist_documents_process()
  RETURNS trigger AS
$BODY$
BEGIN
	IF TG_WHEN='AFTER' AND TG_OP='DELETE' THEN
		DELETE FROM attachments WHERE id = OLD.template_att_id;
		DELETE FROM attachments WHERE id = OLD.document_att_id;
			
		RETURN OLD;
	END IF;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION specialist_documents_process() OWNER TO ;

