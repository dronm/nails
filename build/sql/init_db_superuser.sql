-- SUPERUSER CODE
CREATE USER nails WITH PASSWORD '221d207e625043a9889e45d7b4a5a9e1';
CREATE DATABASE nails OWNER nails;
GRANT ALL PRIVILEGES ON DATABASE nails TO nails;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO nails;
grant pg_execute_server_program to nails;
CREATE EXTENSION pgcrypto;

CREATE USER wa_nails WITH PASSWORD 'ebaf2871600047748b2eb66e1fa0652e';
CREATE DATABASE wa_nails OWNER wa_nails;
GRANT ALL PRIVILEGES ON DATABASE wa_nails TO wa_nails;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO wa_nails;

pg_restore --clean -h 192.168.1.177 -U nails -d nails nails.bckp 
