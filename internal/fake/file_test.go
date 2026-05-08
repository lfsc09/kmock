package fake

import (
	"math/rand/v2"
	"testing"

	"github.com/lfsc09/kmock/internal/randkit"
	"github.com/stretchr/testify/suite"
)

type MockFileTestSuite struct {
	suite.Suite
	seeds []uint64
	file  *File
}

func TestMockFileTestSuite(t *testing.T) {
	suite.Run(t, new(MockFileTestSuite))
}

func (suite *MockFileTestSuite) SetupSuite() {
	suite.seeds = []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	suite.file = &File{
		Rng: rand.New(rand.NewPCG(suite.seeds[0], suite.seeds[1])),
	}
}

/*
	VALID STATES
*/

func (suite *MockFileTestSuite) TestName() {
	name := suite.file.Name()
	suite.NotEmpty(name, "Name should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockFileTestSuite) TestExtension() {
	extension := suite.file.Extension()
	suite.NotEmpty(extension, "Extension should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockFileTestSuite) TestNameWithExtension() {
	nameWithExtension := suite.file.NameWithExtension()
	suite.NotEmpty(nameWithExtension, "NameWithExtension should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}

func (suite *MockFileTestSuite) TestMimeType() {
	mimeType := suite.file.MimeType()
	suite.NotEmpty(mimeType, "MimeType should return a non-empty string [seeds: %d, %d]", suite.seeds[0], suite.seeds[1])
}
