CREATE EXTENSION IF NOT EXISTS citext;
CREATE TABLE if NOT EXISTS motoristas (
    id bigserial PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email citext NOT NULL,
    username VARCHAR(255) NOT NULL,
    password bytea NOT NULL,
    created_at TIMESTAMP with time zone NOT NULL DEFAULT NOW()
);