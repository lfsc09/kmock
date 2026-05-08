package fake

import (
	"fmt"
	"math/rand/v2"

	"github.com/lfsc09/kmock/internal/randexp"
	"github.com/lfsc09/kmock/internal/randkit"
)

type Person struct {
	Rng *rand.Rand
}

// Name generates a random full name for the specified locale using the provided random number generator.
// It randomly decides whether to include a middle name based on a 30% chance.
// It returns an error if the locale is not supported.
func (p Person) Name(locale string) (string, error) {
	if _, ok := availableLocales[locale]; !ok {
		return "", fmt.Errorf("%w: %s", ErrLocaleNotSupported, locale)
	}
	if randkit.RandomBool(p.Rng, 0.3) {
		firstName, err := p.FirstName(locale)
		if err != nil {
			return "", err
		}
		middleName, err := p.MiddleName(locale)
		if err != nil {
			return "", err
		}
		lastName, err := p.LastName(locale)
		if err != nil {
			return "", err
		}
		return firstName + " " + middleName + " " + lastName, nil
	}
	firstName, err := p.FirstName(locale)
	if err != nil {
		return "", err
	}
	lastName, err := p.LastName(locale)
	if err != nil {
		return "", err
	}
	return firstName + " " + lastName, nil
}

// FirstName generates a random first name for the specified locale using the provided random number generator.
// It returns an error if the locale is not supported.
func (p Person) FirstName(locale string) (string, error) {
	if _, ok := availableLocales[locale]; !ok {
		return "", fmt.Errorf("%w: %s", ErrLocaleNotSupported, locale)
	}
	return randkit.PickFromList(p.Rng, personFirstName[locale]), nil
}

// MiddleName generates a random middle name for the specified locale using the provided random number generator.
// It returns an error if the locale is not supported.
func (p Person) MiddleName(locale string) (string, error) {
	if _, ok := availableLocales[locale]; !ok {
		return "", fmt.Errorf("%w: %s", ErrLocaleNotSupported, locale)
	}
	return randkit.PickFromList(p.Rng, personMiddleName[locale]), nil
}

// LastName generates a random last name for the specified locale using the provided random number generator.
// It returns an error if the locale is not supported.
func (p Person) LastName(locale string) (string, error) {
	if _, ok := availableLocales[locale]; !ok {
		return "", fmt.Errorf("%w: %s", ErrLocaleNotSupported, locale)
	}
	return randkit.PickFromList(p.Rng, personLastName[locale]), nil
}

// Phone generates a random phone number using the provided random number generator.
func (p Person) Phone() string {
	randomPhoneData := randkit.PickFromList(p.Rng, personPhone)
	return "+" + randomPhoneData.countryCode + " " + randkit.PickFromList(p.Rng, randomPhoneData.format)
}

// Email generates a random email address using the provided random number generator.
// It returns an error if the Username method fails to generate a username.
func (p Person) Email() (string, error) {
	username, err := p.Username()
	if err != nil {
		return "", err
	}
	return username + "@" + randkit.PickFromList(p.Rng, personEmailDomain), nil
}

// Username generates a random username by combining a base username with a random suffix.
// It returns an error if randexp fails to create its generator.
func (p Person) Username() (string, error) {
	sulfixTemplates := []string{
		"[a-zA-Z0-9]{3,16}",
		"[a-zA-Z0-9][a-zA-Z0-9_-]{1,18}[a-zA-Z0-9]",
		"[a-z0-9][a-z0-9-]{0,38}",
	}
	randExp, err := randexp.NewRandexpGenerator(randkit.PickFromList(p.Rng, sulfixTemplates))
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	return randExp.Generate(p.Rng), nil
}

// Password generates a random password based on the specified strength level (weak, medium, strong) using the provided random number generator.
// It returns an error if randexp fails to create its generator.
func (p Person) Password(strength string) (string, error) {
	var template string
	switch strength {
	case "weak":
		template = "[a-z]{8}"
	case "medium":
		template = "[a-zA-Z0-9]{8}"
	case "strong":
		template = "[a-zA-Z0-9!@#$%^&*()_+]{12}"
	default:
		template = "[a-zA-Z0-9]{8}"
	}
	randExp, err := randexp.NewRandexpGenerator(template)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	return randExp.Generate(p.Rng), nil
}

// JobTitle generates a random job title for the specified locale using the provided random number generator.
// It returns an error if the locale is not supported.
func (p Person) JobTitle(locale string) (string, error) {
	if _, ok := availableLocales[locale]; !ok {
		return "", fmt.Errorf("%w: %s", ErrLocaleNotSupported, locale)
	}
	return randkit.PickFromList(p.Rng, personJobTitle[locale]), nil
}

// CPF generates a random CPF (Cadastro de Pessoas Físicas) number for a person using the provided random number generator.
func (p Person) CPFValid() string {
	cpf := make([]int, 9)
	// Generate the first 9 digits
	for i := range 9 {
		cpf[i] = randkit.RandomIntegerBetween(p.Rng, 0, 9)
	}
	// Multipliers for checksum digits
	multipliers1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	multipliers2 := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	// Calculate the first checksum digit
	cpf = append(cpf, cpfChecksum(cpf[:9], multipliers1))
	// Calculate the second checksum digit
	cpf = append(cpf, cpfChecksum(cpf[:10], multipliers2))
	return fmt.Sprintf("%03d.%03d.%03d-%02d",
		cpf[0]*100+cpf[1]*10+cpf[2],
		cpf[3]*100+cpf[4]*10+cpf[5],
		cpf[6]*100+cpf[7]*10+cpf[8],
		cpf[9]*10+cpf[10],
	)
}

// RuntimeDocs provides runtime documentation for the Person struct and its methods
func (p Person) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "Person",
		Methods: map[string]RunTimeDocsMethod{
			"Name": {
				Name:        "Name",
				Description: "Generates a random full name based on the specified locale",
				Params:      []string{"locale"},
			},
			"FirstName": {
				Name:        "FirstName",
				Description: "Generates a random first name based on the specified locale",
				Params:      []string{"locale"},
			},
			"MiddleName": {
				Name:        "MiddleName",
				Description: "Generates a random middle name based on the specified locale",
				Params:      []string{"locale"},
			},
			"LastName": {
				Name:        "LastName",
				Description: "Generates a random last name based on the specified locale",
				Params:      []string{"locale"},
			},
			"Phone": {
				Name:        "Phone",
				Description: "Generates a random phone number",
				Params:      []string{},
			},
			"Email": {
				Name:        "Email",
				Description: "Generates a random email address",
				Params:      []string{},
			},
			"Username": {
				Name:        "Username",
				Description: "Generates a random username",
				Params:      []string{},
			},
			"Password": {
				Name:        "Password",
				Description: "Generates a random password based on the specified strength level (weak, medium, strong)",
				Params:      []string{"strength"},
			},
			"JobTitle": {
				Name:        "JobTitle",
				Description: "Generates a random job title based on the specified locale",
				Params:      []string{"locale"},
			},
			"CPFValid": {
				Name:        "CPFValid",
				Description: "Generates a random valid CPF number",
				Params:      []string{},
			},
		},
	}
}

// cpfChecksum calculates the check digit for a CPF number using the provided digits and multipliers.
// It returns the calculated check digit.
func cpfChecksum(digits []int, multipliers []int) int {
	sum := 0
	for i := range digits {
		sum += digits[i] * multipliers[i]
	}
	checkDigit := (sum * 10) % 11
	if checkDigit == 10 {
		checkDigit = 0
	}
	return checkDigit
}
