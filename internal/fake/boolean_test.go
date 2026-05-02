package fake

import (
	"fmt"
	"kmock/internal/randkit"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockBooleanTestSuite struct {
	suite.Suite
	seeds   []uint64
	rng     *rand.Rand
	boolean *Boolean
}

func TestMockBooleanTestSuite(t *testing.T) {
	suite.Run(t, new(MockBooleanTestSuite))
}

func (suite *MockBooleanTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.rng = rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1]))
	suite.boolean = &Boolean{}
}

/*
	VALID STATES
*/

func (suite *MockBooleanTestSuite) TestRandom() {
	result := suite.boolean.Random(suite.rng)
	assert.IsType(suite.T(), true, result, fmt.Sprintf("Random should return a boolean value [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockBooleanTestSuite) TestRandomWithProbability() {
	probability := 0.7
	result := suite.boolean.RandomWithProbability(suite.rng, probability)
	assert.IsType(suite.T(), true, result, fmt.Sprintf("RandomWithProbability should return a boolean value for probability: %f [seeds: %d, %d]", probability, suite.seeds[0], suite.seeds[1]))
}
