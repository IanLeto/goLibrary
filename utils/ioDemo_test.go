package utils

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type IOSuite struct {
	suite.Suite
}

func (s *IOSuite) SetupTest() {
}

// TestMarshal :
func (s *IOSuite) TestHelloWorld() {
	//KeepWrite("goLibrary/utils/test.sh", "goLibrary/utils/test.txt")
}

// TestMarshal :
func (s *IOSuite) TestOpeFile() {
	fmt.Println(MakeFileName("test", 1, "", "use-", ".py"))
} // TestMarshal :
func (s *IOSuite) TestTransferFilePath() {
	s.Equal(GetFilePath("utils/path/index.html"), TransferFilePath("index.html", "", "www.baidu.com"))
}

func (s *IOSuite) TestScan() {
	file, err := os.OpenFile("./test.txt", 0, 0777)
	s.NoError(err)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		fmt.Println(scan.Text())
	}
}

func TestIOConfiguration(t *testing.T) {
	suite.Run(t, new(IOSuite))
}
