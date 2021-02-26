package eventBusDemo

func PublishDemo(conf string) {
	GlobalEvent.Publish("demo", conf)
}
