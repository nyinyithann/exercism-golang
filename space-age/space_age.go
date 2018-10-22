package space

// EarthSecondsPerYear is amount of seconds per earth year
const EarthSecondsPerYear float64 = 31557600

// Planet represents planet in solar system
type Planet string

var planets = map[Planet]float64{
	"Earth":   EarthSecondsPerYear,
	"Mercury": 0.2408467 * EarthSecondsPerYear,
	"Venus":   0.61519726 * EarthSecondsPerYear,
	"Mars":    1.8808158 * EarthSecondsPerYear,
	"Jupiter": 11.862615 * EarthSecondsPerYear,
	"Saturn":  29.447498 * EarthSecondsPerYear,
	"Uranus":  84.016846 * EarthSecondsPerYear,
	"Neptune": 164.79132 * EarthSecondsPerYear,
}

// Age calculates planet year based on given number of seconds
func Age(seconds float64, planet Planet) float64 {
	if v, ok := planets[planet]; ok {
		return seconds / v
	}
	return 0
}
