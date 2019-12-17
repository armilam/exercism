// An exercism exercise to learn math and data structures.
package space

// A type that specifies a planet. It defines only the name of the planet as a string.
type Planet string

// Ratio of Earth to planet years
var earthToPlanetYearRatios = map[Planet]float64 {
	"Mercury": 0.2408467,
	"Venus": 0.61519726,
	"Earth": 1.0,
	"Mars": 1.8808158,
	"Jupiter": 11.862615,
	"Saturn": 29.447498,
	"Uranus": 84.016846,
	"Neptune": 164.79132,
}

// Returns the age in years on the given planet represented by the seconds.
func Age(seconds float64, planet Planet) float64 {
	return EarthYears(seconds) / earthToPlanetYearRatios[planet]
}

// Returns the age in Earth years represented by the seconds.
func EarthYears(seconds float64) float64 {
	return seconds / 31557600
}
