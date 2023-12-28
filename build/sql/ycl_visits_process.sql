-- Function: ycl_visits_process()

-- DROP FUNCTION ycl_visits_process();

CREATE OR REPLACE FUNCTION ycl_visits_process()
  RETURNS trigger AS
$BODY$
BEGIN
	IF TG_WHEN='BEFORE' AND (TG_OP='INSERT' OR TG_OP='UPDATE') THEN
		SELECT
			(SELECT sp.id
			FROM specialists AS sp
			WHERE sp.ycl_staff_id = (s.rec->>'staff_id')::int)
		INTO
			NEW.specialist_id
		FROM
		(SELECT jsonb_array_elements(NEW.data->'records') AS rec LIMIT 1) AS s;
			
		RETURN NEW;
		
	END IF;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION ycl_visits_process() OWNER TO ;

