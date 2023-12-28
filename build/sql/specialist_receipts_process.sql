-- Function: specialist_receipts_process()

-- DROP FUNCTION specialist_receipts_process();

CREATE OR REPLACE FUNCTION specialist_receipts_process()
  RETURNS trigger AS
$BODY$
DECLARE
	v_spec text;
BEGIN
	IF TG_WHEN='AFTER' AND TG_OP='DELETE' THEN
		DELETE FROM attachments WHERE (ref->'keys'->>'id')::int = OLD.id AND ref->>'dataType' = 'specialist_receipts';
			
		RETURN OLD;		
	END IF;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION specialist_receipts_process() OWNER TO ;

