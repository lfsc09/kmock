package fake

import (
	"fmt"
	"kmock/internal/randkit"
	"math/rand/v2"
)

type File struct{}

// Name generates a random file name with a prefix and a random integer, without an extension.
func (f File) Name(rng *rand.Rand) string {
	randomFileExtensionPrefix := randkit.PickFromList(rng, file).prefix
	return fmt.Sprintf("%s-%d", randomFileExtensionPrefix, randkit.RandomIntegerBetween(rng, 1, 1_000_000))
}

// Extension generates a random file extension from a predefined list of file extensions.
func (f File) Extension(rng *rand.Rand) string {
	return randkit.PickFromList(rng, file).ext
}

// NameWithExtension generates a random file name with a prefix, a random integer, and an extension.
func (f File) NameWithExtension(rng *rand.Rand) string {
	randomFileExtension := randkit.PickFromList(rng, file)
	return fmt.Sprintf("%s-%d.%s", randomFileExtension.prefix, randkit.RandomIntegerBetween(rng, 1, 1_000_000), randomFileExtension.ext)
}

// MimeType generates a random MIME type based on a predefined list of file types and their associated MIME types.
func (f File) MimeType(rng *rand.Rand) string {
	return randkit.PickFromList(rng, file).mime
}

// RuntimeDocs provides runtime documentation for the Company struct and its methods
func (f File) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "File",
		Methods: map[string]RunTimeDocsMethod{
			"Name": {
				Name:        "Name",
				Description: "Generates a random file name with a prefix and a random integer, without an extension",
				Params:      []string{},
			},
			"Extension": {
				Name:        "Extension",
				Description: "Generates a random file extension from a predefined list of file extensions",
				Params:      []string{},
			},
			"NameWithExtension": {
				Name:        "NameWithExtension",
				Description: "Generates a random file name with a prefix, a random integer, and an extension",
				Params:      []string{},
			},
			"MimeType": {
				Name:        "MimeType",
				Description: "Generates a random MIME type based on a predefined list of file types and their associated MIME types",
				Params:      []string{},
			},
		},
	}
}
