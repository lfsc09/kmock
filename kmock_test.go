package kmock

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MockKmockTestSuite struct {
	suite.Suite
}

func TestMockKmockTestSuite(t *testing.T) {
	suite.Run(t, new(MockKmockTestSuite))
}

func (suite *MockKmockTestSuite) TestKMockInitialization() {
	km1 := New()
	km2 := New()

	suite.NotNil(km1, "KMock instance should not be nil")
	suite.NotNil(km2, "KMock instance should not be nil")
	suite.NotEqual(km1.seeds, km2.seeds, "KMock instances should have different seeds")
	suite.NotEqual(km1.rng, km2.rng, "KMock instances should have different random number generators")
}

func (suite *MockKmockTestSuite) TestKMockAddress() {
	km := New()
	suite.NotNil(km.Address, "Address should not be nil")
	suite.NotEmpty(km.Address.Country(), "Country should not be empty")
	suite.NotEmpty(km.Address.CountryCode(), "CountryCode should not be empty")
	result, err := km.Address.State(EN_US)
	suite.NoError(err, "State should not return an error for EN_US locale")
	suite.NotEmpty(result, "State should not be empty for EN_US locale")
	result, err = km.Address.StateCode(EN_US)
	suite.NoError(err, "StateCode should not return an error for EN_US locale")
	suite.NotEmpty(result, "StateCode should not be empty for EN_US locale")
	result, err = km.Address.City(EN_US)
	suite.NoError(err, "City should not return an error for EN_US locale")
	suite.NotEmpty(result, "City should not be empty for EN_US locale")
	result, err = km.Address.CityFromState(EN_US, "Invalid State Code")
	suite.Error(err, "CityFromState should return an error for EN_US locale and invalid state code")
	suite.Empty(result, "CityFromState should return an empty string for EN_US locale and invalid state code")
	result, err = km.Address.Neighborhood(EN_US)
	suite.NoError(err, "Neighborhood should not return an error for EN_US locale")
	suite.NotEmpty(result, "Neighborhood should not be empty for EN_US locale")
	result, err = km.Address.StreetName(EN_US)
	suite.NoError(err, "StreetName should not return an error for EN_US locale")
	suite.NotEmpty(result, "StreetName should not be empty for EN_US locale")
	suite.NotEmpty(km.Address.StreetNumber(), "StreetNumber should not be empty")
	result, err = km.Address.StreetComplement(EN_US)
	suite.NoError(err, "StreetComplement should not return an error for EN_US locale")
	suite.NotEmpty(result, "StreetComplement should not be empty for EN_US locale")
	result, err = km.Address.ZipCode(EN_US)
	suite.NoError(err, "ZipCode should not return an error for EN_US locale")
	suite.NotEmpty(result, "ZipCode should not be empty for EN_US locale")
	suite.NotEmpty(km.Address.Latitude(), "Latitude should not be empty")
	suite.NotEmpty(km.Address.Longitude(), "Longitude should not be empty")
	suite.NotNil(km.Address.RuntimeDocs(), "RuntimeDocs should not be nil")
}

func (suite *MockKmockTestSuite) TestKMockBoolean() {
	km := New()
	suite.NotNil(km.Boolean, "Boolean should not be nil")
	suite.IsType(true, km.Boolean.Random(), "Random should not be empty")
	suite.IsType(true, km.Boolean.RandomWithProbability(0.7), "RandomWithProbability should not be empty")
	suite.NotNil(km.Boolean.RuntimeDocs(), "RuntimeDocs should not be nil")
}

func (suite *MockKmockTestSuite) TestKMockCar() {
	km := New()
	suite.NotNil(km.Car, "Car should not be nil")
	suite.NotEmpty(km.Car.Brand(), "Brand should not be empty")
	suite.NotEmpty(km.Car.Model(), "Model should not be empty")
	result, err := km.Car.LicensePlate(EN_US)
	suite.NoError(err, "LicensePlate should not return an error for EN_US locale")
	suite.NotEmpty(result, "LicensePlate should not be empty")
	suite.NotEmpty(km.Car.Color(), "Color should not be empty")
	suite.NotNil(km.Car.RuntimeDocs(), "RuntimeDocs should not be nil")
}

func (suite *MockKmockTestSuite) TestKMockCompany() {
	km := New()
	suite.NotNil(km.Company, "Company should not be nil")
	result, err := km.Company.Name(EN_US)
	suite.NoError(err, "Name should not return an error for EN_US locale")
	suite.NotEmpty(result, "Name should not be empty for EN_US locale")
	result, err = km.Company.Dba(EN_US)
	suite.NoError(err, "Dba should not return an error for EN_US locale")
	suite.NotEmpty(result, "Dba should not be empty for EN_US locale")
	result, err = km.Company.Industry(EN_US)
	suite.NoError(err, "Industry should not return an error for EN_US locale")
	suite.NotEmpty(result, "Industry should not be empty for EN_US locale")
	result, err = km.Company.Suffix(EN_US)
	suite.NoError(err, "Suffix should not return an error for EN_US locale")
	suite.NotEmpty(result, "Suffix should not be empty for EN_US locale")
	suite.NotEmpty(km.Company.EIN(), "EIN should not be empty")
	suite.NotEmpty(km.Company.CNPJLegacyValid(), "CNPJLegacyValid should not be empty")
	suite.NotEmpty(km.Company.IE(), "IE should not be empty")
	suite.NotEmpty(km.Company.CNAE(), "CNAE should not be empty")
	suite.NotNil(km.Company.RuntimeDocs(), "RuntimeDocs should not be nil")
}

func (suite *MockKmockTestSuite) TestKMockCurrency() {
	km := New()
	suite.NotNil(km.Currency, "Currency should not be nil")
	suite.NotEmpty(km.Currency.Name(), "Name should not be empty")
	suite.NotEmpty(km.Currency.Code(), "Code should not be empty")
	suite.NotEmpty(km.Currency.Symbol(), "Symbol should not be empty")
	rName, rCode, rSymbol := km.Currency.Full()
	suite.NotEmpty(rName, "Full Name should not be empty")
	suite.NotEmpty(rCode, "Full Code should not be empty")
	suite.NotEmpty(rSymbol, "Full Symbol should not be empty")
	suite.NotNil(km.Currency.RuntimeDocs(), "RuntimeDocs should not be nil")
}

func (suite *MockKmockTestSuite) TestKMockDate() {
	km := New()
	suite.NotNil(km.Date, "Date should not be nil")
	suite.NotEmpty(km.Date.Date("", "", ""), "Date should not be empty")
	suite.NotEmpty(km.Date.Time("", "", ""), "Time should not be empty")
	suite.NotEmpty(km.Date.DateTime("", "", ""), "DateTime should not be empty")
	suite.NotEmpty(km.Date.Now(""), "Now should not be empty")
	suite.NotNil(km.Date.RuntimeDocs(), "RuntimeDocs should not be nil")
}

func (suite *MockKmockTestSuite) TestKMockFile() {
	km := New()
	suite.NotNil(km.File, "File should not be nil")
	suite.NotEmpty(km.File.Name(), "Name should not be empty")
	suite.NotEmpty(km.File.Extension(), "Extension should not be empty")
	suite.NotEmpty(km.File.NameWithExtension(), "NameWithExtension should not be empty")
	suite.NotEmpty(km.File.MimeType(), "MimeType should not be empty")
	suite.NotNil(km.File.RuntimeDocs(), "RuntimeDocs should not be nil")
}

func (suite *MockKmockTestSuite) TestKMockFinance() {
	km := New()
	suite.NotNil(km.Finance, "Finance should not be nil")
	suite.NotEmpty(km.Finance.CreditCardVendor(), "CreditCardVendor should not be empty")
	suite.NotEmpty(km.Finance.CreditCardNumber(), "CreditCardNumber should not be empty")
	suite.NotEmpty(km.Finance.CreditCardCVV(), "CreditCardCVV should not be empty")
	suite.NotEmpty(km.Finance.CreditCardExpirationDate(), "CreditCardExpirationDate should not be empty")
	result, err := km.Finance.CreditCardHolder(EN_US)
	suite.NoError(err, "CreditCardHolder should not return an error for EN_US locale")
	suite.NotEmpty(result, "CreditCardHolder should not be empty")
	suite.NotNil(km.Finance.RuntimeDocs(), "RuntimeDocs should not be nil")
}

func (suite *MockKmockTestSuite) TestKMockID() {
	km := New()
	suite.NotNil(km.ID, "ID should not be nil")
	result, err := km.ID.UUIDv4()
	suite.NoError(err, "UUIDv4 should not return an error")
	suite.NotEmpty(result, "UUIDv4 should not be empty")
	result, err = km.ID.UUIDv6()
	suite.NoError(err, "UUIDv6 should not return an error")
	suite.NotEmpty(result, "UUIDv6 should not be empty")
	result, err = km.ID.UUIDv7()
	suite.NoError(err, "UUIDv7 should not return an error")
	suite.NotEmpty(result, "UUIDv7 should not be empty")
	suite.Equal(km.ID.SequentialID(0), 1, "SequentialID should be 1 when starting from 0")
	suite.NotNil(km.ID.RuntimeDocs(), "RuntimeDocs should not be nil")
}

func (suite *MockKmockTestSuite) TestKMockInternet() {
	km := New()
	suite.NotNil(km.Internet, "Internet should not be nil")
	suite.NotEmpty(km.Internet.Domain(), "Domain should not be empty")
	suite.NotEmpty(km.Internet.DomainWithSubdomain(), "DomainWithSubdomain should not be empty")
	suite.NotEmpty(km.Internet.Ipv4(), "Ipv4 should not be empty")
	suite.NotEmpty(km.Internet.Ipv6(), "Ipv6 should not be empty")
	suite.NotEmpty(km.Internet.MacAddress(), "MacAddress should not be empty")
	suite.NotEmpty(km.Internet.Url(), "Url should not be empty")
	suite.NotEmpty(km.Internet.UserAgent(), "UserAgent should not be empty")
	suite.NotNil(km.Internet.RuntimeDocs(), "RuntimeDocs should not be nil")
}

func (suite *MockKmockTestSuite) TestKMockLorem() {
	km := New()
	suite.NotNil(km.Lorem, "Lorem should not be nil")
	suite.NotEmpty(km.Lorem.Word(), "Word should not be empty")
	suite.NotEmpty(km.Lorem.Sentence(0), "Sentence should not be empty")
	suite.NotEmpty(km.Lorem.Paragraph(0), "Paragraph should not be empty")
	suite.NotNil(km.Lorem.RuntimeDocs(), "RuntimeDocs should not be nil")
}

func (suite *MockKmockTestSuite) TestKMockNumber() {
	km := New()
	suite.NotNil(km.Number, "Number should not be nil")
	suite.NotEmpty(km.Number.Int(), "Int should not be empty")
	suite.NotEmpty(km.Number.IntBetween(1, 10), "IntBetween should not be empty")
	suite.NotEmpty(km.Number.Float(), "Float should not be empty")
	suite.NotEmpty(km.Number.FloatBetween(2, 1.5, 5.5), "FloatBetween should not be empty")
	suite.NotNil(km.Number.RuntimeDocs(), "RuntimeDocs should not be nil")
}

func (suite *MockKmockTestSuite) TestKMockPerson() {
	km := New()
	suite.NotNil(km.Person, "Person should not be nil")
	result, err := km.Person.Name(EN_US)
	suite.NoError(err, "Name should not return an error for EN_US locale")
	suite.NotEmpty(result, "Name should not be empty for EN_US locale")
	result, err = km.Person.FirstName(EN_US)
	suite.NoError(err, "FirstName should not return an error for EN_US locale")
	suite.NotEmpty(result, "FirstName should not be empty for EN_US locale")
	result, err = km.Person.MiddleName(EN_US)
	suite.NoError(err, "MiddleName should not return an error for EN_US locale")
	suite.NotEmpty(result, "MiddleName should not be empty for EN_US locale")
	result, err = km.Person.LastName(EN_US)
	suite.NoError(err, "LastName should not return an error for EN_US locale")
	suite.NotEmpty(result, "LastName should not be empty for EN_US locale")
	suite.NotEmpty(km.Person.Phone(), "Phone should not be empty")
	result, err = km.Person.Email()
	suite.NoError(err, "Email should not return an error")
	suite.NotEmpty(result, "Email should not be empty")
	result, err = km.Person.Username()
	suite.NoError(err, "Username should not return an error")
	suite.NotEmpty(result, "Username should not be empty")
	result, err = km.Person.Password(PWD_WEAK)
	suite.NoError(err, "Password (weak) should not return an error")
	suite.NotEmpty(result, "Password (weak) should not be empty")
	result, err = km.Person.Password(PWD_MEDIUM)
	suite.NoError(err, "Password (medium) should not return an error")
	suite.NotEmpty(result, "Password (medium) should not be empty")
	result, err = km.Person.Password(PWD_STRONG)
	suite.NoError(err, "Password (strong) should not return an error")
	suite.NotEmpty(result, "Password (strong) should not be empty")
	result, err = km.Person.JobTitle(EN_US)
	suite.NoError(err, "JobTitle should not return an error for EN_US locale")
	suite.NotEmpty(result, "JobTitle should not be empty for EN_US locale")
	suite.NotEmpty(km.Person.CPFValid(), "CPFValid should not be empty")
	suite.NotNil(km.Person.RuntimeDocs(), "RuntimeDocs should not be nil")
}

func (suite *MockKmockTestSuite) TestKMockRegex() {
	km := New()
	suite.NotNil(km.Regex, "Regex should not be nil")
	result, err := km.Regex.Generate("\\d{5}")
	suite.NoError(err, "Generate should not return an error for valid regex pattern")
	suite.NotEmpty(result, "Generate should not be empty for valid regex pattern")
	suite.NotNil(km.Regex.RuntimeDocs(), "RuntimeDocs should not be nil")
}
