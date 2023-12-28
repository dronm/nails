    CREATE TRIGGER ycl_transactions_trigger_before
    BEFORE INSERT OR UPDATE
    ON public.ycl_transactions
    FOR EACH ROW
    EXECUTE PROCEDURE public.ycl_transactions_process();

