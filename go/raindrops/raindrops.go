package raindrops

import (
	"strconv"
)

// Convert receives an int and return a string fallowing the rule:
//	-If the number has 3 as a factor, output 'Pling'.
//	-If the number has 5 as a factor, output 'Plang'.
//	-If the number has 7 as a factor, output 'Plong'.
//	-If the number does not have 3, 5, or 7 as a factor, just pass the number's digits straight through.
func Convert(number int) string {
	str := ""
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			if i == 3 {
				str = "Pling"
			}
			if i == 5 {
				str += "Plang"
			}
			if i == 7 {
				str += "Plong"
			}
		}
	}
	if len(str) == 0 {
		str = strconv.Itoa(number)
	}
	return str
}
