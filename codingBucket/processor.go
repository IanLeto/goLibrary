package codingBucket

type Processor interface {
	Process(value []byte, next chan []byte) error
}
