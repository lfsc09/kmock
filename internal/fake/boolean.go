package fake

import (
	"kmock/internal/randkit"
	"math/rand/v2"
)

type Boolean struct{}

// Random generates a random boolean value with a default probability of 0.5 (50% chance of being true).
func (b Boolean) Random(rng *rand.Rand) bool {
	return randkit.RandomBool(rng, 0.5)
}

// RandomWithProbability generates a random boolean value with the specified probability of being true.
func (b Boolean) RandomWithProbability(rng *rand.Rand, probability float64) bool {
	return randkit.RandomBool(rng, probability)
}

// RuntimeDocs provides runtime documentation for the Boolean struct and its methods
func (b Boolean) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "Boolean",
		Methods: map[string]RunTimeDocsMethod{
			"Random": {
				Name:        "Random",
				Description: "Generates a random boolean",
				Params:      []string{},
			},
			"RandomWithProbability": {
				Name:        "RandomWithProbability",
				Description: "Generates a random boolean value with the specified probability",
				Params:      []string{"probability"},
			},
		},
	}
}
