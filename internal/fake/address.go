package fake

import (
	"kmock/internal/randkit"
	"math/rand/v2"
)

type Address struct{}

// Country generates a random country name using the provided random number generator.
func (a Address) Country(rng *rand.Rand) string {
	return randkit.PickFromList(rng, country).name
}

// CountryCode generates a random country code using the provided random number generator.
func (a Address) CountryCode(rng *rand.Rand) string {
	return randkit.PickFromList(rng, country).code
}

// State generates a random state name for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) State(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(rng, state[locale]).name
}

// StateCode generates a random state code for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) StateCode(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(rng, state[locale]).code
}

// City generates a random city name for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) City(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	randomStateCode := randkit.PickFromList(rng, state[locale]).code
	return randkit.PickFromList(rng, city[locale][randomStateCode])
}

// CityFromState generates a random city name for the specified locale and state code using the provided random number generator.
// It panics if the locale or state is not supported.
func (a Address) CityFromState(rng *rand.Rand, locale, stateCode string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	if _, ok := city[locale][stateCode]; !ok {
		panic("state not supported for locale: " + stateCode)
	}
	return randkit.PickFromList(rng, city[locale][stateCode])
}

// Neighborhood generates a random neighborhood name for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) Neighborhood(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(rng, neighborhood[locale])
}

// StreetName generates a random street name for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) StreetName(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(rng, streetName[locale])
}

// StreetNumber generates a random street number using the provided random number generator, following common formatting patterns.
func (a Address) StreetNumber(rng *rand.Rand) string {
	templates := []string{"#####", "####", "###", "##"}
	return randkit.RandomStringTemplate(rng, randkit.PickFromList(rng, templates))
}

// StreetComplement generates a random street complement for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) StreetComplement(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	templates := map[string][]string{
		"en-US": {"Apt. \\d\\d\\d", "Suite \\d\\d\\d", "Floor \\d", "Unit \\d\\d\\d", "Building \\d"},
		"pt-BR": {"Apto. \\d\\d\\d", "Sala \\d\\d\\d", "Andar \\d", "Unidade \\d\\d\\d", "Bloco \\d"},
	}
	return randkit.RandomStringTemplate(rng, randkit.PickFromList(rng, templates[locale]))
}

// ZipCode generates a random zip code for the specified locale using the provided random number generator, following common formatting patterns.
// It panics if the locale is not supported.
func (a Address) ZipCode(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	templates := map[string][]string{
		"en-US": {"\\d\\d\\d\\d\\d", "\\d\\d\\d\\d\\d-\\d\\d\\d\\d"},
		"pt-BR": {"\\d\\d\\d\\d\\d-\\d\\d\\d"},
	}
	return randkit.RandomStringTemplate(rng, randkit.PickFromList(rng, templates[locale]))
}

// Latitude generates a random latitude value between -90 and 90 degrees using the provided random number generator.
func (a Address) Latitude(rng *rand.Rand) float64 {
	return randkit.RandomFloatBetween(rng, 6, true, -90, 90)
}

// Longitude generates a random longitude value between -180 and 180 degrees using the provided random number generator.
func (a Address) Longitude(rng *rand.Rand) float64 {
	return randkit.RandomFloatBetween(rng, 6, true, -180, 180)
}

// RuntimeDocs provides runtime documentation for the Address struct and its methods
func (a Address) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "Address",
		Methods: map[string]RunTimeDocsMethod{
			"Country": {
				Name:        "Country",
				Description: "Generates a random country name",
				Params:      []string{},
			},
			"CountryCode": {
				Name:        "CountryCode",
				Description: "Generates a random country code",
				Params:      []string{},
			},
			"State": {
				Name:        "State",
				Description: "Generates a random state name for the specified locale",
				Params:      []string{"locale"},
			},
			"StateCode": {
				Name:        "StateCode",
				Description: "Generates a random state code for the specified locale",
				Params:      []string{"locale"},
			},
			"City": {
				Name:        "City",
				Description: "Generates a random city name for the specified locale",
				Params:      []string{"locale"},
			},
			"CityFromState": {
				Name:        "CityFromState",
				Description: "Generates a random city name for the specified locale and state code",
				Params:      []string{"locale", "stateCode"},
			},
			"Neighborhood": {
				Name:        "Neighborhood",
				Description: "Generates a random neighborhood name for the specified locale",
				Params:      []string{"locale"},
			},
			"StreetName": {
				Name:        "StreetName",
				Description: "Generates a random street name for the specified locale",
				Params:      []string{"locale"},
			},
			"StreetNumber": {
				Name:        "StreetNumber",
				Description: "Generates a random street number",
				Params:      []string{},
			},
			"StreetComplement": {
				Name:        "StreetComplement",
				Description: "Generates a random street complement for the specified locale",
				Params:      []string{"locale"},
			},
			"ZipCode": {
				Name:        "ZipCode",
				Description: "Generates a random zip code for the specified locale",
				Params:      []string{"locale"},
			},
			"Latitude": {
				Name:        "Latitude",
				Description: "Generates a random latitude value between -90 and 90 degrees",
				Params:      []string{},
			},
			"Longitude": {
				Name:        "Longitude",
				Description: "Generates a random longitude value between -180 and 180 degrees",
				Params:      []string{},
			},
		},
	}
}
