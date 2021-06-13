package backoffDemo

import (
	"github.com/cenkalti/backoff"
)

// 声明
func ExponentialRetry(n int, fn func() error) error {
	bf := backoff.NewExponentialBackOff()
	return backoff.Retry(fn, backoff.WithMaxRetries(bf, uint64(n)))
}

// 使用
func Run()  {

}