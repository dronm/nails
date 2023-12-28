-- Function: specialist_period_salaries_process()

-- DROP FUNCTION specialist_period_salaries_process();

CREATE OR REPLACE FUNCTION specialist_period_salaries_process()
  RETURNS trigger AS
$BODY$
BEGIN
	IF TG_WHEN='BEFORE' AND TG_OP='DELETE' THEN
		DELETE FROM specialist_period_salary_details WHERE specialist_period_salary_id = OLD.id;		
			
		RETURN OLD;
		
	END IF;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION specialist_period_salaries_process() OWNER TO ;

