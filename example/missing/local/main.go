package main

type language int

const (
	golang language = iota
	swift
)

func main() {
	lang := golang
	switch lang {
	case golang:
		println("golang")
	default:
		println("default")
	}
}
