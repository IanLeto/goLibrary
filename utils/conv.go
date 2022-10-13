package utils

import (
	"bufio"
	"fmt"
	"github.com/cstockton/go-conv"
	"os"
	"strings"
)

type stringConverter interface {
	String() (string, error)
}

func StringToArr(sql string) []string {
	return strings.Split(sql, ",")
}

func ArrToString(sql []string) string {
	return strings.Join(sql, ",")
}

// 强制转换
func AnyToString(any interface{}) string {
	res, _ := conv.String(any)
	return res
}

// AnyToStringOrigin 原生方法
func AnyToStringOrigin(from interface{}, file string) (res string, err error) {
	fileObj, err := os.OpenFile("./testsuits", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	writerObj := bufio.NewWriter(fileObj)

	defer func() {
		nums, err := writerObj.WriteString(res)
		if err != nil {
			panic(err)
		}
		fmt.Println(nums)
	}()

	switch T := from.(type) {
	case string:
		return T, nil
	case stringConverter:
		return T.String()
	case []byte:
		return string(T), nil
	case *[]byte:
		// @TODO Maybe validate the bytes are valid runes
		return string(*T), nil
	case *string:
		return *T, nil
	default:
		return fmt.Sprintf("%v", from), nil
	}

}
