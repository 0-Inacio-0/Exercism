package accumulate

// Converter is a func that modifies a string
type converter func(string) string

// Accumulate applies a converter function to a slice of strings and returns the converted slice.
func Accumulate(in []string, cnv converter) (out []string) {
	for _, str := range in {
		out = append(out, cnv(str))
	}
	return out
}
