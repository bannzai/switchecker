package main

import "github.com/bannzai/switchecker/example/valid/with_external/thirdparty"

func main() {
	lang := thirdparty.Golang
	switch lang {
	case thirdparty.Golang:
		println("golang")
	case thirdparty.Swift:
		println("swift")
	default:
		println("default")
	}
}
