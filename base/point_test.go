package base

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PointSuite struct {
	suite.Suite
}

type test struct {
	Key string `json:"key"`
}

func (s *PointSuite) SetupTest() {

}

// TestMarshal :
func (s *PointSuite) TestHelloWorld() {
	byte := []byte(string("{\"key\":\"value\"}"))
	var (
		a1  test
		a2  *test
		a3  = test{}
		a4  = &test{}
		err error
	)

	err = json.Unmarshal(byte, a1)
	s.NoError(err)
	fmt.Println(a1)
	err = json.Unmarshal(byte, a2)
	s.NoError(err)
	fmt.Println(a2)
	err = json.Unmarshal(byte, a3)
	s.NoError(err)
	fmt.Println(a3)
	err = json.Unmarshal(byte, a4)
	s.NoError(err)
	fmt.Println(a4)
	err = json.Unmarshal(byte, &a1)
	s.NoError(err)
	fmt.Println(a1)
	err = json.Unmarshal(byte, &a2)
	s.NoError(err)
	fmt.Println(a2)
	err = json.Unmarshal(byte, &a3)
	s.NoError(err)
	fmt.Println(a3)
	err = json.Unmarshal(byte, &a4)
	s.NoError(err)
	fmt.Println(a4)

}

func TestRunPoint(t *testing.T) {
	suite.Run(t, new(PointSuite))
}
