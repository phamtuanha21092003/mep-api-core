ALTER TABLE "user" ADD role_id uuid;

CREATE INDEX idx_user_role_id ON "user" (role_id) WHERE deleted_at IS NULL;

CREATE TABLE role (
  	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
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
    level INT NOT NULL,
    description varchar(255) NOT NULL
);

CREATE UNIQUE INDEX idx_permission_name_unique ON permission (name) WHERE deleted_at IS NULL;
CREATE INDEX idx_permission_level ON permission(level);

CREATE TABLE role_permission (
  role_id uuid NOT NULL,
  permission_id INTEGER NOT NULL,
  PRIMARY KEY (role_id, permission_id)
);
