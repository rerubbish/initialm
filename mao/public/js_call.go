package public

import (
	"encoding/json"
	"log/slog"
	"mao/cmd/msg"
	"syscall/js"
)

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

// 这个方法接收前端传递过来的json字符串
func GenZip(this js.Value, args []js.Value) interface{} {
	// 将data转换为字符串
	if len(args) < 1 {
		return js.ValueOf("参数不足，需要1个json字符串")
	}
	dataStr := args[0].String()
	// 反序列化为ExData
	var exData msg.ExData
	err := json.Unmarshal([]byte(dataStr), &exData)
	if err != nil {
		slog.Error("json反序列化失败", err.Error())
		return js.ValueOf("json反序列化失败")
	}
	slog.Info("GenZip", "exData", exData)
	return js.ValueOf("success")
}
