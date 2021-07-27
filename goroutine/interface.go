package goroutine

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"time"
)

// 接口访问已经ctx 控制原型

func TimeoutModel() {
	var (
		ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		resCh  = make(chan string)
	)
	defer close(resCh)
	go func(ctx2 context.Context) {
		time.Sleep(10 * time.Second)
		select {
		case resCh <- "result" + cast.ToString(rand.Intn(12)):
		case <-ctx.Done():
			return
		}

	}(ctx)

	select {
	case v, ok := <-resCh:
		if !ok {
			panic("写入失败")
		}
		fmt.Printf("结果%s", v)
	case <-ctx.Done():
		fmt.Println("超时关闭")
		return
	}
}
