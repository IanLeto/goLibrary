package io_test

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

// OpenFile函数flag参数
// 打开方式	说明
// O_RDONLY	只读方式打开
// O_WRONLY	只写方式打开
// O_RDWR	读写方式打开
// O_APPEND	追加方式打开
// O_CREATE	不存在，则创建
// O_EXCL	如果文件存在，且标定了O_CREATE的话，则产生一个错误
// O_TRUNG	如果文件存在，且它成功地被打开为只写或读写方式，将其长度裁剪唯一。（覆盖）
// O_NOCTTY	如果文件名代表一个终端设备，则不把该设备设为调用进程的控制设备
// O_NONBLOCK	如果文件名代表一个FIFO,或一个块设备，字符设备文件，则在以后的文件及I/O操作中置为非阻塞模式。
// O_SYNC	当进行一系列写操作时，每次都要等待上次的I/O操作完成再进行。
func (s *IOSuite) SetupTest() {
	x, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {

	}
	defer func() { _ = x.Close() }()
}

// TestMarshal :
func (s *IOSuite) TestIOReader() {
	file, err := os.OpenFile("x.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	s.NoError(err)
	defer func() { _ = file.Close() }()
	reader := bufio.NewReader(file)
	buf := make([]byte, 1024)
	//读取 Reader 对象中的内容到 []byte 类型的 buf 中
	_, err = reader.Read(buf)
	fmt.Println(string(buf))
}

func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(IOSuite))
}
