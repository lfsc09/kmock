package fake

import (
	"math/rand/v2"

	"github.com/lfsc09/kmock/internal/randkit"
)

type Lorem struct {
	Rng *rand.Rand
}

// Word generates a random word from a predefined list of words.
func (l Lorem) Word() string {
	return randkit.PickFromList(l.Rng, loremWord)
}

// Sentence generates a random sentence with the specified number of words.
// If the word count is less than or equal to 0, it generates a sentence with a random number of words between 5 and 25.
func (l Lorem) Sentence(wordCount int) string {
	if wordCount <= 0 {
		wordCount = randkit.RandomIntegerBetween(l.Rng, 5, 25)
	}
	sb := randkit.GetStringBuilder()
	defer randkit.PutStringBuilder(sb)
	for i := range wordCount {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(l.Word())
	}
	sb.WriteByte('.')
	return sb.String()
}

// Paragraph generates a random paragraph with the specified number of sentences.
// If the sentence count is less than or equal to 0, it generates a paragraph of 1 sentence.
func (l Lorem) Paragraph(sentenceCount int) string {
	if sentenceCount <= 0 {
		sentenceCount = 1
	}
	sb := randkit.GetStringBuilder()
	defer randkit.PutStringBuilder(sb)
	for i := range sentenceCount {
		if i > 0 {
			sb.WriteByte('\n')
		}
		sb.WriteString(l.Sentence(0))
	}
	return sb.String()
}

// RuntimeDocs provides runtime documentation for the Lorem struct and its methods
func (l Lorem) RuntimeDocs() []*RunTimeDocs {
	return []*RunTimeDocs{
		{
			Domain:      "Lorem",
			Method:      "Word",
			Description: "Generates a random lorem word",
			Params:      []string{},
		},
		{
			Domain:      "Lorem",
			Method:      "Sentence",
			Description: "Generates a random lorem sentence with the specified number of words",
			Params:      []string{"number"},
		},
		{
			Domain:      "Lorem",
			Method:      "Paragraph",
			Description: "Generates a random lorem paragraph with the specified number of sentences",
			Params:      []string{"number"},
		},
	}
}
