	
--	DROP TRIGGER attachments_trigger_before ON public.attachments;
/*
	CREATE TRIGGER attachments_trigger_before
	BEFORE INSERT OR UPDATE
	ON public.attachments
	FOR EACH ROW
	EXECUTE PROCEDURE public.attachments_process();
*/

--	DROP TRIGGER attachments_trigger_after ON public.attachments;

	CREATE TRIGGER attachments_trigger_after
	AFTER DELETE
	ON public.attachments
	FOR EACH ROW
	EXECUTE PROCEDURE public.attachments_process();

