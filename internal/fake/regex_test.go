package fake

import (
	"math/rand/v2"
	"testing"

	"github.com/lfsc09/kmock/internal/randkit"
	"github.com/stretchr/testify/suite"
)

type MockRegexTestSuite struct {
	suite.Suite
	seeds []uint64
	regex *Regex
}

func TestMockRegexTestSuite(t *testing.T) {
	suite.Run(t, new(MockRegexTestSuite))
}

func (suite *MockRegexTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.regex = &Regex{
		Rng: rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1])),
	}
}

/*
	VALID VALUES
*/

func (suite *MockRegexTestSuite) TestGenerate() {
	inputPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	outputPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	for range 3 {
		result, err := suite.regex.Generate(inputPattern)
		suite.NoError(err, "Generate should not return an error for a valid regex pattern [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
		suite.Regexp(outputPattern, result, "Generated string should match the regex pattern [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	}
	suite.Equal(1, len(suite.regex.cache), "Cache should contain one entry after generating with the same pattern multiple times [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockRegexTestSuite) TestClearCache() {
	pattern := `\\d{3}-\\d{2}-\\d{4}`
	for range 3 {
		_, err := suite.regex.Generate(pattern)
		suite.NoError(err, "Generate should not return an error for a valid regex pattern [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	}
	suite.Equal(1, len(suite.regex.cache), "Cache should contain one entry after generating with the same pattern multiple times [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.regex.ClearCache()
	suite.Equal(0, len(suite.regex.cache), "Cache should be empty after calling ClearCache [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}
