    CREATE TRIGGER document_templates_trigger_before
    BEFORE DELETE
    ON public.document_templates
    FOR EACH ROW
    EXECUTE PROCEDURE public.document_templates_process();

