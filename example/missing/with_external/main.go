package main

import (
	"fmt"

	"github.com/bannzai/switchecker/example/missing/with_external/thirdparty"
)

func main() {
	lang := thirdparty.Golang
	switch lang {
	case thirdparty.Golang:
		fmt.Println("golang")
	default:
		println("default")
	}
}
