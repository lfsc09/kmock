package fake

import (
	"kmock/internal/randkit"
	"math/rand/v2"
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
// It panics if the locale is not supported.
func (c Car) LicensePlate(locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	templates := map[string][]string{
		"en-US": {"\\L\\L\\L-\\d\\d\\d\\d", "\\L\\L\\L-\\A\\A\\A", "\\l\\L\\L\\L\\l\\l\\l"},
		"pt-BR": {"\\L\\L\\L-\\d\\d\\d\\d", "\\L\\L\\L-\\d\\L\\d\\d"},
	}
	return randkit.RandomStringTemplate(c.Rng, randkit.PickFromList(c.Rng, templates[locale]))
}

// Color generates a random car color from a predefined list of colors.
func (c Car) Color() string {
	colors := []string{"Red", "Green", "Blue", "Black", "White", "Silver", "Yellow", "Purple", "Orange"}
	return randkit.PickFromList(c.Rng, colors)
}

func (c Car) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "Car",
		Methods: map[string]RunTimeDocsMethod{
			"Brand": {
				Name:        "Brand",
				Description: "Generates a random car brand",
				Params:      []string{},
			},
			"Model": {
				Name:        "Model",
				Description: "Generates a random car model based on the brand",
				Params:      []string{},
			},
			"LicensePlate": {
				Name:        "LicensePlate",
				Description: "Generates a random license plate based on the specified locale",
				Params:      []string{"locale"},
			},
			"Color": {
				Name:        "Color",
				Description: "Generates a random car color",
				Params:      []string{},
			},
		},
	}
}
