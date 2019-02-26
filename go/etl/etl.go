package etl

import "strings"

// Transform inverts the map from a value equal to a slice of string to a string equal a value.
func Transform(in map[int][]string) (out map[string]int) {
	out = make(map[string]int)
	for scr, ele := range in {
		for _, str := range ele {
			out[strings.ToLower(str)] = scr
		}
	}
	return out
}
