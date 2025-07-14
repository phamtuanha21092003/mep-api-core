ALTER TABLE "user" ADD role_id INTEGER;

CREATE INDEX idx_user_role_id ON "user" (role_id) WHERE deleted_at IS NULL;

CREATE TABLE role (
    id SERIAL PRIMARY KEY NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID NOT NULL,
    deleted_by UUID,
    name varchar(255) NOT NULL
);

CREATE UNIQUE INDEX idx_role_name_unique ON role (name) WHERE deleted_at IS NULL;

CREATE TABLE permission (
    id SERIAL PRIMARY KEY NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID NOT NULL,
    deleted_by UUID,
    name varchar(255) NOT NULL,
    method varchar(5) NOT NULL,
    url varchar(255) NOT NULL
);

CREATE UNIQUE INDEX idx_permission_name_unique ON permission (name) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_permission_method_url_unique ON permission (method, url) WHERE deleted_at IS NULL;

CREATE TABLE role_permission (
  role_id INTEGER NOT NULL,
  permission_id INTEGER NOT NULL,
  PRIMARY KEY (role_id, permission_id)
);
