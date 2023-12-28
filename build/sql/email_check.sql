-- Function: email_check(in_to_addr text, in_name_full text, in_key text)

-- DROP FUNCTION email_check(in_to_addr text, in_name_full text, in_key text);

-- Все функции по одному шаблону: возвращают структуру


CREATE OR REPLACE FUNCTION email_check(in_to_addr text, in_name_full text, in_key text)
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
		WHERE t.notif_type = 'email_check'::notif_types AND t.notif_provider = 'email'
		LIMIT 1
	)	
	SELECT
		(SELECT v->>'from_addr' FROM email) AS from_addr,
		(SELECT v->>'from_name' FROM email) AS from_name,
		(SELECT v->>'reply_name' FROM email) AS reply_name,
		(SELECT v->>'sender_addr' FROM email) AS sender_addr,
		in_to_addr AS to_addr,
		in_name_full AS to_name,
		coalesce(templates_text(
			ARRAY[
				ROW('key', in_key)::template_value,
				ROW('name', in_name_full)::template_value
			],
			(SELECT v FROM tmpl)
		), '') AS body,
		coalesce((SELECT s FROM tmpl),'') AS subject
	;
$BODY$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION email_check(in_to_addr text, in_name_full text, in_key text) OWNER TO ;

