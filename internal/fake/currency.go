package fake

import (
	"kmock/internal/randkit"
	"math/rand/v2"
)

type Currency struct{}

// Name generates a random currency name using the provided random number generator.
func (c Currency) Name(rng *rand.Rand) string {
	return randkit.PickFromList(rng, currency).name
}

// Code generates a random currency code using the provided random number generator.
func (c Currency) Code(rng *rand.Rand) string {
	return randkit.PickFromList(rng, currency).code
}

// Symbol generates a random currency symbol using the provided random number generator.
func (c Currency) Symbol(rng *rand.Rand) string {
	return randkit.PickFromList(rng, currency).symbol
}

func (c Currency) Full(rng *rand.Rand) (string, string, string) {
	pick := randkit.PickFromList(rng, currency)
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
