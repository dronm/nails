-- Function: email_new_account(in_specialist_id int, in_pwd text)

-- DROP FUNCTION email_new_account(in_specialist_id int, in_pwd text);

-- Все функции по одному шаблону: возвращают структуру


CREATE OR REPLACE FUNCTION email_new_account(in_specialist_id int, in_pwd text)
  RETURNS TABLE(
	from_addr text,
	from_name text,
	reply_name text,
	sender_addr text,
	to_addr text,
	to_name text,		
	body text,	
	subject text
  )  AS
$BODY$
	WITH
	email AS (SELECT const_email_val() AS v),
	tmpl AS (SELECT
			t.template AS v,
			(SELECT
				fields.f->'fields'->>'descr'
			FROM (
				SELECT json_array_elements(t.provider_values->'rows') AS f
			) AS fields	
			WHERE fields.f->'fields'->>'id'='subject'
			LIMIT 1
			) AS s
			
		FROM notif_templates t
		WHERE t.notif_type = 'new_account'::notif_types AND t.notif_provider = 'email'
		LIMIT 1
	)	
	SELECT
		(SELECT v->>'from_addr' FROM email) AS from_addr,
		(SELECT v->>'from_name' FROM email) AS from_name,
		(SELECT v->>'reply_name' FROM email) AS reply_name,
		(SELECT v->>'sender_addr' FROM email) AS sender_addr,
		ct.email AS to_addr,
		sp.name AS to_name,
		coalesce(templates_text(
			ARRAY[
				ROW('pwd', in_pwd)::template_value,
				ROW('name', sp.name)::template_value,
				ROW('login', '7'||u.name)::template_value
			],
			(SELECT v FROM tmpl)
		), '') AS body,
		coalesce((SELECT s FROM tmpl),'') AS subject
	FROM specialists AS sp
	LEFT JOIN users AS u ON u.id = sp.user_id
	LEFT JOIN entity_contacts AS e_ct ON e_ct.entity_id = sp.id AND e_ct.entity_type = 'specialists'
	LEFT JOIN contacts AS ct ON ct.id = e_ct.contact_id
	WHERE sp.id = in_specialist_id
	;
$BODY$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION email_new_account(in_specialist_id int, in_pwd text) OWNER TO ;

