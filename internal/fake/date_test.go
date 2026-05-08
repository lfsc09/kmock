package fake

import (
	"math/rand/v2"
	"testing"

	"github.com/lfsc09/kmock/internal/randkit"
	"github.com/stretchr/testify/suite"
)

type MockDateTestSuite struct {
	suite.Suite
	seeds []uint64
	date  *Date
}

func TestMockDateTestSuite(t *testing.T) {
	suite.Run(t, new(MockDateTestSuite))
}

func (suite *MockDateTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.date = &Date{
		Rng: rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1])),
	}
}

/*
	VALID STATES
*/

func (suite *MockDateTestSuite) TestDate() {
	date := suite.date.Date("", "", "")
	suite.NotEmpty(date, "Date should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`\d{4}-\d{2}-\d{2}`, date, "Date should match the format YYYY-MM-DD [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	date = suite.date.Date("2000-01-01", "2020-12-31", "")
	suite.NotEmpty(date, "Date should return a non-empty string with custom range [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`\d{4}-\d{2}-\d{2}`, date, "Date should match the format YYYY-MM-DD with custom range [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	date = suite.date.Date("2000-01-01", "2020-12-31", "YYYY/MM/DD")
	suite.NotEmpty(date, "Date should return a non-empty string with custom range [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`\d{4}/\d{2}/\d{2}`, date, "Date should match the format YYYY/MM/DD with custom range [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockDateTestSuite) TestTime() {
	time := suite.date.Time("", "", "")
	suite.NotEmpty(time, "Time should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`\d{2}:\d{2}:\d{2}`, time, "Time should match the format hh:mm:ss [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	time = suite.date.Time("08:00:00", "18:00:00", "")
	suite.NotEmpty(time, "Time should return a non-empty string with custom range [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`\d{2}:\d{2}:\d{2}`, time, "Time should match the format hh:mm:ss with custom range [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	time = suite.date.Time("08:00:00", "18:00:00", "hh:mm")
	suite.NotEmpty(time, "Time should return a non-empty string with custom range and format [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`\d{2}:\d{2}`, time, "Time should match the format hh:mm with custom range [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockDateTestSuite) TestDateTime() {
	dateTime := suite.date.DateTime("", "", "")
	suite.NotEmpty(dateTime, "DateTime should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, dateTime, "DateTime should match the format YYYY-MM-DD hh:mm:ss [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	dateTime = suite.date.DateTime("2000-01-01 08:00:00", "2020-12-31 18:00:00", "")
	suite.NotEmpty(dateTime, "DateTime should return a non-empty string with custom range [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, dateTime, "DateTime should match the format YYYY-MM-DD hh:mm:ss with custom range [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	dateTime = suite.date.DateTime("2000-01-01 08:00:00", "2020-12-31 18:00:00", "YYYY/MM/DD hh:mm")
	suite.NotEmpty(dateTime, "DateTime should return a non-empty string with custom range and format [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`\d{4}/\d{2}/\d{2} \d{2}:\d{2}`, dateTime, "DateTime should match the format YYYY/MM/DD hh:mm with custom range [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockDateTestSuite) TestNow() {
	now := suite.date.Now("")
	suite.NotEmpty(now, "Now should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, now, "Now should match the format YYYY-MM-DD hh:mm:ss [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	now = suite.date.Now("YYYY/MM/DD hh:mm")
	suite.NotEmpty(now, "Now should return a non-empty string with custom format [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`\d{4}/\d{2}/\d{2} \d{2}:\d{2}`, now, "Now should match the format YYYY/MM/DD hh:mm with custom format [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}
