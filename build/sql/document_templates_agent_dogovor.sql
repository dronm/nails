-- Function: document_templates_agent_dogovor(in_specialist_id int)

-- DROP FUNCTION document_templates_agent_dogovor(in_specialist_id int);

CREATE OR REPLACE FUNCTION document_templates_agent_dogovor(in_specialist_id int)
  RETURNS jsonb AS
$$
	WITH
	contr_date AS
		(SELECT sp_st.date_time::date AS d
			FROM specialist_statuses AS sp_st
			WHERE sp_st.specialist_id = in_specialist_id
			AND sp_st.status_type = 'contract_signing'
		)	
	SELECT
		jsonb_build_object(
			'dogovorNum', to_char(sp.id, 'fm000'),
			'dateDay', (SELECT to_char(d, 'DD') FROM contr_date),			
			'dateMonthStr', (SELECT month_rus(d) FROM contr_date),
			'dateYear', (SELECT to_char(d, 'YYYY') FROM contr_date),
			'flName', sp.name,
			'flNameShort', person_init(sp.name),
			'flBirthdate', to_char(sp.birthdate, 'DD/MM/YYYY'),
			'flAddressReg', sp.address_reg,
			'flTel', (SELECT
					format_cel_phone(ct.tel)
				FROM entity_contacts AS e_ct				
				LEFT JOIN contacts AS ct ON ct.id = e_ct.contact_id
				WHERE e_ct.entity_type = 'specialists' AND e_ct.entity_id = in_specialist_id
				), 
			'flPassSeries', sp.passport->>'series',
			'flPassNum', sp.passport->>'num',
			'flPassIssueDate', to_char((sp.passport->>'issue_date')::date, 'DD/MM/YYYY'),
			'flPassIssueBody', sp.passport->>'issue_body',
			'flPassDepCode', sp.passport->>'dep_code',
			'flInn', sp.inn,
			'bankBik', sp.bank_bik,
			'bankAccNum', sp.bank_acc
			-- Это делаем при вызове функции когда надо
			--'flSign', '   '
		)
	FROM specialists AS sp
	WHERE sp.id = in_specialist_id;
$$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION document_templates_agent_dogovor(in_specialist_id int) OWNER TO ;
