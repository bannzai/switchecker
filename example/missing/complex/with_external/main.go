package main

import "github.com/bannzai/switchecker/example/missing/complex/with_external/thirdparty"

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
	case 5:
		println("literal")
	default:
		println("default")
	}
}
