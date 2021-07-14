package goroutine

import (
	"fmt"
	"sync"
	"time"
)

// 三种常见的不安全并发模型 注意，并发安全与否与是否为指针操作无关
func UnsafeDemo() int {
	unsafe := struct {
		a   int
		res []int
	}{
		a:   0,
		res: []int{1, 2, 3},
	}

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			wg.Done()
			unsafe.a += 1
		}()
	}
	wg.Wait()
	return unsafe.a
}

func UnsafeDemo2() int {
	ss := 0
	unsafe := &struct {
		a   int
		res []int
	}{
		a:   0,
		res: []int{1, 2, 3},
	}

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			wg.Done()
			unsafe.a += 1
			ss += 1
		}()
	}
	wg.Wait()
	return ss
}

func UnsafeDemo3() int {
	unsafe := struct {
		a   int
		res []int
	}{
		a:   0,
		res: []int{1, 2, 3},
	}

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(*struct {
			a   int
			res []int
		}) {
			wg.Done()
			unsafe.a += 1
		}(&unsafe)
	}
	wg.Wait()
	return unsafe.a
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func UnsafeDemo4() {
	ch := make(chan int)
	res := []int{}
	wg := sync.WaitGroup{}
	l := makeRange(0, 1000)
	wg.Add(len(l))
	go func() {
		for i := range ch {
			res = append(res, i)
			wg.Done()
		}
	}()

	for _, v := range l {
		go func(s int) {
			time.Sleep(1 * time.Second)
			fmt.Println(s)
			ch <- s
		}(v)
	}

	wg.Wait()
	fmt.Println(">?")
	num := make(map[int]bool)
	for _, v := range res {
		if !num[v] {
			num[v] = true
		} else {
			fmt.Println(v)
		}
	}
}
