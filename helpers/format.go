package helpers

import "strings"

func ToSlug(text string) string {
	text = strings.ToLower(text)
	return strings.ReplaceAll(text, " ", "-")
}
