package fake

import (
	"math/rand/v2"

	"github.com/lfsc09/kmock/internal/randkit"
)

type Number struct {
	Rng *rand.Rand
}

// Int generates a random integer using the provided random number generator.
func (n Number) Int() int {
	return randkit.RandomInteger[int](n.Rng)
}

// IntBetween generates a random integer between min and max (inclusive) using the provided random number generator.
func (n Number) IntBetween(min, max int) int {
	return randkit.RandomIntegerBetween(n.Rng, min, max)
}

// Float generates a random float64 using the provided random number generator.
func (n Number) Float() float64 {
	return randkit.RandomFloat(n.Rng)
}

// FloatBetween generates a random float64 between min and max with the specified number of decimal places using the provided random number generator.
func (n Number) FloatBetween(decimals int, min, max float64) float64 {
	return randkit.RandomFloatBetween(n.Rng, decimals, true, min, max)
}

// RuntimeDocs returns documentation for the Number struct and its methods.
func (n Number) RuntimeDocs() []*RunTimeDocs {
	return []*RunTimeDocs{
		{
			Domain:      "Number",
			Method:      "Int",
			Description: "Generates a random integer",
			Params:      []string{},
		},
		{
			Domain:      "Number",
			Method:      "IntBetween",
			Description: "Generates a random integer between min and max (inclusive)",
			Params:      []string{"min", "max"},
		},
		{
			Domain:      "Number",
			Method:      "Float",
			Description: "Generates a random float64",
			Params:      []string{},
		},
		{
			Domain:      "Number",
			Method:      "FloatBetween",
			Description: "Generates a random float64 between min and max with the specified number of decimal places",
			Params:      []string{"decimals", "min", "max"},
		},
	}
}
