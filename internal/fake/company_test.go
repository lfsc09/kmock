package fake

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/lfsc09/kmock/internal/randkit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockCompanyTestSuite struct {
	suite.Suite
	seeds   []uint64
	company *Company
}

func TestMockCompanyTestSuite(t *testing.T) {
	suite.Run(t, new(MockCompanyTestSuite))
}

func (suite *MockCompanyTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.company = &Company{
		Rng: rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1])),
	}
}

/*
	INVALID STATES
*/

func (suite *MockCompanyTestSuite) TestNameWithUnsupportedLocale() {
	_, err := suite.company.Name("unsupported-locale")
	assert.Error(suite.T(), err, fmt.Sprintf("Name should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockCompanyTestSuite) TestDbaWithUnsupportedLocale() {
	_, err := suite.company.Dba("unsupported-locale")
	assert.Error(suite.T(), err, fmt.Sprintf("Dba should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockCompanyTestSuite) TestIndustryWithUnsupportedLocale() {
	_, err := suite.company.Industry("unsupported-locale")
	assert.Error(suite.T(), err, fmt.Sprintf("Industry should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockCompanyTestSuite) TestSuffixWithUnsupportedLocale() {
	_, err := suite.company.Suffix("unsupported-locale")
	assert.Error(suite.T(), err, fmt.Sprintf("Suffix should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

/*
	VALID STATES
*/

func (suite *MockCompanyTestSuite) TestName() {
	for locale := range availableLocales {
		name, err := suite.company.Name(locale)
		assert.NoError(suite.T(), err, fmt.Sprintf("Name should not return an error for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
		assert.NotEmpty(suite.T(), name, fmt.Sprintf("Name should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockCompanyTestSuite) TestDba() {
	for locale := range availableLocales {
		dba, err := suite.company.Dba(locale)
		assert.NoError(suite.T(), err, fmt.Sprintf("Dba should not return an error for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
		assert.NotEmpty(suite.T(), dba, fmt.Sprintf("Dba should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockCompanyTestSuite) TestIndustry() {
	for locale := range availableLocales {
		industry, err := suite.company.Industry(locale)
		assert.NoError(suite.T(), err, fmt.Sprintf("Industry should not return an error for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
		assert.NotEmpty(suite.T(), industry, fmt.Sprintf("Industry should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockCompanyTestSuite) TestSuffix() {
	for locale := range availableLocales {
		suffix, err := suite.company.Suffix(locale)
		assert.NoError(suite.T(), err, fmt.Sprintf("Suffix should not return an error for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
		assert.NotEmpty(suite.T(), suffix, fmt.Sprintf("Suffix should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1]))
	}
}

func (suite *MockCompanyTestSuite) TestEIN() {
	ein := suite.company.EIN()
	assert.NotEmpty(suite.T(), ein, fmt.Sprintf("EIN should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockCompanyTestSuite) TestCNPJLegacyValid() {
	cnpj := suite.company.CNPJLegacyValid()
	assert.NotEmpty(suite.T(), cnpj, fmt.Sprintf("CNPJLegacyValid should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
	assert.Regexp(suite.T(), `^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`, cnpj, fmt.Sprintf("CNPJLegacyValid should return a string in the format 00.000.000/0000-00 [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockCompanyTestSuite) TestIE() {
	ie := suite.company.IE()
	assert.NotEmpty(suite.T(), ie, fmt.Sprintf("IE should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockCompanyTestSuite) TestCNAE() {
	cnae := suite.company.CNAE()
	assert.NotEmpty(suite.T(), cnae, fmt.Sprintf("CNAE should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}
