package utils

import (
	"regexp"
	"strings"
)

func NormalizeNameToUsername(name string) string {
	// Replace spaces with underscores
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")

	// Remove all non-alphanumeric or underscore characters
	re := regexp.MustCompile(`[^a-z0-9_]+`)
	return re.ReplaceAllString(name, "")
}

func GenerateUserIdentity(base string) string {
	return "@" + NormalizeNameToUsername(base)
}
