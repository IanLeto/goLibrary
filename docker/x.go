package main

import (
	"bufio"
	"os"
)

func main() {
	buf := []byte{}
	for i := 0; i < 29999999; i++ {
		buf = append(buf, byte(1))
	}
	f, _ := os.OpenFile("/Users/ian/go/src/goLibrary/docker/testdata", os.O_WRONLY|os.O_CREATE, 0666)
	writer := bufio.NewWriter(f)
	writer.Write(buf)

}
