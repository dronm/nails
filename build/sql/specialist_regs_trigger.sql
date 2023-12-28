    CREATE TRIGGER specialist_regs_trigger_before
    BEFORE DELETE
    ON public.specialist_regs
    FOR EACH ROW
    EXECUTE PROCEDURE public.specialist_regs_process();

