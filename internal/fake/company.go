package fake

import (
	"fmt"
	"kmock/internal/randkit"
	"math/rand/v2"
)

type Company struct{}

// Name generates a random company name based on the specified locale.
// It panics if the locale is not supported.
func (c Company) Name(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(rng, companyName[locale])
}

// Dba generates a random "doing business as" (Nome fantasia) name for a company.
// It panics if the locale is not supported.
func (c Company) Dba(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(rng, companyDba[locale])
}

// Industry generates a random industry name for a company.
// It panics if the locale is not supported.
func (c Company) Industry(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(rng, companyIndustry[locale])
}

// Suffix generates a random company suffix (e.g., Inc., LLC, etc.).
// It panics if the locale is not supported.
func (c Company) Suffix(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(rng, companySuffix[locale])
}

// EIN generates a random Employer Identification Number (EIN) for a company.
func (c Company) EIN(rng *rand.Rand) string {
	return randkit.RandomStringTemplate(rng, "\\d\\d-\\d\\d\\d\\d\\d\\d\\d")
}

// CNPJLegacyValid generates a random valid legacy (00.000.000/0000-00) Brazilian CNPJ (Cadastro Nacional da Pessoa Jurídica) number for a company.
func (c Company) CNPJLegacyValid(rng *rand.Rand) string {
	cnpj := make([]int, 12)

	// Generate the first 12 digits
	for i := range 12 {
		cnpj[i] = rng.IntN(10)
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
func (c Company) IE(rng *rand.Rand) string {
	templates := []string{
		"\\d\\d\\d.\\d\\d\\d.\\d\\d\\d.\\d\\d\\d",
		"\\d\\d\\d.\\d\\d\\d.\\d\\d\\d",
		"\\d\\d\\d/\\d\\d\\d\\d\\d\\d\\d",
	}
	return randkit.RandomStringTemplate(rng, randkit.PickFromList(rng, templates))
}

// CNAE generates a random Brazilian CNAE (Classificação Nacional de Atividades Econômicas) code for a company.
func (c Company) CNAE(rng *rand.Rand) string {
	return randkit.RandomStringTemplate(rng, "\\d\\d\\d\\d-\\d/\\d\\d")
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
