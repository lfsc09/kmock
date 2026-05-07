package fake

import (
	"kmock/internal/randkit"
	"math/rand/v2"
)

type Address struct {
	Rng *rand.Rand
}

// Country generates a random country name using the provided random number generator.
func (a Address) Country() string {
	return randkit.PickFromList(a.Rng, country).name
}

// CountryCode generates a random country code using the provided random number generator.
func (a Address) CountryCode() string {
	return randkit.PickFromList(a.Rng, country).code
}

// State generates a random state name for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) State(locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(a.Rng, state[locale]).name
}

// StateCode generates a random state code for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) StateCode(locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(a.Rng, state[locale]).code
}

// City generates a random city name for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) City(locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	randomStateCode := randkit.PickFromList(a.Rng, state[locale]).code
	return randkit.PickFromList(a.Rng, city[locale][randomStateCode])
}

// CityFromState generates a random city name for the specified locale and state code using the provided random number generator.
// It panics if the locale or state is not supported.
func (a Address) CityFromState(locale, stateCode string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	if _, ok := city[locale][stateCode]; !ok {
		panic("state not supported for locale: " + stateCode)
	}
	return randkit.PickFromList(a.Rng, city[locale][stateCode])
}

// Neighborhood generates a random neighborhood name for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) Neighborhood(locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(a.Rng, neighborhood[locale])
}

// StreetName generates a random street name for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) StreetName(locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(a.Rng, streetName[locale])
}

// StreetNumber generates a random street number using the provided random number generator, following common formatting patterns.
func (a Address) StreetNumber() string {
	templates := []string{"#####", "####", "###", "##"}
	return randkit.RandomStringTemplate(a.Rng, randkit.PickFromList(a.Rng, templates))
}

// StreetComplement generates a random street complement for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) StreetComplement(locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	templates := map[string][]string{
		"en-US": {"Apt. \\d\\d\\d", "Suite \\d\\d\\d", "Floor \\d", "Unit \\d\\d\\d", "Building \\d"},
		"pt-BR": {"Apto. \\d\\d\\d", "Sala \\d\\d\\d", "Andar \\d", "Unidade \\d\\d\\d", "Bloco \\d"},
	}
	return randkit.RandomStringTemplate(a.Rng, randkit.PickFromList(a.Rng, templates[locale]))
}

// ZipCode generates a random zip code for the specified locale using the provided random number generator, following common formatting patterns.
// It panics if the locale is not supported.
func (a Address) ZipCode(locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	templates := map[string][]string{
		"en-US": {"\\d\\d\\d\\d\\d", "\\d\\d\\d\\d\\d-\\d\\d\\d\\d"},
		"pt-BR": {"\\d\\d\\d\\d\\d-\\d\\d\\d"},
	}
	return randkit.RandomStringTemplate(a.Rng, randkit.PickFromList(a.Rng, templates[locale]))
}

// Latitude generates a random latitude value between -90 and 90 degrees using the provided random number generator.
func (a Address) Latitude() float64 {
	return randkit.RandomFloatBetween(a.Rng, 6, true, -90, 90)
}

// Longitude generates a random longitude value between -180 and 180 degrees using the provided random number generator.
func (a Address) Longitude() float64 {
	return randkit.RandomFloatBetween(a.Rng, 6, true, -180, 180)
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
