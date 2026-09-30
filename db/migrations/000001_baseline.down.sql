--
-- 000001_baseline.down.sql
--
-- Drops the baseline schema. Order respects foreign key dependencies.
-- The golang-migrate bookkeeping table (schema_migrations) is managed by
-- the migrator itself and is intentionally not dropped here.
--

DROP TABLE IF EXISTS activity_log;
DROP TABLE IF EXISTS contents;
DROP TABLE IF EXISTS pages;
DROP TABLE IF EXISTS committee_members;
DROP TABLE IF EXISTS positions;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS careers;
DROP TABLE IF EXISTS educations;
DROP TABLE IF EXISTS profiles;
DROP TABLE IF EXISTS password_reset_tokens;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS role_has_permissions;
DROP TABLE IF EXISTS model_has_roles;
DROP TABLE IF EXISTS model_has_permissions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS jobs;
DROP TABLE IF EXISTS job_batches;
DROP TABLE IF EXISTS failed_jobs;
DROP TABLE IF EXISTS cache_locks;
DROP TABLE IF EXISTS cache;
DROP TABLE IF EXISTS migrations;
