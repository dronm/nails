-- View: users_profile

-- DROP VIEW users_profile;

CREATE OR REPLACE VIEW users_profile AS 
	SELECT
	 	u.id,
	 	u.name,
	 	u.name_full,
	 	u.locale_id,
	 	time_zone_locales_ref(tzl) AS time_zone_locales_ref
	 	
		,att.content_info || jsonb_build_object('dataBase64',encode(att.content_data, 'base64')) AS photo
				 	
 	FROM users AS u
 	LEFT JOIN time_zone_locales AS tzl ON tzl.id=u.time_zone_locale_id
 	LEFT JOIN attachments AS att ON att.ref->>'dataType' = 'users' AND (att.ref->'keys'->>'id')::int = u.id AND att.content_info->>'id' = 'photo'
	;
ALTER TABLE users_profile OWNER TO ;


