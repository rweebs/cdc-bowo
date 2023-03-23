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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: t; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t (
    id integer NOT NULL,
    message text,
    created_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.t OWNER TO postgres;

--
-- Name: t_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_id_seq OWNER TO postgres;

--
-- Name: t_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_id_seq OWNED BY public.t.id;


--
-- Name: test_numeric; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.test_numeric (
    id integer NOT NULL,
    "numeric" numeric(5,2),
    "decimal" numeric,
    "real" real,
    double_precision double precision
);


ALTER TABLE public.test_numeric OWNER TO postgres;

--
-- Name: test_numeric_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.test_numeric_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.test_numeric_id_seq OWNER TO postgres;

--
-- Name: test_numeric_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.test_numeric_id_seq OWNED BY public.test_numeric.id;


--
-- Name: test_temporal; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.test_temporal (
    id integer NOT NULL,
    "time" time without time zone DEFAULT now(),
    time_mili time(3) without time zone DEFAULT now(),
    time_with_time_zone time with time zone DEFAULT now(),
    "timestamp" timestamp without time zone DEFAULT now(),
    timestamp_mili timestamp(3) without time zone DEFAULT now(),
    timestamp_with_time_zone timestamp with time zone DEFAULT now(),
    date date DEFAULT now()
);


ALTER TABLE public.test_temporal OWNER TO postgres;

--
-- Name: test_temporal_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.test_temporal_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.test_temporal_id_seq OWNER TO postgres;

--
-- Name: test_temporal_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.test_temporal_id_seq OWNED BY public.test_temporal.id;


--
-- Name: t id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t ALTER COLUMN id SET DEFAULT nextval('public.t_id_seq'::regclass);


--
-- Name: test_numeric id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_numeric ALTER COLUMN id SET DEFAULT nextval('public.test_numeric_id_seq'::regclass);


--
-- Name: test_temporal id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_temporal ALTER COLUMN id SET DEFAULT nextval('public.test_temporal_id_seq'::regclass);


--
-- Data for Name: t; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t (id, message, created_at) FROM stdin;
270076	test	2023-02-08 07:53:03.518356
270077	test	2023-02-08 08:19:41.816962
\.


--
-- Data for Name: test_numeric; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.test_numeric (id, "numeric", "decimal", "real", double_precision) FROM stdin;
1	1.10	1.1	1.1	1.1
2	1.10	1.1	1.1	1.1
3	1.10	1.1	1.1	1.1
4	1.10	1.1	1.1	1.1
5	1.10	1.1	1.1	1.1
6	1.11	1.1	1.1	1.1
7	1.10	1.1	1.1	1.1
8	1.11	1.11111	1.11111	1.11111
9	1.11	1.11111	1.11111	1.11111
10	1.11	1.11111	1.11111	NaN
11	1.11	1.11111	1.11111	Infinity
12	1.11	1.11111	1.11111	Infinity
13	1.11	1.11111	1.11111	Infinity
14	1.11	1.11111	1.11111	Infinity
15	1.11	1.11111	1.11111	Infinity
16	1.11	1.11111	1.11111	Infinity
17	1.11	1.11111	1.11111	Infinity
18	1.11	1.11111	1.11111	Infinity
\.


--
-- Data for Name: test_temporal; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.test_temporal (id, "time", time_mili, time_with_time_zone, "timestamp", timestamp_mili, timestamp_with_time_zone, date) FROM stdin;
44	07:05:48.000129	07:05:48	02:47:09.863519+00	2023-02-10 07:05:48.000129	2023-02-10 07:05:48	2023-02-10 07:05:48.000129+00	2023-02-10
46	06:50:36.14781	06:50:36.148	06:50:36.14781+00	2023-02-23 06:50:36.14781	2023-02-23 06:50:36.148	2023-02-23 06:50:36.14781+00	2023-02-23
47	06:51:03.72789	06:51:03.728	06:51:03.72789+00	2023-02-23 06:51:03.72789	2023-02-23 06:51:03.728	2023-02-23 06:51:03.72789+00	2023-02-23
48	06:51:14.250308	06:51:14.25	06:51:14.250308+00	2023-02-23 06:51:14.250308	2023-02-23 06:51:14.25	2023-02-23 06:51:14.250308+00	2023-02-23
49	06:53:12.443598	06:53:12.444	06:53:12.443598+00	2023-02-23 06:53:12.443598	2023-02-23 06:53:12.444	2023-02-23 06:53:12.443598+00	2023-02-23
51	07:08:21.490913	07:08:21.491	07:08:21.490913+00	2023-02-23 07:08:21.490913	2023-02-23 07:08:21.491	2023-02-23 07:08:21.490913+00	2023-02-23
52	12:19:05.904239	12:19:05.904	12:19:05.904239+00	2023-02-25 12:19:05.904239	2023-02-25 12:19:05.904	2023-02-25 12:19:05.904239+00	2023-02-25
53	12:20:33.522694	12:20:33.523	12:20:33.522694+00	2023-02-25 12:20:33.522694	2023-02-25 12:20:33.523	2023-02-25 12:20:33.522694+00	2023-02-25
54	01:01:19.271221	01:01:19.271	01:01:19.271221+00	2023-02-27 01:01:19.271221	2023-02-27 01:01:19.271	2023-02-27 01:01:19.271221+00	2023-02-27
55	01:01:27.802982	01:01:27.803	01:01:27.802982+00	2023-02-27 01:01:27.802982	2023-02-27 01:01:27.803	2023-02-27 01:01:27.802982+00	2023-02-27
56	01:01:30.222598	01:01:30.223	01:01:30.222598+00	2023-02-27 01:01:30.222598	2023-02-27 01:01:30.223	2023-02-27 01:01:30.222598+00	2023-02-27
57	01:01:31.189841	01:01:31.19	01:01:31.189841+00	2023-02-27 01:01:31.189841	2023-02-27 01:01:31.19	2023-02-27 01:01:31.189841+00	2023-02-27
58	01:01:31.984726	01:01:31.985	01:01:31.984726+00	2023-02-27 01:01:31.984726	2023-02-27 01:01:31.985	2023-02-27 01:01:31.984726+00	2023-02-27
59	01:39:25.133019	01:39:25.133	01:39:25.133019+00	2023-02-27 01:39:25.133019	2023-02-27 01:39:25.133	2023-02-27 01:39:25.133019+00	2023-02-27
60	07:19:04.38374	07:19:04.384	07:19:04.38374+00	2023-02-27 07:19:04.38374	2023-02-27 07:19:04.384	2023-02-27 07:19:04.38374+00	2023-02-27
45	07:21:22.798069	07:21:22.798	02:47:09.863519+00	2023-02-10 07:21:22.798069	2023-02-10 07:21:22.798	2023-02-10 07:21:22.798069+00	2023-02-10
\.


--
-- Name: t_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_id_seq', 270077, true);


--
-- Name: test_numeric_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.test_numeric_id_seq', 18, true);


--
-- Name: test_temporal_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.test_temporal_id_seq', 60, true);


--
-- Name: t t_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t
    ADD CONSTRAINT t_pkey PRIMARY KEY (created_at);


--
-- Name: test_numeric test_numeric_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_numeric
    ADD CONSTRAINT test_numeric_pkey PRIMARY KEY (id);


--
-- Name: test_temporal test_temporal_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_temporal
    ADD CONSTRAINT test_temporal_pkey PRIMARY KEY (id);


--
-- Name: dbz_publication; Type: PUBLICATION; Schema: -; Owner: postgres
--

CREATE PUBLICATION dbz_publication FOR ALL TABLES WITH (publish = 'insert, update, delete, truncate');


ALTER PUBLICATION dbz_publication OWNER TO postgres;

--
-- Name: pub; Type: PUBLICATION; Schema: -; Owner: postgres
--

CREATE PUBLICATION pub FOR ALL TABLES WITH (publish = 'insert, update, delete, truncate');


ALTER PUBLICATION pub OWNER TO postgres;

--
-- PostgreSQL database dump complete
--

