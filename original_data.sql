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
-- Data for Name: distributors; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.distributors (did, name) FROM stdin;
3	test
4	test
5	test
6	test
\.


--
-- Data for Name: t; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t (id, text) FROM stdin;
1	test
2	test
3	test
4	test
\.


--
-- Data for Name: t2; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t2 (id, text, created, bool) FROM stdin;
32260	390ae0cb748a169b4a56b741d8d8bab7	2023-02-02 01:08:21.140309	t
32261	3d7c2ecab057b712c871db479f515942	2023-02-02 01:09:03.546782	t
32262	94fb6c6f6a13ff1b65f12f5c811ba092	2023-02-02 01:11:16.803443	t
32263	3d0d34138e906b3dbe7ab614f5f5710f	2023-02-02 01:14:20.108846	t
32264	1e205843d5348e8c6a3498fbae6a9225	2023-02-02 01:17:14.26728	t
32265	9c142ded1efde3e8d2e4a813dd5336cd	2023-02-02 01:18:37.926224	t
32266	fb6f0a0815d86909b53bae85c97282e9	2023-02-02 01:20:59.412283	t
32267	4b254c183c8d860380b5a9aea6a7e65b	2023-02-02 01:25:50.075218	t
32268	19db0b0b0e82d85048c892fb54e1513b	2023-02-02 01:26:05.6985	t
32269	0b9ac5c3b3cc5256bd069e3d612a74b2	2023-02-02 01:26:39.075757	t
32270	ac29d06ad9acde764a9706a283964c17	2023-02-02 01:28:53.41117	t
32271	d91fff94520c33b56bd17d8371354a8e	2023-02-02 01:29:07.300012	t
32272	9a099b80d32da11624379ae2810597c4	2023-02-02 01:30:32.642846	t
32273	73f49dc326b0623f1c45ffd31bf2fcc3	2023-02-02 03:09:38.626323	t
32274	f09021588c71c25b786fac30319d08f2	2023-02-02 03:25:02.982311	t
32275	bcf063fc5047e8fef9cbecc1a94d3623	2023-02-02 03:50:25.210219	t
32276	9cbf7c5391ac1ec9392274e47bb5fd4b	2023-02-02 03:50:37.893142	t
\.


--
-- Data for Name: t3; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t3 (id, text) FROM stdin;
3	34ab636b773ecef74fcc3acd02d98a1d
4	c061b1804aa29249b78591e9378ea459
5	2d01671c3dcc4de9613c98e33a84ad7b
6	d87b23d8d8240d0329da13723ca7ee45
7	f7d7f30c94c078074ef4c752f12464e1
8	a85dbf4ba0e635f24f999422c8b91a4a
9	380d4f81cf9d2d19d19cd4638505c40a
10	bfa8298342e2def29fead94c71ff5b38
11	cb48c51c2e18e160f17ab27ceb679f31
12	8f58404ccf061996b9836aec4ec8b433
13	09d0e80c640bb264596469dc61e0cc59
14	d51a06936d287f720e15d491760a7733
15	550c78e4df7090f3d83ec5091e6fae30
16	f2dbc8fe1315d6b977f8ea2380c6ef73
17	5ca0e7086fb0a9e89287519ffb74f860
18	59cd0de7c04eaa1335cd715515d43cc2
19	072c32536359fe9b77e75bf5ff2ea7df
20	7a4722557ae5b006125d1d54e3875c1f
21	170d2ec115a1d3a24afba3eb5e6b6d73
22	5834a9e949040ab48ba13f6dac7cffb8
23	15c3b21acb0bd57217707e390053e866
24	7cc2449d8c5cf46a14565f04124f4824
25	1704393763602498a495a2db36a499b6
26	099cb96767ff42f2df1eaafa359775db
27	44ff285f58965ee1bdb31ea0f4528e31
28	0185a663f03203d1119e6eaa13fce573
29	e0ebaf587fc35a637b5c5c9e3c14d3d2
30	3b8b5654aa93a19f47881a60ecdc95d2
31	29eb6ca019e882c7e38bbe3fc1d4d9e1
\.


--
-- Data for Name: tl; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tl (id, text) FROM stdin;
6	test
7	test
8	test
9	test
\.


--
-- Data for Name: t; Type: TABLE DATA; Schema: temp; Owner: postgres
--

COPY temp.t (id, text) FROM stdin;
1	test
2	test
4	test
\.


--
-- Name: t2_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t2_id_seq', 32276, true);


--
-- Name: t3_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t3_id_seq', 31, true);


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

CREATE PUBLICATION dbz_publication FOR ALL TABLES WITH (publish = 'insert, update, delete, truncate');


ALTER PUBLICATION dbz_publication OWNER TO postgres;

--
-- PostgreSQL database dump complete
--

