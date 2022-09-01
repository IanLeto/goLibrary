package utils

import (
	"github.com/cstockton/go-conv"
	"strings"
)

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
