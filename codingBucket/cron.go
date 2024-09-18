package codingBucket

type CronJobTask interface {
	Do(err error)
}
