package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var verbose bool

const argumentsSeparator = ","

func main() {
	var source string
	var target string
	flag.StringVar(&source, "source", "*.go", "Source go files are containing enum definition. Multiple specifications can be specified separated by ,. e.g) *.go,pkg/**/*.go ")
	flag.StringVar(&target, "target", "*.go", "Target go files are containing to use enum. Multiple specifications can be specified separated by ,. e.g) *.go,pkg/**/*.go ")
	flag.BoolVar(&verbose, "verbose", false, "Enabled verbose log")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	debugf("source file search start. path: %s\n", source)
	splitedSource := strings.Split(source, argumentsSeparator)
	sources := []string{}
	for _, source := range splitedSource {
		foundPaths, err := filepath.Glob(cwd + "/" + source)
		if err != nil {
			log.Fatal(err)
		}
		sources = append(sources, foundPaths...)
	}
	debugf("source file search end: %v\n", sources)

	debugf("target file search start. path: %s\n", target)
	splitedTarget := strings.Split(target, argumentsSeparator)
	targets := []string{}
	for _, target := range splitedTarget {
		foundPaths, err := filepath.Glob(cwd + "/" + target)
		if err != nil {
			log.Fatal(err)
		}
		targets = append(targets, foundPaths...)
	}
	debugf("target file search end: %v\n", targets)

	debugf("enum parse start: \n")
	enums := parse(sources)
	debugf("enum parse end: %v\n", enums)

	if err := check(enums, targets); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\033[32mSuccesfull!! $ switchecker -source=%s -target=%s\033[0m\n", source, target)
}
