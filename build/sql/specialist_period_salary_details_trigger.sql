-- Trigger: specialist_period_salary_details_trigger_before on public.specialist_period_salary_details


-- DROP TRIGGER specialist_period_salary_details_trigger_before ON public.specialist_period_salary_details;

CREATE TRIGGER specialist_period_salary_details_trigger_before
  BEFORE INSERT OR UPDATE OR DELETE
  ON public.specialist_period_salary_details
  FOR EACH ROW
  EXECUTE PROCEDURE public.specialist_period_salary_details_process();


/*
-- Trigger: specialist_period_salary_details_trigger_after on public.specialist_period_salary_details

 --DROP TRIGGER specialist_period_salary_details_trigger_after ON public.specialist_period_salary_details;

CREATE TRIGGER specialist_period_salary_details_trigger_after
  AFTER DELETE OR INSERT OR UPDATE
  ON public.specialist_period_salary_details
  FOR EACH ROW
  EXECUTE PROCEDURE public.specialist_period_salary_details_process();
*/
