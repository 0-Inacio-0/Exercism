package accumulate

// Converter is a func that modifies a string
type converter func(string) string

// Accumulate applies a converter function to a slice of strings and returns the converted slice.
func Accumulate(convIn []string, conv converter) []string {
	convOut := make([]string, len(convIn))
	for i, str := range convIn {
		convOut[i] = conv(str)
	}
	return convOut
}
