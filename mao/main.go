package main

import (
	"log/slog"
	"mao/assets"
	"mao/lib"
)

func init() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}
func main() {
	content, err := assets.StaticFs.ReadFile("banner.txt")
	if err == nil {
		lib.Banner(content, nil)
	}
}
