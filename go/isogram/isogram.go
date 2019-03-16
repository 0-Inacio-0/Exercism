package isogram

import "strings"

// IsIsogram checks if the string does not repeat any letter, spaces and hyphens are allowed to appear multiple times.
func IsIsogram(str string) bool {
	out := true
	str = strings.ToLower(str)
	for _, r := range str {
		if r != '-' && r != ' ' && strings.Count(str, string(r)) > 1 {
			out = false
			break
		}
	}
	return out
}
