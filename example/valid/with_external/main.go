package main

import "./thirdparty"

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
