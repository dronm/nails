--Refrerece type
CREATE OR REPLACE FUNCTION bank_payments_ref(bank_payments)
  RETURNS json AS
$$
	SELECT json_build_object(
		'keys',json_build_object(
			'id',$1.id    
			),	
		'descr', $1.document_num::text || ' ' ||to_char($1.document_date, 'DD/MM/YY'),
		'dataType','bank_payments'
	);
$$
  LANGUAGE sql VOLATILE COST 100;
ALTER FUNCTION bank_payments_ref(bank_payments) OWNER TO ;	

