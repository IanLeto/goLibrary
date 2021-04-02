package config_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/config"
	_ "goLibrary/config"
	"goLibrary/testsuites"
	"testing"
)

// TestInitConfigSuit :
type TestInitConfigSuit struct {
	suite.Suite
	testsuites.ConfigSuit
}

func (s *TestInitConfigSuit) SetupTest() {
	s.InitConfigSuit()
}

// TestMarshal :
func (s *TestInitConfigSuit) TestConfig() {
	s.Equal("debug", config.BaseConfig.RunMode)
}

// TestViperConfiguration :
func TestConfiguration(t *testing.T) {
	suite.Run(t, new(TestInitConfigSuit))
}
