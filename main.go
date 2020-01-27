package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var verbose bool

func main() {
	var source string
	var target string
	flag.StringVar(&source, "source", "*.go", "Source go files are containing enum definition")
	flag.StringVar(&target, "target", "*.go", "Target go files are containing to use enum")
	flag.BoolVar(&verbose, "verbose", false, "Enabled verbose log")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	debugf("source file search start. path: %s\n", source)
	sources, err := filepath.Glob(cwd + "/" + source)
	if err != nil {
		log.Fatal(err)
	}
	debugf("source file search end: %v\n", sources)
	debugf("target file search start. path: %s\n", target)
	targets, err := filepath.Glob(cwd + "/" + target)
	if err != nil {
		log.Fatal(err)
	}

	debugf("enum parse start: \n")
	enums := parse(sources)
	for _, targetPath := range targets {
		if err := check(enums, targetPath); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("\033[32mSuccesfull!!\033[0m")
}
