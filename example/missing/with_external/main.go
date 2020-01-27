package main

import "./thirdparty"

func main() {
	lang := thirdparty.Golang
	switch lang {
	case thirdparty.Golang:
		println("golang")
	default:
		println("default")
	}
}
