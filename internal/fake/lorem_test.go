package fake

import (
	"math/rand/v2"
	"testing"

	"github.com/lfsc09/kmock/internal/randkit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockLoremTestSuite struct {
	suite.Suite
	seeds []uint64
	lorem *Lorem
}

func TestMockLoremTestSuite(t *testing.T) {
	suite.Run(t, new(MockLoremTestSuite))
}

func (suite *MockLoremTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.lorem = &Lorem{
		Rng: rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1])),
	}
}

/*
	VALID STATES
*/

func (suite *MockLoremTestSuite) TestWord() {
	word := suite.lorem.Word()
	assert.NotEmpty(suite.T(), word, "Word should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockLoremTestSuite) TestSentence() {
	sentence := suite.lorem.Sentence(5)
	assert.NotEmpty(suite.T(), sentence, "Sentence should return a non-empty string when a positive word count is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	sentence = suite.lorem.Sentence(0)
	assert.NotEmpty(suite.T(), sentence, "Sentence should return a non-empty string when a non-positive word count is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	sentence = suite.lorem.Sentence(-5)
	assert.NotEmpty(suite.T(), sentence, "Sentence should return a non-empty string when a negative word count is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockLoremTestSuite) TestParagraph() {
	paragraph := suite.lorem.Paragraph(3)
	assert.NotEmpty(suite.T(), paragraph, "Paragraph should return a non-empty string when a positive sentence count is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	paragraph = suite.lorem.Paragraph(0)
	assert.NotEmpty(suite.T(), paragraph, "Paragraph should return a non-empty string when a non-positive sentence count is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	paragraph = suite.lorem.Paragraph(-3)
	assert.NotEmpty(suite.T(), paragraph, "Paragraph should return a non-empty string when a negative sentence count is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}
