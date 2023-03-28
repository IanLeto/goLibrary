package gotemplate_test

import (
	"github.com/stretchr/testify/suite"
	"html/template"
	"io/ioutil"
	"os"
	"testing"
)

type TemplateTextSuite struct {
	suite.Suite
	Data []byte
}

func (s *TemplateTextSuite) SetupTest() {
	var err error
	s.Data, err = ioutil.ReadFile("./file.yml")
	s.NoError(err)
}

// TestMarshal :
func (s *TemplateTextSuite) TestHelloWorld() {
	// hello 必须是结构体，不能是纯字符串
	var hello = struct {
		Hello string
	}{
		Hello: "world",
	}
	template, err := template.New("manifest").Parse(string(s.Data))
	s.NoError(err)
	err = template.Execute(os.Stdout, hello)
	s.NoError(err)

}

func TestTempalteConfiguration(t *testing.T) {
	suite.Run(t, new(TemplateTextSuite))
}
