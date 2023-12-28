-- FUNCTION: public.month_period_rus(in_date date)

-- DROP FUNCTION IF EXISTS public.month_period_rus(in_date date);

CREATE OR REPLACE FUNCTION public.month_period_rus(in_date date)
    RETURNS text
    LANGUAGE sql
    COST 100
    IMMUTABLE
AS $BODY$
	SELECT
		(SELECT unnest(ARRAY['Январь', 'Февраль', 'Март', 'Апрель', 'Май',
			'Июнь', 'Июль', 'Август', 'Сентябрь', 'Октябрь',
			'Ноябрь', 'Декабрь'])
		LIMIT 1 OFFSET date_part('month', $1) - 1		
		) || ' ' ||date_part('year', in_date)::text
	;
$BODY$;

ALTER FUNCTION public.month_period_rus(in_date date) OWNER TO ;

