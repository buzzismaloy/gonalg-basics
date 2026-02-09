package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	printAllFiles("../")
}

func printAllFiles(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}

	for _, f := range files {
		filename := filepath.Join(path, f.Name())
		fmt.Println(filename)

		if f.IsDir() {
			printAllFiles(filename)
		}
	}
}
