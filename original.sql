--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6
-- Dumped by pg_dump version 14.7 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: temp; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA temp;


ALTER SCHEMA temp OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: distributors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.distributors (
    did numeric(3,0) NOT NULL,
    name character varying(40)
);


ALTER TABLE public.distributors OWNER TO postgres;

--
-- Name: t; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t (
    id integer NOT NULL,
    text text
);


ALTER TABLE public.t OWNER TO postgres;

--
-- Name: t2; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t2 (
    id integer NOT NULL,
    text text NOT NULL,
    created timestamp without time zone DEFAULT now(),
    bool boolean
);

ALTER TABLE ONLY public.t2 REPLICA IDENTITY FULL;


ALTER TABLE public.t2 OWNER TO postgres;

--
-- Name: t2_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t2_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t2_id_seq OWNER TO postgres;

--
-- Name: t2_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t2_id_seq OWNED BY public.t2.id;


--
-- Name: t3; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t3 (
    id integer NOT NULL,
    text text NOT NULL
);


ALTER TABLE public.t3 OWNER TO postgres;

--
-- Name: t3_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t3_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t3_id_seq OWNER TO postgres;

--
-- Name: t3_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t3_id_seq OWNED BY public.t3.id;


--
-- Name: tl; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tl (
    id integer NOT NULL,
    text text
);


ALTER TABLE public.tl OWNER TO postgres;

--
-- Name: t; Type: TABLE; Schema: temp; Owner: postgres
--

CREATE TABLE temp.t (
    id integer NOT NULL,
    text text
);


ALTER TABLE temp.t OWNER TO postgres;

--
-- Name: t2 id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t2 ALTER COLUMN id SET DEFAULT nextval('public.t2_id_seq'::regclass);


--
-- Name: t3 id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t3 ALTER COLUMN id SET DEFAULT nextval('public.t3_id_seq'::regclass);


--
-- Name: distributors distributors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.distributors
    ADD CONSTRAINT distributors_pkey PRIMARY KEY (did);


--
-- Name: t3 t3_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t3
    ADD CONSTRAINT t3_pkey PRIMARY KEY (id, text);


--
-- Name: t t_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t
    ADD CONSTRAINT t_pkey PRIMARY KEY (id);


--
-- Name: tl tl_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tl
    ADD CONSTRAINT tl_pkey PRIMARY KEY (id);


--
-- Name: t t_pkey; Type: CONSTRAINT; Schema: temp; Owner: postgres
--

ALTER TABLE ONLY temp.t
    ADD CONSTRAINT t_pkey PRIMARY KEY (id);


--
-- Name: dbz_publication; Type: PUBLICATION; Schema: -; Owner: postgres
--

ALTER PUBLICATION dbz_publication OWNER TO postgres;

--
-- PostgreSQL database dump complete
--

