package main

import (
	"github.com/bannzai/switchecker/example/missing/with_external/thirdparty"
)

func main() {
	lang := thirdparty.Golang
	switch lang {
	case thirdparty.Golang:
		println("golang")
	default:
		println("default")
	}
}
