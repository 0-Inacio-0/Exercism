// Package triangle helps with triangles operations
package triangle

import "math"

// Kind is an int tha indicates which type the triangle is
type Kind int

const (
	// NaT == not a triangle
	NaT = 0
	// Equ == equilateral
	Equ = 1
	// Iso == isosceles
	Iso = 2
	// Sca == scalene
	Sca = 3
)

// KindFromSides returns the kind of triangle, not a triangle, equilateral, isosceles or scalene.
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	var k Kind
	for _, l := range sides {
		if l == math.NaN() || l == math.Inf(1) || l == math.Inf(-1) {
			return NaT
		}
	}
	if (a+b) >= c && (a+c) >= b && (b+c) >= a && (a != 0 && b != 0 && c != 0) {
		if a == b && b == c {
			k = Equ
		} else if a == b || b == c || c == a {
			k = Iso
		} else {
			k = Sca
		}
	} else {
		k = 0
	}
	return k
}
