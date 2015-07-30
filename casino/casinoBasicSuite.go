package casino_new

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CasinoBasicSuite struct {
	suite.Suite
}

func (suite *CasinoBasicSuite) AssertEquals(expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return suite.Equal(expected, actual, msgAndArgs...)
}

func (suite *CasinoBasicSuite) AssertFalse(value bool, msgAndArgs ...interface{}) bool {
	return assert.False(suite.T(), value, msgAndArgs...)
}

func (suite *CasinoBasicSuite) AssertTrue(value bool, msgAndArgs ...interface{}) bool {
	return assert.True(suite.T(), value, msgAndArgs...)
}

func (suite *CasinoBasicSuite) AssertNotNil(object interface{}, msgAndArgs ...interface{}) bool {
	return assert.NotNil(suite.T(), object, msgAndArgs...)
}