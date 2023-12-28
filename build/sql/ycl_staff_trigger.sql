    CREATE TRIGGER ycl_staff_trigger_before
    BEFORE INSERT OR DELETE
    ON public.ycl_staff
    FOR EACH ROW
    EXECUTE PROCEDURE public.ycl_staff_process();

