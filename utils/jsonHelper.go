package utils

import "github.com/valyala/fastjson"

type JsonHelper struct {
	*fastjson.Parser
	*fastjson.Value
}

func NewJsonHelperString(v string) *JsonHelper {
	helper := &JsonHelper{}
	if v == "" {
		return helper
	}
	res, err := helper.Parse(v)
	NoErr(err)
	helper.Value = res
	return helper
}

func NewJsonHelperByte(v []byte) *JsonHelper {
	helper := &JsonHelper{}
	if v == nil {
		return helper
	}
	res, err := helper.ParseBytes(v)
	NoErr(err)
	helper.Value = res
	return helper
}
