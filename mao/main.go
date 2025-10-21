package main

import (
	"mao/assets"
	"mao/lib"
)

func main() {
	content, err := assets.StaticFs.ReadFile("banner.txt")
	if err == nil {
		lib.Banner(content, nil)
	}

}
