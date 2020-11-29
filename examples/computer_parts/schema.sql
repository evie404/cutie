DROP TABLE IF EXISTS make;

CREATE TABLE make (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,

    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);

DROP TABLE IF EXISTS cpu;

CREATE TABLE cpu (
    id bigint NOT NULL,
    make_id bigint NOT NULL,
    name character varying(255) NOT NULL,
    cores int NOT NULL,
    clock_speed_ghz DECIMAL(4,2) NOT NULL,

    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);

DROP TABLE IF EXISTS harddrives;

CREATE TABLE harddrives (
    id bigint NOT NULL,
    make_id bigint NOT NULL,
    name character varying(255) NOT NULL,
    size_bytes bigint NOT NULL,
    rpm bigint NOT NULL,

    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);
