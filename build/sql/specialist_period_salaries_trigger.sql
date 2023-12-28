    CREATE TRIGGER specialist_period_salaries_trigger_before
    BEFORE DELETE
    ON public.specialist_period_salaries
    FOR EACH ROW
    EXECUTE PROCEDURE public.specialist_period_salaries_process();
