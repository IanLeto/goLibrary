package viperDemo_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/viperDemo"
	"testing"
)

// ViperConfigurationSuite :
type ViperDemoConfigurationSuite struct {
	suite.Suite
	*viperDemo.Configuration
}

func (v *ViperDemoConfigurationSuite) SetupTest() {
	v.Configuration = viperDemo.InitConfiguration()
}

// TestMarshal :
func (s *ViperDemoConfigurationSuite) TestConfig() {
	s.Equal("localhost", s.Configuration.Backend.MySql.Address)
}

// TestViperConfiguration :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(ViperDemoConfigurationSuite))
}
