-- Function: email_docs_for_sign(in_specialist_id int)

 DROP FUNCTION email_docs_for_sign(in_specialist_id int);

-- Все функции по одному шаблону: возвращают структуру

/*
 * нет такой функции, зачем отпралять неподписанные документы мастеру? только с подписью!
 */

/*
CREATE OR REPLACE FUNCTION email_docs_for_sign(in_specialist_id int)
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
		WHERE t.notif_type = 'docs_for_sign'::notif_types AND t.notif_provider = 'email'
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
				ROW('name', sp.name)::template_value
			],
			(SELECT v FROM tmpl)
		), '') AS body,
		coalesce((SELECT s FROM tmpl),'') AS subject
		
	FROM specialists_dialog AS sp
	LEFT JOIN entity_contacts AS e_ct ON e_ct.entity_type = 'specialists' AND e_ct.entity_id = sp.id
	LEFT JOIN contacts AS ct ON ct.id = e_ct.contact_id
	WHERE sp.id = in_specialist_id
	;
$BODY$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION email_docs_for_sign(in_specialist_id int) OWNER TO ;
*/
