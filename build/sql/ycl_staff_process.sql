-- Function: ycl_staff_process()

-- DROP FUNCTION ycl_staff_process();

CREATE OR REPLACE FUNCTION ycl_staff_process()
  RETURNS trigger AS
$BODY$
BEGIN
	IF TG_WHEN='BEFORE' AND TG_OP='INSERT' OR TG_OP='UPDATE' THEN
		IF NEW.data IS NOT NULL AND  coalesce(NEW.data->>'id', '')<> '' THEN
			NEW.id = (NEW.data->>'id')::int;
		END IF;
		
		IF NEW.data IS NOT NULL AND  coalesce(NEW.data->>'name', '')<> '' THEN
			NEW.name = NEW.data->>'name';
			
		ELSIF TG_OP='INSERT' THEN
			NEW.name = '<>';
		END IF;
		
		RETURN NEW;
		
	END IF;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION ycl_staff_process() OWNER TO ;

