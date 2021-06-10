package response

type BaseRes struct {
}
type ARes struct {
	*BaseRes
}

func Hello(res BaseRes) {

}

// 这里是一种错误的认知 组合继承 和 继承完全不一样
func UsageDemo() {
	// Hello(ARes{nil})
}
