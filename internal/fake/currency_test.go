package fake

import (
	"math/rand/v2"
	"testing"

	"github.com/lfsc09/kmock/internal/randkit"
	"github.com/stretchr/testify/suite"
)

type MockCurrencyTestSuite struct {
	suite.Suite
	seeds    []uint64
	currency *Currency
}

func TestMockCurrencyTestSuite(t *testing.T) {
	suite.Run(t, new(MockCurrencyTestSuite))
}

func (suite *MockCurrencyTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.currency = &Currency{
		Rng: rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1])),
	}
}

/*
	VALID STATES
*/

func (suite *MockCurrencyTestSuite) TestName() {
	name := suite.currency.Name()
	suite.NotEmpty(name, "Name should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockCurrencyTestSuite) TestCode() {
	code := suite.currency.Code()
	suite.NotEmpty(code, "Code should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockCurrencyTestSuite) TestSymbol() {
	symbol := suite.currency.Symbol()
	suite.NotEmpty(symbol, "Symbol should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockCurrencyTestSuite) TestFull() {
	name, code, symbol := suite.currency.Full()
	suite.NotEmpty(name, "Full should return a non-empty name string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.NotEmpty(code, "Full should return a non-empty code string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.NotEmpty(symbol, "Full should return a non-empty symbol string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}
