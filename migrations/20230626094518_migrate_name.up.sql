DROP TABLE IF EXISTS audit CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE audit
(
    audit_id     UUID PRIMARY KEY    DEFAULT uuid_generate_v4(),
    user_id      VARCHAR(50),
    time         timestamp
   ,data         JSONB
);
