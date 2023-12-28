-- Function: specialist_works_process()

-- DROP FUNCTION specialist_works_process();

CREATE OR REPLACE FUNCTION specialist_works_process()
  RETURNS trigger AS
$BODY$
BEGIN
	IF TG_WHEN='BEFORE' AND TG_OP='DELETE' THEN
		DELETE FROM attachments WHERE (ref->'keys'->>'id')::int = OLD.id AND ref->>'dataType' = 'specialist_works';
			
		RETURN OLD;
		
	END IF;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION specialist_works_process() OWNER TO ;

