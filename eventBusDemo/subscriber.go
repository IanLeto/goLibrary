package eventBusDemo

import (
	"fmt"
)

func ReceiveDemo() {
	//defer PublishDemo("key")

	err := GlobalEvent.Subscribe("demo", func(key string) {
		fmt.Println("get publish", key)
	})
	PublishDemo("Value")
	if err != nil {
		panic(err)
	}

	//err := GlobalEvent.SubscribeAsync("hh", func() {},false)
}
