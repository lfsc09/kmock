package fake

import (
	"github.com/google/uuid"
)

type ID struct {
	seqIdState int
}

// UUIDv4 generates a random UUID version 4 string.
func (i ID) UUIDv4() string {
	randomUuidv4, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return randomUuidv4.String()
}

// UUIDv6 generates a random UUID version 6 string.
func (i ID) UUIDv6() string {
	randomUuidv6, err := uuid.NewV6()
	if err != nil {
		panic(err)
	}
	return randomUuidv6.String()
}

// UUIDv7 generates a random UUID version 7 string.
func (i ID) UUIDv7() string {
	randomUuidv7, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	return randomUuidv7.String()
}

// SequentialID generates a sequential integer ID starting from the specified value.
// If the provided startFrom value is negative, it defaults to 0.
// Each call to SequentialID increments the internal state and returns the next sequential ID.
func (i *ID) SequentialID(startFrom int) int {
	if startFrom < 0 {
		startFrom = 0
	}
	i.seqIdState = startFrom
	i.seqIdState++
	return i.seqIdState
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
