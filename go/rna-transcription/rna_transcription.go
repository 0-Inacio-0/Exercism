package strand

// ToRNA transforms a dna into a rna
func ToRNA(dna string) (rna string) {
	for _, ncl := range dna {
		if ncl == 'G' {
			rna += "C"
		} else if ncl == 'C' {
			rna += "G"
		} else if ncl == 'T' {
			rna += "A"
		} else if ncl == 'A' {
			rna += "U"
		}
	}
	return rna
}
