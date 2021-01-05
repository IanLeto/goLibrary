package syntaxDemo

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func ErrorGroupDemo() {
	var eg errgroup.Group
	for i := 0; i < 30; i++ {
		eg.Go(func() error {
			time.Sleep(1 * time.Second)
			if i < 20 {
				fmt.Printf("error occured: %d", i)
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}

func ErrorGroupDemoRun() {
	ErrorGroupDemo()
}
