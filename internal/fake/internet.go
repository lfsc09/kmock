package fake

import (
	"fmt"
	"kmock/internal/randkit"
	"math/rand/v2"
)

type Internet struct{}

// Domain generates a random domain name using the provided random number generator.
func (i Internet) Domain(rng *rand.Rand) string {
	randomDomain := randkit.PickFromList(rng, webDomain)
	randomTld := randkit.PickFromList(rng, webTLD)
	return fmt.Sprintf("%s.%s", randomDomain, randomTld)
}

// DomainWithSubdomain generates a random domain name with a subdomain using the provided random number generator.
func (i Internet) DomainWithSubdomain(rng *rand.Rand) string {
	randomSubdomain := randkit.PickFromList(rng, webSubdomain)
	randomDomain := randkit.PickFromList(rng, webDomain)
	randomTld := randkit.PickFromList(rng, webTLD)
	return fmt.Sprintf("%s.%s.%s", randomSubdomain, randomDomain, randomTld)
}

// Ipv4 generates a random IPv4 address using the provided random number generator.
func (i Internet) Ipv4(rng *rand.Rand) string {
	segments := make([]string, 4)
	for j := 0; j < 4; j++ {
		segments[j] = fmt.Sprintf("%d", randkit.RandomIntegerBetween(rng, 0, 255))
	}
	return fmt.Sprintf("%s.%s.%s.%s", segments[0], segments[1], segments[2], segments[3])
}

// Ipv6 generates a random IPv6 address using the provided random number generator.
func (i Internet) Ipv6(rng *rand.Rand) string {
	segments := make([]string, 8)
	for j := 0; j < 8; j++ {
		segments[j] = fmt.Sprintf("%x", randkit.RandomIntegerBetween(rng, 0, 65535))
	}
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%s", segments[0], segments[1], segments[2], segments[3], segments[4], segments[5], segments[6], segments[7])
}

// MacAddress generates a random MAC address using the provided random number generator.
func (i Internet) MacAddress(rng *rand.Rand) string {
	segments := make([]string, 6)
	for j := 0; j < 6; j++ {
		segments[j] = fmt.Sprintf("%02X", randkit.RandomIntegerBetween(rng, 0, 255))
	}
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s", segments[0], segments[1], segments[2], segments[3], segments[4], segments[5])
}

// Url generates a random URL using the provided random number generator.
func (i Internet) Url(rng *rand.Rand) string {
	prefixTemplates := []string{"http://", "https://", "wwww."}
	randomDomain := randkit.PickFromList(rng, webDomain)
	randomTld := randkit.PickFromList(rng, webTLD)
	var randomSubdomain string
	if randkit.RandomBool(rng, 0.5) {
		randomSubdomain = randkit.PickFromList(rng, webSubdomain) + "."
	}
	return fmt.Sprintf("%s%s%s.%s", randkit.PickFromList(rng, prefixTemplates), randomSubdomain, randomDomain, randomTld)
}

// UserAgent generates a random user agent string using the provided random number generator.
func (i Internet) UserAgent(rng *rand.Rand) string {
	return randkit.PickFromList(rng, webUserAgent)
}

// RuntimeDocs provides runtime documentation for the Internet struct and its methods
func (i Internet) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "Internet",
		Methods: map[string]RunTimeDocsMethod{
			"Domain": {
				Name:        "Domain",
				Description: "Generates a random domain name",
				Params:      []string{},
			},
			"DomainWithSubdomain": {
				Name:        "DomainWithSubdomain",
				Description: "Generates a random domain name with a subdomain",
				Params:      []string{},
			},
			"Ipv4": {
				Name:        "Ipv4",
				Description: "Generates a random IPv4 address",
				Params:      []string{},
			},
			"Ipv6": {
				Name:        "Ipv6",
				Description: "Generates a random IPv6 address",
				Params:      []string{},
			},
			"MacAddress": {
				Name:        "MacAddress",
				Description: "Generates a random MAC address",
				Params:      []string{},
			},
			"Url": {
				Name:        "Url",
				Description: "Generates a random URL",
				Params:      []string{},
			},
			"UserAgent": {
				Name:        "UserAgent",
				Description: "Generates a random user agent string",
				Params:      []string{},
			},
		},
	}
}
