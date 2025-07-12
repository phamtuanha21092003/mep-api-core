CREATE TABLE "user" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULl,
    password varchar(128) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    last_login timestamptz NULL,
    is_superuser bool NOT NULL,
    username varchar(150) NOT NULL,
    first_name varchar(150) NOT NULL,
    last_name varchar(150) NOT NULL,
    email varchar(255) NOT NULL,
    avatar varchar(255) NULL,
    is_active bool NOT NULL
);

CREATE INDEX idx_username_like ON "user" USING btree (username varchar_pattern_ops);

CREATE UNIQUE INDEX idx_unique_email ON "user" (email) WHERE deleted_at IS NULL;