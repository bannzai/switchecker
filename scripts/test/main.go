package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	exampleRootDirectory := pwd + "/../../example/"
	var currentPath string
	err = filepath.Walk(exampleRootDirectory, func(path string, info os.FileInfo, err error) error {
		currentPath = path
		if info.IsDir() {
			return nil
		}
		if "test.sh" != filepath.Base(path) {
			return nil
		}
		if err := os.Chdir(filepath.Dir(path)); err != nil {
			return err
		}
		defer os.Chdir(pwd)

		fmt.Printf("exec test.sh for %s\n", path)
		cmd := exec.Command("./test.sh")
		if err := cmd.Start(); err != nil {
			return err
		}
		if err := cmd.Wait(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(fmt.Sprintf("cause error of %v, path is %s", err, currentPath))
	}
}
