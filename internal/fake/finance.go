package fake

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/lfsc09/kmock/internal/randkit"
)

type Finance struct {
	Rng *rand.Rand
}

func (f Finance) CreditCardVendor() string {
	return randkit.PickFromList(f.Rng, creditCardVendor).name
}

func (f Finance) CreditCardNumber() string {
	randomVendor := randkit.PickFromList(f.Rng, creditCardVendor)
	randomVendorBin := randkit.PickFromList(f.Rng, randomVendor.bin)
	randomVendorLength := randkit.PickFromList(f.Rng, randomVendor.length)
	// Generate the credit card number with the correct length and prefix
	ccNumber := randomVendorBin
	for len(ccNumber) < randomVendorLength-1 {
		ccNumber += string(rune('0' + randkit.RandomIntegerBetween(f.Rng, 0, 9)))
	}
	// Calculate the checksum digit using the Luhn algorithm
	checksumDigit := luhnChecksum(ccNumber)
	return ccNumber + string(rune('0'+checksumDigit))
}

// CreditCardCVV generates a random credit card CVV (Card Verification Value) consisting of 3 digits using the provided random number generator.
func (f Finance) CreditCardCVV() string {
	return randkit.RandomStringTemplate(f.Rng, "\\d\\d\\d")
}

// CreditCardExpirationDate generates a random credit card expiration date in the format "MM/YYYY" using the provided random number generator.
func (f Finance) CreditCardExpirationDate() string {
	month := randkit.RandomIntegerBetween(f.Rng, 1, 12)
	year := randkit.RandomIntegerBetween(f.Rng, 2000, 2038)
	return fmt.Sprintf("%02d/%d", month, year)
}

// CreditCardHolder generates a random credit card holder name using the provided random number generator.
// It returns an error if the locale is not supported.
func (f Finance) CreditCardHolder(locale string) (string, error) {
	if _, ok := availableLocales[locale]; !ok {
		return "", fmt.Errorf("%w: %s", ErrLocaleNotSupported, locale)
	}
	var middleNameInitial string
	if randkit.RandomBool(f.Rng, 0.3) {
		middleName, err := Person{Rng: f.Rng}.MiddleName(locale)
		if err != nil {
			return "", err
		}
		// Use only the first letter
		middleNameInitial = string(middleName[0]) + "."
	}
	personFirstName, err := Person{Rng: f.Rng}.FirstName(locale)
	if err != nil {
		return "", err
	}
	personLastName, err := Person{Rng: f.Rng}.LastName(locale)
	if err != nil {
		return "", err
	}
	var personName string
	if middleNameInitial != "" {
		personName = fmt.Sprintf("%s %s %s", personFirstName, middleNameInitial, personLastName)
	} else {
		personName = fmt.Sprintf("%s %s", personFirstName, personLastName)
	}
	// Ensure the credit card holder name does not exceed 26 characters (common limit for credit card names)
	if len(personName) > 26 {
		personName = personName[:26]
	}
	return strings.ToUpper(personName), nil
}

// RuntimeDocs provides runtime documentation for the Finance struct and its methods
func (f Finance) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "Finance",
		Methods: map[string]RunTimeDocsMethod{
			"CreditCardVendor": {
				Name:        "CreditCardVendor",
				Description: "Generates a random credit card vendor",
				Params:      []string{},
			},
			"CreditCardNumber": {
				Name:        "CreditCardNumber",
				Description: "Generates a random credit card number with a valid Luhn checksum",
				Params:      []string{},
			},
			"CreditCardCVV": {
				Name:        "CreditCardCVV",
				Description: "Generates a random 3-digit credit card CVV",
				Params:      []string{},
			},
			"CreditCardExpirationDate": {
				Name:        "CreditCardExpirationDate",
				Description: "Generates a random credit card expiration date in the format MM/YYYY",
				Params:      []string{},
			},
			"CreditCardHolder": {
				Name:        "CreditCardHolder",
				Description: "Generates a random credit card holder name",
				Params:      []string{"locale"},
			},
		},
	}
}

// Luhn algorithm implementation to calculate the checksum digit for credit card numbers.
func luhnChecksum(number string) int {
	sum := 0
	double := false
	// Process digits from right to left
	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	return sum % 10
}
