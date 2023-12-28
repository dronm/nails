-- Function: wa_new_specialist()

-- DROP FUNCTION wa_new_specialist();

-- Все функции по одному шаблону: возвращают структуру

--Сообщение для всех админов с телефонами

CREATE OR REPLACE FUNCTION wa_new_specialist()
  RETURNS TABLE(
	tel text,
	body text
  )  AS
$BODY$
	WITH
	tmpl AS (SELECT
			t.template AS v
		FROM notif_templates t
		WHERE t.notif_type = 'new_specialist'::notif_types AND t.notif_provider = 'wa'
		LIMIT 1
	)	
	SELECT
		'7'||ct.tel AS tel,
		coalesce(templates_text(
			ARRAY[
				ROW('name', u.name_full)::template_value
			],
			(SELECT v FROM tmpl)
		), '') AS body
	FROM users AS u
	LEFT JOIN entity_contacts AS e_ct ON e_ct.entity_type = 'users' AND e_ct.entity_id = u.id
	LEFT JOIN contacts AS ct ON ct.id = e_ct.contact_id
	WHERE u.role_id = 'admin' AND coalesce(ct.tel, '') <>''
	;
$BODY$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION wa_new_specialist() OWNER TO ;

