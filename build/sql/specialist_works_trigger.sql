    CREATE TRIGGER specialist_works_trigger_before
    BEFORE DELETE
    ON public.specialist_works
    FOR EACH ROW
    EXECUTE PROCEDURE public.specialist_works_process();

