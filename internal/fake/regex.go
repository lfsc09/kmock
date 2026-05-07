package fake

import (
	"kmock/internal/randexp"
	"math/rand/v2"
)

type Regex struct {
	Rng   *rand.Rand
	cache map[string]*randexp.RandexpGenerator
}

// Generate produces a random string that matches the provided regular expression pattern using the given random number generator.
// It utilizes a cache to store compiled regular expression generators for improved performance on repeated patterns.
func (r *Regex) Generate(pattern string) string {
	if r.cache == nil {
		r.cache = make(map[string]*randexp.RandexpGenerator, 8)
	}
	var randexpInstance *randexp.RandexpGenerator
	var ok bool
	if randexpInstance, ok = r.cache[pattern]; !ok {
		// If the pattern is not in the cache, create a new randexp generator and store it in the cache
		var err error
		randexpInstance, err = randexp.NewRandexpGenerator(pattern)
		if err != nil {
			// Handle error appropriately
			panic(err)
		}
		r.cache[pattern] = randexpInstance
	}
	return randexpInstance.Generate(r.Rng)
}

// ClearCache empties the cache of compiled regular expression generators, allowing them to be garbage collected.
func (r *Regex) ClearCache() {
	r.cache = make(map[string]*randexp.RandexpGenerator, 8)
}

// RuntimeDocs provides runtime documentation for the Regex struct and its methods
func (r Regex) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "Regex",
		Methods: map[string]RunTimeDocsMethod{
			"Generate": {
				Name:        "Generate",
				Description: "Generates a random string that matches the provided regular expression pattern",
				Params:      []string{"pattern"},
			},
		},
	}
}
