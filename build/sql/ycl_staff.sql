-- VIEW: public.ycl_staff_list

--DROP VIEW public.ycl_staff_list;

CREATE OR REPLACE VIEW public.ycl_staff_list AS
	SELECT
		t.id
		,t.name
		,t.data->>'specialization' AS specialization
		,t.data->>'avatar' AS avatar
		,t.data->>'avatar_big' AS avatar_big
		,(t.data->>'rating')::numeric(15,2) AS rating
		,(t.data->>'votes_count')::int AS votes_count
	FROM public.ycl_staff AS t
	ORDER BY t.name
	;
	
ALTER VIEW public.ycl_staff_list OWNER TO ;
