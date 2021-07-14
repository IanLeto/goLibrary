package goroutine

import (
	"fmt"
	"goLibrary/utils"
	"sync"
	"time"
)

func HowToCount() {
	list := utils.MakeRange(1, 999)
	var once = sync.Once{}
	var wg = sync.WaitGroup{}
	for _, i := range list {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			once.Do(func() {
				fmt.Println(i)
			})
		}(i)
	}
	wg.Wait()
}
