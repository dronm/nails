-- Function: wa_docs_for_sign(in_specialist_id int)

-- DROP FUNCTION wa_docs_for_sign(in_specialist_id int);

-- Все функции по одному шаблону: возвращают структуру


CREATE OR REPLACE FUNCTION wa_docs_for_sign(in_specialist_id int)
  RETURNS TABLE(
	tel text,
	body text
  )  AS
$BODY$
	WITH
	tmpl AS (SELECT
			t.template AS v
		FROM notif_templates t
		WHERE t.notif_type = 'docs_for_sign'::notif_types AND t.notif_provider = 'wa'
		LIMIT 1
	)	
	SELECT
		'7'||sp.tel AS tel,
		coalesce(templates_text(
			ARRAY[
				ROW('name', sp.name)::template_value
			],
			(SELECT v FROM tmpl)
		), '') AS body
	FROM specialists_dialog AS sp
	WHERE sp.id = in_specialist_id
	;
$BODY$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION wa_docs_for_sign(in_specialist_id int) OWNER TO ;

