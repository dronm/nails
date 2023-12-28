-- Function: specialist_period_salaries_fill(in_specialist_period_salary_id int)

-- DROP FUNCTION specialist_period_salaries_fill(in_specialist_period_salary_id int);

CREATE OR REPLACE FUNCTION specialist_period_salaries_fill(in_specialist_period_salary_id int)
  RETURNS void AS
$$
	DELETE FROM specialist_period_salary_details WHERE specialist_period_salary_id = in_specialist_period_salary_id;
	
	INSERT INTO specialist_period_salary_details
	(specialist_period_salary_id, specialist_id, studio_id, period,
	hours, agent_percent, work_total, work_total_salary, debet, kredit, rent_price, rent_total, total
	)
	WITH
	header AS (
		SELECT
			t.studio_id,
			t.period::date,
			(date_trunc('month', period) + '1 month'::interval - '1 day'::interval)::date period_end,
			std.hour_rent_price AS rent_price
			
		FROM specialist_period_salaries AS t
		LEFT JOIN studios AS std ON std.id = t.studio_id
		WHERE t.id = in_specialist_period_salary_id
	)
	SELECT
		in_specialist_period_salary_id,
		trans.specialist_id,
		(SELECT studio_id FROM header),
		(SELECT period FROM header),
		
		spt.agent_percent,
		
		sum(trans.len_hour) AS hours,
		sum(trans.amount) AS work_total,
		
		sum(trans.amount) - round(sum(trans.amount) * spt.agent_percent /100,2) AS work_total_salary,
		
		(SELECT
			sum(t.total)
		FROM specialist_salary_debets AS t
		WHERE t.specialist_id = trans.specialist_id AND (t.date_time::date BETWEEN (SELECT period FROM header) AND (SELECT period_end FROM header))
		) AS debet,
		
		(SELECT
			sum(t.total)
		FROM specialist_salary_kredits AS t
		WHERE t.specialist_id = trans.specialist_id AND (t.date_time::date BETWEEN (SELECT period FROM header) AND (SELECT period_end FROM header))
		) AS kredit,
		
		(SELECT rent_price FROM header) AS rent_price,
		(SELECT rent_price FROM header) * sum(trans.len_hour) AS rent_total,
		
		--work_total - 
		sum(trans.amount) - round(sum(trans.amount) * spt.agent_percent /100,2) - 
		(SELECT
			sum(t.total)
		FROM specialist_salary_debets AS t
		WHERE t.specialist_id = trans.specialist_id AND (t.date_time::date BETWEEN (SELECT period FROM header) AND (SELECT period_end FROM header))
		) - 		
		(SELECT
			sum(t.total)
		FROM specialist_salary_kredits AS t
		WHERE t.specialist_id = trans.specialist_id AND (t.date_time::date BETWEEN (SELECT period FROM header) AND (SELECT period_end FROM header))
		) - 
		((SELECT rent_price FROM header) * sum(trans.len_hour) )
		AS total
				
	FROM ycl_transactions_doc_all_list AS trans
	LEFT JOIN specialists AS sp ON sp.id = trans.specialist_id
	LEFT JOIN specialities AS spt ON spt.id = sp.speciality_id
	WHERE trans.date::date BETWEEN (SELECT period FROM header) AND (SELECT period_end FROM header)
		AND trans.specialist_id IS NOT NULL
	GROUP BY trans.specialist_id, spt.agent_percent	
	;
$$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION specialist_period_salaries_fill(in_specialist_period_salary_id int) OWNER TO ;
