package fake

import (
	"math/rand/v2"

	"github.com/lfsc09/kmock/internal/randkit"
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

func (c Currency) RuntimeDocs() []*RunTimeDocs {
	return []*RunTimeDocs{
		{
			Domain:      "Currency",
			Method:      "Name",
			Description: "Generates a random currency name",
			Params:      []string{},
		},
		{
			Domain:      "Currency",
			Method:      "Code",
			Description: "Generates a random currency code",
			Params:      []string{},
		},
		{
			Domain:      "Currency",
			Method:      "Symbol",
			Description: "Generates a random currency symbol",
			Params:      []string{},
		},
	}
}
