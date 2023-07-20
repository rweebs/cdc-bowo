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
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: day_menus; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.day_menus (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    image character varying(255),
    day character varying(255),
    menu_id text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.day_menus OWNER TO postgres;

--
-- Name: menus; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.menus (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    title character varying(255),
    price bigint,
    description character varying(255),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.menus OWNER TO postgres;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    is_morning boolean DEFAULT false,
    is_noon boolean DEFAULT false,
    is_afternoon boolean DEFAULT false,
    menu_id uuid,
    count bigint,
    remaining bigint DEFAULT 0,
    status text,
    amount bigint,
    address character varying(255),
    user_id bigint,
    lat numeric(10,8),
    lng numeric(11,8),
    upload character varying(255),
    start_date text,
    end_date text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- Name: day_menus day_menus_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.day_menus
    ADD CONSTRAINT day_menus_pkey PRIMARY KEY (id);


--
-- Name: menus menus_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.menus
    ADD CONSTRAINT menus_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: idx_day_menus_day; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_day_menus_day ON public.day_menus USING btree (day);


--
-- Name: idx_day_menus_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_day_menus_id ON public.day_menus USING btree (id);


--
-- Name: idx_menus_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_menus_id ON public.menus USING btree (id);


--
-- Name: idx_menus_title; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_menus_title ON public.menus USING btree (title);


--
-- Name: idx_transactions_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_transactions_id ON public.transactions USING btree (id);


--
-- Name: transactions fk_transactions_menu; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_menu FOREIGN KEY (menu_id) REFERENCES public.menus(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

