package goroutine_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
)

// bench

type ConcurrencySuite struct {
	suite.Suite
}

func (s *ConcurrencySuite) SetupTest() {

}

// TestMarshal : pool 语法
func (s *ConcurrencySuite) TestPool() {
	var (
		a   int
		wg  sync.WaitGroup
		num int // 并发数量
	)
	// 创建一个pool
	myPool := &sync.Pool{New: func() interface{} {
		a += 1
		fmt.Println("创建新实例")
		return make([]byte, 1024)
	}}
	// 初始化pool
	myPool.Put(myPool.New())
	myPool.Put(myPool.New())
	// 并发数量
	num = 100000
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			// mypool 在调用get 时会先看看有没有已经有在pool中的实例
			// if 有， 捞出来用，没有就新建一个
			number := myPool.Get().([]byte)
			// 用完放回到池子中
			defer myPool.Put(number)
			defer wg.Done()
		}()
	}
	wg.Wait()
	// 可以看到，我们每次产生的新对象数量不一样
	fmt.Println(a)
}

// 语法2
func (s *ConcurrencySuite) TestPool2() {
	var (
		wg     sync.WaitGroup
		num    int // 并发数量
		conNum int
	)
	// 我们看看， 每次的对象是否改变
	// 返回匿名结构体
	myPool := &sync.Pool{New: func() interface{} {
		conNum += 1
		return struct {
			a int // 注意 这个a 是并发不安全，此处为举例
		}{a: 1}
	}}
	// 初始化pool
	myPool.Put(myPool.New())
	myPool.Put(myPool.New())
	// 并发数量
	num = 100000
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			// mypool 在调用get 时会先看看有没有已经有在pool中的实例
			// if 有， 捞出来用，没有就新建一个
			obj := myPool.Get().(struct{ a int })
			obj.a += 1
			defer myPool.Put(obj)
			defer wg.Done()
		}()
	}
	wg.Wait()
	// 我们一个new 了 多少个对像
	fmt.Println(conNum)
	x := conNum
	wg.Add(x)
	for i := 0; i < x; i++ {
		go func() {
			obj := myPool.Get().(struct{ a int })
			fmt.Println(obj.a) // 可以看到，我们的obj.a 的值是我们无法预知的
			//defer myPool.Put(obj) // 拿出来不放回去, 结果更直观
			defer wg.Done()
		}()

	}
	wg.Wait()

}

// 测试pool 内存情况
func BenchmarkMemoryPool(b *testing.B) {
	var ()
	myPool := &sync.Pool{New: func() interface{} {
		return struct {
			a int
		}{a: 1}
	}}
	// 初始化pool
	for i := 0; i < 10; i++ {
		myPool.Put(myPool.New())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		obj := myPool.Get().(struct{ a int })
		obj.a += 1
		myPool.Put(obj)
	}
}

// go test 使用 这个文件 run benchmark  run匹配到memorypool 的bench； 打印内存情况
// go test -v concurrency_test.go -test.bench MemoryPool -test.run MemoryPool -benchmem
// 测试pool 内存情况
// 结论： 2 远快于1 但这只是单线程
func BenchmarkMemoryPool2(b *testing.B) {

	for i := 0; i < b.N; i++ {
		obj := struct{ a int }{}
		obj.a += 1
	}
}

func TestConcurrency(t *testing.T) {
	suite.Run(t, new(ConcurrencySuite))
}
