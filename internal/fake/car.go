package fake

import (
	"fmt"
	"math/rand/v2"

	"github.com/lfsc09/kmock/internal/randkit"
)

type Car struct {
	Rng *rand.Rand
}

// Brand generates a random car brand from a predefined list of brands.
func (c Car) Brand() string {
	return randkit.PickFromList(c.Rng, carBrand)
}

// Model generates a random car model based on the brand. It first selects a random brand and then picks a model from the list of models associated with that brand.
func (c Car) Model() string {
	randomBrand := c.Brand()
	return randkit.PickFromList(c.Rng, carModel[randomBrand])
}

// LicensePlate generates a random license plate based on the specified locale.
// It returns an error if the locale is not supported.
func (c Car) LicensePlate(locale string) (string, error) {
	if _, ok := availableLocales[locale]; !ok {
		return "", fmt.Errorf("%w: %s", ErrLocaleNotSupported, locale)
	}
	templates := map[string][]string{
		"en-US": {"\\L\\L\\L-\\d\\d\\d\\d", "\\L\\L\\L-\\A\\A\\A", "\\l\\L\\L\\L\\l\\l\\l"},
		"pt-BR": {"\\L\\L\\L-\\d\\d\\d\\d", "\\L\\L\\L-\\d\\L\\d\\d"},
	}
	return randkit.RandomStringTemplate(c.Rng, randkit.PickFromList(c.Rng, templates[locale])), nil
}

// Color generates a random car color from a predefined list of colors.
func (c Car) Color() string {
	colors := []string{"Red", "Green", "Blue", "Black", "White", "Silver", "Yellow", "Purple", "Orange"}
	return randkit.PickFromList(c.Rng, colors)
}

func (c Car) RuntimeDocs() []*RunTimeDocs {
	return []*RunTimeDocs{
		{
			Domain:      "Car",
			Method:      "Brand",
			Description: "Generates a random car brand",
			Params:      []string{},
		},
		{
			Domain:      "Car",
			Method:      "Model",
			Description: "Generates a random car model based on the brand",
			Params:      []string{},
		},
		{
			Domain:      "Car",
			Method:      "LicensePlate",
			Description: "Generates a random license plate based on the specified locale",
			Params:      []string{"locale"},
		},
		{
			Domain:      "Car",
			Method:      "Color",
			Description: "Generates a random car color",
			Params:      []string{},
		},
	}
}
