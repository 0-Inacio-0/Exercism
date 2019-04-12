package isogram

import "strings"

// IsIsogram checks if the string does not repeat any letter, spaces and hyphens are allowed to appear multiple times.
func IsIsogram(str string) bool {
	str = strings.ToLower(str)
	for _, r := range str {
		if r == '-' || r == ' ' {
			continue
		}
		if strings.Count(str, string(r)) > 1 {
			return false
		}
	}
	return true
}
