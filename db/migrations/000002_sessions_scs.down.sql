--
-- 000002_sessions_scs.down.sql
--
-- Restores the Laravel-era sessions table shape from the baseline
-- (000001). Sessions written by scs are dropped; they are ephemeral.
--

DROP TABLE IF EXISTS sessions;

CREATE TABLE sessions (
    id            character varying(255) NOT NULL,
    user_id       bigint,
    ip_address    character varying(45),
    user_agent    text,
    payload       text NOT NULL,
    last_activity integer NOT NULL,
    CONSTRAINT sessions_pkey PRIMARY KEY (id)
);

CREATE INDEX sessions_last_activity_index ON sessions (last_activity);
CREATE INDEX sessions_user_id_index ON sessions (user_id);
