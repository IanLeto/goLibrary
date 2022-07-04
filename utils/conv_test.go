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
		{ori: []string{"1", "2", "3"}, except: "[\"1\",\"2\",\"3\"]"},
		{ori: "[\"1\",\"2\",\"3\"]", except: []string{"1", "2", "3"}},
		{ori: "{\"data\":{\"k\":\"v\"}}", except: map[string]interface{}{"data": map[string]string{"k": "v"}}},
		{ori: map[string]interface{}{"data": map[string]string{"k": "v"}}, except: "{\"data\":{\"k\":\"v\"}}"},
	}
	s.Equal(cases[0].except, utils.ArrToString([]string{"1", "2", "3"}))
	s.Equal(cases[1].except, utils.StringToArr(utils.AnyToString(cases[1].ori.(string))))
	s.Equal(cases[2].except, utils.StringToMap(""))
	s.Equal(cases[3].except, utils.MapToString(map[string]interface{}{}))

}

func TestConvConfiguration(t *testing.T) {
	suite.Run(t, new(ConvSuite))
}
