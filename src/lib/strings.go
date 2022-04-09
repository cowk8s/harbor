package lib

import "strings"

// TrimsLineBreaks trims line breaks in string.
func TrimLineBreaks(s string) string {
	escaped := strings.ReplaceAll(s, "\n", "")
	escaped = strings.ReplaceAll(escaped, "\r", "")
	return escaped
}
