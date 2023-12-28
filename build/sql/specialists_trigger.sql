/*
	DROP TRIGGER specialists_trigger_before ON public.specialists;
	
	CREATE TRIGGER specialists_trigger_before
	BEFORE DELETE OR INSERT
	ON public.specialists
	FOR EACH ROW
	EXECUTE PROCEDURE public.specialists_process();
*/    
	 DROP TRIGGER specialists_trigger_after ON public.specialists;
	CREATE TRIGGER specialists_trigger_after
	AFTER DELETE OR INSERT
	ON public.specialists
	FOR EACH ROW
	EXECUTE PROCEDURE public.specialists_process();

