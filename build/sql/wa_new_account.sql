-- Function: wa_new_account(in_specialist_id int, in_pwd text)

-- DROP FUNCTION wa_new_account(in_specialist_id int, in_pwd text);

-- Все функции по одному шаблону: возвращают структуру


CREATE OR REPLACE FUNCTION wa_new_account(in_specialist_id int, in_pwd text)
  RETURNS TABLE(
	tel text,
	body text
  )  AS
$BODY$
	WITH
	tmpl AS (SELECT
			t.template AS v
		FROM notif_templates t
		WHERE t.notif_type = 'new_account'::notif_types AND t.notif_provider = 'wa'
		LIMIT 1
	)	
	SELECT
		'7'||sp.tel AS tel,
		coalesce(templates_text(
			ARRAY[
				ROW('name', sp.name)::template_value,
				ROW('pwd', in_pwd)::template_value,
				ROW('login', '7'||u.name)::template_value
			],
			(SELECT v FROM tmpl)
		), '') AS body
	FROM specialists_dialog AS sp
	LEFT JOIN users AS u ON u.id = (sp.users_ref->'keys'->>'id')::int
	WHERE sp.id = in_specialist_id
	;
$BODY$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION wa_new_account(in_specialist_id int, in_pwd text) OWNER TO ;

