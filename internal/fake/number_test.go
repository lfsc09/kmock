package fake

import (
	"math"
	"math/rand/v2"
	"testing"

	"github.com/lfsc09/kmock/internal/randkit"
	"github.com/stretchr/testify/suite"
)

type MockNumberTestSuite struct {
	suite.Suite
	seeds  []uint64
	number *Number
}

func TestMockNumberTestSuite(t *testing.T) {
	suite.Run(t, new(MockNumberTestSuite))
}

func (suite *MockNumberTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.number = &Number{
		Rng: rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1])),
	}
}

/*
	VALID STATES
*/

func (suite *MockNumberTestSuite) TestInt() {
	result := suite.number.Int()
	suite.NotNil(result, "Int should return a non-nil value [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockNumberTestSuite) TestIntBetween() {
	result := suite.number.IntBetween(1, 5)
	suite.GreaterOrEqual(result, 1, "IntBetween should return a value greater than or equal to min [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.LessOrEqual(result, 5, "IntBetween should return a value less than or equal to max [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	result = suite.number.IntBetween(5, 1)
	suite.GreaterOrEqual(result, 1, "IntBetween should return a value greater than or equal to min [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.LessOrEqual(result, 5, "IntBetween should return a value less than or equal to max [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	result = suite.number.IntBetween(3, 3)
	suite.Equal(3, result, "IntBetween should return the same value when min and max are equal [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	result = suite.number.IntBetween(0, 0)
	suite.GreaterOrEqual(result, 0, "IntBetween should return a value greater than or equal to min [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.LessOrEqual(result, math.MaxInt, "IntBetween should return a value less than or equal to max [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockNumberTestSuite) TestFloat() {
	result := suite.number.Float()
	suite.NotNil(result, "Float should return a non-nil value [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockNumberTestSuite) TestFloatBetween() {
	result := suite.number.FloatBetween(2, 1.0, 5.0)
	suite.GreaterOrEqual(result, 1.0, "FloatBetween should return a value greater than or equal to min [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.LessOrEqual(result, 5.0, "FloatBetween should return a value less than or equal to max [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	result = suite.number.FloatBetween(2, 5.0, 1.0)
	suite.GreaterOrEqual(result, 1.0, "FloatBetween should return a value greater than or equal to min [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.LessOrEqual(result, 5.0, "FloatBetween should return a value less than or equal to max [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	result = suite.number.FloatBetween(2, 3.0, 3.0)
	suite.Equal(3.0, result, "FloatBetween should return the same value when min and max are equal [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	result = suite.number.FloatBetween(2, 1.55, 1.60)
	suite.GreaterOrEqual(result, 1.55, "FloatBetween should return a value greater than or equal to min [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.LessOrEqual(result, 1.60, "FloatBetween should return a value less than or equal to max [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	result = suite.number.FloatBetween(2, 1.555, 1.555)
	suite.Equal(1.56, result, "FloatBetween should return the rounded value when min and max are equal and decimals is greater than the number of decimal places in the value [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}
