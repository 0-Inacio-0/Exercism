package space

type planet struct {
	name string
	// Time that a planet takes to orbit around the sun
	orbitalPeriod float64
}

// One year in seconds
const earthYear = 31557600

var planets = []planet{{"Earth", earthYear}, {"Mercury", earthYear * 0.2408467}, {"Venus", earthYear * 0.61519726}, {"Mars", earthYear * 1.8808158},
	{"Jupiter", earthYear * 11.862615}, {"Saturn", earthYear * 29.447498}, {"Uranus", earthYear * 84.016846}, {"Neptune", earthYear * 164.79132}}

// Receives the time and planet and returns the age in that planet
func Age(sec float64, planetName string) float64 {
	age := 0.0
	for i := range planets {
		if planets[i].name == planetName {
			age = sec / planets[i].orbitalPeriod
		}
	}
	return age
}
