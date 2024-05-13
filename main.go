package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func main() {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	fmt.Printf("CWD: %s\n", wd)

	// Create a filesystem interface for the current working directory
	dir := os.DirFS(wd)
	var nodeModulesPaths []string
	// Walk through the directory tree starting from the current working directory
	err = fs.WalkDir(dir, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %s\n", path, err)
			return err
		}

		// fmt.Printf("Path: %s, IsDir: %t\n", path, d.IsDir())
		// Check if the path contains "node_modules"
		if strings.HasSuffix(path, "node_modules") && d.IsDir() {
			// fmt.Printf("Found directory: %s\n", path)
			toRemove := wd + "/" + path
			nodeModulesPaths = append(nodeModulesPaths, toRemove)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking directory:", err)
	}

	for _, path := range nodeModulesPaths {
		fmt.Printf("Deleting directory: %s...\n", path)
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Printf("Error deleting directory %s: %s\n", path, err)
		} else {
			fmt.Printf("Deleted directory: %s\n", path)
		}
	}

	fmt.Println("Done!")
	os.Exit(0)
}
