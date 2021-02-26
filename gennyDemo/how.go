package gennyDemo

// 首先 我们实现了一个任意类型
//func SaySYMBOL(input string) TYPE {
//	res, _ := conv.SYMBOL(input)
//	return res
//}
//然后 把这串代码变成类似于模板的东西 how.tpl

// go 生成器  命令  使用哪个文件 使用哪个包 -out 输出什么文件 gen 生成什么类型
//go:generate genny -in how.tpl -pkg ${GOPACKAGE} -out maphelper_string.go gen "SYMBOL=String TYPE=string"


//bug
//Failed to goimports the generated code: how.tpl:10:2: expected declaration, found 'return'
//how.go:12: running "genny": exit status 4

// 解决