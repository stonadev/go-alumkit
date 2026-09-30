--
-- 000002_sessions_scs.up.sql
--
-- Replaces the Laravel-era sessions table (id/payload/last_activity)
-- with the schema expected by github.com/alexedwards/scs/v2
-- postgresstore: token text PK, data bytea, expiry timestamptz.
--
-- Existing rows are PHP-serialized session blobs that Go cannot
-- decode, so the table is dropped rather than converted. Sessions are
-- ephemeral; users log in again once after cutover.
--

DROP TABLE IF EXISTS sessions;

CREATE TABLE sessions (
    token  text PRIMARY KEY,
    data   bytea,
    expiry timestamptz NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);
