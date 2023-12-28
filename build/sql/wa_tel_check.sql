-- Function: wa_tel_check(in_tel text, in_key text)

-- DROP FUNCTION wa_tel_check(in_tel text, in_key text);

-- Все функции по одному шаблону: возвращают структуру


CREATE OR REPLACE FUNCTION wa_tel_check(in_tel text, in_key text)
  RETURNS TABLE(
	tel text,
	body text
  )  AS
$BODY$
	WITH
	tmpl AS (SELECT
			t.template AS v
		FROM notif_templates t
		WHERE t.notif_type = 'tel_check'::notif_types AND t.notif_provider = 'wa'
		LIMIT 1
	)	
	SELECT
		in_tel AS tel,
		coalesce(templates_text(
			ARRAY[
				ROW('key', in_key)::template_value
			],
			(SELECT v FROM tmpl)
		), '') AS body
	;
$BODY$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION wa_tel_check(in_tel text, in_key text) OWNER TO ;

