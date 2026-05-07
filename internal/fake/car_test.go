package fake

import (
	"fmt"
	"kmock/internal/randkit"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockCarTestSuite struct {
	suite.Suite
	seeds []uint64
	car   *Car
}

func TestMockCarTestSuite(t *testing.T) {
	suite.Run(t, new(MockCarTestSuite))
}

func (suite *MockCarTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	rng := rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1]))
	suite.car = &Car{
		Rng: rng,
	}
}

/*
	INVALID STATES
*/

func (suite *MockCarTestSuite) TestLicensePlateWithUnsupportedLocale() {
	assert.Panics(suite.T(), func() {
		suite.car.LicensePlate("unsupported-locale")
	}, fmt.Sprintf("LicensePlate should panic when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

/*
	VALID STATES
*/

func (suite *MockCarTestSuite) TestBrand() {
	brand := suite.car.Brand()
	assert.NotEmpty(suite.T(), brand, fmt.Sprintf("Brand should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockCarTestSuite) TestModel() {
	model := suite.car.Model()
	assert.NotEmpty(suite.T(), model, fmt.Sprintf("Model should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockCarTestSuite) TestLicensePlate() {
	for locale := range availableLocales {
		licensePlate := suite.car.LicensePlate(locale)
		assert.NotEmpty(suite.T(), licensePlate, fmt.Sprintf("LicensePlate should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockCarTestSuite) TestColor() {
	color := suite.car.Color()
	assert.NotEmpty(suite.T(), color, fmt.Sprintf("Color should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}
