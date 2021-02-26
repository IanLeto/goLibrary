package utils

func CheckPanicError(err error) {
	if err != nil {
		panic(err)
	}
}
