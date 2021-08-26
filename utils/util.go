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
