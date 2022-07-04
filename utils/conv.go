package utils

import (
	"encoding/json"
	"github.com/cstockton/go-conv"
	"github.com/spf13/cast"
	"strings"
)

func StringToArr(sql string) []string {
	return cast.ToStringSlice(sql)
}

func ArrToString(sql []string) string {
	return strings.Join(sql, ",")
}

func MapToString(v map[string]interface{}) string {
	//return strings.Join(sql, ",")
	return ""
}
func StringToMap(v string) map[string]interface{} {
	//return strings.Join(sql, ",")
	return map[string]interface{}{}
}

func PayloadToString(payload interface{}) string {
	res, err := json.Marshal(payload)
	if err != nil {
		return ""
	}
	return string(res)
}

// 强制转换
func AnyToString(any interface{}) string {
	res, _ := conv.String(any)
	return res
}
