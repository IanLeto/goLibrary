package gennyDemo

import (
	"fmt"
	"github.com/cstockton/go-conv"
)

// why 啥需要泛型
// if 我们想实现一个类型转化的helper
func stringHelper(a interface{}) string {
	res, _ := conv.String(a)
	return res
}

// 我们需要写好多个helper
func intHelper(a interface{}) int {
	res, _ := conv.Int(a)
	return res
}

// 好tmd 难受
func UseHelper() {
	var r string
	r = stringHelper("xx")
	fmt.Println(r)
}

// 如果我们偷懒用interface 来替换这种
func helper(a interface{}) interface{} {
	return a
}

// 调用

func UserHelper2() {
	var r string
	// 我们在使用的时候就必须要强转
	r, err := conv.String(helper("xx"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// 如果转错了，比如说传入了无法变成字符串的东西
	r, err = conv.String(helper(struct {
	}{}))
	// interface 接口的缺陷GG

}

// 在java 中 可以用泛型结局，go中无泛型
