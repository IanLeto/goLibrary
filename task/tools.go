package task

import (
	"goLibrary/utils"
	"time"
)

func NewCronTask(t Task, ticker time.Ticker) error {

	utils.NoErr(t.Start())
	for {
		select {
		case <-ticker.C:
			utils.NoErr(t.Start())
		}
	}
}
