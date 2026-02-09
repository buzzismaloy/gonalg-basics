package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func containsDot(s string) bool {
	return strings.Contains(s, ".")
}

func main() {
	fmt.Println("All files without filter:\n")
	printAllFiles("../")

	fmt.Println("\n\nAll files with filter:\n")
	printAllFilesWithFilterClosure("../", ".go")

	fmt.Println("\n\nAll files with filter function(path is . and filter searches for dot in the filename:\n")
	printFilesWithFuncFilter(".", containsDot)
}

func printFilesWithFuncFilter(path string, predicate func(string) bool) {
	var walk func(string)

	walk = func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}

		for _, f := range files {
			filename := filepath.Join(path, f.Name())

			if predicate(filename) {
				fmt.Println(filename)
			}

			if f.IsDir() {
				walk(filename)
			}
		}
	}

	walk(path)

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
