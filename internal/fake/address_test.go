package fake

import (
	"fmt"
	"kmock/internal/randkit"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockAddressTestSuite struct {
	suite.Suite
	seeds   []uint64
	rng     *rand.Rand
	address *Address
}

func TestMockAddressTestSuite(t *testing.T) {
	suite.Run(t, new(MockAddressTestSuite))
}

func (suite *MockAddressTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.rng = rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1]))
	suite.address = &Address{}
}

/*
	INVALID STATES
*/

func (suite *MockAddressTestSuite) TestStateWithUnsupportedLocale() {
	assert.Panics(suite.T(), func() {
		suite.address.State(suite.rng, "unsupported-locale")
	}, fmt.Sprintf("State should panic when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestStateCodeWithUnsupportedLocale() {
	assert.Panics(suite.T(), func() {
		suite.address.StateCode(suite.rng, "unsupported-locale")
	}, fmt.Sprintf("StateCode should panic when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestCityWithUnsupportedLocale() {
	assert.Panics(suite.T(), func() {
		suite.address.City(suite.rng, "unsupported-locale")
	}, fmt.Sprintf("City should panic when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestCityFromStateWithUnsupportedLocale() {
	assert.Panics(suite.T(), func() {
		suite.address.CityFromState(suite.rng, "unsupported-locale", "CA")
	}, fmt.Sprintf("CityFromState should panic when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestCityFromStateWithUnsupportedState() {
	for locale := range availableLocales {
		assert.Panics(suite.T(), func() {
			suite.address.CityFromState(suite.rng, locale, "Unsupported State Code")
		}, fmt.Sprintf("CityFromState should panic when an unsupported state is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestNeighborhoodWithUnsupportedLocale() {
	assert.Panics(suite.T(), func() {
		suite.address.Neighborhood(suite.rng, "unsupported-locale")
	}, fmt.Sprintf("Neighborhood should panic when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestStreetNameWithUnsupportedLocale() {
	assert.Panics(suite.T(), func() {
		suite.address.StreetName(suite.rng, "unsupported-locale")
	}, fmt.Sprintf("StreetName should panic when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestStreetComplementWithUnsupportedLocale() {
	assert.Panics(suite.T(), func() {
		suite.address.StreetComplement(suite.rng, "unsupported-locale")
	}, fmt.Sprintf("StreetComplement should panic when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestZipCodeWithUnsupportedLocale() {
	assert.Panics(suite.T(), func() {
		suite.address.ZipCode(suite.rng, "unsupported-locale")
	}, fmt.Sprintf("ZipCode should panic when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

/*
	VALID STATES
*/

func (suite *MockAddressTestSuite) TestCountry() {
	country := suite.address.Country(suite.rng)
	assert.NotEmpty(suite.T(), country, fmt.Sprintf("Country should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestCountryCode() {
	code := suite.address.CountryCode(suite.rng)
	assert.NotEmpty(suite.T(), code, fmt.Sprintf("CountryCode should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestState() {
	for locale := range availableLocales {
		state := suite.address.State(suite.rng, locale)
		assert.NotEmpty(suite.T(), state, fmt.Sprintf("State should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestStateCode() {
	for locale := range availableLocales {
		code := suite.address.StateCode(suite.rng, locale)
		assert.NotEmpty(suite.T(), code, fmt.Sprintf("StateCode should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestCity() {
	for locale := range availableLocales {
		city := suite.address.City(suite.rng, locale)
		assert.NotEmpty(suite.T(), city, fmt.Sprintf("City should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestCityFromState() {
	for locale := range availableLocales {
		pickedStateCode := suite.address.StateCode(suite.rng, locale)
		city := suite.address.CityFromState(suite.rng, locale, pickedStateCode)
		assert.NotEmpty(suite.T(), city, fmt.Sprintf("CityFromState should return a non-empty string for locale: %s and state code: %s [seeds: %d, %d]", locale, pickedStateCode, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestNeighborhood() {
	for locale := range availableLocales {
		neighborhood := suite.address.Neighborhood(suite.rng, locale)
		assert.NotEmpty(suite.T(), neighborhood, fmt.Sprintf("Neighborhood should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestStreetName() {
	for locale := range availableLocales {
		streetName := suite.address.StreetName(suite.rng, locale)
		assert.NotEmpty(suite.T(), streetName, fmt.Sprintf("StreetName should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestStreetNumber() {
	streetNumber := suite.address.StreetNumber(suite.rng)
	assert.NotEmpty(suite.T(), streetNumber, fmt.Sprintf("StreetNumber should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestStreetComplement() {
	for locale := range availableLocales {
		complement := suite.address.StreetComplement(suite.rng, locale)
		assert.NotEmpty(suite.T(), complement, fmt.Sprintf("StreetComplement should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestZipCode() {
	for locale := range availableLocales {
		zipCode := suite.address.ZipCode(suite.rng, locale)
		assert.NotEmpty(suite.T(), zipCode, fmt.Sprintf("ZipCode should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestLatitude() {
	latitude := suite.address.Latitude(suite.rng)
	suite.Assert().GreaterOrEqual(latitude, -90.0, fmt.Sprintf("Latitude should be greater than or equal to -90 [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
	suite.Assert().LessOrEqual(latitude, 90.0, fmt.Sprintf("Latitude should be less than or equal to 90 [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestLongitude() {
	longitude := suite.address.Longitude(suite.rng)
	suite.Assert().GreaterOrEqual(longitude, -180.0, fmt.Sprintf("Longitude should be greater than or equal to -180 [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
	suite.Assert().LessOrEqual(longitude, 180.0, fmt.Sprintf("Longitude should be less than or equal to 180 [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}
