package utils

import (
	"bufio"
	"fmt"
	"github.com/cstockton/go-conv"
	"goLibrary/utils/path"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

//OpenFile函数flag参数
//打开方式	说明
//O_RDONLY	只读方式打开
//O_WRONLY	只写方式打开
//O_RDWR	读写方式打开
//O_APPEND	追加方式打开
//O_CREATE	不存在，则创建
//O_EXCL	如果文件存在，且标定了O_CREATE的话，则产生一个错误
//O_TRUNG	如果文件存在，且它成功地被打开为只写或读写方式，将其长度裁剪唯一。（覆盖）
//O_NOCTTY	如果文件名代表一个终端设备，则不把该设备设为调用进程的控制设备
//O_NONBLOCK	如果文件名代表一个FIFO,或一个块设备，字符设备文件，则在以后的文件及I/O操作中置为非阻塞模式。
//O_SYNC	当进行一系列写操作时，每次都要等待上次的I/O操作完成再进行。

// 如何在go 中使用IO
func IanIODemo(input, output string) {
	// 初始化输入输出
	var (
		err        error
		inputFile  *os.File
		res        = make([]string, 0)
		outputFile *os.File
	)

	// 如果该方法传入参数无用，那么说明输入是从stdin
	if input == "" {
		inputFile = os.Stdin
	} else {
		inputFile, err = os.Open(input)
		if err != nil {
			panic(err)
		}
		defer func() {
			CheckPanicError(inputFile.Close())
		}()
	}
	// 初始化一个reader
	reader := bufio.NewReader(inputFile)
	// 按行读取
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		l, err := conv.String(line)
		res = append(res, l)
	}

	if output == "" {
		outputFile = os.Stdout
	} else {
		outputFile, err = os.OpenFile(output, os.O_RDWR|os.O_APPEND, 0777)
		if err != nil {
			panic(err)
		}
		defer func() {
			CheckPanicError(outputFile.Close())
		}()
	}

	for _, v := range res {
		_, err = io.WriteString(outputFile, v+"\n")
	}

}

// 将cmd 的输出不断地写入file
func BoWrite() {
	c1 := exec.Command("/bin/bash", path.GetFilePath("goLibrary/utils/test.sh"))
	outputFile, err := os.OpenFile(path.GetFilePath("goLibrary/utils/test.txt"), os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriterSize(outputFile, 10)
	r, err := c1.StdoutPipe()
	_, err = io.Copy(writer, r)
	if err != nil {
		panic(err)
	}
	c1.Start()
	c1.Wait()
	fmt.Println("done")
}

// 判断所给路径文件/文件夹是否存在

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {

			return true

		}
		return false
	}
	return true
}

// 拼接文件名称
func FileUrlADD(path1, path2 string) string {
	return filepath.Join(path1, path2)
}

// 创建文件目录
func MakeDirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// 创建文件
func CreatFile(path string) error {
	_, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	return err
}

// 创建 if 不存在

func CreatOrOver(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)

}

// 根据条件生成文件名称
func MakeFileName(name string, id int, path string, prefix, suffix string) string {

	fileName := fmt.Sprintf("%s-%s-%s", strconv.Itoa(id), name, time.Now().Format("2006-01-02 15:04:05"))
	if prefix != "" {
		fileName = fmt.Sprintf("%s%s", prefix, fileName)
	}
	if suffix != "" {
		fileName = fmt.Sprintf("%s%s", fileName, suffix)
	}
	if path != "" {
		fileName = filepath.Join(path, fileName)
	}
	return fileName
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {

		return false

	}
	return s.IsDir()
}

// 判断所给路径是否为文件

func IsFile(path string) bool {

	return !IsDir(path)

}

