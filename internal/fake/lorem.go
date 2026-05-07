package fake

import (
	"kmock/internal/randkit"
	"math/rand/v2"
)

type Lorem struct {
	Rng *rand.Rand
}

// Word generates a random word from a predefined list of words.
func (l Lorem) Word() string {
	return randkit.PickFromList(l.Rng, loremWord)
}

// Sentence generates a random sentence with the specified number of words.
func (l Lorem) Sentence(wordCount int) string {
	sb := randkit.GetStringBuilder()
	defer randkit.PutStringBuilder(sb)
	for i := 0; i < wordCount; i++ {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(l.Word())
	}
	sb.WriteByte('.')
	return sb.String()
}

// Paragraph generates a random paragraph with the specified number of sentences.
func (l Lorem) Paragraph(sentenceCount int) string {
	sb := randkit.GetStringBuilder()
	defer randkit.PutStringBuilder(sb)
	for i := 0; i < sentenceCount; i++ {
		if i > 0 {
			sb.WriteByte('\n')
		}
		sb.WriteString(l.Sentence(randkit.RandomIntegerBetween(l.Rng, 3, 15)))
	}
	return sb.String()
}

// RuntimeDocs provides runtime documentation for the Lorem struct and its methods
func (l Lorem) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "Lorem",
		Methods: map[string]RunTimeDocsMethod{
			"Word": {
				Name:        "Word",
				Description: "Generates a random lorem word",
				Params:      []string{},
			},
			"Sentence": {
				Name:        "Sentence",
				Description: "Generates a random lorem sentence with the specified number of words",
				Params:      []string{"number"},
			},
			"Paragraph": {
				Name:        "Paragraph",
				Description: "Generates a random lorem paragraph with the specified number of sentences",
				Params:      []string{"number"},
			},
		},
	}
}
