-- Function: document_templates_format_num(in_num numeric(15,2))

-- DROP FUNCTION document_templates_format_num(in_num numeric(15,2));

CREATE OR REPLACE FUNCTION document_templates_format_num(in_num numeric(15,2))
  RETURNS text AS
$$
	SELECT
		CASE
			WHEN in_num = 0 THEN '0,00'
			ELSE
				replace(
					replace( 
						trim( to_char(in_num, '999,999D99') ) 
					, ',', ' ') --thousand
				,'.',',') --decimal
		END
	;
$$
  LANGUAGE sql IMMUTABLE
  COST 100;
ALTER FUNCTION document_templates_format_num(in_num numeric(15,2)) OWNER TO ;
