-- Function: specialist_period_salary_details_docs_to_bank(in_ids integer[])

-- DROP FUNCTION specialist_period_salary_details_docs_to_bank(in_ids integer[]);

CREATE OR REPLACE FUNCTION specialist_period_salary_details_docs_to_bank(in_ids integer[])
  RETURNS SETOF bank_payments AS
$$
DECLARE
	v_firm_cnt int;
BEGIN
	-- проверим, что нет разных организаций
	SELECT
		count(DISTINCT studios.firm_id)
	INTO 
		v_firm_cnt
	FROM specialist_period_salary_details AS t
	LEFT JOIN studios ON studios.id = t.studio_id
	WHERE t.id =ANY(in_ids);
	
	IF v_firm_cnt > 1 THEN
		RAISE EXCEPTION 'В списке документов есть разные организации';
	END IF;
	
	RETURN QUERY
		INSERT INTO bank_payments
		(document_date, document_num, document_total, document_comment,
		payer_acc, payer_bank_acc, payer_bank_bik, payer_bank, payer_bank_place, payer, payer_inn,
		rec_acc, rec_bank_acc, rec_bank_bik, rec_bank, rec_bank_place, rec, rec_inn,
		specialist_id, specialist_period_salary_detail_id
		)
		SELECT
			now(),
			coalesce((SELECT max(document_num) FROM bank_payments), 1) + 1,
			det.total,
			replace(
				replace(const_specialist_pay_comment_template_val(), '[sum]', trim(replace(to_char(det.total, '999999.99'), '.', '-')))
			,'[fio]', person_init(sp.name)),
			
			firms.bank_acc,
			payer_bn.korshet,
			firms.bank_bik,
			payer_bn.name,
			payer_bn.gor,
			firms.name_full,
			firms.inn,
			
			sp.bank_acc,
			rec_bn.korshet,
			sp.bank_bik,
			rec_bn.name,
			rec_bn.gor,
			sp.name,
			sp.inn,
			
			sp.id,
			det.id
			
		FROM specialist_period_salary_details AS det
		LEFT JOIN specialists AS sp ON sp.id = det.specialist_id
		LEFT JOIN studios ON studios.id = det.studio_id
		LEFT JOIN firms ON firms.id = studios.firm_id
		LEFT JOIN banks.banks AS payer_bn ON payer_bn.bik = firms.bank_bik
		LEFT JOIN banks.banks AS rec_bn ON rec_bn.bik = sp.bank_bik
		WHERE det.id =ANY(in_ids)
		RETURNING bank_payments.*
		;		
	
END;
$$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION specialist_period_salary_details_docs_to_bank(in_ids integer[]) OWNER TO ;
