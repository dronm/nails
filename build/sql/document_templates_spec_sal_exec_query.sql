-- Function: document_templates_spec_sal_exec_query(in_salary_detail_id int, in_query text)

-- DROP FUNCTION document_templates_spec_sal_exec_query(in_salary_detail_id int, in_query text);

--Query MUST be a function with one arg - in_salary_detail_id and MUST return jsonb
CREATE OR REPLACE FUNCTION document_templates_spec_sal_exec_query(in_salary_detail_id int, in_query text)
  RETURNS jsonb AS
$$
DECLARE
	v_res jsonb;
BEGIN
	EXECUTE replace(in_query, '%d'::text, in_salary_detail_id::text) INTO v_res;
	RETURN v_res;
END;
$$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION document_templates_spec_sal_exec_query(in_salary_detail_id int, in_query text) OWNER TO ;
