package fake

import (
	"fmt"
	"math/rand/v2"
	"strings"
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
	assert.True(suite.T(), cnpjLegacyIsValid(cnpj), fmt.Sprintf("CNPJLegacyValid should return a valid CNPJ number [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockCompanyTestSuite) TestCNPJLegacyInvalid() {
	cnpj := suite.company.CNPJLegacyInvalid()
	assert.NotEmpty(suite.T(), cnpj, fmt.Sprintf("CNPJLegacyInvalid should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
	assert.Regexp(suite.T(), `^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`, cnpj, fmt.Sprintf("CNPJLegacyInvalid should return a string in the format 00.000.000/0000-00 [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
	assert.False(suite.T(), cnpjLegacyIsValid(cnpj), fmt.Sprintf("CNPJLegacyInvalid should return an invalid CNPJ number [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockCompanyTestSuite) TestCNPJAlphanumericValid() {
	cnpj := suite.company.CNPJAlphanumericValid()
	assert.NotEmpty(suite.T(), cnpj, fmt.Sprintf("CNPJAlphanumericValid should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
	assert.Regexp(suite.T(), `^[A-Z0-9]{2}\.[A-Z0-9]{3}\.[A-Z0-9]{3}/[A-Z0-9]{4}-\d{2}$`, cnpj, fmt.Sprintf("CNPJAlphanumericValid should return a string in the format XX.XXX.XXX/XXXX-00 with alphanumeric characters [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
	assert.True(suite.T(), cnpjAlphanumericIsValid(cnpj), fmt.Sprintf("CNPJAlphanumericValid should return a valid alphanumeric CNPJ number [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockCompanyTestSuite) TestCNPJAlphanumericInvalid() {
	cnpj := suite.company.CNPJAlphanumericInvalid()
	assert.NotEmpty(suite.T(), cnpj, fmt.Sprintf("CNPJAlphanumericInvalid should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
	assert.Regexp(suite.T(), `^[A-Z0-9]{2}\.[A-Z0-9]{3}\.[A-Z0-9]{3}/[A-Z0-9]{4}-\d{2}$`, cnpj, fmt.Sprintf("CNPJAlphanumericInvalid should return a string in the format XX.XXX.XXX/XXXX-00 with alphanumeric characters [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
	assert.False(suite.T(), cnpjAlphanumericIsValid(cnpj), fmt.Sprintf("CNPJAlphanumericInvalid should return an invalid alphanumeric CNPJ number [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockCompanyTestSuite) TestIE() {
	ie := suite.company.IE()
	assert.NotEmpty(suite.T(), ie, fmt.Sprintf("IE should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

func (suite *MockCompanyTestSuite) TestCNAE() {
	cnae := suite.company.CNAE()
	assert.NotEmpty(suite.T(), cnae, fmt.Sprintf("CNAE should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1]))
}

// cnpjLegacyIsValid is a helper function to validate legacy CNPJ numbers. It checks the format and calculates the checksum digits to ensure the CNPJ is valid.
func cnpjLegacyIsValid(cnpj string) bool {
	s := strings.NewReplacer(".", "", "/", "", "-", "").Replace(cnpj)
	if len(s) != 14 {
		return false
	}
	digits := make([]int, 14)
	for i, ch := range s {
		digits[i] = int(ch - '0')
	}
	multipliers1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	multipliers2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	checkDigit := func(d []int, mults []int) int {
		sum := 0
		for i, m := range mults {
			sum += d[i] * m
		}
		if mod := sum % 11; mod >= 2 {
			return 11 - mod
		}
		return 0
	}
	return digits[12] == checkDigit(digits, multipliers1) &&
		digits[13] == checkDigit(digits[:13], multipliers2)
}

// cnpjAlphanumericIsValid is a helper function to validate alphanumeric CNPJ numbers. It checks the format and calculates the checksum digits to ensure the CNPJ is valid.
func cnpjAlphanumericIsValid(cnpj string) bool {
	s := strings.NewReplacer(".", "", "/", "", "-", "").Replace(cnpj)
	if len(s) != 14 {
		return false
	}
	b := []byte(s)
	multipliers1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	multipliers2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	checkDigit := func(bs []byte, mults []int) byte {
		sum := 0
		for i, m := range mults {
			sum += (int(bs[i]) - 48) * m
		}
		if mod := sum % 11; mod >= 2 {
			return byte(48 + 11 - mod)
		}
		return byte(48)
	}
	return b[12] == checkDigit(b, multipliers1) &&
		b[13] == checkDigit(b[:13], multipliers2)
}
