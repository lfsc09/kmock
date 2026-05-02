package fake

import (
	"kmock/internal/randkit"
	"math/rand/v2"
)

var (
	brand = []string{
		"Audi", "BMW", "Chevrolet", "Chrysler", "Citroën",
		"Dodge", "Ferrari", "Fiat", "Ford", "Honda",
		"Hyundai", "Jaguar", "Jeep", "Kia", "Lamborghini",
		"Land Rover", "Lexus", "Maserati", "Mazda", "Mercedes-Benz",
		"Mitsubishi", "Nissan", "Peugeot", "Porsche", "Renault",
		"Subaru", "Tesla", "Toyota", "Volkswagen", "Volvo",
	}

	model = map[string][]string{
		"Audi":          {"A3", "A4", "A6", "Q5", "Q7"},
		"BMW":           {"3 Series", "5 Series", "X3", "X5", "M3"},
		"Chevrolet":     {"Silverado", "Equinox", "Malibu", "Camaro", "Tahoe"},
		"Chrysler":      {"300", "Pacifica", "Voyager", "Aspen", "Sebring"},
		"Citroën":       {"C3", "C4", "C5 Aircross", "Berlingo", "Jumper"},
		"Dodge":         {"Charger", "Challenger", "Durango", "Ram 1500", "Journey"},
		"Ferrari":       {"488 GTB", "F8 Tributo", "Roma", "SF90 Stradale", "Portofino"},
		"Fiat":          {"500", "Pulse", "Fastback", "Doblo", "Toro"},
		"Ford":          {"F-150", "Mustang", "Explorer", "Escape", "Bronco"},
		"Honda":         {"Civic", "Accord", "CR-V", "Pilot", "HR-V"},
		"Hyundai":       {"Elantra", "Sonata", "Tucson", "Santa Fe", "Kona"},
		"Jaguar":        {"XE", "XF", "F-Pace", "E-Pace", "I-Pace"},
		"Jeep":          {"Wrangler", "Cherokee", "Grand Cherokee", "Compass", "Renegade"},
		"Kia":           {"Forte", "Optima", "Sportage", "Sorento", "Telluride"},
		"Lamborghini":   {"Huracán", "Urus", "Aventador", "Gallardo", "Murciélago"},
		"Land Rover":    {"Defender", "Discovery", "Range Rover", "Freelander", "Evoque"},
		"Lexus":         {"IS", "ES", "RX", "NX", "GX"},
		"Maserati":      {"Ghibli", "Quattroporte", "Levante", "GranTurismo", "GranCabrio"},
		"Mazda":         {"Mazda3", "Mazda6", "CX-5", "CX-9", "MX-5 Miata"},
		"Mercedes-Benz": {"C-Class", "E-Class", "S-Class", "GLC", "GLE"},
		"Mitsubishi":    {"Outlander", "Eclipse Cross", "Pajero", "L200", "ASX"},
		"Nissan":        {"Altima", "Sentra", "Rogue", "Murano", "Frontier"},
		"Peugeot":       {"208", "308", "3008", "5008", "508"},
		"Porsche":       {"911", "Cayenne", "Macan", "Panamera", "Taycan"},
		"Renault":       {"Clio", "Megane", "Duster", "Sandero", "Captur"},
		"Subaru":        {"Impreza", "Legacy", "Outback", "Forester", "Crosstrek"},
		"Tesla":         {"Model 3", "Model S", "Model X", "Model Y", "Cybertruck"},
		"Toyota":        {"Camry", "Corolla", "RAV4", "Highlander", "Tacoma"},
		"Volkswagen":    {"Golf", "Jetta", "Passat", "Tiguan", "Touareg"},
		"Volvo":         {"S60", "S90", "XC40", "XC60", "XC90"},
	}
)

type Car struct{}

// Brand generates a random car brand from a predefined list of brands.
func (c Car) Brand(rng *rand.Rand) string {
	return randkit.PickFromList(rng, brand)
}

// Model generates a random car model based on the brand. It first selects a random brand and then picks a model from the list of models associated with that brand.
func (c Car) Model(rng *rand.Rand) string {
	randomBrand := c.Brand(rng)
	return randkit.PickFromList(rng, model[randomBrand])
}

// LicensePlate generates a random license plate based on the specified locale. It uses predefined templates for each supported locale to create realistic license plate formats.
func (c Car) LicensePlate(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	templates := map[string][]string{
		"en-US": {"!!!-####", "!!!-&&&", "?!!!???"},
		"pt-BR": {"!!!-####", "!!!-#!##"},
	}
	return randkit.RandomStringTemplate(rng, randkit.PickFromList(rng, templates[locale]))
}

// Color generates a random car color from a predefined list of colors.
func (c Car) Color(rng *rand.Rand) string {
	colors := []string{"Red", "Green", "Blue", "Black", "White", "Silver", "Yellow", "Purple", "Orange"}
	return randkit.PickFromList(rng, colors)
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
