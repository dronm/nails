-- VIEW: banks.banks_list

--DROP VIEW banks.banks_list;

CREATE OR REPLACE VIEW banks.banks_list AS
	SELECT
		t.bik
		,t.codegr
		,t.gr_descr
		,t.name
		,t.korshet
		,t.adres
		,t.gor
		,t.tgroup
	FROM banks.banks AS t
	
	ORDER BY codegr ASC,bik ASC
	;
	
ALTER VIEW banks.banks_list OWNER TO ;
