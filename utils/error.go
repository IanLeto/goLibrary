package utils

import "github.com/pkg/errors"

func CheckPanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

type causeError interface {
	error
	Cause() error
}

// 以防万一的处理异常方式
func RecoverError(fn func(error)) {
	v := recover()
	switch err := v.(type) {
	case nil:
		return
	case causeError:
		fn(err)
	case error:
		fn(errors.Errorf("%v", err))
	default:
		fn(errors.Errorf("%v", v))
	}
}
