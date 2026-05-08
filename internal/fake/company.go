package fake

import (
	"fmt"
	"math/rand/v2"

	"github.com/lfsc09/kmock/internal/randkit"
)

type Company struct {
	Rng *rand.Rand
}

// Name generates a random company name based on the specified locale.
// It returns an error if the locale is not supported.
func (c Company) Name(locale string) (string, error) {
	if _, ok := availableLocales[locale]; !ok {
		return "", fmt.Errorf("%w: %s", ErrLocaleNotSupported, locale)
	}
	return randkit.PickFromList(c.Rng, companyName[locale]), nil
}

// Dba generates a random "doing business as" (Nome fantasia) name for a company.
// It returns an error if the locale is not supported.
func (c Company) Dba(locale string) (string, error) {
	if _, ok := availableLocales[locale]; !ok {
		return "", fmt.Errorf("%w: %s", ErrLocaleNotSupported, locale)
	}
	return randkit.PickFromList(c.Rng, companyDba[locale]), nil
}

// Industry generates a random industry name for a company.
// It returns an error if the locale is not supported.
func (c Company) Industry(locale string) (string, error) {
	if _, ok := availableLocales[locale]; !ok {
		return "", fmt.Errorf("%w: %s", ErrLocaleNotSupported, locale)
	}
	return randkit.PickFromList(c.Rng, companyIndustry[locale]), nil
}

// Suffix generates a random company suffix (e.g., Inc., LLC, etc.).
// It returns an error if the locale is not supported.
func (c Company) Suffix(locale string) (string, error) {
	if _, ok := availableLocales[locale]; !ok {
		return "", fmt.Errorf("%w: %s", ErrLocaleNotSupported, locale)
	}
	return randkit.PickFromList(c.Rng, companySuffix[locale]), nil
}

// EIN generates a random Employer Identification Number (EIN) for a company.
func (c Company) EIN() string {
	return randkit.RandomStringTemplate(c.Rng, "\\d\\d-\\d\\d\\d\\d\\d\\d\\d")
}

// CNPJLegacyValid generates a random valid legacy (00.000.000/0000-00) Brazilian CNPJ (Cadastro Nacional da Pessoa Jurídica) number for a company.
func (c Company) CNPJLegacyValid() string {
	cnpj := make([]int, 12)
	// Generate the first 12 digits
	for i := range 12 {
		cnpj[i] = randkit.RandomIntegerBetween(c.Rng, 0, 9)
	}
	// Multipliers for checksum digits
	multipliers1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	multipliers2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	// Calculate the first checksum digit
	cnpj = append(cnpj[:12], cnpjChecksum(cnpj[:12], multipliers1))
	// Calculate the second checksum digit
	cnpj = append(cnpj[:13], cnpjChecksum(cnpj[:13], multipliers2))
	return fmt.Sprintf("%02d.%03d.%03d/%04d-%02d",
		cnpj[0]*10+cnpj[1],
		cnpj[2]*100+cnpj[3]*10+cnpj[4],
		cnpj[5]*100+cnpj[6]*10+cnpj[7],
		cnpj[8]*1000+cnpj[9]*100+cnpj[10]*10+cnpj[11],
		cnpj[12]*10+cnpj[13],
	)
}

// IE generates a random Brazilian IE (Inscrição Estadual) number for a company.
func (c Company) IE() string {
	templates := []string{
		"\\d\\d\\d.\\d\\d\\d.\\d\\d\\d.\\d\\d\\d",
		"\\d\\d\\d.\\d\\d\\d.\\d\\d\\d",
		"\\d\\d\\d/\\d\\d\\d\\d\\d\\d\\d",
	}
	return randkit.RandomStringTemplate(c.Rng, randkit.PickFromList(c.Rng, templates))
}

// CNAE generates a random Brazilian CNAE (Classificação Nacional de Atividades Econômicas) code for a company.
func (c Company) CNAE() string {
	return randkit.RandomStringTemplate(c.Rng, "\\d\\d\\d\\d-\\d/\\d\\d")
}

// RuntimeDocs provides runtime documentation for the Company struct and its methods
func (c Company) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "Company",
		Methods: map[string]RunTimeDocsMethod{
			"Name": {
				Name:        "Name",
				Description: "Generates a random company name based on the specified locale",
				Params:      []string{"locale"},
			},
			"Dba": {
				Name:        "Dba",
				Description: "Generates a random 'doing business as' (Nome fantasia) name for a company",
				Params:      []string{"locale"},
			},
			"Industry": {
				Name:        "Industry",
				Description: "Generates a random industry name for a company",
				Params:      []string{"locale"},
			},
			"Suffix": {
				Name:        "Suffix",
				Description: "Generates a random company suffix (e.g., Inc., LLC, etc.)",
				Params:      []string{"locale"},
			},
			"EIN": {
				Name:        "EIN",
				Description: "Generates a random Employer Identification Number (EIN) for a company",
				Params:      []string{},
			},
			"CNPJLegacyValid": {
				Name:        "CNPJLegacyValid",
				Description: "Generates a random valid legacy (00.000.000/0000-00) Brazilian CNPJ (Cadastro Nacional da Pessoa Jurídica) number for a company",
				Params:      []string{},
			},
			"IE": {
				Name:        "IE",
				Description: "Generates a random Brazilian IE (Inscrição Estadual) number for a company",
				Params:      []string{},
			},
			"CNAE": {
				Name:        "CNAE",
				Description: "Generates a random Brazilian CNAE (Classificação Nacional de Atividades Econômicas) code for a company",
				Params:      []string{},
			},
		},
	}
}

// cnpjChecksum calculates the checksum digit for a CNPJ number based on the provided digits and multipliers.
// It sums the products of the digits and their corresponding multipliers, then calculates the modulus 11 of the sum.
func cnpjChecksum(digits []int, multipliers []int) int {
	sum := 0
	for i, m := range multipliers {
		sum += digits[i] * m
	}
	mod := sum % 11
	if mod < 2 {
		return 0
	}
	return 11 - mod
}
