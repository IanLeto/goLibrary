package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

var BaseTestFilePath = GetFilePath("utils/path")

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

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../..", "goLibrary")
)

func IncludeString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getRootPath() string {
	path, _ := os.Getwd()
	return filepath.Join(path, "")
}

func GetFilePath(path string) string {
	return filepath.Join(Root, path)
}

//
func Wget(url, target, logOut string, retry string, limit int) error {
	cmd := exec.Command("wget", "-O", target, "-o", logOut, "-t", retry, url)
	return cmd.Start()
}

// 文件名， 文件路径， 文件基础路径， wget 日志路径， wget 日志名， retry
// 返回 该文件的绝对路径
func TransferFilePath(fileName, filePath, url string) (path string) {
	address := filepath.Join(url, filePath, fileName)
	// 目标文件目录
	dir := MakeFileName(fileName, 1, BaseTestFilePath, "", "")
	// 创建目标文件目录
	_ = os.MkdirAll(dir, os.ModePerm)
	// 目标文件绝对路径
	abs := filepath.Join(dir, fileName)

	_ = Wget(address, filepath.Join(dir, fileName), filepath.Join(dir, "wgetLog"), "1", 1)
	return abs
}

func trace2() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}

func TimeCost() func() {
	start := time.Now()
	return func() {
		fmt.Println("cost", time.Since(start))
		trace2()
	}
}

func RmEle(tar int, origin []int) []int {
	target := origin[:0]
	for _, item := range origin {
		if item != tar {
			target = append(target, item)
		}
	}
	return target
}

func Consisten(a []int) []int {

	if len(a) < 3 {

	}
	res := make([]int, 0)
	sort.Ints(a)
	for i, j := 0, 1; i < len(a)-1; i, j = i+1, j+1 {
		if x := a[j] - a[i]; x > 1 {
			for z := 0; z < x-1; z++ {
				res = append(res, a[i]+1+z)
			}
		}
	}
	return res
}

func JustSee(s interface{}) {
	res, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
