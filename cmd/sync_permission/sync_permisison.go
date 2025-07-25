package sync_permission

import (
	"fmt"
	"log"
	"strings"

	"github.com/phamtuanha21092003/mep-api-core/platform/database"
)

// TODO: thinking solution remove all permisison?
func SyncPermission(db *database.SqlxDatabase) {
	permissions := []string{"collection.product.variant.image"}

	for _, permission := range permissions {
		if err := insertPermissionsRecursively(db, permission); err != nil {
			log.Println("Error:", err)
		}
	}
}

func insertPermissionsRecursively(db *database.SqlxDatabase, permission string) error {
	parts := strings.Split(permission, ".")

	var currentPermission string
	for i := 0; i < len(parts); i++ {
		if i == 0 {
			currentPermission = parts[0]
		} else {
			currentPermission = currentPermission + "." + parts[i]
		}

		err := insertPermissions(db, currentPermission, i+1)
		if err != nil {
			return fmt.Errorf("failed on %s: %w", currentPermission, err)
		}
	}

	return nil
}

func insertPermissions(db *database.SqlxDatabase, permission string, level int) error {
	actions := []string{"view", "create", "update", "delete", "*"}
	for _, action := range actions {
		permKey := fmt.Sprintf("%s:%s", permission, action)
		description := generateDescription(permission, action)

		_, err := db.Exec(`
			INSERT INTO permission (name, description, level, created_at, updated_at)
			VALUES ($1, $2, $3, NOW(), NOW())
			ON CONFLICT (name) WHERE deleted_at IS NULL DO NOTHING
		`, permKey, description, level)
		if err != nil {
			return fmt.Errorf("failed to insert %s: %w", permKey, err)
		}

		log.Println("Inserted:", permKey)
	}

	return nil
}

func generateDescription(path, action string) string {
	resources := strings.Split(path, ".")
	resourceStr := strings.Join(resources, " > ")

	switch action {
	case "view":
		return fmt.Sprintf("View %s", resourceStr)

	case "create":
		return fmt.Sprintf("Create %s", resourceStr)

	case "update":
		return fmt.Sprintf("Update %s", resourceStr)

	case "delete":
		return fmt.Sprintf("Delete %s", resourceStr)

	case "*":
		return fmt.Sprintf("Full access to %s", resourceStr)

	default:
		return fmt.Sprintf("%s %s", strings.Title(action), resourceStr)
	}
}
