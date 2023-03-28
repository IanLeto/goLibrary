package utils

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type BytesSuite struct {
	suite.Suite
}

func (s *BytesSuite) SetupTest() {
	file, err := os.OpenFile("testBytes", os.O_WRONLY|os.O_CREATE, 0666)
	s.NoError(err)
	_, err = file.Write([]byte("hello world"))
	s.NoError(err)

}

func (s *BytesSuite) TestMySQL() {
	file, err := os.Open("testBytes")
	s.NoError(err)
	defer func() { _ = file.Close() }()
	scanner := bufio.NewScanner(file)
	var data []string
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}
	s.NoError(scanner.Err())
	fmt.Println(data)

}

func TestBytesSuite(t *testing.T) {
	suite.Run(t, new(BytesSuite))

}
