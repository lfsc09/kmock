package fake

import (
	"kmock/internal/randkit"
	"math/rand/v2"
)

type Number struct{}

// Int generates a random integer using the provided random number generator.
func (n Number) Int(rng *rand.Rand) int {
	return randkit.RandomInteger[int](rng)
}

// IntBetween generates a random integer between min and max (inclusive) using the provided random number generator.
func (n Number) IntBetween(rng *rand.Rand, min, max int) int {
	if min > max {
		min, max = max, min
	}
	if min == max {
		return min
	}
	return randkit.RandomIntegerBetween(rng, min, max)
}

// Float generates a random float64 using the provided random number generator.
func (n Number) Float(rng *rand.Rand) float64 {
	return randkit.RandomFloat(rng)
}

// FloatBetween generates a random float64 between min and max with the specified number of decimal places using the provided random number generator.
func (n Number) FloatBetween(rng *rand.Rand, decimals int, min, max float64) float64 {
	return randkit.RandomFloatBetween(rng, decimals, true, min, max)
}

// RuntimeDocs returns documentation for the Number struct and its methods.
func (n Number) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "Number",
		Methods: map[string]RunTimeDocsMethod{
			"Int": {
				Name:        "Int",
				Description: "Generates a random integer",
				Params:      []string{},
			},
			"IntBetween": {
				Name:        "IntBetween",
				Description: "Generates a random integer between min and max (inclusive)",
				Params:      []string{"min", "max"},
			},
			"Float": {
				Name:        "Float",
				Description: "Generates a random float64",
				Params:      []string{},
			},
			"FloatBetween": {
				Name:        "FloatBetween",
				Description: "Generates a random float64 between min and max with the specified number of decimal places",
				Params:      []string{"decimals", "min", "max"},
			},
		},
	}
}
