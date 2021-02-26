package utils

import (
	"bufio"
	"github.com/cstockton/go-conv"
	"io"
	"os"
)

func IOByte(data []byte) {

}

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
