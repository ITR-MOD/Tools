package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ITR-MOD/Tools/libs"
)

// TODO, make this not only work on images

// processImages walks through the directory structure and processes image files.
func processImages(exePath string, exeArgs []string, root string) error {
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

			// Build the command with the executable and arguments
			args := append(exeArgs, path) // Pass the image file path as the last argument
			cmd := exec.Command(exePath, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			// Run the executable
			if err := cmd.Run(); err != nil {
				fmt.Printf("Error processing image %s: %v\n", path, err)
			}
		}

		return nil
	})
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: program.exe <exe> <root dir> [args...]")
		os.Exit(1)
	}

	exePath := os.Args[1]
	rootDir := os.Args[2]
	exeArgs := os.Args[3:] // Capture additional arguments to pass to the executable

	rootDir = filepath.Clean(rootDir)

	// Resolve the full path of the executable
	fullExePath, err := exec.LookPath(exePath)
	if err != nil {
		fmt.Printf("Executable not found: %s\n", exePath)
		os.Exit(1)
	}

	// Check if the root directory exists
	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		fmt.Printf("Directory does not exist: %s\n", rootDir)
		os.Exit(1)
	}

	// Process images
	if err := processImages(fullExePath, exeArgs, rootDir); err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
		os.Exit(1)
	}
}
