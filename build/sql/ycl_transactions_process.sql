-- Function: ycl_transactions_process()

-- DROP FUNCTION ycl_transactions_process();

CREATE OR REPLACE FUNCTION ycl_transactions_process()
  RETURNS trigger AS
$BODY$
BEGIN
	IF TG_WHEN='BEFORE' AND (TG_OP='INSERT' OR TG_OP='UPDATE') THEN
		BEGIN
			NEW.date = (NEW.data->>'date')::timestamp;
			NEW.amount = (NEW.data->>'amount')::numeric(15,2);
			NEW.seance_length = coalesce(NEW.data->'record'->>'seance_length', '0')::int;
			NEW.document_id = (NEW.data->>'document_id')::int;
			NEW.record_id = (NEW.data->>'record_id')::int;
			NEW.staff_id = (NEW.data->'record'->>'staff_id')::int;
		EXCEPTION
			WHEN OTHERS THEN
				NULL;
		END;
		
	
		SELECT
			sp.id
		INTO
			NEW.specialist_id
		FROM specialists AS sp
		WHERE sp.ycl_staff_id = NEW.staff_id
		;
		
		RETURN NEW;
		
	END IF;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION ycl_transactions_process() OWNER TO ;

