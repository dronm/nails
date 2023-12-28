-- Function: document_templates_salary(in_salary_detail_id int)

-- DROP FUNCTION document_templates_salary(v int);

CREATE OR REPLACE FUNCTION document_templates_salary(in_salary_detail_id int)
  RETURNS jsonb AS
$$
	WITH
	contr_date AS
		(SELECT
			sp_st.date_time::date AS d
		FROM specialist_statuses AS sp_st
		WHERE sp_st.specialist_id = in_salary_detail_id
		AND sp_st.status_type = 'contract_signing'
		)	
	SELECT
		jsonb_build_object(
			'dateDay', to_char(sal_h.date_time::date, 'DD'),			
			'dateMonthStr', month_rus(sal_h.date_time::date),
			'dateYear', to_char(sal_h.date_time, 'YYYY'),
			
			'aktNum', to_char(sal_h.id, 'fm00000') || '-' || to_char(sal.line_num, 'fm00'),
			'aktDateDay', to_char(sal_h.date_time::date, 'DD'),			
			'aktDateMonthStr', month_rus(sal_h.date_time::date),
			'aktDateYear', to_char(sal_h.date_time::date, 'YYYY'),
			
			'dogNum', to_char(sp.id, 'fm000'),
			'dogDateDay', to_char(stat.date_time::date, 'DD'),			
			'dogDateMonthStr', month_rus(stat.date_time::date),
			'dogDateYear', to_char(stat.date_time::date, 'YYYY'),
			'flName', sp.name,
			'flNameShort', person_init(sp.name),
			'perFromDateDay', to_char(sal.period, 'DD'),
			'perFromDateMonthStr', month_rus(sal.period),
			'perFromDateYear', to_char(sal.period, 'YYYY'),
			'perToDateDay', to_char(last_month_day(sal.period), 'DD'),
			'perToDateMonthStr', month_rus(sal.period),
			'perToDateYear', to_char(sal.period, 'YYYY'),
			'cust',
				(SELECT
					count(sb.*)
				FROM (
					SELECT
						DISTINCT ycl.client||ycl.client_phone
					FROM ycl_transactions_list AS ycl
					WHERE ycl.specialist_id = sp.id
						AND ycl.date BETWEEN sal.period AND last_month_day(sal.period)
					GROUP BY ycl.client,ycl.client_phone
				) AS sb
				),
			
			'serv',
				(SELECT count(ycl.*)
				FROM ycl_transactions_doc_all_list AS ycl
				WHERE ycl.specialist_id = sp.id
					AND ycl.date BETWEEN sal.period AND last_month_day(sal.period)
				),
			
			'hour', sal.hours,
			'total', document_templates_format_num(sal.total),
			'workTotal', document_templates_format_num(sal.work_total),
			'dkTotal', document_templates_format_num(sal.debet - sal.kredit),
			'rent', document_templates_format_num(sal.rent_total),
			'agentTotal', document_templates_format_num(sal.work_total - sal.work_total_salary)
		)
	FROM specialist_period_salary_details AS sal
	LEFT JOIN specialists AS sp ON sp.id = sal.specialist_id
	LEFT JOIN specialist_statuses AS stat ON stat.specialist_id = sal.specialist_id AND stat.status_type = 'contract_signing'
	LEFT JOIN specialist_period_salaries AS sal_h ON sal_h.id = sal.specialist_period_salary_id
	WHERE sal.id = in_salary_detail_id
	;
$$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION document_templates_salary(in_salary_detail_id int) OWNER TO ;
