--
-- 000001_baseline.up.sql
--
-- Baseline schema export from the existing PHP production database
-- (alumkit-go). All 23 production tables. This migration only runs on a
-- fresh database; new schema changes go in 000002+ migrations.
--
-- Regenerate with:
--   pg_dump --schema-only -U <user> alumkit-go
-- then apply the idempotency transforms (IF NOT EXISTS guards; strip
-- role-specific ownership/grant lines and the search_path configuration
-- line) before committing.
--
--
-- PostgreSQL database dump
--


-- Dumped from database version 16.14 (Homebrew)
-- Dumped by pg_dump version 16.14 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: activity_log; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.activity_log (
    id bigint NOT NULL,
    log_name character varying(255),
    description text NOT NULL,
    subject_type character varying(255),
    subject_id bigint,
    causer_type character varying(255),
    causer_id bigint,
    event character varying(255),
    properties json,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    batch_uuid uuid
);



--
-- Name: activity_log_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.activity_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: activity_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.activity_log_id_seq OWNED BY public.activity_log.id;


--
-- Name: cache; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.cache (
    key character varying(255) NOT NULL,
    value text NOT NULL,
    expiration bigint NOT NULL
);



--
-- Name: cache_locks; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.cache_locks (
    key character varying(255) NOT NULL,
    owner character varying(255) NOT NULL,
    expiration bigint NOT NULL
);



--
-- Name: careers; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.careers (
    id bigint NOT NULL,
    profile_id bigint NOT NULL,
    job_title character varying(255) NOT NULL,
    company character varying(255) NOT NULL,
    employment_type character varying(255) NOT NULL,
    industry character varying(255),
    location character varying(255),
    start_year integer NOT NULL,
    start_month smallint,
    is_current boolean DEFAULT false NOT NULL,
    end_year integer,
    end_month smallint,
    description text,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);



--
-- Name: careers_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.careers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: careers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.careers_id_seq OWNED BY public.careers.id;


--
-- Name: committee_members; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.committee_members (
    id bigint NOT NULL,
    position_id bigint,
    user_id bigint,
    name character varying(255),
    photo_path character varying(255),
    sort_order integer DEFAULT 0 NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);



--
-- Name: committee_members_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.committee_members_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: committee_members_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.committee_members_id_seq OWNED BY public.committee_members.id;


--
-- Name: contents; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.contents (
    id bigint NOT NULL,
    owner character varying(255) NOT NULL,
    type character varying(255) NOT NULL,
    fields json NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);



--
-- Name: contents_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.contents_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: contents_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.contents_id_seq OWNED BY public.contents.id;


--
-- Name: educations; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.educations (
    id bigint NOT NULL,
    profile_id bigint NOT NULL,
    level character varying(255) NOT NULL,
    institution character varying(255) NOT NULL,
    student_id character varying(255),
    subject character varying(255) NOT NULL,
    is_current boolean DEFAULT false NOT NULL,
    start_year integer NOT NULL,
    start_month smallint,
    end_year integer,
    end_month smallint,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);



--
-- Name: educations_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.educations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: educations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.educations_id_seq OWNED BY public.educations.id;


--
-- Name: failed_jobs; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.failed_jobs (
    id bigint NOT NULL,
    uuid character varying(255) NOT NULL,
    connection character varying(255) NOT NULL,
    queue character varying(255) NOT NULL,
    payload text NOT NULL,
    exception text NOT NULL,
    failed_at timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);



--
-- Name: failed_jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.failed_jobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: failed_jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.failed_jobs_id_seq OWNED BY public.failed_jobs.id;


--
-- Name: job_batches; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.job_batches (
    id character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    total_jobs integer NOT NULL,
    pending_jobs integer NOT NULL,
    failed_jobs integer NOT NULL,
    failed_job_ids text NOT NULL,
    options text,
    cancelled_at integer,
    created_at integer NOT NULL,
    finished_at integer
);



--
-- Name: jobs; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.jobs (
    id bigint NOT NULL,
    queue character varying(255) NOT NULL,
    payload text NOT NULL,
    attempts smallint NOT NULL,
    reserved_at integer,
    available_at integer NOT NULL,
    created_at integer NOT NULL
);



--
-- Name: jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.jobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.jobs_id_seq OWNED BY public.jobs.id;


--
-- Name: migrations; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.migrations (
    id integer NOT NULL,
    migration character varying(255) NOT NULL,
    batch integer NOT NULL
);



--
-- Name: migrations_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.migrations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: migrations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.migrations_id_seq OWNED BY public.migrations.id;


--
-- Name: model_has_permissions; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.model_has_permissions (
    permission_id bigint NOT NULL,
    model_type character varying(255) NOT NULL,
    model_id bigint NOT NULL
);



--
-- Name: model_has_roles; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.model_has_roles (
    role_id bigint NOT NULL,
    model_type character varying(255) NOT NULL,
    model_id bigint NOT NULL
);



--
-- Name: pages; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.pages (
    id uuid NOT NULL,
    title character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    meta_title character varying(255),
    meta_description text,
    is_published boolean DEFAULT false NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);



--
-- Name: password_reset_tokens; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.password_reset_tokens (
    email character varying(255) NOT NULL,
    token character varying(255) NOT NULL,
    created_at timestamp(0) without time zone
);



--
-- Name: permissions; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.permissions (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    guard_name character varying(255) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);



--
-- Name: permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.permissions_id_seq OWNED BY public.permissions.id;


--
-- Name: positions; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.positions (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);



--
-- Name: positions_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.positions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: positions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.positions_id_seq OWNED BY public.positions.id;


--
-- Name: posts; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.posts (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    title character varying(255) NOT NULL,
    body text NOT NULL,
    thumbnail character varying(255),
    published_at timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);



--
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.posts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;


--
-- Name: profiles; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.profiles (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    photo_path character varying(255),
    date_of_birth date,
    gender character varying(255),
    blood_group character varying(255),
    present_address character varying(255),
    permanent_address character varying(255),
    social_links json,
    website character varying(255),
    emergency_contact json,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    local_names json
);



--
-- Name: profiles_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.profiles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: profiles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.profiles_id_seq OWNED BY public.profiles.id;


--
-- Name: role_has_permissions; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.role_has_permissions (
    permission_id bigint NOT NULL,
    role_id bigint NOT NULL
);



--
-- Name: roles; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.roles (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    guard_name character varying(255) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);



--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: sessions; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.sessions (
    id character varying(255) NOT NULL,
    user_id bigint,
    ip_address character varying(45),
    user_agent text,
    payload text NOT NULL,
    last_activity integer NOT NULL
);



--
-- Name: users; Type: TABLE; Schema: public; Owner: baaku
--

CREATE TABLE IF NOT EXISTS public.users (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    state character varying(255) DEFAULT 'unverified'::character varying NOT NULL,
    email_verified_at timestamp(0) without time zone,
    phone character varying(255),
    phone_verified_at timestamp(0) without time zone,
    password character varying(255) NOT NULL,
    remember_token character varying(100),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    two_factor_secret text,
    two_factor_recovery_codes text,
    two_factor_confirmed_at timestamp(0) without time zone
);



--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: baaku
--

CREATE SEQUENCE IF NOT EXISTS public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baaku
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: activity_log id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.activity_log ALTER COLUMN id SET DEFAULT nextval('public.activity_log_id_seq'::regclass);


--
-- Name: careers id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.careers ALTER COLUMN id SET DEFAULT nextval('public.careers_id_seq'::regclass);


--
-- Name: committee_members id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.committee_members ALTER COLUMN id SET DEFAULT nextval('public.committee_members_id_seq'::regclass);


--
-- Name: contents id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.contents ALTER COLUMN id SET DEFAULT nextval('public.contents_id_seq'::regclass);


--
-- Name: educations id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.educations ALTER COLUMN id SET DEFAULT nextval('public.educations_id_seq'::regclass);


--
-- Name: failed_jobs id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.failed_jobs ALTER COLUMN id SET DEFAULT nextval('public.failed_jobs_id_seq'::regclass);


--
-- Name: jobs id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.jobs ALTER COLUMN id SET DEFAULT nextval('public.jobs_id_seq'::regclass);


--
-- Name: migrations id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.migrations ALTER COLUMN id SET DEFAULT nextval('public.migrations_id_seq'::regclass);


--
-- Name: permissions id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.permissions ALTER COLUMN id SET DEFAULT nextval('public.permissions_id_seq'::regclass);


--
-- Name: positions id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.positions ALTER COLUMN id SET DEFAULT nextval('public.positions_id_seq'::regclass);


--
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);


--
-- Name: profiles id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.profiles ALTER COLUMN id SET DEFAULT nextval('public.profiles_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: baaku
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: activity_log activity_log_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'activity_log_pkey') THEN
ALTER TABLE ONLY public.activity_log
    ADD CONSTRAINT activity_log_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: cache_locks cache_locks_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'cache_locks_pkey') THEN
ALTER TABLE ONLY public.cache_locks
    ADD CONSTRAINT cache_locks_pkey PRIMARY KEY (key);
    END IF;
END $alumkit_baseline$;


--
-- Name: cache cache_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'cache_pkey') THEN
ALTER TABLE ONLY public.cache
    ADD CONSTRAINT cache_pkey PRIMARY KEY (key);
    END IF;
END $alumkit_baseline$;


--
-- Name: careers careers_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'careers_pkey') THEN
ALTER TABLE ONLY public.careers
    ADD CONSTRAINT careers_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: committee_members committee_members_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'committee_members_pkey') THEN
ALTER TABLE ONLY public.committee_members
    ADD CONSTRAINT committee_members_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: contents contents_owner_type_unique; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'contents_owner_type_unique') THEN
ALTER TABLE ONLY public.contents
    ADD CONSTRAINT contents_owner_type_unique UNIQUE (owner, type);
    END IF;
END $alumkit_baseline$;


--
-- Name: contents contents_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'contents_pkey') THEN
ALTER TABLE ONLY public.contents
    ADD CONSTRAINT contents_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: educations educations_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'educations_pkey') THEN
ALTER TABLE ONLY public.educations
    ADD CONSTRAINT educations_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: failed_jobs failed_jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'failed_jobs_pkey') THEN
ALTER TABLE ONLY public.failed_jobs
    ADD CONSTRAINT failed_jobs_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: failed_jobs failed_jobs_uuid_unique; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'failed_jobs_uuid_unique') THEN
ALTER TABLE ONLY public.failed_jobs
    ADD CONSTRAINT failed_jobs_uuid_unique UNIQUE (uuid);
    END IF;
END $alumkit_baseline$;


--
-- Name: job_batches job_batches_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'job_batches_pkey') THEN
ALTER TABLE ONLY public.job_batches
    ADD CONSTRAINT job_batches_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: jobs jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'jobs_pkey') THEN
ALTER TABLE ONLY public.jobs
    ADD CONSTRAINT jobs_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: migrations migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'migrations_pkey') THEN
ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: model_has_permissions model_has_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'model_has_permissions_pkey') THEN
ALTER TABLE ONLY public.model_has_permissions
    ADD CONSTRAINT model_has_permissions_pkey PRIMARY KEY (permission_id, model_id, model_type);
    END IF;
END $alumkit_baseline$;


--
-- Name: model_has_roles model_has_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'model_has_roles_pkey') THEN
ALTER TABLE ONLY public.model_has_roles
    ADD CONSTRAINT model_has_roles_pkey PRIMARY KEY (role_id, model_id, model_type);
    END IF;
END $alumkit_baseline$;


--
-- Name: pages pages_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'pages_pkey') THEN
ALTER TABLE ONLY public.pages
    ADD CONSTRAINT pages_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: pages pages_slug_unique; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'pages_slug_unique') THEN
ALTER TABLE ONLY public.pages
    ADD CONSTRAINT pages_slug_unique UNIQUE (slug);
    END IF;
END $alumkit_baseline$;


--
-- Name: password_reset_tokens password_reset_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'password_reset_tokens_pkey') THEN
ALTER TABLE ONLY public.password_reset_tokens
    ADD CONSTRAINT password_reset_tokens_pkey PRIMARY KEY (email);
    END IF;
END $alumkit_baseline$;


--
-- Name: permissions permissions_name_guard_name_unique; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'permissions_name_guard_name_unique') THEN
ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_name_guard_name_unique UNIQUE (name, guard_name);
    END IF;
END $alumkit_baseline$;


--
-- Name: permissions permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'permissions_pkey') THEN
ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: positions positions_name_unique; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'positions_name_unique') THEN
ALTER TABLE ONLY public.positions
    ADD CONSTRAINT positions_name_unique UNIQUE (name);
    END IF;
END $alumkit_baseline$;


--
-- Name: positions positions_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'positions_pkey') THEN
ALTER TABLE ONLY public.positions
    ADD CONSTRAINT positions_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'posts_pkey') THEN
ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: profiles profiles_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'profiles_pkey') THEN
ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: profiles profiles_user_id_unique; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'profiles_user_id_unique') THEN
ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_user_id_unique UNIQUE (user_id);
    END IF;
END $alumkit_baseline$;


--
-- Name: role_has_permissions role_has_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'role_has_permissions_pkey') THEN
ALTER TABLE ONLY public.role_has_permissions
    ADD CONSTRAINT role_has_permissions_pkey PRIMARY KEY (permission_id, role_id);
    END IF;
END $alumkit_baseline$;


--
-- Name: roles roles_name_guard_name_unique; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'roles_name_guard_name_unique') THEN
ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_name_guard_name_unique UNIQUE (name, guard_name);
    END IF;
END $alumkit_baseline$;


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'roles_pkey') THEN
ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'sessions_pkey') THEN
ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: users users_email_unique; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'users_email_unique') THEN
ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_unique UNIQUE (email);
    END IF;
END $alumkit_baseline$;


--
-- Name: users users_phone_unique; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'users_phone_unique') THEN
ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_phone_unique UNIQUE (phone);
    END IF;
END $alumkit_baseline$;


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'users_pkey') THEN
ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
    END IF;
END $alumkit_baseline$;


--
-- Name: activity_log_log_name_index; Type: INDEX; Schema: public; Owner: baaku
--

CREATE INDEX IF NOT EXISTS activity_log_log_name_index ON public.activity_log USING btree (log_name);


--
-- Name: cache_expiration_index; Type: INDEX; Schema: public; Owner: baaku
--

CREATE INDEX IF NOT EXISTS cache_expiration_index ON public.cache USING btree (expiration);


--
-- Name: cache_locks_expiration_index; Type: INDEX; Schema: public; Owner: baaku
--

CREATE INDEX IF NOT EXISTS cache_locks_expiration_index ON public.cache_locks USING btree (expiration);


--
-- Name: causer; Type: INDEX; Schema: public; Owner: baaku
--

CREATE INDEX IF NOT EXISTS causer ON public.activity_log USING btree (causer_type, causer_id);


--
-- Name: contents_owner_index; Type: INDEX; Schema: public; Owner: baaku
--

CREATE INDEX IF NOT EXISTS contents_owner_index ON public.contents USING btree (owner);


--
-- Name: contents_type_index; Type: INDEX; Schema: public; Owner: baaku
--

CREATE INDEX IF NOT EXISTS contents_type_index ON public.contents USING btree (type);


--
-- Name: failed_jobs_connection_queue_failed_at_index; Type: INDEX; Schema: public; Owner: baaku
--

CREATE INDEX IF NOT EXISTS failed_jobs_connection_queue_failed_at_index ON public.failed_jobs USING btree (connection, queue, failed_at);


--
-- Name: jobs_queue_index; Type: INDEX; Schema: public; Owner: baaku
--

CREATE INDEX IF NOT EXISTS jobs_queue_index ON public.jobs USING btree (queue);


--
-- Name: model_has_permissions_model_id_model_type_index; Type: INDEX; Schema: public; Owner: baaku
--

CREATE INDEX IF NOT EXISTS model_has_permissions_model_id_model_type_index ON public.model_has_permissions USING btree (model_id, model_type);


--
-- Name: model_has_roles_model_id_model_type_index; Type: INDEX; Schema: public; Owner: baaku
--

CREATE INDEX IF NOT EXISTS model_has_roles_model_id_model_type_index ON public.model_has_roles USING btree (model_id, model_type);


--
-- Name: sessions_last_activity_index; Type: INDEX; Schema: public; Owner: baaku
--

CREATE INDEX IF NOT EXISTS sessions_last_activity_index ON public.sessions USING btree (last_activity);


--
-- Name: sessions_user_id_index; Type: INDEX; Schema: public; Owner: baaku
--

CREATE INDEX IF NOT EXISTS sessions_user_id_index ON public.sessions USING btree (user_id);


--
-- Name: subject; Type: INDEX; Schema: public; Owner: baaku
--

CREATE INDEX IF NOT EXISTS subject ON public.activity_log USING btree (subject_type, subject_id);


--
-- Name: careers careers_profile_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'careers_profile_id_foreign') THEN
ALTER TABLE ONLY public.careers
    ADD CONSTRAINT careers_profile_id_foreign FOREIGN KEY (profile_id) REFERENCES public.profiles(id) ON DELETE CASCADE;
    END IF;
END $alumkit_baseline$;


--
-- Name: committee_members committee_members_position_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'committee_members_position_id_foreign') THEN
ALTER TABLE ONLY public.committee_members
    ADD CONSTRAINT committee_members_position_id_foreign FOREIGN KEY (position_id) REFERENCES public.positions(id) ON DELETE SET NULL;
    END IF;
END $alumkit_baseline$;


--
-- Name: committee_members committee_members_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'committee_members_user_id_foreign') THEN
ALTER TABLE ONLY public.committee_members
    ADD CONSTRAINT committee_members_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL;
    END IF;
END $alumkit_baseline$;


--
-- Name: educations educations_profile_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'educations_profile_id_foreign') THEN
ALTER TABLE ONLY public.educations
    ADD CONSTRAINT educations_profile_id_foreign FOREIGN KEY (profile_id) REFERENCES public.profiles(id) ON DELETE CASCADE;
    END IF;
END $alumkit_baseline$;


--
-- Name: model_has_permissions model_has_permissions_permission_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'model_has_permissions_permission_id_foreign') THEN
ALTER TABLE ONLY public.model_has_permissions
    ADD CONSTRAINT model_has_permissions_permission_id_foreign FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE;
    END IF;
END $alumkit_baseline$;


--
-- Name: model_has_roles model_has_roles_role_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'model_has_roles_role_id_foreign') THEN
ALTER TABLE ONLY public.model_has_roles
    ADD CONSTRAINT model_has_roles_role_id_foreign FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE;
    END IF;
END $alumkit_baseline$;


--
-- Name: posts posts_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'posts_user_id_foreign') THEN
ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
    END IF;
END $alumkit_baseline$;


--
-- Name: profiles profiles_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'profiles_user_id_foreign') THEN
ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
    END IF;
END $alumkit_baseline$;


--
-- Name: role_has_permissions role_has_permissions_permission_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'role_has_permissions_permission_id_foreign') THEN
ALTER TABLE ONLY public.role_has_permissions
    ADD CONSTRAINT role_has_permissions_permission_id_foreign FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE;
    END IF;
END $alumkit_baseline$;


--
-- Name: role_has_permissions role_has_permissions_role_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: baaku
--

DO $alumkit_baseline$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'role_has_permissions_role_id_foreign') THEN
ALTER TABLE ONLY public.role_has_permissions
    ADD CONSTRAINT role_has_permissions_role_id_foreign FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE;
    END IF;
END $alumkit_baseline$;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: pg_database_owner
--



--
-- PostgreSQL database dump complete
--


