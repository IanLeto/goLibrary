package base

import (
	"context"
	"fmt"
	"time"
)

func RunSelect(ctx context.Context) {
	// cancel 提前取消的意思
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("超时")
		}

	}

	//time.Sleep(1000 * time.Second)

}
