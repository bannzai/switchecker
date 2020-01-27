package main

import "fmt"

func debugf(format string, arguments ...interface{}) {
	fmt.Printf(format, arguments...)
}
