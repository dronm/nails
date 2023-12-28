    CREATE TRIGGER specialist_documents_trigger_after
    AFTER DELETE
    ON public.specialist_documents
    FOR EACH ROW
    EXECUTE PROCEDURE public.specialist_documents_process();

