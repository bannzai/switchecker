package main

import "github.com/bannzai/switchecker/example/with_thirdparty/thirdparty"

func main() {
	lang := thirdpartyGolang
	switch lang {
	case thirdparty.Golang:
		println("golang")
	case thirdparty.Swift:
		println("swift")
	default:
		println("default")
	}
}
