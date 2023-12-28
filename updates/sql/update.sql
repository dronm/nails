-- ******************* update 10/11/2023 15:13:06 ******************

	-- Adding new type
	CREATE TYPE locales AS ENUM (
	
		'ru'			
				
	);
	ALTER TYPE locales OWNER TO nails;
		
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_locales_val(locales,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='ru'::locales AND $2='ru'::locales THEN 'Русский'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_locales_val(locales,locales) OWNER TO nails;		
	
	-- Adding new type
	CREATE TYPE role_types AS ENUM (
	
		'admin'			
	,
		'specialist'			
	,
		'accountant'			
				
	);
	ALTER TYPE role_types OWNER TO nails;
		
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_role_types_val(role_types,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='admin'::role_types AND $2='ru'::locales THEN 'Администратор'
		WHEN $1='specialist'::role_types AND $2='ru'::locales THEN 'Мастер'
		WHEN $1='accountant'::role_types AND $2='ru'::locales THEN 'Бухгалтер'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_role_types_val(role_types,locales) OWNER TO nails;		
	
	-- Adding new type
	CREATE TYPE data_types AS ENUM (
	
		'users'			
				
	);
	ALTER TYPE data_types OWNER TO nails;
		
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_data_types_val(data_types,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='users'::data_types AND $2='ru'::locales THEN 'Пользователи'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_data_types_val(data_types,locales) OWNER TO nails;		
	
	-- Adding new type
	CREATE TYPE mail_types AS ENUM (
	
		'password_recover'			
				
	);
	ALTER TYPE mail_types OWNER TO nails;
		
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_mail_types_val(mail_types,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='password_recover'::mail_types AND $2='ru'::locales THEN 'Восстановление пароля'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_mail_types_val(mail_types,locales) OWNER TO nails;		
	
	-- Adding new type
	CREATE TYPE notif_types AS ENUM (
	
		'new_specialist'			
				
	);
	ALTER TYPE notif_types OWNER TO nails;
		
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_notif_types_val(notif_types,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='new_specialist'::notif_types AND $2='ru'::locales THEN 'Добавлен самозанятый'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_notif_types_val(notif_types,locales) OWNER TO nails;		
	

-- ******************* update 10/11/2023 15:13:54 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.time_zone_locales
	(id serial NOT NULL,descr  varchar(100) NOT NULL,name  varchar(50) NOT NULL,hour_dif  numeric(5,1) NOT NULL,CONSTRAINT time_zone_locales_pkey PRIMARY KEY (id)
	);
	
	ALTER TABLE public.time_zone_locales OWNER TO nails;
	
	

-- ******************* update 10/11/2023 15:15:00 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.users
	(id serial NOT NULL,name  varchar(100) NOT NULL,name_full text,role_id role_types NOT NULL,pwd  varchar(32),create_dt timestampTZ
			DEFAULT CURRENT_TIMESTAMP,banned bool
			DEFAULT FALSE,time_zone_locale_id int REFERENCES time_zone_locales(id),locale_id locales,CONSTRAINT users_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS users_name_index;
	CREATE UNIQUE INDEX users_name_index
	ON users(lower(name));

	DROP INDEX IF EXISTS users_role_id_idx;
	CREATE INDEX users_role_id_idx
	ON users(role_id);

	ALTER TABLE public.users OWNER TO nails;
	



-- ******************* update 11/11/2023 06:54:04 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.logins
	(id serial NOT NULL,date_time_in timestampTZ,date_time_out timestampTZ,ip  varchar(15),session_id  varchar(128),user_id int,pub_key  varchar(15),set_date_time timestampTZ,headers jsonb,user_agent jsonb,CONSTRAINT logins_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS logins_session_idx;
	CREATE INDEX logins_session_idx
	ON logins(session_id);

	DROP INDEX IF EXISTS logins_user_idx;
	CREATE INDEX logins_user_idx
	ON logins(user_id);

	DROP INDEX IF EXISTS logins_pub_key_idx;
	CREATE INDEX logins_pub_key_idx
	ON logins(pub_key);

	ALTER TABLE public.logins OWNER TO nails;
	
	

-- ******************* update 11/11/2023 06:54:41 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.login_device_bans
	(user_id int NOT NULL,hash  varchar(32) NOT NULL,create_dt timestampTZ
			DEFAULT CURRENT_TIMESTAMP,CONSTRAINT login_device_bans_pkey PRIMARY KEY (user_id,hash)
	);
	
	ALTER TABLE public.login_device_bans OWNER TO nails;
	
	

-- ******************* update 11/11/2023 07:19:30 ******************

	-- ********** constant value table  doc_per_page_count *************
	CREATE TABLE IF NOT EXISTS const_doc_per_page_count
	(name text  NOT NULL, descr text, val int,
		val_type text,ctrl_class text,ctrl_options json, view_class text,view_options json,
	CONSTRAINT const_doc_per_page_count_pkey PRIMARY KEY (name));
	ALTER TABLE const_doc_per_page_count OWNER TO nails;
	INSERT INTO const_doc_per_page_count (name,descr,val,val_type,ctrl_class,ctrl_options,view_class,view_options) VALUES (
		'Количество документов на странице'
		,'Количество документов на странице в журнале документов'
		,60
		,'Int'
		,NULL
		,NULL
		,NULL
		,NULL
	);
	
		--constant get value
		
	CREATE OR REPLACE FUNCTION const_doc_per_page_count_val()
	RETURNS int AS
	$BODY$
		
		SELECT val::int AS val FROM const_doc_per_page_count LIMIT 1;
		
	$BODY$
	LANGUAGE sql STABLE COST 100;
	ALTER FUNCTION const_doc_per_page_count_val() OWNER TO nails;
	
	--constant set value
	CREATE OR REPLACE FUNCTION const_doc_per_page_count_set_val(Int)
	RETURNS void AS
	$BODY$
		UPDATE const_doc_per_page_count SET val=$1;
	$BODY$
	LANGUAGE sql VOLATILE COST 100;
	ALTER FUNCTION const_doc_per_page_count_set_val(Int) OWNER TO nails;
	
	--edit view: all keys and descr
	CREATE OR REPLACE VIEW const_doc_per_page_count_view AS
	SELECT
		'doc_per_page_count'::text AS id
		,t.name
		,t.descr
	,
	t.val::text AS val
	
	,t.val_type::text AS val_type
	,t.ctrl_class::text
	,t.ctrl_options::json
	,t.view_class::text
	,t.view_options::json
	FROM const_doc_per_page_count AS t
	;
	ALTER VIEW const_doc_per_page_count_view OWNER TO nails;
	
	

-- ******************* update 11/11/2023 07:19:58 ******************

	-- ********** constant value table  grid_refresh_interval *************
	CREATE TABLE IF NOT EXISTS const_grid_refresh_interval
	(name text  NOT NULL, descr text, val int,
		val_type text,ctrl_class text,ctrl_options json, view_class text,view_options json,
	CONSTRAINT const_grid_refresh_interval_pkey PRIMARY KEY (name));
	ALTER TABLE const_grid_refresh_interval OWNER TO nails;
	INSERT INTO const_grid_refresh_interval (name,descr,val,val_type,ctrl_class,ctrl_options,view_class,view_options) VALUES (
		'Период обновления таблиц'
		,'Период обновления таблиц в секундах'
		,15
		,'Int'
		,NULL
		,NULL
		,NULL
		,NULL
	);
	
		--constant get value
		
	CREATE OR REPLACE FUNCTION const_grid_refresh_interval_val()
	RETURNS int AS
	$BODY$
		
		SELECT val::int AS val FROM const_grid_refresh_interval LIMIT 1;
		
	$BODY$
	LANGUAGE sql STABLE COST 100;
	ALTER FUNCTION const_grid_refresh_interval_val() OWNER TO nails;
	
	--constant set value
	CREATE OR REPLACE FUNCTION const_grid_refresh_interval_set_val(Int)
	RETURNS void AS
	$BODY$
		UPDATE const_grid_refresh_interval SET val=$1;
	$BODY$
	LANGUAGE sql VOLATILE COST 100;
	ALTER FUNCTION const_grid_refresh_interval_set_val(Int) OWNER TO nails;
	
	--edit view: all keys and descr
	CREATE OR REPLACE VIEW const_grid_refresh_interval_view AS
	SELECT
		'grid_refresh_interval'::text AS id
		,t.name
		,t.descr
	,
	t.val::text AS val
	
	,t.val_type::text AS val_type
	,t.ctrl_class::text
	,t.ctrl_options::json
	,t.view_class::text
	,t.view_options::json
	FROM const_grid_refresh_interval AS t
	;
	ALTER VIEW const_grid_refresh_interval_view OWNER TO nails;
	
	

-- ******************* update 11/11/2023 07:59:19 ******************
	
		ALTER TABLE public.users ADD COLUMN photo bytea;



-- ******************* update 11/11/2023 09:02:31 ******************

	
	CREATE OR REPLACE VIEW constants_list_view AS
	
	SELECT *
	FROM const_doc_per_page_count_view
	UNION ALL
	
	SELECT *
	FROM const_grid_refresh_interval_view
	ORDER BY name;
	ALTER VIEW constants_list_view OWNER TO nails;
	
	

-- ******************* update 13/11/2023 08:45:21 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.user_operations
	(user_id int NOT NULL,operation_id  varchar(36) NOT NULL,operation text,status text,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP,error_text text,comment_text text,date_time_end timestampTZ,end_wal_lsn text,CONSTRAINT user_operations_pkey PRIMARY KEY (user_id,operation_id)
	);
	
	ALTER TABLE public.user_operations OWNER TO nails;
	
	

-- ******************* update 13/11/2023 09:44:21 ******************
ALTER TABLE public.contacts ADD COLUMN tel_confirmed boolean DEFAULT false;



-- ******************* update 13/11/2023 09:50:34 ******************
ALTER TABLE IF EXISTS public.entity_contacts
    OWNER TO nails;
-- Index: entity_contacts_contact_idx

-- DROP INDEX IF EXISTS public.entity_contacts_contact_idx;

CREATE INDEX IF NOT EXISTS entity_contacts_contact_idx
    ON public.entity_contacts USING btree
    (entity_type ASC NULLS LAST, contact_id ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: entity_contacts_id_idx

-- DROP INDEX IF EXISTS public.entity_contacts_id_idx;

CREATE UNIQUE INDEX IF NOT EXISTS entity_contacts_id_idx
    ON public.entity_contacts USING btree
    (entity_type ASC NULLS LAST, entity_id ASC NULLS LAST, contact_id ASC NULLS LAST)
    TABLESPACE pg_default;


-- ******************* update 13/11/2023 11:40:22 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10000',
	'User_Controller',
	'get_list',
	'UserList',
	'Справочники',
	'Пользователи',
	FALSE
	);
	
	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10001',
	'LoginDevice_Controller',
	'get_list',
	'LoginDeviceList',
	'Справочники',
	'Устройства входа пользователей',
	FALSE
	);
	
	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10002',
	'Contact_Controller',
	'get_list',
	'ContactList',
	'Справочники',
	'Контакты',
	FALSE
	);
	
	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10003',
	'Post_Controller',
	'get_list',
	'PostList',
	'Справочники',
	'Должности',
	FALSE
	);
	
	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10004',
	'Bank_Controller',
	'get_list',
	'BankList',
	'Справочники',
	'Банки',
	FALSE
	);
	
	

-- ******************* update 13/11/2023 12:21:38 ******************

ALTER TABLE users DROP COLUMN photo CASCADE;


-- ******************* update 13/11/2023 15:02:19 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.firms
	(id serial NOT NULL,name text NOT NULL,inn  varchar(12),name_full text,legal_address text,post_address text,kpp  varchar(10),ogrn  varchar(15),okpo  varchar(20),okved text,CONSTRAINT firms_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS firms_name_idx;
	CREATE UNIQUE INDEX firms_name_idx
	ON firms(lower(name));

	DROP INDEX IF EXISTS firms_inn_idx;
	CREATE INDEX firms_inn_idx
	ON firms(inn);

	DROP INDEX IF EXISTS firms_ogrn_idx;
	CREATE UNIQUE INDEX firms_ogrn_idx
	ON firms(ogrn);

	ALTER TABLE public.firms OWNER TO nails;
	
	


-- ******************* update 13/11/2023 15:08:53 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10005',
	'Firm_Controller',
	'get_list',
	'FirmList',
	'Справочники',
	'Организации',
	FALSE
	);
	
	

-- ******************* update 13/11/2023 15:12:28 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.studios
	(id serial NOT NULL,name text NOT NULL,firm_id int NOT NULL REFERENCES firms(id),CONSTRAINT studios_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS firms_name_idx;
	CREATE UNIQUE INDEX firms_name_idx
	ON studios(lower(name));

	ALTER TABLE public.studios OWNER TO nails;
	



-- ******************* update 13/11/2023 15:34:12 ******************
	
		ALTER TABLE public.studios ADD COLUMN equipments jsonb;



-- ******************* update 13/11/2023 15:41:06 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialists
	(id serial NOT NULL,name text NOT NULL,inn  varchar(12) NOT NULL,bank_bik  varchar(9),bank_account  varchar(20),studio_id int NOT NULL REFERENCES studios(id),CONSTRAINT specialists_pkey PRIMARY KEY (id)
	);
	
	ALTER TABLE public.specialists OWNER TO nails;
	
	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10007',
	'Specialist_Controller',
	'get_list',
	'SpecialistList',
	'Справочники',
	'Специалисты',
	FALSE
	);
	



-- ******************* update 13/11/2023 15:49:51 ******************

	-- Adding new type
	CREATE TYPE specialist_statuses AS ENUM (
	
		'contract_signed'			
	,
		'contract_terminated'			
				
	);
	ALTER TYPE specialist_statuses OWNER TO nails;
		
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_specialist_statuses_val(specialist_statuses,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='contract_signed'::specialist_statuses AND $2='ru'::locales THEN 'Заключен контракт'
		WHEN $1='contract_terminated'::specialist_statuses AND $2='ru'::locales THEN 'Расторгнут контракт'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_specialist_statuses_val(specialist_statuses,locales) OWNER TO nails;		
	

-- ******************* update 13/11/2023 15:52:04 ******************

	-- Adding new type
	CREATE TYPE specialist_status_types AS ENUM (
	
		'contract_signed'			
	,
		'contract_terminated'			
				
	);
	ALTER TYPE specialist_status_types OWNER TO nails;
		
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_specialist_status_types_val(specialist_status_types,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='contract_signed'::specialist_status_types AND $2='ru'::locales THEN 'Заключен контракт'
		WHEN $1='contract_terminated'::specialist_status_types AND $2='ru'::locales THEN 'Расторгнут контракт'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_specialist_status_types_val(specialist_status_types,locales) OWNER TO nails;		
	

-- ******************* update 13/11/2023 15:56:15 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_statuses
	(id serial NOT NULL,specialist_id int NOT NULL REFERENCES specialists(id),status_type specialist_status_types NOT NULL,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,CONSTRAINT specialist_statuses_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_statuses_specialist;
	CREATE UNIQUE INDEX specialist_statuses_specialist
	ON specialist_statuses(specialist_id,date_time);

	ALTER TABLE public.specialist_statuses OWNER TO nails;
	
	


-- ******************* update 13/11/2023 16:05:08 ******************
	
		ALTER TABLE public.specialists ADD COLUMN birthdate date,ADD COLUMN address_reg text,ADD COLUMN passport jsonb;



-- ******************* update 17/11/2023 12:09:10 ******************

	-- Adding new type
	CREATE TYPE notif_providers AS ENUM (
	
		'email'			
	,
		'sms'			
	,
		'wa'			
	,
		'tm'			
	,
		'vb'			
				
	);
	ALTER TYPE notif_providers OWNER TO nails;
		
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_notif_providers_val(notif_providers,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='email'::notif_providers AND $2='ru'::locales THEN 'Электронная почта'
		WHEN $1='sms'::notif_providers AND $2='ru'::locales THEN 'СМС'
		WHEN $1='wa'::notif_providers AND $2='ru'::locales THEN 'WhatsUp'
		WHEN $1='tm'::notif_providers AND $2='ru'::locales THEN 'Telegram'
		WHEN $1='vb'::notif_providers AND $2='ru'::locales THEN 'Viber'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_notif_providers_val(notif_providers,locales) OWNER TO nails;		
	

-- ******************* update 17/11/2023 12:13:43 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.notif_templates
	(id serial NOT NULL,notif_provider notif_providers NOT NULL,notif_type notif_types NOT NULL,template text NOT NULL,comment_text text NOT NULL,fields json NOT NULL,provider_values json,CONSTRAINT notif_templates_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS notif_templates_type_idx;
	CREATE UNIQUE INDEX notif_templates_type_idx
	ON notif_templates(notif_provider,notif_type);

	ALTER TABLE public.notif_templates OWNER TO nails;
	
	

-- ******************* update 17/11/2023 12:30:00 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10008',
	'NotifTemplate_Controller',
	'get_list',
	'NotifTemplateList',
	'Справочники',
	'Шаблоны уведомлений',
	FALSE
	);
	
	

-- ******************* update 20/11/2023 13:00:24 ******************
	
		ALTER TABLE public.specialists ADD COLUMN user_id int NOT NULL REFERENCES users(id);
		


-- ******************* update 20/11/2023 13:06:40 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.confirmation_status
	(id serial NOT NULL,ref jsonb NOT NULL,field  varchar(50) NOT NULL,secret text NOT NULL,confirmed bool
			DEFAULT FALSE NOT NULL,tries int
			DEFAULT 0 NOT NULL,try_date_time timestampTZ,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,CONSTRAINT confirmation_status_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS confirmation_status_ref_field_idx;
	CREATE UNIQUE INDEX confirmation_status_ref_field_idx
	ON confirmation_status(field,(ref->>'dataType'),((ref->'keys'->>'id')::int));

	ALTER TABLE public.confirmation_status OWNER TO nails;
	
	

-- ******************* update 20/11/2023 13:20:45 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_regs
	(user_operation_id varchar(36),fields jsonb NOT NULL,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,inn_checked bool
			DEFAULT FALSE NOT NULL,CONSTRAINT specialist_regs_pkey PRIMARY KEY (user_operation_id)
	);
	
	ALTER TABLE public.specialist_regs OWNER TO nails;
	



-- ******************* update 21/11/2023 13:23:20 ******************
	
		ALTER TABLE public.specialist_regs ADD COLUMN inn_fns_ok bool
			DEFAULT FALSE NOT NULL;
		
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20000',
	'SpecialistReg_Controller',
	'get_list',
	'SpecialistRegList',
	'Журналы',
	'Регистрация самозанятых',
	FALSE
	);



-- ******************* update 21/11/2023 13:27:01 ******************
	
		ALTER TABLE public.specialist_regs ADD COLUMN email_confirmed bool
			DEFAULT FALSE NOT NULL;



-- ******************* update 21/11/2023 13:27:05 ******************
	
		ALTER TABLE public.specialist_regs ADD COLUMN tel_confirmed bool
			DEFAULT FALSE NOT NULL;



-- ******************* update 21/11/2023 13:30:39 ******************
DROP TABLE public.specialist_regs CASCADE;
/*	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_regs
	(id serial NOT NULL,user_operation_id  varchar(36) NOT NULL,fields jsonb NOT NULL,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,inn_checked bool
			DEFAULT FALSE NOT NULL,inn_fns_ok bool
			DEFAULT FALSE NOT NULL,CONSTRAINT specialist_regs_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_regs_operation_id;
	CREATE UNIQUE INDEX specialist_regs_operation_id
	ON specialist_regs(user_operation_id);

	ALTER TABLE public.specialist_regs OWNER TO nails;
	
*/	


-- ******************* update 21/11/2023 13:31:18 ******************
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_regs
	(id serial NOT NULL,user_operation_id  varchar(36) NOT NULL,fields jsonb NOT NULL,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,inn_checked bool
			DEFAULT FALSE NOT NULL,
			inn_fns_ok bool DEFAULT FALSE NOT NULL,
			tel_confirmed bool DEFAULT FALSE NOT NULL,
			email_confirmed bool DEFAULT FALSE NOT NULL,
			CONSTRAINT specialist_regs_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_regs_operation_id;
	CREATE UNIQUE INDEX specialist_regs_operation_id
	ON specialist_regs(user_operation_id);

	ALTER TABLE public.specialist_regs OWNER TO nails;
	



-- ******************* update 21/11/2023 17:20:39 ******************

					ALTER TYPE notif_types ADD VALUE 'tel_check';
					
					ALTER TYPE notif_types ADD VALUE 'email_check';
					
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_notif_types_val(notif_types,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='new_specialist'::notif_types AND $2='ru'::locales THEN 'Добавлен самозанятый'
		WHEN $1='tel_check'::notif_types AND $2='ru'::locales THEN 'Проверка телефона'
		WHEN $1='email_check'::notif_types AND $2='ru'::locales THEN 'Проверка электронной почты'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_notif_types_val(notif_types,locales) OWNER TO nails;		
	

-- ******************* update 21/11/2023 17:40:20 ******************

	-- ********** constant value table  email *************
	CREATE TABLE IF NOT EXISTS const_email
	(name text  NOT NULL, descr text, val jsonb,
		val_type text,ctrl_class text,ctrl_options json, view_class text,view_options json,
	CONSTRAINT const_email_pkey PRIMARY KEY (name));
	ALTER TABLE const_email OWNER TO nails;
	INSERT INTO const_email (name,descr,val,val_type,ctrl_class,ctrl_options,view_class,view_options) VALUES (
		'Данные электронной почты'
		,'Данные электронной почты для отправки писем'
		,NULL
		,'JSONB'
		,NULL
		,NULL
		,NULL
		,NULL
	);
	
		--constant get value
		
	CREATE OR REPLACE FUNCTION const_email_val()
	RETURNS jsonb AS
	$BODY$
		
		SELECT val::jsonb AS val FROM const_email LIMIT 1;
		
	$BODY$
	LANGUAGE sql STABLE COST 100;
	ALTER FUNCTION const_email_val() OWNER TO nails;
	
	--constant set value
	CREATE OR REPLACE FUNCTION const_email_set_val(JSONB)
	RETURNS void AS
	$BODY$
		UPDATE const_email SET val=$1;
	$BODY$
	LANGUAGE sql VOLATILE COST 100;
	ALTER FUNCTION const_email_set_val(JSONB) OWNER TO nails;
	
	--edit view: all keys and descr
	CREATE OR REPLACE VIEW const_email_view AS
	SELECT
		'email'::text AS id
		,t.name
		,t.descr
	,
	t.val::text AS val
	
	,t.val_type::text AS val_type
	,t.ctrl_class::text
	,t.ctrl_options::json
	,t.view_class::text
	,t.view_options::json
	FROM const_email AS t
	;
	ALTER VIEW const_email_view OWNER TO nails;
	
	
	
	CREATE OR REPLACE VIEW constants_list_view AS
	
	SELECT *
	FROM const_doc_per_page_count_view
	UNION ALL
	
	SELECT *
	FROM const_grid_refresh_interval_view
	UNION ALL
	
	SELECT *
	FROM const_email_view
	ORDER BY name;
	ALTER VIEW constants_list_view OWNER TO nails;
	
	

-- ******************* update 22/11/2023 14:13:19 ******************
DROP TABLE public.specialist_regs;
/*
	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_regs
	(id serial NOT NULL,user_operation_id  varchar(36) NOT NULL,name text NOT NULL,inn  varchar(12) NOT NULL,bank_bik  varchar(9),bank_account  varchar(20),studio_id int NOT NULL REFERENCES studios(id),birthdate date,address_reg text,passport json,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,inn_checked bool
			DEFAULT FALSE NOT NULL,inn_fns_ok bool
			DEFAULT FALSE NOT NULL,CONSTRAINT specialist_regs_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_regs_operation_id;
	CREATE UNIQUE INDEX specialist_regs_operation_id
	ON specialist_regs(user_operation_id);

	ALTER TABLE public.specialist_regs OWNER TO nails;
*/


-- ******************* update 22/11/2023 14:13:23 ******************
	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_regs
	(id serial NOT NULL,user_operation_id  varchar(36) NOT NULL,name text NOT NULL,inn  varchar(12) NOT NULL,bank_bik  varchar(9),bank_account  varchar(20),studio_id int NOT NULL REFERENCES studios(id),birthdate date,address_reg text,passport json,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,inn_checked bool
			DEFAULT FALSE NOT NULL,inn_fns_ok bool
			DEFAULT FALSE NOT NULL,CONSTRAINT specialist_regs_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_regs_operation_id;
	CREATE UNIQUE INDEX specialist_regs_operation_id
	ON specialist_regs(user_operation_id);

	ALTER TABLE public.specialist_regs OWNER TO nails;



-- ******************* update 22/11/2023 14:22:27 ******************
	ALTER TABLE specialist_regs ADD COLUMN banks_ref jsonb;
	/*
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_regs
	(id serial NOT NULL,user_operation_id  varchar(36) NOT NULL,name text NOT NULL,inn  varchar(12) NOT NULL,bank_bik  varchar(9),bank_account  varchar(20),studio_id int NOT NULL REFERENCES studios(id),birthdate date,address_reg text,passport json,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,inn_checked bool
			DEFAULT FALSE NOT NULL,inn_fns_ok bool
			DEFAULT FALSE NOT NULL,CONSTRAINT specialist_regs_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_regs_operation_id;
	CREATE UNIQUE INDEX specialist_regs_operation_id
	ON specialist_regs(user_operation_id);

	ALTER TABLE public.specialist_regs OWNER TO nails;
*/


-- ******************* update 22/11/2023 14:23:07 ******************
	ALTER TABLE specialist_regs ADD COLUMN name_full text;
	/*
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_regs
	(id serial NOT NULL,user_operation_id  varchar(36) NOT NULL,name text NOT NULL,inn  varchar(12) NOT NULL,bank_bik  varchar(9),bank_account  varchar(20),studio_id int NOT NULL REFERENCES studios(id),birthdate date,address_reg text,passport json,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,inn_checked bool
			DEFAULT FALSE NOT NULL,inn_fns_ok bool
			DEFAULT FALSE NOT NULL,CONSTRAINT specialist_regs_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_regs_operation_id;
	CREATE UNIQUE INDEX specialist_regs_operation_id
	ON specialist_regs(user_operation_id);

	ALTER TABLE public.specialist_regs OWNER TO nails;
*/


-- ******************* update 22/11/2023 14:23:50 ******************
	ALTER TABLE specialist_regs ADD COLUMN tel varchar(11);
	/*
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_regs
	(id serial NOT NULL,user_operation_id  varchar(36) NOT NULL,name text NOT NULL,inn  varchar(12) NOT NULL,bank_bik  varchar(9),bank_account  varchar(20),studio_id int NOT NULL REFERENCES studios(id),birthdate date,address_reg text,passport json,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,inn_checked bool
			DEFAULT FALSE NOT NULL,inn_fns_ok bool
			DEFAULT FALSE NOT NULL,CONSTRAINT specialist_regs_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_regs_operation_id;
	CREATE UNIQUE INDEX specialist_regs_operation_id
	ON specialist_regs(user_operation_id);

	ALTER TABLE public.specialist_regs OWNER TO nails;
*/


-- ******************* update 22/11/2023 14:26:09 ******************
	ALTER TABLE specialist_regs ADD COLUMN email varchar(50);
	/*
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_regs
	(id serial NOT NULL,user_operation_id  varchar(36) NOT NULL,name text NOT NULL,inn  varchar(12) NOT NULL,bank_bik  varchar(9),bank_account  varchar(20),studio_id int NOT NULL REFERENCES studios(id),birthdate date,address_reg text,passport json,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,inn_checked bool
			DEFAULT FALSE NOT NULL,inn_fns_ok bool
			DEFAULT FALSE NOT NULL,CONSTRAINT specialist_regs_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_regs_operation_id;
	CREATE UNIQUE INDEX specialist_regs_operation_id
	ON specialist_regs(user_operation_id);

	ALTER TABLE public.specialist_regs OWNER TO nails;
*/


-- ******************* update 22/11/2023 14:39:01 ******************
	ALTER TABLE specialist_regs ADD COLUMN bank_acc varchar(20);
	/*
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_regs
	(id serial NOT NULL,user_operation_id  varchar(36) NOT NULL,name text NOT NULL,inn  varchar(12) NOT NULL,bank_bik  varchar(9),bank_account  varchar(20),studio_id int NOT NULL REFERENCES studios(id),birthdate date,address_reg text,passport json,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,inn_checked bool
			DEFAULT FALSE NOT NULL,inn_fns_ok bool
			DEFAULT FALSE NOT NULL,CONSTRAINT specialist_regs_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_regs_operation_id;
	CREATE UNIQUE INDEX specialist_regs_operation_id
	ON specialist_regs(user_operation_id);

	ALTER TABLE public.specialist_regs OWNER TO nails;
*/


-- ******************* update 22/11/2023 14:46:34 ******************
	ALTER TABLE specialist_regs ALTER COLUMN studio_id DROP NOT NULL;
	/*
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_regs
	(id serial NOT NULL,user_operation_id  varchar(36) NOT NULL,name text NOT NULL,inn  varchar(12) NOT NULL,bank_bik  varchar(9),bank_account  varchar(20),studio_id int NOT NULL REFERENCES studios(id),birthdate date,address_reg text,passport json,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,inn_checked bool
			DEFAULT FALSE NOT NULL,inn_fns_ok bool
			DEFAULT FALSE NOT NULL,CONSTRAINT specialist_regs_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_regs_operation_id;
	CREATE UNIQUE INDEX specialist_regs_operation_id
	ON specialist_regs(user_operation_id);

	ALTER TABLE public.specialist_regs OWNER TO nails;
*/


-- ******************* update 23/11/2023 07:24:43 ******************
	
		ALTER TABLE public.specialist_regs ADD COLUMN email_sent bool
			DEFAULT FALSE, ADD COLUMN tel_sent bool
			DEFAULT FALSE, ADD COLUMN passport_uploaded bool
			DEFAULT FALSE;



-- ******************* update 23/11/2023 13:56:36 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'30000',
	'Mgateway_Controller',
	NULL,
	'Mgateway',
	'Формы',
	'Управление сообщениями',
	FALSE
	);
	
	

-- ******************* update 24/11/2023 11:48:17 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.document_templates
	(id serial NOT NULL,name text,sql_query text,cell_matching jsonb,CONSTRAINT document_templates_pkey PRIMARY KEY (id)
	);
	
	ALTER TABLE public.document_templates OWNER TO nails;
	
	

-- ******************* update 24/11/2023 12:08:40 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10009',
	'DocumentTemplate_Controller',
	'get_list',
	'DocumentTemplateList',
	'Справочники',
	'Шаблоны документов',
	FALSE
	);
	
	

-- ******************* update 24/11/2023 12:44:37 ******************
	
		ALTER TABLE public.document_templates ADD COLUMN fields jsonb;
		--ALTER TABLE public.document_templates DROP COLUMN cell_matching jsonb;
		--ALTER TABLE public.document_templates DROP COLUMN sql_query jsonb;


-- ******************* update 24/11/2023 12:44:59 ******************
	
		--ALTER TABLE public.document_templates ADD COLUMN fields jsonb;
		ALTER TABLE public.document_templates DROP COLUMN cell_matching;
		--ALTER TABLE public.document_templates DROP COLUMN sql_query;


-- ******************* update 24/11/2023 12:45:05 ******************
	
		--ALTER TABLE public.document_templates ADD COLUMN fields jsonb;
		--ALTER TABLE public.document_templates DROP COLUMN cell_matching;
		ALTER TABLE public.document_templates DROP COLUMN sql_query;


-- ******************* update 24/11/2023 12:45:29 ******************
	
		--ALTER TABLE public.document_templates ADD COLUMN fields jsonb;
		--ALTER TABLE public.document_templates DROP COLUMN cell_matching;
		ALTER TABLE public.document_templates ADD COLUMN sql_query text;


-- ******************* update 25/11/2023 08:13:36 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10006',
	'Studio_Controller',
	'get_list',
	'StudioList',
	'Справочники',
	'Салоны',
	FALSE
	);
	
	

-- ******************* update 25/11/2023 09:29:53 ******************
	
		ALTER TABLE public.specialists ADD COLUMN equipments jsonb;



-- ******************* update 25/11/2023 09:37:46 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_documents
	(id serial NOT NULL,specialist_id int NOT NULL REFERENCES specialists(id),template_att_id int NOT NULL REFERENCES attachments(id),dicument_att_id int NOT NULL REFERENCES attachments(id),date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,sign_date_time timestampTZ,open_date_time timestampTZ,need_signing bool,CONSTRAINT specialist_documents_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_documents_specialist_idx;
	CREATE INDEX specialist_documents_specialist_idx
	ON specialist_documents(specialist_id,date_time);

	ALTER TABLE public.specialist_documents OWNER TO nails;



-- ******************* update 25/11/2023 09:42:14 ******************
DROP TABLE public.specialist_documents;
	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_documents
	(id serial NOT NULL,specialist_id int NOT NULL REFERENCES specialists(id),template_att_id int NOT NULL REFERENCES attachments(id),document_att_id int NOT NULL REFERENCES attachments(id),date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,
			sign_date_time timestampTZ,
			open_date_time timestampTZ,
			need_signing bool,
			sign_img bytea,
			CONSTRAINT specialist_documents_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_documents_specialist_idx;
	CREATE INDEX specialist_documents_specialist_idx
	ON specialist_documents(specialist_id,date_time);

	ALTER TABLE public.specialist_documents OWNER TO nails;



-- ******************* update 26/11/2023 08:24:26 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_documents_on_register
	(id serial NOT NULL,document_template_id int NOT NULL REFERENCES document_templates(id),date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,need_signing bool,CONSTRAINT specialist_documents_on_register_pkey PRIMARY KEY (id)
	);
	
	ALTER TABLE public.specialist_documents_on_register OWNER TO nails;
	
	



-- ******************* update 26/11/2023 08:53:09 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10010',
	'SpecialistDocumentOnRegister_Controller',
	'get_list',
	'SpecialistDocumentOnRegisterList',
	'Справочники',
	'Шаблоны документов для отправки при регистрации',
	FALSE
	);
	
	

-- ******************* update 26/11/2023 09:57:48 ******************

					ALTER TYPE specialist_status_types ADD VALUE 'contract_signing';
					
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_specialist_status_types_val(specialist_status_types,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='contract_signing'::specialist_status_types AND $2='ru'::locales THEN 'Подписание контракта'
		WHEN $1='contract_signed'::specialist_status_types AND $2='ru'::locales THEN 'Заключен контракт'
		WHEN $1='contract_terminated'::specialist_status_types AND $2='ru'::locales THEN 'Расторгнут контракт'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_specialist_status_types_val(specialist_status_types,locales) OWNER TO nails;		
	
	DROP INDEX IF EXISTS specialist_statuses_specialist;
	CREATE UNIQUE INDEX specialist_statuses_specialist
	ON specialist_statuses(specialist_id,date_time);

--Refrerece type
CREATE OR REPLACE FUNCTION specialists_ref(specialists)
  RETURNS json AS
$$
	SELECT json_build_object(
		'keys',json_build_object(
			'id',$1.id    
			),	
		'descr',$1.name,
		'dataType','specialists'
	);
$$
  LANGUAGE sql VOLATILE COST 100;
ALTER FUNCTION specialists_ref(specialists) OWNER TO nails;	
	

-- ******************* update 26/11/2023 10:50:32 ******************

					ALTER TYPE data_types ADD VALUE 'specialists';
					
					ALTER TYPE data_types ADD VALUE 'specialist_regs';
					
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_data_types_val(data_types,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='users'::data_types AND $2='ru'::locales THEN 'Пользователи'
		WHEN $1='specialists'::data_types AND $2='ru'::locales THEN 'Мастера'
		WHEN $1='specialist_regs'::data_types AND $2='ru'::locales THEN 'Регистрация мастера'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_data_types_val(data_types,locales) OWNER TO nails;		
	

-- ******************* update 27/11/2023 09:58:46 ******************

	DROP INDEX IF EXISTS attachments_idx;
	CREATE UNIQUE INDEX attachments_idx
	ON attachments((ref->>'dataType'),((ref->'keys'->>'id')::int),(content_info->>'id'));


-- ******************* update 27/11/2023 11:44:30 ******************

	DROP INDEX IF EXISTS specialists_inn_idx;
	CREATE UNIQUE INDEX specialists_inn_idx
	ON specialists(inn);



-- ******************* update 27/11/2023 12:26:02 ******************
	
		ALTER TABLE public.specialist_regs DROP COLUMN bank_account;



-- ******************* update 27/11/2023 12:26:32 ******************
	
		ALTER TABLE public.specialist_regs DROP COLUMN bank_bik;



-- ******************* update 28/11/2023 07:41:45 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_works
	(id serial NOT NULL,specialist_id int NOT NULL REFERENCES specialists(id),studio_id int NOT NULL REFERENCES studios(id),date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,CONSTRAINT specialist_works_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_works_specialist_idx;
	CREATE INDEX specialist_works_specialist_idx
	ON specialist_works(specialist_id);

	DROP INDEX IF EXISTS specialist_works_studio_idx;
	CREATE INDEX specialist_works_studio_idx
	ON specialist_works(studio_id);

	DROP INDEX IF EXISTS specialist_works_date_time_idx;
	CREATE INDEX specialist_works_date_time_idx
	ON specialist_works(date_time);

	ALTER TABLE public.specialist_works OWNER TO nails;
	



-- ******************* update 28/11/2023 07:51:27 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20001',
	'SpecialistWork_Controller',
	'get_list',
	'SpecialistWorkList',
	'Журналы',
	'Работы самозанятых',
	FALSE
	);
	
	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20002',
	'SpecialistDocument_Controller',
	'get_list',
	'SpecialistDocumentList',
	'Журналы',
	'Документы самозанятых',
	FALSE
	);
	
	

-- ******************* update 28/11/2023 10:11:49 ******************

					ALTER TYPE notif_types ADD VALUE 'docs_for_sign';
					
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_notif_types_val(notif_types,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='new_specialist'::notif_types AND $2='ru'::locales THEN 'Добавлен самозанятый'
		WHEN $1='tel_check'::notif_types AND $2='ru'::locales THEN 'Проверка телефона'
		WHEN $1='email_check'::notif_types AND $2='ru'::locales THEN 'Проверка электронной почты'
		WHEN $1='docs_for_sign'::notif_types AND $2='ru'::locales THEN 'Документы для подписания'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_notif_types_val(notif_types,locales) OWNER TO nails;		
	

-- ******************* update 28/11/2023 10:20:40 ******************

					ALTER TYPE notif_types ADD VALUE 'signed_docs';
					
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_notif_types_val(notif_types,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='new_specialist'::notif_types AND $2='ru'::locales THEN 'Добавлен самозанятый'
		WHEN $1='tel_check'::notif_types AND $2='ru'::locales THEN 'Проверка телефона'
		WHEN $1='email_check'::notif_types AND $2='ru'::locales THEN 'Проверка электронной почты'
		WHEN $1='docs_for_sign'::notif_types AND $2='ru'::locales THEN 'Документы для подписания'
		WHEN $1='signed_docs'::notif_types AND $2='ru'::locales THEN 'Подписанные документы'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_notif_types_val(notif_types,locales) OWNER TO nails;		
	

-- ******************* update 29/11/2023 09:40:49 ******************

					ALTER TYPE notif_types ADD VALUE 'new_account';
					
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_notif_types_val(notif_types,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='new_specialist'::notif_types AND $2='ru'::locales THEN 'Добавлен самозанятый'
		WHEN $1='tel_check'::notif_types AND $2='ru'::locales THEN 'Проверка телефона'
		WHEN $1='email_check'::notif_types AND $2='ru'::locales THEN 'Проверка электронной почты'
		WHEN $1='docs_for_sign'::notif_types AND $2='ru'::locales THEN 'Документы для подписания'
		WHEN $1='signed_docs'::notif_types AND $2='ru'::locales THEN 'Подписанные документы'
		WHEN $1='new_account'::notif_types AND $2='ru'::locales THEN 'Новый аккаунт'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_notif_types_val(notif_types,locales) OWNER TO nails;		
	

-- ******************* update 29/11/2023 10:36:41 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20003',
	'SpecialistDocument_Controller',
	'get_list',
	'SpecialistDocumentForSignList',
	'Журналы',
	'Документы на подпись',
	FALSE
	);
	
	

-- ******************* update 29/11/2023 10:41:33 ******************

		UPDATE views SET
			c='Specialist_Controller',
			f='get_list',
			t='SpecialistList',
			section='Справочники',
			descr='Мастера',
			limited=FALSE
		WHERE id='10007';
	
		UPDATE views SET
			c='SpecialistWork_Controller',
			f='get_list',
			t='SpecialistWorkList',
			section='Журналы',
			descr='Работы',
			limited=FALSE
		WHERE id='20001';
	

-- ******************* update 29/11/2023 11:37:27 ******************

		UPDATE views SET
			c='SpecialistDocument_Controller',
			f='get_for_sign_list',
			t='SpecialistDocumentForSignList',
			section='Журналы',
			descr='Документы на подпись',
			limited=FALSE
		WHERE id='20003';
	

-- ******************* update 29/11/2023 12:10:26 ******************
	
		ALTER TABLE public.specialist_documents ADD COLUMN name text;



-- ******************* update 29/11/2023 15:00:10 ******************
	
		ALTER TABLE public.specialist_documents ADD COLUMN document_template_id int REFERENCES document_templates(id);



-- ******************* update 29/11/2023 15:01:11 ******************
	
		ALTER TABLE public.specialist_documents ALTER COLUMN document_template_id SET NOT NULL;



-- ******************* update 30/11/2023 12:09:55 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.equipment_types
	(id serial NOT NULL,name text NOT NULL,CONSTRAINT equipment_types_pkey PRIMARY KEY (id)
	);
	
	ALTER TABLE public.equipment_types OWNER TO nails;
	
	

-- ******************* update 30/11/2023 12:18:04 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20004',
	'EquipmentType_Controller',
	'get_list',
	'EquipmentTypeList',
	'Журналы',
	'Виды оборудования',
	FALSE
	);
	
	

-- ******************* update 30/11/2023 12:30:23 ******************
/*
	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialities
	(id serial NOT NULL,name text NOT NULL,CONSTRAINT specialities_pkey PRIMARY KEY (id)
	);
	
	ALTER TABLE public.specialities OWNER TO nails;
	
	
	
	-- ********** Adding new table from model **********
	CREATE TABLE public.items
	(id serial NOT NULL,name text NOT NULL,comment_text text NOT NULL,price  numeric(15,2),CONSTRAINT items_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS items_name_idx;
	CREATE UNIQUE INDEX items_name_idx
	ON items(lower(name));

	ALTER TABLE public.items OWNER TO nails;
	
	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20005',
	'Item_Controller',
	'get_list',
	'ItemList',
	'Журналы',
	'Номенклатура',
	FALSE
	);
	
*/	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20006',
	'Speciality_Controller',
	'get_list',
	'SpecialityList',
	'Журналы',
	'Специальности',
	FALSE
	);
	
	


-- ******************* update 30/11/2023 14:50:24 ******************
ALTER TABLE items ALTER COLUMN comment_text DROP NOT NULL;


-- ******************* update 30/11/2023 15:04:21 ******************
	
		ALTER TABLE public.specialities ADD COLUMN equipments jsonb;
		

-- ******************* update 05/12/2023 09:36:56 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.ycl_staff
	(id serial NOT NULL,name text,data jsonb,CONSTRAINT ycl_staff_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS ycl_staff_name_idx;
	CREATE INDEX ycl_staff_name_idx
	ON ycl_staff(lower(name));

	ALTER TABLE public.ycl_staff OWNER TO nails;
	
	

-- ******************* update 05/12/2023 09:40:34 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.ycl_transactions
	(id serial NOT NULL,created_at timestampTZ NOT NULL,data jsonb,specialist_id int NOT NULL REFERENCES specialists(id),CONSTRAINT ycl_transactions_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS ycl_transactions_specialist_idx;
	CREATE INDEX ycl_transactions_specialist_idx
	ON ycl_transactions(specialist_id,created_at);

	ALTER TABLE public.ycl_transactions OWNER TO nails;
	



-- ******************* update 05/12/2023 09:56:19 ******************
	
		ALTER TABLE public.specialists ADD COLUMN ycl_staff_id int REFERENCES ycl_staff(id);



-- ******************* update 05/12/2023 09:57:04 ******************
	CREATE UNIQUE INDEX specialists_ycl_staff_idx
	ON specialists(ycl_staff_id);



-- ******************* update 05/12/2023 11:17:39 ******************
	
		ALTER TABLE public.specialist_works ADD COLUMN admin_rate int;



-- ******************* update 05/12/2023 15:59:12 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.ycl_visits
	(id serial NOT NULL,created_at timestampTZ NOT NULL,data jsonb,specialist_id int NOT NULL REFERENCES specialists(id),CONSTRAINT ycl_visits_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS ycl_visits_specialist_idx;
	CREATE INDEX ycl_visits_specialist_idx
	ON ycl_visits(specialist_id,created_at);

	ALTER TABLE public.ycl_visits OWNER TO nails;



-- ******************* update 05/12/2023 17:08:51 ******************
	
		ALTER TABLE public.specialist_works ADD COLUMN ycl_document_id int;
		
	CREATE INDEX specialist_works_ycl_document_id_idx
	ON specialist_works(ycl_document_id);



-- ******************* update 05/12/2023 17:12:36 ******************
	DROP INDEX ycl_transactions_document_idx;
	CREATE INDEX ycl_transactions_document_idx
	ON ycl_transactions((data->>'document_id'));

DROP INDEX ycl_transactions_date_idx;
	CREATE INDEX ycl_transactions_date_idx
	ON ycl_transactions((data->>'date'));



-- ******************* update 05/12/2023 18:00:42 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20007',
	'YClStaff_Controller',
	'get_list',
	'YClStaffList',
	'Журналы',
	'Сотрудники yclients',
	FALSE
	);
	
	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20008',
	'YClTransaction_Controller',
	'get_list',
	'YClTransactionList',
	'Журналы',
	'Транзакции yclients',
	FALSE
	);
	
	

-- ******************* update 08/12/2023 08:40:38 ******************

	-- ********** constant value table  join_contract *************
	CREATE TABLE IF NOT EXISTS const_join_contract
	(name text  NOT NULL, descr text, val text,
		val_type text,ctrl_class text,ctrl_options json, view_class text,view_options json,
	CONSTRAINT const_join_contract_pkey PRIMARY KEY (name));
	ALTER TABLE const_join_contract OWNER TO nails;
	INSERT INTO const_join_contract (name,descr,val,val_type,ctrl_class,ctrl_options,view_class,view_options) VALUES (
		'Договор присоединения'
		,'Договор присоединения'
		,NULL
		,'Text'
		,NULL
		,NULL
		,NULL
		,NULL
	);
	
		--constant get value
		
	CREATE OR REPLACE FUNCTION const_join_contract_val()
	RETURNS text AS
	$BODY$
		
		SELECT val::text AS val FROM const_join_contract LIMIT 1;
		
	$BODY$
	LANGUAGE sql STABLE COST 100;
	ALTER FUNCTION const_join_contract_val() OWNER TO nails;
	
	--constant set value
	CREATE OR REPLACE FUNCTION const_join_contract_set_val(Text)
	RETURNS void AS
	$BODY$
		UPDATE const_join_contract SET val=$1;
	$BODY$
	LANGUAGE sql VOLATILE COST 100;
	ALTER FUNCTION const_join_contract_set_val(Text) OWNER TO nails;
	
	--edit view: all keys and descr
	CREATE OR REPLACE VIEW const_join_contract_view AS
	SELECT
		'join_contract'::text AS id
		,t.name
		,t.descr
	,
	t.val::text AS val
	
	,t.val_type::text AS val_type
	,t.ctrl_class::text
	,t.ctrl_options::json
	,t.view_class::text
	,t.view_options::json
	FROM const_join_contract AS t
	;
	ALTER VIEW const_join_contract_view OWNER TO nails;
	
	
	
	CREATE OR REPLACE VIEW constants_list_view AS
	
	SELECT *
	FROM const_doc_per_page_count_view
	UNION ALL
	
	SELECT *
	FROM const_grid_refresh_interval_view
	UNION ALL
	
	SELECT *
	FROM const_email_view
	UNION ALL
	
	SELECT *
	FROM const_join_contract_view
	ORDER BY name;
	ALTER VIEW constants_list_view OWNER TO nails;
	
	

-- ******************* update 12/12/2023 07:05:04 ******************
	
		ALTER TABLE public.specialists ADD COLUMN agent_percent  numeric(15,2);



-- ******************* update 12/12/2023 07:07:07 ******************
	
		ALTER TABLE public.specialities ADD COLUMN agent_percent  numeric(15,2);
		

-- ******************* update 12/12/2023 07:19:08 ******************
	
		ALTER TABLE public.specialists ADD COLUMN spciality_id int REFERENCES specialities(id);



-- ******************* update 12/12/2023 07:24:20 ******************
	ALTER TABLE public.specialists DROP COLUMN spciality_id;
--		ALTER TABLE public.specialists ADD COLUMN speciality_id int REFERENCES specialities(id);



-- ******************* update 12/12/2023 07:24:24 ******************
		ALTER TABLE public.specialists ADD COLUMN speciality_id int REFERENCES specialities(id);



-- ******************* update 12/12/2023 07:48:55 ******************
	
		ALTER TABLE public.studios ADD COLUMN hour_rent_price  numeric(15,2)
			DEFAULT 0;



-- ******************* update 12/12/2023 08:17:05 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.salary_debets
	(id serial NOT NULL,name text NOT NULL,CONSTRAINT salary_debets_pkey PRIMARY KEY (id)
	);
	
	ALTER TABLE public.salary_debets OWNER TO nails;
	
	
/*	
	-- ********** Adding new table from model **********
	CREATE TABLE public.salary_kredits
	(id serial NOT NULL,name text NOT NULL,CONSTRAINT salary_kredits_pkey PRIMARY KEY (id)
	);
	
	ALTER TABLE public.salary_kredits OWNER TO nails;
	
	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10011',
	'SalaryDebet_Controller',
	'get_list',
	'SalaryDebetList',
	'Справочники',
	'Начисления',
	FALSE
	);
	
	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10012',
	'SalaryKredit_Controller',
	'get_list',
	'SalaryKreditList',
	'Справочники',
	'Удержания',
	FALSE
	);
	
*/


-- ******************* update 12/12/2023 08:17:14 ******************
	
	-- ********** Adding new table from model **********
	CREATE TABLE public.salary_kredits
	(id serial NOT NULL,name text NOT NULL,CONSTRAINT salary_kredits_pkey PRIMARY KEY (id)
	);
	
	ALTER TABLE public.salary_kredits OWNER TO nails;
	
	
/*	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10011',
	'SalaryDebet_Controller',
	'get_list',
	'SalaryDebetList',
	'Справочники',
	'Начисления',
	FALSE
	);
	
	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10012',
	'SalaryKredit_Controller',
	'get_list',
	'SalaryKreditList',
	'Справочники',
	'Удержания',
	FALSE
	);
	
*/


-- ******************* update 12/12/2023 08:17:31 ******************
	

	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10011',
	'SalaryDebet_Controller',
	'get_list',
	'SalaryDebetList',
	'Справочники',
	'Начисления',
	FALSE
	);
	
	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10012',
	'SalaryKredit_Controller',
	'get_list',
	'SalaryKreditList',
	'Справочники',
	'Удержания',
	FALSE
	);
	



-- ******************* update 12/12/2023 08:18:16 ******************

		UPDATE views SET
			c='Studio_Controller',
			f='get_list',
			t='StudioList',
			section='Справочники',
			descr='Студии',
			limited=FALSE
		WHERE id='10006';
	

-- ******************* update 12/12/2023 08:30:57 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_salary_debets
	(id serial NOT NULL,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,specialist_id int NOT NULL REFERENCES specialists(id),salary_debet_id int NOT NULL REFERENCES salary_debets(id),total  numeric(15,2) NOT NULL,CONSTRAINT specialist_salary_debets_pkey PRIMARY KEY (id)
	);



-- ******************* update 12/12/2023 08:36:24 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_salary_kredits
	(id serial NOT NULL,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,specialist_id int NOT NULL REFERENCES specialists(id),salary_kredit_id int NOT NULL REFERENCES salary_kredits(id),total  numeric(15,2) NOT NULL,CONSTRAINT specialist_salary_kredits_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_salary_kredits_specialist_idx;
	CREATE INDEX specialist_salary_kredits_specialist_idx
	ON specialist_salary_kredits(specialist_id);

	DROP INDEX IF EXISTS specialist_salary_kredits_date_time_idx;
	CREATE INDEX specialist_salary_kredits_date_time_idx
	ON specialist_salary_kredits(date_time);

	ALTER TABLE public.specialist_salary_kredits OWNER TO nails;



-- ******************* update 12/12/2023 09:17:35 ******************

	-- ********** constant value table  person_tax *************
	CREATE TABLE IF NOT EXISTS const_person_tax
	(name text  NOT NULL, descr text, val int,
		val_type text,ctrl_class text,ctrl_options json, view_class text,view_options json,
	CONSTRAINT const_person_tax_pkey PRIMARY KEY (name));
	ALTER TABLE const_person_tax OWNER TO nails;
	INSERT INTO const_person_tax (name,descr,val,val_type,ctrl_class,ctrl_options,view_class,view_options) VALUES (
		'Налог на самлзанятого,%'
		,'Процент налога с самозанятого'
		,6
		,'Int'
		,NULL
		,NULL
		,NULL
		,NULL
	);
	
		--constant get value
		
	CREATE OR REPLACE FUNCTION const_person_tax_val()
	RETURNS int AS
	$BODY$
		
		SELECT val::int AS val FROM const_person_tax LIMIT 1;
		
	$BODY$
	LANGUAGE sql STABLE COST 100;
	ALTER FUNCTION const_person_tax_val() OWNER TO nails;
	
	--constant set value
	CREATE OR REPLACE FUNCTION const_person_tax_set_val(Int)
	RETURNS void AS
	$BODY$
		UPDATE const_person_tax SET val=$1;
	$BODY$
	LANGUAGE sql VOLATILE COST 100;
	ALTER FUNCTION const_person_tax_set_val(Int) OWNER TO nails;
	
	--edit view: all keys and descr
	CREATE OR REPLACE VIEW const_person_tax_view AS
	SELECT
		'person_tax'::text AS id
		,t.name
		,t.descr
	,
	t.val::text AS val
	
	,t.val_type::text AS val_type
	,t.ctrl_class::text
	,t.ctrl_options::json
	,t.view_class::text
	,t.view_options::json
	FROM const_person_tax AS t
	;
	ALTER VIEW const_person_tax_view OWNER TO nails;
	
	
	
	CREATE OR REPLACE VIEW constants_list_view AS
	
	SELECT *
	FROM const_doc_per_page_count_view
	UNION ALL
	
	SELECT *
	FROM const_grid_refresh_interval_view
	UNION ALL
	
	SELECT *
	FROM const_email_view
	UNION ALL
	
	SELECT *
	FROM const_join_contract_view
	UNION ALL
	
	SELECT *
	FROM const_person_tax_view
	ORDER BY name;
	ALTER VIEW constants_list_view OWNER TO nails;
	
	

-- ******************* update 12/12/2023 09:31:56 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_period_salaries
	(id serial NOT NULL,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,period date NOT NULL,total  numeric(15,2),CONSTRAINT specialist_period_salaries_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_period_salaries_period_idx;
	CREATE INDEX specialist_period_salaries_period_idx
	ON specialist_period_salaries(period);

	ALTER TABLE public.specialist_period_salaries OWNER TO nails;
	
	


-- ******************* update 12/12/2023 10:03:06 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20011',
	'SpecialistPeriodSalary_Controller',
	'get_list',
	'SpecialistPeriodSalaryList',
	'Журналы',
	'Зарплата сотрудников',
	FALSE
	);
	
	

-- ******************* update 13/12/2023 14:56:24 ******************
	
		ALTER TABLE public.specialist_period_salaries ADD COLUMN work_total  numeric(15,2),ADD COLUMN hours int,ADD COLUMN debet numeric(15,2),ADD COLUMN kredit numeric(15,2),
		ADD COLUMN rent_total numeric(15,2);



-- ******************* update 13/12/2023 15:03:35 ******************
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_period_salary_details
	(id serial NOT NULL,specialist_period_salary_id int NOT NULL REFERENCES specialist_period_salaries(id),specialist_id int NOT NULL REFERENCES specialists(id),studio_id int NOT NULL REFERENCES studios(id),period date NOT NULL,hours int,work_total  numeric(15,2),debet  numeric(15,2),kredit  numeric(15,2),rent_price  numeric(15,2),rent_total  numeric(15,2),total  numeric(15,2),CONSTRAINT specialist_period_salary_details_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_period_salaries_period_idx;
	CREATE INDEX specialist_period_salaries_period_idx
	ON specialist_period_salary_details(period);

	ALTER TABLE public.specialist_period_salary_details OWNER TO nails;



-- ******************* update 13/12/2023 15:11:42 ******************
	
		ALTER TABLE public.specialist_period_salary_details ADD COLUMN line_num int NOT NULL;
		
	DROP INDEX IF EXISTS specialist_period_salaries_period_idx;
	CREATE INDEX specialist_period_salaries_period_idx
	ON specialist_period_salary_details(period);

	DROP INDEX IF EXISTS specialist_period_salaries_specialist_idx;
	CREATE INDEX specialist_period_salaries_specialist_idx
	ON specialist_period_salary_details(specialist_id);

	DROP INDEX IF EXISTS specialist_period_salaries_period_idx;
	CREATE UNIQUE INDEX specialist_period_salaries_period_idx
	ON specialist_period_salary_details(specialist_period_salary_id,line_num);



-- ******************* update 13/12/2023 16:16:52 ******************
	
		ALTER TABLE public.specialist_period_salaries ADD COLUMN work_total_salary  numeric(15,2);



-- ******************* update 13/12/2023 16:18:25 ******************
	
		ALTER TABLE public.specialist_period_salary_details ADD COLUMN work_total_salary  numeric(15,2);



-- ******************* update 13/12/2023 16:27:16 ******************
	
		ALTER TABLE public.specialist_period_salaries ADD COLUMN studio_id int NOT NULL REFERENCES studios(id);
		
	CREATE INDEX specialist_period_salaries_studio_id_idx
	ON specialist_period_salaries(studio_id);



-- ******************* update 14/12/2023 17:19:53 ******************
	
		ALTER TABLE public.specialist_period_salary_details ADD COLUMN agent_percent  numeric(15,2);



-- ******************* update 15/12/2023 07:53:41 ******************
	
		ALTER TABLE public.ycl_transactions ADD COLUMN document_id int,ADD COLUMN date_time timestampTZ,ADD COLUMN amount  numeric(15,2),ADD COLUMN hours  numeric(15,2);
		
	DROP INDEX IF EXISTS ycl_transactions_specialist_idx;
	CREATE INDEX ycl_transactions_specialist_idx
	ON ycl_transactions(specialist_id);

	DROP INDEX IF EXISTS ycl_transactions_document_idx;
	CREATE INDEX ycl_transactions_document_idx
	ON ycl_transactions(document_id);

	DROP INDEX IF EXISTS ycl_transactions_date_time_idx;
	CREATE INDEX ycl_transactions_date_time_idx
	ON ycl_transactions(date_time);



-- ******************* update 15/12/2023 08:07:08 ******************
	
		ALTER TABLE public.ycl_transactions ADD COLUMN seance_length  int;



-- ******************* update 15/12/2023 08:07:17 ******************
	
		ALTER TABLE public.ycl_transactions DROP COLUMN hours;



-- ******************* update 15/12/2023 08:09:07 ******************
	
		ALTER TABLE public.ycl_transactions ADD COLUMN record_id int;



-- ******************* update 15/12/2023 08:12:00 ******************
	
		ALTER TABLE public.ycl_transactions ADD COLUMN staff_id int;



-- ******************* update 15/12/2023 08:33:13 ******************
	
		ALTER TABLE public.ycl_transactions ADD COLUMN record_inf_updated bool default FALSE;



-- ******************* update 15/12/2023 11:02:19 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.specialist_receipts
	(id serial NOT NULL,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,document_date_time timestampTZ,document_total  numeric(15,2)
			DEFAULT 0,document_parsed bool
			DEFAULT FALSE,specialist_id int NOT NULL REFERENCES specialists(id),specialist_period_salary_detail_id int NOT NULL REFERENCES specialist_period_salary_details(id),CONSTRAINT specialist_receipts_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS specialist_receipts_specialist_idx;
	CREATE INDEX specialist_receipts_specialist_idx
	ON specialist_receipts(specialist_id);

	DROP INDEX IF EXISTS specialist_receipts_salary_detail_idx;
	CREATE INDEX specialist_receipts_salary_detail_idx
	ON specialist_receipts(specialist_period_salary_detail_id);

	ALTER TABLE public.specialist_receipts OWNER TO nails;



-- ******************* update 15/12/2023 11:08:41 ******************
	
		ALTER TABLE public.specialist_receipts ADD COLUMN document_error text;



-- ******************* update 18/12/2023 17:01:55 ******************
	
		ALTER TABLE public.specialist_receipts ADD COLUMN qrextr_request_id text;



-- ******************* update 18/12/2023 17:50:41 ******************
	
		ALTER TABLE public.specialist_receipts ADD COLUMN operation_id  varchar(36);



-- ******************* update 18/12/2023 17:55:00 ******************
	CREATE INDEX specialist_receipts_qrextr_request_id_idx
	ON specialist_receipts(qrextr_request_id);


-- ******************* update 19/12/2023 04:02:16 ******************

	-- ********** constant value table  specialist_services *************
	CREATE TABLE IF NOT EXISTS const_specialist_services
	(name text  NOT NULL, descr text, val text,
		val_type text,ctrl_class text,ctrl_options json, view_class text,view_options json,
	CONSTRAINT const_specialist_services_pkey PRIMARY KEY (name));
	ALTER TABLE const_specialist_services OWNER TO nails;
	INSERT INTO const_specialist_services (name,descr,val,val_type,ctrl_class,ctrl_options,view_class,view_options) VALUES (
		'Услуги самозанятого'
		,'Все возможные усдуги самозанятого для проверки чека'
		,
			'маник, педик'
			
		,'Text'
		,NULL
		,NULL
		,NULL
		,NULL
	);
	
		--constant get value
		
	CREATE OR REPLACE FUNCTION const_specialist_services_val()
	RETURNS text AS
	$BODY$
		
		SELECT val::text AS val FROM const_specialist_services LIMIT 1;
		
	$BODY$
	LANGUAGE sql STABLE COST 100;
	ALTER FUNCTION const_specialist_services_val() OWNER TO nails;
	
	--constant set value
	CREATE OR REPLACE FUNCTION const_specialist_services_set_val(Text)
	RETURNS void AS
	$BODY$
		UPDATE const_specialist_services SET val=$1;
	$BODY$
	LANGUAGE sql VOLATILE COST 100;
	ALTER FUNCTION const_specialist_services_set_val(Text) OWNER TO nails;
	
	--edit view: all keys and descr
	CREATE OR REPLACE VIEW const_specialist_services_view AS
	SELECT
		'specialist_services'::text AS id
		,t.name
		,t.descr
	,
	t.val::text AS val
	
	,t.val_type::text AS val_type
	,t.ctrl_class::text
	,t.ctrl_options::json
	,t.view_class::text
	,t.view_options::json
	FROM const_specialist_services AS t
	;
	ALTER VIEW const_specialist_services_view OWNER TO nails;
	
	
	
	CREATE OR REPLACE VIEW constants_list_view AS
	
	SELECT *
	FROM const_doc_per_page_count_view
	UNION ALL
	
	SELECT *
	FROM const_grid_refresh_interval_view
	UNION ALL
	
	SELECT *
	FROM const_email_view
	UNION ALL
	
	SELECT *
	FROM const_join_contract_view
	UNION ALL
	
	SELECT *
	FROM const_person_tax_view
	UNION ALL
	
	SELECT *
	FROM const_specialist_services_view
	ORDER BY name;
	ALTER VIEW constants_list_view OWNER TO nails;
	
	


-- ******************* update 19/12/2023 04:38:13 ******************
	
		ALTER TABLE public.specialist_receipts ADD COLUMN document_href text;



-- ******************* update 19/12/2023 05:27:11 ******************
	
		ALTER TABLE public.document_templates ADD COLUMN need_signing bool;
		

-- ******************* update 19/12/2023 05:41:01 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.template_batches
	(id serial NOT NULL,name text,CONSTRAINT template_batches_pkey PRIMARY KEY (id)
	);
	
	ALTER TABLE public.template_batches OWNER TO nails;
	
	
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'10013',
	'TemplateBatch_Controller',
	'get_list',
	'TemplateBatchList',
	'Справочники',
	'Пакеты шаблонов',
	FALSE
	);
	
	

-- ******************* update 19/12/2023 05:49:14 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.template_batch_items
	(id serial NOT NULL,template_batch_id int NOT NULL REFERENCES template_batches(id),template_id int NOT NULL REFERENCES document_templates(id),CONSTRAINT template_batch_items_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS template_batch_items_template_idx;
	CREATE UNIQUE INDEX template_batch_items_template_idx
	ON template_batch_items(template_batch_id,template_id);

	ALTER TABLE public.template_batch_items OWNER TO nails;
	



-- ******************* update 19/12/2023 09:24:51 ******************

	-- Adding new type
	CREATE TYPE template_batch_types AS ENUM (
	
		'specialist_registration'			
	,
		'specialist_salary'			
				
	);
	ALTER TYPE template_batch_types OWNER TO nails;
		
	/* type get function */
	CREATE OR REPLACE FUNCTION enum_template_batch_types_val(template_batch_types,locales)
	RETURNS text AS $$
		SELECT
		CASE
		WHEN $1='specialist_registration'::template_batch_types AND $2='ru'::locales THEN 'Регистрация мастера'
		WHEN $1='specialist_salary'::template_batch_types AND $2='ru'::locales THEN 'Расчет зарплаты'
		ELSE ''
		END;		
	$$ LANGUAGE sql;	
	ALTER FUNCTION enum_template_batch_types_val(template_batch_types,locales) OWNER TO nails;		
	

-- ******************* update 19/12/2023 09:26:48 ******************
	
		ALTER TABLE public.template_batches ADD COLUMN template_batch_type template_batch_types NOT NULL;
		

-- ******************* update 19/12/2023 09:27:48 ******************
	
		ALTER TABLE public.template_batches ALTER COLUMN template_batch_type DROP NOT NULL;
		


-- ******************* update 19/12/2023 11:42:17 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20013',
	'SpecialistPeriodSalaryDetail_Controller',
	'get_list',
	'SpecialistPeriodSalaryDetailList',
	'Журналы',
	'Зарплата сотрудника',
	FALSE
	);
	
	

-- ******************* update 21/12/2023 01:37:34 ******************

	
	-- ********** Adding new table from model **********
	CREATE TABLE public.bank_payments
	(id serial NOT NULL,date_time timestampTZ
			DEFAULT CURRENT_TIMESTAMP NOT NULL,document_date timestampTZ,document_num int,document_total  numeric(15,2)
			DEFAULT 0,document_comment text,payer_acc  varchar(20),payer_bank_acc  varchar(20),payer_bank_bik  varchar(9),payer_bank text,payer_bank_place text,rec_acc  varchar(20),rec_bank_acc  varchar(20),rec_bank_bik  varchar(9),rec_bank text,rec_bank_place text,specialist_id int NOT NULL REFERENCES specialists(id),specialist_period_salary_detail_id int NOT NULL REFERENCES specialist_period_salary_details(id),CONSTRAINT bank_payments_pkey PRIMARY KEY (id)
	);
	
	DROP INDEX IF EXISTS bank_payments_specialist_idx;
	CREATE INDEX bank_payments_specialist_idx
	ON bank_payments(specialist_id);

	DROP INDEX IF EXISTS bank_payments_document_date_idx;
	CREATE INDEX bank_payments_document_date_idx
	ON bank_payments(document_date);

	DROP INDEX IF EXISTS bank_payments_document_num_idx;
	CREATE INDEX bank_payments_document_num_idx
	ON bank_payments(document_num);

	DROP INDEX IF EXISTS bank_payments_specialist_salary_idx;
	CREATE INDEX bank_payments_specialist_salary_idx
	ON bank_payments(specialist_period_salary_detail_id);

	ALTER TABLE public.bank_payments OWNER TO nails;



-- ******************* update 21/12/2023 01:58:29 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20014',
	'BankPayment_Controller',
	'get_list',
	'BankPaymentList',
	'Журналы',
	'Платежные поручения',
	FALSE
	);
	
	

-- ******************* update 21/12/2023 02:03:37 ******************
	
		ALTER TABLE public.firms ADD COLUMN bank_acc  varchar(20),ADD COLUMN bank_bik  varchar(9);



-- ******************* update 21/12/2023 09:47:49 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20015',
	'SpecialistPeriodSalaryDetail_Controller',
	'get_for_pay_list',
	'SpecialistPeriodSalaryDetailList',
	'Журналы',
	'Формирование п/п в банк',
	FALSE
	);
	
	

-- ******************* update 21/12/2023 09:50:37 ******************
DELETE FROM views WHERE id= '20015';
	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'20015',
	'SpecialistPeriodSalaryDetail_Controller',
	'get_for_pay_list',
	'SpecialistPeriodSalaryDetailForPayList',
	'Журналы',
	'Формирование п/п в банк',
	FALSE
	);
	
	


-- ******************* update 21/12/2023 15:48:46 ******************

	-- ********** constant value table  specialist_pay_comment_template *************
	CREATE TABLE IF NOT EXISTS const_specialist_pay_comment_template
	(name text  NOT NULL, descr text, val text,
		val_type text,ctrl_class text,ctrl_options json, view_class text,view_options json,
	CONSTRAINT const_specialist_pay_comment_template_pkey PRIMARY KEY (name));
	ALTER TABLE const_specialist_pay_comment_template OWNER TO nails;
	INSERT INTO const_specialist_pay_comment_template (name,descr,val,val_type,ctrl_class,ctrl_options,view_class,view_options) VALUES (
		'Шаблон назначения платежа для самозанятого'
		,'Используется при формировании п/п в банк на оплату самозанятым. Принимает следующие поля: [sum]-сумма, [fio]-ФИО мастера'
		,
			'Оплата самозанятому, сумма [sum], без НДС'
			
		,'Text'
		,NULL
		,NULL
		,NULL
		,NULL
	);
	
		--constant get value
		
	CREATE OR REPLACE FUNCTION const_specialist_pay_comment_template_val()
	RETURNS text AS
	$BODY$
		
		SELECT val::text AS val FROM const_specialist_pay_comment_template LIMIT 1;
		
	$BODY$
	LANGUAGE sql STABLE COST 100;
	ALTER FUNCTION const_specialist_pay_comment_template_val() OWNER TO nails;
	
	--constant set value
	CREATE OR REPLACE FUNCTION const_specialist_pay_comment_template_set_val(Text)
	RETURNS void AS
	$BODY$
		UPDATE const_specialist_pay_comment_template SET val=$1;
	$BODY$
	LANGUAGE sql VOLATILE COST 100;
	ALTER FUNCTION const_specialist_pay_comment_template_set_val(Text) OWNER TO nails;
	
	--edit view: all keys and descr
	CREATE OR REPLACE VIEW const_specialist_pay_comment_template_view AS
	SELECT
		'specialist_pay_comment_template'::text AS id
		,t.name
		,t.descr
	,
	t.val::text AS val
	
	,t.val_type::text AS val_type
	,t.ctrl_class::text
	,t.ctrl_options::json
	,t.view_class::text
	,t.view_options::json
	FROM const_specialist_pay_comment_template AS t
	;
	ALTER VIEW const_specialist_pay_comment_template_view OWNER TO nails;
	
	
	
	CREATE OR REPLACE VIEW constants_list_view AS
	
	SELECT *
	FROM const_doc_per_page_count_view
	UNION ALL
	
	SELECT *
	FROM const_grid_refresh_interval_view
	UNION ALL
	
	SELECT *
	FROM const_email_view
	UNION ALL
	
	SELECT *
	FROM const_join_contract_view
	UNION ALL
	
	SELECT *
	FROM const_person_tax_view
	UNION ALL
	
	SELECT *
	FROM const_specialist_services_view
	UNION ALL
	
	SELECT *
	FROM const_specialist_pay_comment_template_view
	ORDER BY name;
	ALTER VIEW constants_list_view OWNER TO nails;
	
	


-- ******************* update 21/12/2023 16:34:06 ******************
	
		ALTER TABLE public.bank_payments ADD COLUMN payer text, ADD COLUMN rec text;



-- ******************* update 21/12/2023 17:28:18 ******************
	
		ALTER TABLE public.bank_payments ADD COLUMN payer_inn  varchar(12),ADD COLUMN rec_inn  varchar(12);



-- ******************* update 25/12/2023 08:05:40 ******************
	
		ALTER TABLE public.template_batches ADD COLUMN studio_id int REFERENCES studios(id);



-- ******************* update 25/12/2023 08:11:39 ******************
	
		ALTER TABLE public.template_batch_items ADD COLUMN studio_id int REFERENCES studios(id);
		
	DROP INDEX IF EXISTS template_batch_items_template_idx;
	CREATE UNIQUE INDEX template_batch_items_template_idx
	ON template_batch_items(template_batch_id,studio_id,template_id);



-- ******************* update 26/12/2023 10:15:31 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'30001',
	NULL,
	NULL,
	'SpecialistRegTemplateBatch',
	'Формы',
	'Формирование комплекта документов для мастера',
	FALSE
	);
	
	

-- ******************* update 26/12/2023 13:03:57 ******************

	
	-- Adding menu item
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'30002',
	NULL,
	NULL,
	'SpecialistSalaryTemplateBatch',
	'Формы',
	'Формирование комплекта документов по з/ап для мастера',
	FALSE
	);
	
	

-- ******************* update 26/12/2023 13:04:21 ******************

	
	-- Adding menu item
	DELETE FROM views where id = '30002';
	INSERT INTO views
	(id,c,f,t,section,descr,limited)
	VALUES (
	'30002',
	NULL,
	NULL,
	'SpecialistSalaryTemplateBatch',
	'Формы',
	'Формирование комплекта документов по з/п для мастера',
	FALSE
	);
	
	


-- ******************* update 26/12/2023 14:59:02 ******************
	
		ALTER TABLE public.document_templates ADD COLUMN sign_image_name text;