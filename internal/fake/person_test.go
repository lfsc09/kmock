package fake

import (
	"math/rand/v2"
	"testing"

	"github.com/lfsc09/kmock/internal/randkit"
	"github.com/stretchr/testify/suite"
)

type MockPersonTestSuite struct {
	suite.Suite
	seeds  []uint64
	person *Person
}

func TestMockPersonTestSuite(t *testing.T) {
	suite.Run(t, new(MockPersonTestSuite))
}

func (suite *MockPersonTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.person = &Person{
		Rng: rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1])),
	}
}

/*
	INVALID STATES
*/

func (suite *MockPersonTestSuite) TestNameWithUnsupportedLocale() {
	_, err := suite.person.Name("unsupported-locale")
	suite.Error(err, "Name should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockPersonTestSuite) TestFirstNameWithUnsupportedLocale() {
	_, err := suite.person.FirstName("unsupported-locale")
	suite.Error(err, "FirstName should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockPersonTestSuite) TestMiddleNameWithUnsupportedLocale() {
	_, err := suite.person.MiddleName("unsupported-locale")
	suite.Error(err, "MiddleName should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockPersonTestSuite) TestLastNameWithUnsupportedLocale() {
	_, err := suite.person.LastName("unsupported-locale")
	suite.Error(err, "LastName should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockPersonTestSuite) TestJobTitleWithUnsupportedLocale() {
	_, err := suite.person.JobTitle("unsupported-locale")
	suite.Error(err, "JobTitle should return an error when an unsupported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

/*
	VALID STATES
*/

func (suite *MockPersonTestSuite) TestName() {
	for locale := range availableLocales {
		name, err := suite.person.Name(locale)
		suite.NoError(err, "Name should not return an error when a supported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
		suite.NotEmpty(name, "Name should return a non-empty string when a supported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	}
}

func (suite *MockPersonTestSuite) TestFirstName() {
	for locale := range availableLocales {
		firstName, err := suite.person.FirstName(locale)
		suite.NoError(err, "FirstName should not return an error when a supported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
		suite.NotEmpty(firstName, "FirstName should return a non-empty string when a supported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	}
}

func (suite *MockPersonTestSuite) TestMiddleName() {
	for locale := range availableLocales {
		middleName, err := suite.person.MiddleName(locale)
		suite.NoError(err, "MiddleName should not return an error when a supported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
		suite.NotEmpty(middleName, "MiddleName should return a non-empty string when a supported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	}
}

func (suite *MockPersonTestSuite) TestLastName() {
	for locale := range availableLocales {
		lastName, err := suite.person.LastName(locale)
		suite.NoError(err, "LastName should not return an error when a supported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
		suite.NotEmpty(lastName, "LastName should return a non-empty string when a supported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	}
}

func (suite *MockPersonTestSuite) TestPhone() {
	phone := suite.person.Phone()
	suite.NotEmpty(phone, "Phone should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockPersonTestSuite) TestEmail() {
	email, err := suite.person.Email()
	suite.NoError(err, "Email should not return an error [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.NotEmpty(email, "Email should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockPersonTestSuite) TestUsername() {
	username, err := suite.person.Username()
	suite.NoError(err, "Username should not return an error [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.NotEmpty(username, "Username should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockPersonTestSuite) TestPassword() {
	weak, err := suite.person.Password(PWD_WEAK)
	suite.NoError(err, "Password (weak) should not return an error for valid strength [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.NotEmpty(weak, "Password (weak) should return a non-empty string for valid strength [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	medium, err := suite.person.Password(PWD_MEDIUM)
	suite.NoError(err, "Password (medium) should not return an error for valid strength [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.NotEmpty(medium, "Password (medium) should return a non-empty string for valid strength [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	strong, err := suite.person.Password(PWD_STRONG)
	suite.NoError(err, "Password (strong) should not return an error for valid strength [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.NotEmpty(strong, "Password (strong) should return a non-empty string for valid strength [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockPersonTestSuite) TestJobTitle() {
	for locale := range availableLocales {
		jobTitle, err := suite.person.JobTitle(locale)
		suite.NoError(err, "JobTitle should not return an error when a supported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
		suite.NotEmpty(jobTitle, "JobTitle should return a non-empty string when a supported locale is provided [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	}
}

func (suite *MockPersonTestSuite) TestCPFValid() {
	cpf := suite.person.CPFValid()
	suite.NotEmpty(cpf, "CPFValid should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`, cpf, "CPFValid should return a string in the format XXX.XXX.XXX-XX [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}
