package fake

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/lfsc09/kmock/internal/randkit"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockAddressTestSuite struct {
	suite.Suite
	seeds   []uint64
	address *Address
}

func TestMockAddressTestSuite(t *testing.T) {
	suite.Run(t, new(MockAddressTestSuite))
}

func (suite *MockAddressTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.address = &Address{
		Rng: rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1])),
	}
}

/*
	INVALID STATES
*/

func (suite *MockAddressTestSuite) TestStateWithUnsupportedLocale() {
	_, err := suite.address.State("unsupported-locale")
	assert.Error(suite.T(), err, fmt.Sprintf("State should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestStateCodeWithUnsupportedLocale() {
	_, err := suite.address.StateCode("unsupported-locale")
	assert.Error(suite.T(), err, fmt.Sprintf("StateCode should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestCityWithUnsupportedLocale() {
	_, err := suite.address.City("unsupported-locale")
	assert.Error(suite.T(), err, fmt.Sprintf("City should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestCityFromStateWithUnsupportedLocale() {
	_, err := suite.address.CityFromState("unsupported-locale", "CA")
	assert.Error(suite.T(), err, fmt.Sprintf("CityFromState should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestCityFromStateWithUnsupportedState() {
	for locale := range availableLocales {
		_, err := suite.address.CityFromState(locale, "Unsupported State Code")
		assert.Error(suite.T(), err, fmt.Sprintf("CityFromState should return an error when an unsupported state is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestNeighborhoodWithUnsupportedLocale() {
	_, err := suite.address.Neighborhood("unsupported-locale")
	assert.Error(suite.T(), err, fmt.Sprintf("Neighborhood should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestStreetNameWithUnsupportedLocale() {
	_, err := suite.address.StreetName("unsupported-locale")
	assert.Error(suite.T(), err, fmt.Sprintf("StreetName should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestStreetComplementWithUnsupportedLocale() {
	_, err := suite.address.StreetComplement("unsupported-locale")
	assert.Error(suite.T(), err, fmt.Sprintf("StreetComplement should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestZipCodeWithUnsupportedLocale() {
	_, err := suite.address.ZipCode("unsupported-locale")
	assert.Error(suite.T(), err, fmt.Sprintf("ZipCode should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

/*
	VALID STATES
*/

func (suite *MockAddressTestSuite) TestCountry() {
	country := suite.address.Country()
	assert.NotEmpty(suite.T(), country, fmt.Sprintf("Country should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestCountryCode() {
	code := suite.address.CountryCode()
	assert.NotEmpty(suite.T(), code, fmt.Sprintf("CountryCode should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestState() {
	for locale := range availableLocales {
		state, err := suite.address.State(locale)
		assert.NoError(suite.T(), err, fmt.Sprintf("State should not return an error for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
		assert.NotEmpty(suite.T(), state, fmt.Sprintf("State should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestStateCode() {
	for locale := range availableLocales {
		code, err := suite.address.StateCode(locale)
		assert.NoError(suite.T(), err, fmt.Sprintf("StateCode should not return an error for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
		assert.NotEmpty(suite.T(), code, fmt.Sprintf("StateCode should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestCity() {
	for locale := range availableLocales {
		city, err := suite.address.City(locale)
		assert.NoError(suite.T(), err, fmt.Sprintf("City should not return an error for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
		assert.NotEmpty(suite.T(), city, fmt.Sprintf("City should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestCityFromState() {
	for locale := range availableLocales {
		pickedStateCode, err := suite.address.StateCode(locale)
		assert.NoError(suite.T(), err, fmt.Sprintf("StateCode should not return an error for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
		city, err := suite.address.CityFromState(locale, pickedStateCode)
		assert.NoError(suite.T(), err, fmt.Sprintf("CityFromState should not return an error for locale: %s and state code: %s [seeds: %d, %d]", locale, pickedStateCode, suite.seeds[0], suite.seeds[1]))
		assert.NotEmpty(suite.T(), city, fmt.Sprintf("CityFromState should return a non-empty string for locale: %s and state code: %s [seeds: %d, %d]", locale, pickedStateCode, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestNeighborhood() {
	for locale := range availableLocales {
		neighborhood, err := suite.address.Neighborhood(locale)
		assert.NoError(suite.T(), err, fmt.Sprintf("Neighborhood should not return an error for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
		assert.NotEmpty(suite.T(), neighborhood, fmt.Sprintf("Neighborhood should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestStreetName() {
	for locale := range availableLocales {
		streetName, err := suite.address.StreetName(locale)
		assert.NoError(suite.T(), err, fmt.Sprintf("StreetName should not return an error for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
		assert.NotEmpty(suite.T(), streetName, fmt.Sprintf("StreetName should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestStreetNumber() {
	streetNumber := suite.address.StreetNumber()
	assert.NotEmpty(suite.T(), streetNumber, fmt.Sprintf("StreetNumber should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestStreetComplement() {
	for locale := range availableLocales {
		complement, err := suite.address.StreetComplement(locale)
		assert.NoError(suite.T(), err, fmt.Sprintf("StreetComplement should not return an error for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
		assert.NotEmpty(suite.T(), complement, fmt.Sprintf("StreetComplement should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestZipCode() {
	for locale := range availableLocales {
		zipCode, err := suite.address.ZipCode(locale)
		assert.NoError(suite.T(), err, fmt.Sprintf("ZipCode should not return an error for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
		assert.NotEmpty(suite.T(), zipCode, fmt.Sprintf("ZipCode should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockAddressTestSuite) TestLatitude() {
	latitude := suite.address.Latitude()
	suite.Assert().GreaterOrEqual(latitude, -90.0, fmt.Sprintf("Latitude should be greater than or equal to -90 [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
	suite.Assert().LessOrEqual(latitude, 90.0, fmt.Sprintf("Latitude should be less than or equal to 90 [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockAddressTestSuite) TestLongitude() {
	longitude := suite.address.Longitude()
	suite.Assert().GreaterOrEqual(longitude, -180.0, fmt.Sprintf("Longitude should be greater than or equal to -180 [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
	suite.Assert().LessOrEqual(longitude, 180.0, fmt.Sprintf("Longitude should be less than or equal to 180 [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}
