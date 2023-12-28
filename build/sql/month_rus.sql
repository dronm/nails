-- FUNCTION: public.month_rus(date)

-- DROP FUNCTION IF EXISTS public.month_rus(date);

CREATE OR REPLACE FUNCTION public.month_rus(date)
    RETURNS character varying
    LANGUAGE sql
    COST 100
    VOLATILE
AS $BODY$
	SELECT unnest(ARRAY['Января', 'Февраля', 'Марта', 'Апреля', 'Мая',
		'Июня', 'Июля', 'Августа', 'Сентября', 'Октября',
		'Ноября', 'Декабря'])
	LIMIT 1 OFFSET date_part('month', $1) - 1;
$BODY$;

ALTER FUNCTION public.month_rus(date)
    OWNER TO ;

