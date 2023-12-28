-- Function: specialists_process()

-- DROP FUNCTION specialists_process();

CREATE OR REPLACE FUNCTION specialists_process()
  RETURNS trigger AS
$BODY$
DECLARE
	v_spec text;
BEGIN
	IF TG_WHEN='BEFORE' AND TG_OP='DELETE' THEN
		DELETE FROM specialist_statuses WHERE specialist_id = OLD.id;
		DELETE FROM specialist_documents WHERE specialist_id = OLD.id;
		DELETE FROM attachments WHERE (ref->'keys'->>'id')::int = OLD.id AND ref->>'dataType' = 'specialists';
		DELETE FROM entity_contacts WHERE entity_id = OLD.id AND entity_type = 'specialists';
		DELETE FROM entity_contacts WHERE entity_id = OLD.user_id AND entity_type = 'users';
		DELETE FROM ycl_transactions WHERE specialist_id = OLD.id;
			
		RETURN OLD;
		
	ELSIF TG_WHEN='AFTER' AND TG_OP='DELETE' THEN
		DELETE FROM users WHERE id = OLD.user_id;
		RETURN OLD;
		
	ELSIF TG_WHEN='BEFORE' AND TG_OP='INSERT' THEN	
		SELECT ycl_staff.data->>'specialization' INTO v_spec FROM ycl_staff WHERE ycl_staff.id = NEW.ycl_staff_id;
		SELECT
			spt.id,
			spt.agent_percent
		INTO
			NEW.speciality_id,
			NEW.agent_percent
		FROM specialities AS spt
		WHERE spt.name = trim(v_spec);
		
		IF NEW.speciality_id IS NULL THEN
			RAISE EXCEPTION 'Не найдена специальность: %', v_spec;
		END IF;
		
	ELSIF TG_WHEN='AFTER' AND TG_OP='INSERT' THEN	
		UPDATE ycl_transactions SET amount=amount
		WHERE
			staff_id = NEW.ycl_staff_id
			AND specialist_id IS NULL;
					
		RETURN NEW;
	END IF;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION specialists_process() OWNER TO ;

