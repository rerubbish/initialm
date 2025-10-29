package public

import "syscall/js"

// 定义可被JS调用的函数：处理参数并返回结果
func Add(this js.Value, args []js.Value) interface{} {
	// 检查参数数量和类型（避免panic）
	if len(args) < 2 {
		return js.ValueOf("参数不足，需要2个数字")
	}
	a := args[0].Float() // 转换JS数字为float64
	b := args[1].Float()
	return js.ValueOf(a + b) // 返回结果（转换为JS类型）
}
