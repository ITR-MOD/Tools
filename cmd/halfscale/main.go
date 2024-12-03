package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/ITR-MOD/tools/libs"

	"github.com/nfnt/resize"
)

func downscaleImage(inputPath string, scaleFactor float64) {
	// Open the input image
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the new dimensions
	newWidth := uint(float64(img.Bounds().Dx()) * scaleFactor)
	newHeight := uint(float64(img.Bounds().Dy()) * scaleFactor)

	// Resize the image
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	// Create the output file
	outputPath := inputPath[:len(inputPath)-4] + ".ds.png"
	outFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Encode the resized image and save to the output file
	err = png.Encode(outFile, resizedImg)
	if err != nil {
		log.Fatal(err)
	}

	// Print success message
	fmt.Printf("Image saved to %s\n", outputPath)
}

func main() {
	// Ensure the right number of arguments
	if len(os.Args) <= 2 {
		fmt.Println("Usage: go run script.go path/to-image.png")
		return
	}

	// Parse the arguments to accommodate spaces in the file path
	args := libs.ParsePathArgs()

	for _, arg := range args {
		if libs.IsPathFile(arg) == false {
			fmt.Printf("Invalid file path: %s\n", arg)
			continue
		}
		// Downscale the image by 50%
		downscaleImage(arg, 0.5)

	}
}
