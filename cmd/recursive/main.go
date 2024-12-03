package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ITR-MOD/tools/libs"
)

// processImages walks through the directory structure and processes image files.
func processImages(exePath, root string) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // Return the error to stop walking on fatal errors.
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Process image files
		if libs.IsFileImage(d.Name()) {
			fmt.Printf("Processing image: %s\n", path)
			fmt.Println("Running executable:", exePath)

			// // Execute the provided executable
			// cmd := exec.Command(exePath, path)
			// cmd.Stdout = os.Stdout
			// cmd.Stderr = os.Stderr

			// if err := cmd.Run(); err != nil {
			// 	fmt.Printf("Error processing image %s: %v\n", path, err)
			// }
		}

		return nil
	})
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: program.exe <exe> <root dir>")
		os.Exit(1)
	}

	exePath := os.Args[1]
	rootDir := os.Args[2]

	if _, err := os.Stat(exePath); os.IsNotExist(err) {
		fmt.Printf("Executable does not exist: %s\n", exePath)
		os.Exit(1)
	}

	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		fmt.Printf("Directory does not exist: %s\n", rootDir)
		os.Exit(1)
	}

	if err := processImages(exePath, rootDir); err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
		os.Exit(1)
	}
}
