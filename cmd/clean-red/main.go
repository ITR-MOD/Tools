package main

import (
	"fmt"
	"os"

	"github.com/ITR-MOD/Tools/libs"
)

func main() {
	// Ensure the right number of arguments
	if len(os.Args) <= 1 {
		fmt.Println("Usage: go run script.go path/to-image.png")
		return
	}

	// Parse the arguments to accommodate spaces in the file path
	args := libs.ParsePathArgs()

	for _, arg := range args {
		if !libs.IsPathFile(arg) {
			fmt.Printf("Invalid file path: %s\n", arg)
			continue
		}
		img, err := libs.ReadImage(arg)
		if err != nil {
			fmt.Printf("Failed to read image: %s\n", err)
			continue
		}
		// Remove the green channel from the image
		noGreen := libs.RemoveRedChannel(img)
		outputPath := fmt.Sprintf("%s.no-R.png", arg)
		err = libs.WriteImage(noGreen, outputPath)
		if err != nil {
			fmt.Printf("Failed to write image: %s\n", err)
		} else {
			fmt.Printf("Image written to: %s\n", outputPath)
		}
	}
}
