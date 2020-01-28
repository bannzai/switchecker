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
	err = filepath.Walk(exampleRootDirectory, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("check test path is %s\n", path)
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

		fmt.Printf("exec to test.sh\n")
		cmd := exec.Command("./test.sh")
		if err := cmd.Start(); err != nil {
			fmt.Printf("start failed\n")
			return err
		}
		if err := cmd.Wait(); err != nil {
			fmt.Printf("wait failed\n")
			return err
		}
		fmt.Printf("success\n")
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
