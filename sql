-- Database: project_sanber

-- DROP DATABASE IF EXISTS project_sanber;

CREATE DATABASE project_sanber
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_Indonesia.1252'
    LC_CTYPE = 'English_Indonesia.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Table: public.buku

-- DROP TABLE IF EXISTS public.buku;

CREATE TABLE IF NOT EXISTS public.buku
(
    id integer NOT NULL DEFAULT nextval('buku_id_seq'::regclass),
    judul_buku character varying(256) COLLATE pg_catalog."default",
    author character varying(256) COLLATE pg_catalog."default",
    tahun_terbit bigint,
    publisher character varying(256) COLLATE pg_catalog."default",
    isbn character varying(256) COLLATE pg_catalog."default",
    jumlah_halaman character varying(256) COLLATE pg_catalog."default",
    is_dipinjam character varying(256) COLLATE pg_catalog."default",
    created_at character varying(256) COLLATE pg_catalog."default",
    created_by character varying(256) COLLATE pg_catalog."default",
    updated_by character varying(256) COLLATE pg_catalog."default",
    updated_at character varying(256) COLLATE pg_catalog."default",
    CONSTRAINT buku_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.buku
    OWNER to postgres;

-- Table: public.pinjam_buku

-- DROP TABLE IF EXISTS public.pinjam_buku;

CREATE TABLE IF NOT EXISTS public.pinjam_buku
(
    id integer NOT NULL DEFAULT nextval('pinjam_buku_id_seq'::regclass),
    id_buku bigint,
    id_user_peminjam bigint,
    created_by character varying(256) COLLATE pg_catalog."default",
    updated_by character varying(256) COLLATE pg_catalog."default",
    updated_at character varying(256) COLLATE pg_catalog."default",
    CONSTRAINT pinjam_buku_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.pinjam_buku
    OWNER to postgres;

-- Table: public.user_login

-- DROP TABLE IF EXISTS public.user_login;

CREATE TABLE IF NOT EXISTS public.user_login
(
    id integer NOT NULL DEFAULT nextval('user_login_id_seq'::regclass),
    username character varying(256) COLLATE pg_catalog."default",
    nama character varying(256) COLLATE pg_catalog."default",
    role character varying(256) COLLATE pg_catalog."default",
    token character varying(256) COLLATE pg_catalog."default",
    CONSTRAINT user_login_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.user_login
    OWNER to postgres;

-- Table: public.users

-- DROP TABLE IF EXISTS public.users;

CREATE TABLE IF NOT EXISTS public.users
(
    id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    username character varying(256) COLLATE pg_catalog."default",
    password character varying(256) COLLATE pg_catalog."default",
    email character varying(256) COLLATE pg_catalog."default",
    wa character varying(256) COLLATE pg_catalog."default",
    nama character varying(256) COLLATE pg_catalog."default",
    alamat character varying(256) COLLATE pg_catalog."default",
    role character varying(256) COLLATE pg_catalog."default",
    created_by character varying(256) COLLATE pg_catalog."default",
    created_at character varying(256) COLLATE pg_catalog."default",
    updated_by character varying(256) COLLATE pg_catalog."default",
    updated_at character varying(256) COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;