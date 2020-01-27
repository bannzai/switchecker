package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var source string
	var target string
	flag.StringVar(&source, "source", "./**/*.go", "Source go files are containing enum definition")
	flag.StringVar(&target, "target", "./**/*.go", "Target go files are containing to use enum")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	sources, err := filepath.Glob(cwd + "/" + source)
	if err != nil {
		log.Fatal(err)
	}
	targets, err := filepath.Glob(cwd + "/" + target)
	if err != nil {
		log.Fatal(err)
	}

	enums := parse(sources)
	for _, targetPath := range targets {
		if err := check(enums, targetPath); err != nil {
			log.Fatal(err)
		}
	}
}
