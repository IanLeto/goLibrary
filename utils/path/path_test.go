package path

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PathSuite struct {
	suite.Suite
}

func (s *PathSuite) SetupTest() {
}

// TestMarshal :
func (s *PathSuite) TestHelloWorld() {
	s.Equal("/Users/ian/go/src/goLibrary/redis/redisCli.go",GetFilePath("redis/redisCli.go"))
}

func TestPath(t *testing.T) {
	suite.Run(t, new(PathSuite))
}
