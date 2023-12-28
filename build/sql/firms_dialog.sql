-- VIEW: public.firms_dialog

--DROP VIEW public.firms_dialog;

CREATE OR REPLACE VIEW public.firms_dialog AS
	SELECT
		t.id
		,t.name
		,t.inn
		,t.name_full
		,t.legal_address
		,t.post_address
		,t.kpp
		,t.ogrn
		,t.okpo
		,t.okved
		,t.bank_acc
		,banks.banks_ref(bnk) AS banks_ref
	FROM public.firms AS t
	LEFT JOIN banks.banks AS bnk ON bnk.bik = t.bank_bik
	
	;
	
ALTER VIEW public.firms_dialog OWNER TO ;
