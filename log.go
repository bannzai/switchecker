package main

import "fmt"

func debugf(format string, arguments ...interface{}) {
	if verbose {
		fmt.Printf(format, arguments...)
	}
}
