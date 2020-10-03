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
-- Name: twitter_tweets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.twitter_tweets (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    text character varying NOT NULL,
    source character varying(24) NOT NULL,
    in_reply_to_status_id bigint,
    in_reply_to_user_id bigint,
    in_reply_to_screen_name character varying(24) NOT NULL,
    quoted_status_id bigint,
    is_quote_status boolean NOT NULL,
    twitter_data jsonb NOT NULL,
    twitter_created_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);


ALTER TABLE public.twitter_tweets OWNER TO postgres;

--
-- PostgreSQL database dump complete
--

