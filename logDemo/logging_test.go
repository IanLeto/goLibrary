package logDemo_test

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/suite"
	"goLibrary/logDemo"
	"net/url"
	"unicode/utf8"

	"testing"
)

type TestLogSuit struct {
	suite.Suite
}

func (s *TestLogSuit) SetupTest() {

}

func (s *TestLogSuit) TestConf() {
	str := "Hello, 世界!"

	// 输出原始字符串
	fmt.Println("原始字符串：", str)

	// 输出 UTF-8 编码的字符串
	fmt.Print("UTF-8 编码：")
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%x ", r)
		i += size
	}
	fmt.Println()

	// 输出 Base64 编码的字符串
	fmt.Println("Base64 编码：", base64.StdEncoding.EncodeToString([]byte(str)))

	// 输出 URL 编码的字符串
	fmt.Println("URL 编码：", url.QueryEscape(str))

}

func (s *TestLogSuit) TestHook() {
	//logDemo.UseHook()
	logDemo.UseDivLogFile()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestLogSuit))
}
