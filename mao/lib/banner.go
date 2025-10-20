package lib

import (
	"fmt"
	"html/template"
	"strings"

	"mao/public"
)

func Banner(data any) {
	content, err := public.StaticFs.ReadFile("banner.txt")
	if err != nil {
		return
	}
	banner := string(content)
	if data != nil {
		tpl, err := template.New("banner").Parse(banner)
		if err != nil {
			return
		}
		var result strings.Builder
		if err := tpl.Execute(&result, data); err != nil {
			return
		}
		// 打印Banner内容
		fmt.Println(result.String())
	}
	fmt.Println(banner)
}
