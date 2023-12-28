-- Function: document_templates_application(in_specialist_reg_id int)

-- DROP FUNCTION document_templates_application(in_specialist_reg_id int);

CREATE OR REPLACE FUNCTION document_templates_application(in_specialist_reg_id int)
  RETURNS jsonb AS
$$
	SELECT
		jsonb_build_object(
			'flName', sp.name_full,
			'flNameShort', person_init(sp.name_full),
			'flBirthdate', to_char(sp.birthdate, 'DD/MM/YYYY'),
			'flAddressReg', sp.address_reg,
			'flTel', format_cel_phone(sp.tel),
			'flEmail', sp.email,
			'flPassSeries', sp.passport->>'series',
			'flPassNum', sp.passport->>'num',
			'flPassIssueDate', to_char((sp.passport->>'issue_date')::date, 'DD/MM/YYYY'),
			'flPassIssueBody', sp.passport->>'issue_body',
			'flPassIssueDepCode', sp.passport->>'dep_code',
			'flInn', sp.inn,
			'bankName', sp.banks_ref->>'descr',
			'bankBik', sp.banks_ref->'id'->>'bik',
			'bankAccNum', sp.bank_acc
		)
	FROM specialist_regs AS sp
	WHERE sp.id = in_specialist_reg_id;
$$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION document_templates_application(in_specialist_reg_id int) OWNER TO ;
