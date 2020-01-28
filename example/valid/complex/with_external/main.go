package main

import "github.com/bannzai/switchecker/example/valid/complex/with_external/thirdparty"

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

func a() {
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
