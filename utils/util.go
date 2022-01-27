package utils

import (
	"os"
)

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// 获取本地环境变量

func GetLocalOSEnv(key string) string {
	return os.Getenv(key)

}

// 分批处理slice
func Batch(s []int, limit int) [][]int {

	var batches [][]int

	for limit < len(s) {
		s, batches = s[limit:], append(batches, s[0:limit:limit])
	}
	batches = append(batches, s)
	return batches
}
