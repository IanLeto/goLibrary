package gennyDemo

import (
	"github.com/cstockton/go-conv"
)

// 首先 我们实现了一个任意类型
func SaySYMBOL(input string) TYPE {
	res, _ := conv.SYMBOL(input)
	return res
}