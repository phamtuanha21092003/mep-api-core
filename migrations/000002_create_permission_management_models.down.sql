-- Drop indexes for permission
DROP INDEX IF EXISTS idx_permission_name_unique;
DROP INDEX IF EXISTS idx_permission_method_url_unique;

DROP TABLE IF EXISTS permission;

DROP INDEX IF EXISTS idx_role_name_unique;

DROP TABLE IF EXISTS role;

DROP INDEX IF EXISTS idx_user_role_id;

ALTER TABLE "user" DROP COLUMN IF EXISTS role_id;
