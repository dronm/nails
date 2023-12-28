-- Function: document_templates_process()

-- DROP FUNCTION document_templates_process();

CREATE OR REPLACE FUNCTION document_templates_process()
  RETURNS trigger AS
$BODY$
BEGIN
	IF TG_WHEN='BEFORE' AND TG_OP='DELETE' THEN
		DELETE FROM attachments WHERE ref->>'dataType' = 'document_templates' AND (ref->'keys'->>'id')::int = OLD.id;
			
		RETURN OLD;
		
	END IF;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION document_templates_process() OWNER TO ;

