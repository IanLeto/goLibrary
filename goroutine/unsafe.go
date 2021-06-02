package goroutine

import (
	"sync"
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
		}()
	}
	wg.Wait()
	return unsafe.a
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
