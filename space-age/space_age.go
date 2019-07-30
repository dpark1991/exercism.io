package space

// Planet is a planet
type Planet string

const secondsInYear float64 = 31557600
const (
	mercury float64 = 0.2408467
	venus   float64 = 0.61519726
	mars    float64 = 1.8808158
	jupiter float64 = 11.862615
	saturn  float64 = 29.447498
	uranus  float64 = 84.016846
	neptune float64 = 164.79132
)

// Age returns seconds converted to age on planet
func Age(seconds float64, planet Planet) float64 {
	ageOnEarth := seconds / secondsInYear
	switch {
	case planet == "Earth":
		return ageOnEarth
	case planet == "Mercury":
		return ageOnEarth / mercury
	case planet == "Venus":
		return ageOnEarth / venus
	case planet == "Mars":
		return ageOnEarth / mars
	case planet == "Jupiter":
		return ageOnEarth / jupiter
	case planet == "Saturn":
		return ageOnEarth / saturn
	case planet == "Uranus":
		return ageOnEarth / uranus
	case planet == "Neptune":
		return ageOnEarth / neptune
	}
	return 0
}
