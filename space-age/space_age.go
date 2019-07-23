package space

import "fmt"

type Planet string

const (
	Earth   = "Earth"
	Mercury = "Mercury"
	Venus   = "Venus"
	Mars    = "Mars"
	Jupiter = "Jupiter"
	Saturn  = "Saturn"
	Uranus  = "Uranus"
	Neptune = "Neptune"

	earthOrbitalPeriodInDays = 365.25
)

func Age(seconds float64, planet Planet) float64 {
	var orbitalPeriodInDays float64
	switch planet {
	case Earth:
		orbitalPeriodInDays = earthOrbitalPeriodInDays
	case Mercury:
		orbitalPeriodInDays = earthOrbitalPeriodInDays * 0.2408467
	case Venus:
		orbitalPeriodInDays = earthOrbitalPeriodInDays * 0.61519726
	case Mars:
		orbitalPeriodInDays = earthOrbitalPeriodInDays * 1.8808158
	case Jupiter:
		orbitalPeriodInDays = earthOrbitalPeriodInDays * 11.862615
	case Saturn:
		orbitalPeriodInDays = earthOrbitalPeriodInDays * 29.447498
	case Uranus:
		orbitalPeriodInDays = earthOrbitalPeriodInDays * 84.016846
	case Neptune:
		orbitalPeriodInDays = earthOrbitalPeriodInDays * 164.79132
	default:
		panic(fmt.Errorf("planet %s undefined", planet))
	}
	return seconds / (orbitalPeriodInDays * 24 * 60 * 60)
}
