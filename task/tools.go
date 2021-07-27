package task

import (
	"context"
	"goLibrary/utils"
	"time"
)

func NewCronTask(ctx context.Context, t Task, ticker time.Ticker) error {
	for {
		select {
		case <-ticker.C:
			utils.NoErr(t.Work())
		}
	}
}
