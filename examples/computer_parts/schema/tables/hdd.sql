--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2
-- Dumped by pg_dump version 12.4 (Ubuntu 12.4-1.pgdg18.04+1)

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
-- Name: hdd; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.hdd (
    id bigint NOT NULL,
    make_id bigint NOT NULL,
    name character varying(255) NOT NULL,
    size_bytes bigint NOT NULL,
    rpm bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);


ALTER TABLE public.hdd OWNER TO postgres;

--
-- PostgreSQL database dump complete
--

