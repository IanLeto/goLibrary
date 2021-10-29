package utils

import (
	"encoding/json"
	"github.com/spf13/cast"
	"strings"
)

func StringToArr(sql string) []string {
	return cast.ToStringSlice(sql)
}

func ArrToString(sql []string) string {
	return strings.Join(sql, ",")
}

func StringToMap(sql string) []string {
	//v, _ := json.Marshal(sql)
	return nil
	//return cast.ToString(v)
}

func MapToString(sql []string) string {
	return strings.Join(sql, ",")
}

func PayloadToString(payload interface{}) string {
	res, err := json.Marshal(payload)
	if err != nil {
		return ""
	}
	return string(res)
}
