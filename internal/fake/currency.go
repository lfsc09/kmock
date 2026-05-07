package fake

import (
	"kmock/internal/randkit"
	"math/rand/v2"
)

type Currency struct {
	Rng *rand.Rand
}

// Name generates a random currency name using the provided random number generator.
func (c Currency) Name() string {
	return randkit.PickFromList(c.Rng, currency).name
}

// Code generates a random currency code using the provided random number generator.
func (c Currency) Code() string {
	return randkit.PickFromList(c.Rng, currency).code
}

// Symbol generates a random currency symbol using the provided random number generator.
func (c Currency) Symbol() string {
	return randkit.PickFromList(c.Rng, currency).symbol
}

// Full generates a random currency name, code, and symbol using the provided random number generator and returns them as a tuple.
func (c Currency) Full() (string, string, string) {
	pick := randkit.PickFromList(c.Rng, currency)
	return pick.name, pick.code, pick.symbol
}

func (c Currency) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "Currency",
		Methods: map[string]RunTimeDocsMethod{
			"Name": {
				Name:        "Name",
				Description: "Generates a random currency name",
				Params:      []string{},
			},
			"Code": {
				Name:        "Code",
				Description: "Generates a random currency code",
				Params:      []string{},
			},
			"Symbol": {
				Name:        "Symbol",
				Description: "Generates a random currency symbol",
				Params:      []string{},
			},
		},
	}
}
