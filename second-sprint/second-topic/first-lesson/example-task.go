package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("All files without filter:\n")
	printAllFiles("../")

	fmt.Println("\n\nAll files with filter:\n")
	printAllFilesWithFilterClosure("../", ".go")
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

func printAllFilesWithFilterClosure(path, filter string) {
	var walk func(string)

	walk = func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}

		for _, f := range files {
			filename := filepath.Join(path, f.Name())

			if strings.Contains(filename, filter) {
				fmt.Println(filename)
			}

			if f.IsDir() {
				walk(filename)
			}
		}
	}

	walk(path)
}
