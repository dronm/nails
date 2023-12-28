-- Function: document_templates_spec_reg_exec_query(in_specialist_id int, in_query text)

-- DROP FUNCTION document_templates_spec_reg_exec_query(in_specialist_id int, in_query text);

--Query MUST be a function with one arg - in_specialist_id and MUST return jsonb
CREATE OR REPLACE FUNCTION document_templates_spec_reg_exec_query(in_specialist_id int, in_query text)
  RETURNS jsonb AS
$$
DECLARE
	v_res jsonb;
BEGIN
	EXECUTE replace(in_query, '%d'::text, in_specialist_id::text) INTO v_res;
	RETURN v_res;
END;
$$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION document_templates_spec_reg_exec_query(in_specialist_id int, in_query text) OWNER TO ;
