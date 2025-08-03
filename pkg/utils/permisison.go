package utils

import "strings"

func BuildPermissionHierarchy(permissions []string) []string {
	seen := make(map[string]struct{})
	var result []string

	for _, permission := range permissions {
		parts := strings.Split(permission, ":")
		if len(parts) != 2 {
			continue
		}

		paths := strings.Split(parts[0], ".")
		method := parts[1]

		for i := 1; i <= len(paths); i++ {
			prefix := strings.Join(paths[:i], ".") + ":*"
			if _, exists := seen[prefix]; !exists {
				seen[prefix] = struct{}{}
				result = append(result, prefix)
			}
		}

		final := parts[0] + ":" + method
		if _, exists := seen[final]; !exists {
			seen[final] = struct{}{}
			result = append(result, final)
		}
	}

	return result
}
