package utils_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/config"
	"testing"
)

type RelationSuite struct {
	suite.Suite
}

func (s *RelationSuite) SetupTest() {
}

func BenchmarkStr(b *testing.B) {
	var confi *config.Config
	for i := 0; i < b.N; i++ {
		confi = &config.Config{}
	}
	_ = confi
}

func BenchmarkRefectNew(b *testing.B) {
	var confi interface{}
	for i := 0; i < b.N; i++ {
		confi = &config.Config{}
		confi = confi.(*config.Config)
	}
	_ = confi
}

// mysql 常用场合
func (s *RelationSuite) TestMySQL() {

}

func TestConvBench(t *testing.T) {
	suite.Run(t, new(RelationSuite))

}
