package eventBus

func PublishDemo(conf string) {
	GlobalEvent.Publish("demo", conf)
}
