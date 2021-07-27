package pprof

import (
	"fmt"
	"time"
)



func BeTest() {
	f1()
	f2()
	fmt.Println("done")
}

func f1() {
	time.Sleep(1 * time.Second)
}

func f2() {
	time.Sleep(2 * time.Second)
}
