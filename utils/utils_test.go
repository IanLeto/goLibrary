package utils_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"goLibrary/utils"
	"testing"
)

type TestRateSuit struct {
	suite.Suite
}

func (s *TestRateSuit) SetupTest() {
	//var base float64 = 100000
	//var rate float64 = 6.8 / base
	//for i := 0; i < 365*60; i++ {
	//	base = base * (1 + rate)
	//}
	//fmt.Println(base)
}

// 银行
func (s *TestRateSuit) TestRate() {
	var base float64 = 100000
	var rate float64 = 0.036
	for i := 0; i < 20; i++ {
		base = base * (1 + rate)
	}
	fmt.Println(base)
}

// 通胀
func (s *TestRateSuit) TestRate2() {
	var base float64 = 100000
	var rate float64 = 0.036
	for i := 0; i < 20; i++ {
		base = base * (1 - rate)
	}
	fmt.Println(base)
}

func (s *TestRateSuit) TestEnv() {
	s.Equal(utils.GetLocalOSEnv("CCMODE"), "DEBUG")
}

func (s *TestRateSuit) TestBatch() {
	s.Equal([][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9}}, utils.Batch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))
}

func (s *TestRateSuit) TestFormatJson() {
	t := "1,2,3cxx\n{\"a\":\"1\"}"
	err, res := utils.FormatJson(t)
	s.NoError(err)
	s.Equal(res["a"], "1")
}
func (s *TestRateSuit) TestWget() {
	s.NoError(utils.Wget("www.baidu.com", "/Users/ian/go/src/goLibrary/utils/test1", "wget_logs", "1", 1))
	//testFilePath := fmt.Sprintf("%s", path.PathADD("./", "test"))
	//if utils.Exists(testFilePath) {
	//	err := os.Remove(testFilePath)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//fileObj ,err:= os.OpenFile(testFilePath, os.O_RDWR|os.O_APPEND, 0777)
	//if err != nil {
	//	panic(err)
	//}
}

func (s *TestRateSuit) Test() {
	ori := []int{7, 8, 9, 10, 1, 3, 3, 5, 6}

	s.Equal([]int{7, 8, 9, 10, 1, 3, 5, 6}, utils.RmEle(5, ori))

}
func (s *TestRateSuit) TestConsisten() {
	//ori := []int{7, 8, 9, 10, 1, 3, 3, 5, 6}

	s.Equal([]int{}, utils.Consisten([]int{7, 8, 9, 10, 4, 5, 6}))
	s.Equal([]int{11}, utils.Consisten([]int{7, 8, 9, 10, 4, 12, 5, 6}))

}

func TestRaSuite(t *testing.T) {
	suite.Run(t, new(TestRateSuit))
}
