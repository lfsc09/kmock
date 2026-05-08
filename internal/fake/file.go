package fake

import (
	"fmt"
	"math/rand/v2"

	"github.com/lfsc09/kmock/internal/randkit"
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

// RuntimeDocs provides runtime documentation for the File struct and its methods
func (f File) RuntimeDocs() []*RunTimeDocs {
	return []*RunTimeDocs{
		{
			Domain:      "File",
			Method:      "Name",
			Description: "Generates a random file name with a prefix and a random integer, without an extension",
			Params:      []string{},
		},
		{
			Domain:      "File",
			Method:      "Extension",
			Description: "Generates a random file extension from a predefined list of file extensions",
			Params:      []string{},
		},
		{
			Domain:      "File",
			Method:      "NameWithExtension",
			Description: "Generates a random file name with a prefix, a random integer, and an extension",
			Params:      []string{},
		},
		{
			Domain:      "File",
			Method:      "MimeType",
			Description: "Generates a random MIME type based on a predefined list of file types and their associated MIME types",
			Params:      []string{},
		},
	}
}
