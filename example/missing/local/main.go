package main

import (
	"fmt"
)

type language int

const (
	golang language = iota
	swift
)

func main() {
	lang := golang
	switch lang {
	case golang:
		fmt.Println("golang")
	default:
		println("default")
	}
}
