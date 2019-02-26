// Package triangle helps with triangles operations
package triangle

import "math"

// Kind is an int tha indicates which type the triangle is
type Kind int

const (
	// NaT == not a triangle
	NaT = iota
	// Equ == equilateral
	Equ
	// Iso == isosceles
	Iso
	// Sca == scalene
	Sca
)

// KindFromSides returns the kind of triangle, not a triangle, equilateral, isosceles or scalene.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if math.IsNaN(a+b+c) || math.IsInf(a+b+c, 1) || math.IsInf(a+b+c, -1) || a+b < c || a+c < b || b+c < a || a+b+c == 0 {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || b == c || c == a {
		k = Iso
	} else {
		k = Sca
	}
	return k
}
