package fake

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/lfsc09/kmock/internal/randkit"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockBooleanTestSuite struct {
	suite.Suite
	seeds   []uint64
	boolean *Boolean
}

func TestMockBooleanTestSuite(t *testing.T) {
	suite.Run(t, new(MockBooleanTestSuite))
}

func (suite *MockBooleanTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.boolean = &Boolean{
		Rng: rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1])),
	}
}

/*
	VALID STATES
*/

func (suite *MockBooleanTestSuite) TestRandom() {
	result := suite.boolean.Random()
	assert.IsType(suite.T(), true, result, fmt.Sprintf("Random should return a boolean value [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockBooleanTestSuite) TestRandomWithProbability() {
	probability := 0.7
	result := suite.boolean.RandomWithProbability(probability)
	assert.IsType(suite.T(), true, result, fmt.Sprintf("RandomWithProbability should return a boolean value for probability: %f [seeds: %d, %d]", probability, suite.seeds[0], suite.seeds[1]))
}
