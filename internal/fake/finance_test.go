package fake

import (
	"math/rand/v2"
	"testing"

	"github.com/lfsc09/kmock/internal/randkit"
	"github.com/stretchr/testify/suite"
)

type MockFinanceTestSuite struct {
	suite.Suite
	seeds   []uint64
	finance *Finance
}

func TestMockFinanceTestSuite(t *testing.T) {
	suite.Run(t, new(MockFinanceTestSuite))
}

func (suite *MockFinanceTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.finance = &Finance{
		Rng: rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1])),
	}
}

/*
	INVALID STATES
*/

func (suite *MockFinanceTestSuite) TestCreditCardHolderWithUnsupportedLocale() {
	_, err := suite.finance.CreditCardHolder("unsupported-locale")
	suite.Error(err, "CreditCardHolder should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

/*
	VALID STATES
*/

func (suite *MockFinanceTestSuite) TestCreditCardVendor() {
	vendor := suite.finance.CreditCardVendor()
	suite.NotEmpty(vendor, "CreditCardVendor should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockFinanceTestSuite) TestCreditCardNumber() {
	ccNumber := suite.finance.CreditCardNumber()
	suite.NotEmpty(ccNumber, "CreditCardNumber should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockFinanceTestSuite) TestCreditCardCVV() {
	cvv := suite.finance.CreditCardCVV()
	suite.NotEmpty(cvv, "CreditCardCVV should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockFinanceTestSuite) TestCreditCardExpirationDate() {
	expirationDate := suite.finance.CreditCardExpirationDate()
	suite.NotEmpty(expirationDate, "CreditCardExpirationDate should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockFinanceTestSuite) TestCreditCardHolder() {
	for locale := range availableLocales {
		holder, err := suite.finance.CreditCardHolder(locale)
		suite.NoError(err, "CreditCardHolder should not return an error for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1])
		suite.NotEmpty(holder, "CreditCardHolder should return a non-empty string for locale: %s [seeds: %d, %d]", locale, suite.seeds[0], suite.seeds[1])
	}
}
