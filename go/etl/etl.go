package etl

import "strings"

// Transform inverts the map from a value equal to a slice of string to a string equal a value.
func Transform(in map[int][]string) (out map[string]int) {
	out = make(map[string]int)
	for i, ele := range in {
		for _, str := range ele {
			str = strings.ToLower(str)
			out[str] = i
		}
	}
	return out
}
