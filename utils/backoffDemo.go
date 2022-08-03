package utils

import (
	"github.com/cenkalti/backoff"
	"time"
)

// 声明
func ExponentialRetry(n int, fn func() error) error {
	bf := backoff.NewExponentialBackOff()
	bf.Multiplier = 2.0
	bf.InitialInterval = 3 * time.Second
	return backoff.Retry(fn, backoff.WithMaxRetries(bf, uint64(n)))
}
