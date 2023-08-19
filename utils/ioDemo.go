package utils

import (
	"bufio"
	"fmt"
	"github.com/cstockton/go-conv"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

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
			//CheckPanicError(outputFile.Close())
		}()
	}

	for _, v := range res {
		_, err = io.WriteString(outputFile, v+"\n")
	}

}

// KeepWrite 将cmd 的输出不断地写入file
// 脚本文件，目标文件建
//func KeepWrite(shellPath, filePath string) {
//	cmd := exec.Command("/bin/bash", path.GetFilePath(shellPath))
//	outputFile, err := os.OpenFile(path.GetFilePath(filePath), os.O_RDWR|os.O_APPEND, 0777)
//	NoErr(err)
//	writer := bufio.NewWriterSize(outputFile, 10)
//	stdOut, err := cmd.CombinedOutput()
//	_, err = io.Copy(writer, stdOut)
//	NoErr(err)
//	NoErr(cmd.Start())
//	NoErr(cmd.Wait())
//	fmt.Println("done")
//}

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

// FileUrlADD 拼接文件名称
func FileUrlADD(path1, path2 string) string {
	return filepath.Join(path1, path2)
}

// 创建文件目录
func MakeDirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// CreatFile 创建文件
func CreatFile(path string) error {
	_, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	return err
}

// CreatOrOver 创建 if 不存在
func CreatOrOver(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)

}

// 写入文件
func Write(content, fileName string) {
	fileObj, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	NoErr(err)
	defer NoErr(fileObj.Close())
	writerObj := bufio.NewWriterSize(fileObj, 4096)
	buf := []byte(content)
	if _, err := writerObj.Write(buf); err != nil {
		if err := writerObj.Flush(); err != nil {
			panic(err)
		}
	}
}

func ReadFile() {
	fileObj, err := os.Open("demo.txt")
	if err != nil {
		fmt.Println("文件打开失败：", err)
		return
	}
	//os.OpenFile()
	defer fileObj.Close()
	//一个文件对象本身是实现了io.Reader的 使用bufio.NewReader去初始化一个Reader对象，存在buffer中的，读取一次就会被清空
	reader := bufio.NewReader(fileObj)
	buf := make([]byte, 1024)
	//读取 Reader 对象中的内容到 []byte 类型的 buf 中
	info, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("读取的字节数:" + strconv.Itoa(info))
	//这里的buf是一个[]byte，因此如果需要只输出内容，仍然需要将文件内容的换行符替换掉
	fmt.Println("读取的文件内容:", string(buf))
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

// IsDir 判断所给路径是否为文件夹
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
