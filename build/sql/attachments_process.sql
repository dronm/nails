-- Function: attachments_process()

-- DROP FUNCTION attachments_process();

CREATE OR REPLACE FUNCTION attachments_process()
  RETURNS trigger AS
$BODY$
BEGIN
	IF TG_WHEN='AFTER' AND TG_OP='DELETE' THEN
		--delete files
		PERFORM pg_notify('Attachment.clear_cache', json_build_object(
			'ref', OLD.ref::json,
			'content_id', OLD.content_info->>'id'
		)::text);
		
		--EXECUTE format('COPY (SELECT 1) TO PROGRAM ''rm -f %s/%s''', '/home/andrey/www/nails/CACHE', md5(format('att_%s%s_%s', OLD.ref->>'dataType', OLD.ref->'keys'->>'id', OLD.content_info->>'id')));		
		--EXECUTE format('COPY (SELECT 1) TO PROGRAM ''rm -f %s/%s''', '/home/andrey/www/nails/CACHE', md5(format('prev_%s%s_%s', OLD.ref->>'dataType', OLD.ref->'keys'->>'id', OLD.content_info->>'id')));
		
		RETURN OLD;
		
	ELSIF TG_WHEN='BEFORE' AND (TG_OP='INSERT' OR TG_OP='UPDATE' ) THEN
		IF TG_OP='INSERT' OR NEW.content_data IS NOT NULL THEN
			NEW.content_info = coalesce(NEW.content_info, OLD.content_info) ||
					jsonb_build_object('size', length(NEW.content_data));
		END IF;
		
		RETURN NEW;
	END IF;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION attachments_process() OWNER TO ;

