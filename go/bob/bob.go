// This package contains the answer that bob will give
package bob

import "strings"

// This function receives a message and returns bob answer
func Hey(str string) string {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return "Fine. Be that way!"
	} else if str == strings.ToUpper(str) && str != strings.ToLower(str) && str[len(str)-1:] == "?" {
		return "Calm down, I know what I'm doing!"
	} else if str == strings.ToUpper(str) && str != strings.ToLower(str) {
		return "Whoa, chill out!"
	} else if str[len(str)-1:] == "?" {
		return "Sure."
	}
	return "Whatever."
}
