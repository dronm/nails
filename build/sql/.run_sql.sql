-- Function: specialist_period_salary_details_process()

-- DROP FUNCTION specialist_period_salary_details_process();

CREATE OR REPLACE FUNCTION specialist_period_salary_details_process()
  RETURNS trigger AS
$BODY$
BEGIN
	IF TG_WHEN='BEFORE' AND (TG_OP='INSERT' OR TG_OP='UPDATE') THEN

		IF TG_OP = 'INSERT' THEN
			SELECT
				coalesce(MAX(t.line_num),0)+1
			INTO NEW.line_num
			FROM specialist_period_salary_details AS t
			WHERE t.specialist_period_salary_id = NEW.specialist_period_salary_id;
			
			--default values for studio/period
			SELECT
				sal.period,
				sal.studio_id
			INTO
				NEW.period,
				NEW.studio_id
			FROM specialist_period_salaries AS sal
			WHERE sal.id = NEW.specialist_period_salary_id;
		END IF;
		
		RETURN NEW;
	
	ELSIF TG_WHEN='AFTER' AND (TG_OP='UPDATE' OR TG_OP='INSERT') THEN
		
		--totals
		IF (TG_OP='INSERT' AND coalesce(NEW.total,0)<>0) OR ( TG_OP='UPDATE' AND (coalesce(NEW.total,0)<>coalesce(OLD.total,0)) )
		THEN
			UPDATE specialist_period_salaries SET
				total = sel.total,
				work_total = sel.work_total,
				work_total_salary = sel.work_total_salary,
				rent_total = sel.rent_total,
				hours = sel.hours,
				debet = sel.debet,
				kredit = sel.kredit				
			FROM (
				SELECT
					sum(coalesce(it.total)) AS total,
					sum(coalesce(it.work_total)) AS work_total,
					sum(coalesce(it.work_total_salary)) AS work_total_salary,
					sum(coalesce(it.rent_total)) AS rent_total,
					sum(coalesce(it.hours)) AS hours,
					sum(coalesce(it.debet)) AS debet,
					sum(coalesce(it.kredit)) AS kredit
				FROM specialist_period_salary_details AS it
				WHERE it.specialist_period_salary_id = NEW.specialist_period_salary_id
			) AS sel
			WHERE id = NEW.specialist_period_salary_id;			
		END IF;
		
		RETURN NEW;
	
	ELSIF (TG_WHEN='BEFORE' AND TG_OP='DELETE') THEN
		DELETE FROM specialist_receipts WHERE specialist_period_salary_detail_id = OLD.id;
		DELETE FROM bank_payments WHERE specialist_period_salary_detail_id = OLD.id;
		
		RETURN OLD;
		
	ELSIF (TG_WHEN='AFTER' AND TG_OP='DELETE') THEN
		
		UPDATE specialist_period_salaries SET
			total = total - OLD.total,
			work_total = work_total - OLD.work_total,
			work_total_salary = work_total_salary - OLD.work_total_salary,
			rent_total = rent_total - OLD.rent_total,
			hours = hours - OLD.hours,
			debet = debet - OLD.debet,
			kredit = kredit - OLD.kredit
		WHERE id = NEW.specialist_period_salary_id;
		
		RETURN OLD;
	END IF;
END;
$BODY$
LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION specialist_period_salary_details_process()
  OWNER TO nails;
