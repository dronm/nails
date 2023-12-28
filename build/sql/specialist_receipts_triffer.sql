	-- DROP TRIGGER specialist_receipts_trigger_after ON public.specialist_receipts;
	CREATE TRIGGER specialist_receipts_trigger_after
	AFTER DELETE
	ON public.specialist_receipts
	FOR EACH ROW
	EXECUTE PROCEDURE public.specialist_receipts_process();

