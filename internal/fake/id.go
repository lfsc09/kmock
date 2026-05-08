package fake

import (
	"fmt"

	"github.com/google/uuid"
)

var errUUIDGeneration = fmt.Errorf("failed to generate")

type ID struct{}

// UUIDv4 generates a random UUID version 4 string.
// It returns an error if the UUID generation fails.
func (i ID) UUIDv4() (string, error) {
	randomUuidv4, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("%w UUIDv4: %v", errUUIDGeneration, err)
	}
	return randomUuidv4.String(), nil
}

// UUIDv6 generates a random UUID version 6 string.
// It returns an error if the UUID generation fails.
func (i ID) UUIDv6() (string, error) {
	randomUuidv6, err := uuid.NewV6()
	if err != nil {
		return "", fmt.Errorf("%w UUIDv6: %v", errUUIDGeneration, err)
	}
	return randomUuidv6.String(), nil
}

// UUIDv7 generates a random UUID version 7 string.
// It returns an error if the UUID generation fails.
func (i ID) UUIDv7() (string, error) {
	randomUuidv7, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("%w UUIDv7: %v", errUUIDGeneration, err)
	}
	return randomUuidv7.String(), nil
}

// SequentialID generates a sequential integer ID starting from the specified value.
// If the provided startFrom value is negative, it defaults to 0.
func (i *ID) SequentialID(startFrom int) int {
	if startFrom < 0 {
		startFrom = 0
	}
	startFrom++
	return startFrom
}

// // RuntimeDocs provides runtime documentation for the ID struct and its methods
func (i ID) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "ID",
		Methods: map[string]RunTimeDocsMethod{
			"UUIDv4": {
				Name:        "UUIDv4",
				Description: "Generates a random UUID version 4 string",
				Params:      []string{},
			},
			"UUIDv6": {
				Name:        "UUIDv6",
				Description: "Generates a random UUID version 6 string",
				Params:      []string{},
			},
			"UUIDv7": {
				Name:        "UUIDv7",
				Description: "Generates a random UUID version 7 string",
				Params:      []string{},
			},
			"SequentialID": {
				Name:        "SequentialID",
				Description: "Generates a sequential integer ID starting from the specified value",
				Params:      []string{"start"},
			},
		},
	}
}
