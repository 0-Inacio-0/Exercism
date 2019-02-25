// This package contains the answer that bob will give
package bob

import "strings"

// This function receives a message and returns bob answer
func Hey(remark string) string {
	answers:="Whatever."
	if strings.Contains(remark, "question"){
		answers = "Sure."
	}
	if strings.Contains(remark, "shouting"){
		answers = "Whoa, chill out!"
	}
	if strings.Contains(remark, "shouting") && strings.Contains(remark, "shouting"){
		answers = "Calm down, I know what I'm doing!"
	}
	return answers
}
