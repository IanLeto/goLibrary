package utils_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/utils"

	"testing"
)

type ConvSuite struct {
	suite.Suite
}

func (s *ConvSuite) SetupTest() {
}

// mysql 常用场合
func (s *ConvSuite) TestMySQL() {
	cases := []struct {
		ori    interface{}
		except interface{}
	}{
		{ori: []string{"1", "2", "3"}, except: "1,2,3"},
		{ori: "1,2,3", except: []string{"1", "2", "3"}},
	}
	s.Equal(cases[0].except, utils.ArrToString([]string{"1", "2", "3"}))
	s.Equal(cases[1].except, utils.StringToArr(utils.AnyToString(cases[1].ori.(string))))

}

func TestConvConfiguration(t *testing.T) {
	suite.Run(t, new(ConvSuite))
}
