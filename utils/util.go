package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func FormatJson(input string) (err error, j map[string]interface{}) {
	// 提醒推送

	if err = json.Unmarshal([]byte(input), &j); err != nil {
		// json 处理
		resultList2 := strings.Split(input, "\n")
		str3 := ""
		for i := range resultList2 {
			str3 = StringBuilder([]string{resultList2[len(resultList2)-i-1], "\n", str3})
			if err = ParamIllegalJsonCheck(str3); err == nil {
				input = str3
				break
			}
		}
	}
	if err = json.Unmarshal([]byte(input), &j); err != nil {
		return err, j
	}
	return nil, j
}
func ParamIllegalJsonCheck(param string) error {
	if !json.Valid([]byte(param)) {
		return errors.New("Param is not Json")
	}
	return nil
}

// StringBuilder 高效率字符串拼接
func StringBuilder(p []string) string {
	var b strings.Builder
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}


//
func Wget(url ,fileName ,logOut string, retry string, limit int)  {
	exec.Command("wget", "-a",fmt.Sprintf("%s", fileName), "-o", fmt.Sprintf("%s", logOut), "-t", retry)
}
