package json_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type JsonSuite struct {
	suite.Suite
}

func (s *JsonSuite) SetupTest() {
}

// TestMarshal :
func (s *JsonSuite) TestHelloWorld() {
	data := []byte("{\"key\":\"word\",\"context\": {\"k1\":\"v1\"}}")
	var a = struct {
		Key     string      `json:"key"`
		Context interface{} `json:"context"`
	}{}
	err := json.Unmarshal(data, &a)
	s.NoError(err)
	fmt.Println(a.Context)
}

func TestJSONConfiguration(t *testing.T) {
	suite.Run(t, new(JsonSuite))
}
