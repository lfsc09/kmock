package fake

import (
	"testing"

	"github.com/lfsc09/kmock/internal/randkit"
	"github.com/stretchr/testify/suite"
)

type MockIDTestSuite struct {
	suite.Suite
	seeds []uint64
	id    *ID
}

func TestMockIDTestSuite(t *testing.T) {
	suite.Run(t, new(MockIDTestSuite))
}

func (suite *MockIDTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.id = &ID{}
}

/*
	VALID STATES
*/

func (suite *MockIDTestSuite) TestUUIDv4() {
	uuidv4, err := suite.id.UUIDv4()
	suite.NoError(err, "UUIDv4 should not return an error [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$`, uuidv4, "UUIDv4 should return a valid UUIDv4 string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockIDTestSuite) TestUUIDv6() {
	uuidv6, err := suite.id.UUIDv6()
	suite.NoError(err, "UUIDv6 should not return an error [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`^[a-f0-9]{8}-[a-f0-9]{4}-6[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$`, uuidv6, "UUIDv6 should return a valid UUIDv6 string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockIDTestSuite) TestUUIDv7() {
	uuidv7, err := suite.id.UUIDv7()
	suite.NoError(err, "UUIDv7 should not return an error [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	suite.Regexp(`^[a-f0-9]{8}-[a-f0-9]{4}-7[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$`, uuidv7, "UUIDv7 should return a valid UUIDv7 string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockIDTestSuite) TestSequentialID() {
	startFrom := 100
	for range 5 {
		expectedID := startFrom + 1
		startFrom = suite.id.SequentialID(startFrom)
		suite.Equal(expectedID, startFrom, "SequentialID should return the correct sequential ID [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
	}
}
