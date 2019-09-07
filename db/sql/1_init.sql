-- +migrate Up
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

CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;

COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';


SET default_tablespace = '';

SET default_with_oids = false;

CREATE TABLE public.equipmqnt (
    id integer NOT NULL,
    model_id integer,
    name character varying(200),
    address character varying(2000),
    gate integer
);


ALTER TABLE public.equipmqnt OWNER TO postgres;

CREATE TABLE public.models (
    id integer NOT NULL,
    name character varying(200)
);


ALTER TABLE public.models OWNER TO postgres;

CREATE TABLE public.node (
    id integer NOT NULL,
    model_id integer,
    name character varying(200)
);


ALTER TABLE public.node OWNER TO postgres;

ALTER TABLE ONLY public.equipmqnt
    ADD CONSTRAINT equipmqnt_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.models
    ADD CONSTRAINT models_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.node
    ADD CONSTRAINT node_pkey PRIMARY KEY (id);

-- +migrate Down
DROP TABLE public.equipmqnt;
DROP TABLE public.models;
DROP TABLE public.node;
