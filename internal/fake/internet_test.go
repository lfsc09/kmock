package fake

import (
	"math/rand/v2"
	"testing"

	"github.com/lfsc09/kmock/internal/randkit"
	"github.com/stretchr/testify/suite"
)

type MockInternetTestSuite struct {
	suite.Suite
	seeds    []uint64
	internet *Internet
}

func TestMockInternetTestSuite(t *testing.T) {
	suite.Run(t, new(MockInternetTestSuite))
}

func (suite *MockInternetTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.internet = &Internet{
		Rng: rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1])),
	}
}

/*
	VALID STATES
*/

func (suite *MockInternetTestSuite) TestDomain() {
	domain := suite.internet.Domain()
	suite.NotEmpty(domain, "Domain should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockInternetTestSuite) TestDomainWithSubdomain() {
	domain := suite.internet.DomainWithSubdomain()
	suite.NotEmpty(domain, "DomainWithSubdomain should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockInternetTestSuite) TestIpv4() {
	ipv4 := suite.internet.Ipv4()
	suite.NotEmpty(ipv4, "Ipv4 should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockInternetTestSuite) TestIpv6() {
	ipv6 := suite.internet.Ipv6()
	suite.NotEmpty(ipv6, "Ipv6 should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockInternetTestSuite) TestMacAddress() {
	macAddress := suite.internet.MacAddress()
	suite.NotEmpty(macAddress, "MacAddress should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockInternetTestSuite) TestUrl() {
	url := suite.internet.Url()
	suite.NotEmpty(url, "Url should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockInternetTestSuite) TestUserAgent() {
	userAgent := suite.internet.UserAgent()
	suite.NotEmpty(userAgent, "UserAgent should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}
