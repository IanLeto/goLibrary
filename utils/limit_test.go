package utils_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"os"
	"os/exec"
	"testing"
)

type TestListSuit struct {
	suite.Suite
}

func (s *TestListSuit) SetupTest() {

}

// 银行
func (s *TestListSuit) TestList() {

	//for _, i2 := range []int{1, 2, 3} {
	//	if i2 > 2 {
	//		fmt.Println(111)
	//		continue
	//	}
	//	fmt.Println(22222)
	//}

}

func (s *TestListSuit) TestFastDemo() {
	//fmt.Println('1')
	fmt.Println(os.Environ())
	fmt.Println(1111)
	cmd := exec.Command("bash", "-c", "echo $PATH afjl $a")
	cmd.Stdout = os.Stdout
	cmd.Env = append(cmd.Env, "a=1")
	cmd.Run()
	//fmt.Println(cmd.Stdout)

}

func TestLimitSuite(t *testing.T) {
	suite.Run(t, new(TestListSuit))
}
