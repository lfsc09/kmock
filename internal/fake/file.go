package fake

import (
	"fmt"
	"kmock/internal/randkit"
	"math/rand/v2"
)

type File struct {
	Rng *rand.Rand
}

// Name generates a random file name with a prefix and a random integer, without an extension.
func (f File) Name() string {
	randomFileExtensionPrefix := randkit.PickFromList(f.Rng, file).prefix
	return fmt.Sprintf("%s-%d", randomFileExtensionPrefix, randkit.RandomIntegerBetween(f.Rng, 1, 1_000_000))
}

// Extension generates a random file extension from a predefined list of file extensions.
func (f File) Extension() string {
	return randkit.PickFromList(f.Rng, file).ext
}

// NameWithExtension generates a random file name with a prefix, a random integer, and an extension.
func (f File) NameWithExtension() string {
	randomFileExtension := randkit.PickFromList(f.Rng, file)
	return fmt.Sprintf("%s-%d.%s", randomFileExtension.prefix, randkit.RandomIntegerBetween(f.Rng, 1, 1_000_000), randomFileExtension.ext)
}

// MimeType generates a random MIME type based on a predefined list of file types and their associated MIME types.
func (f File) MimeType() string {
	return randkit.PickFromList(f.Rng, file).mime
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
