package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/ITR-MOD/tools/libs"

	"github.com/nfnt/resize"
)

func downscaleImage(inputPath string, scaleFactor float64, overwrite *bool) {
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
	outputPath := fmt.Sprintf("%s.ds_%.2f.png", inputPath, scaleFactor)
	if *overwrite {
		outputPath = fmt.Sprintf("%s.png", inputPath)
	}

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
	// Define and parse flags
	scaleFactor := flag.Float64("scale", 0.5, "Scaling factor (0.0 < scale <= 1.0)")
	overwrite := flag.Bool("overwrite", false, "Overwrite the original image")
	flag.Parse()

	// Validate the scaling factor
	if *scaleFactor <= 0 || *scaleFactor > 1 {
		log.Fatal("Error: scale must be a positive number less than or equal to 1")
	}

	// Remaining arguments are file paths
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: go run script.go -scale=0.5 path/to-image1.png [path/to-image2.png ...]")
		return
	}

	for _, arg := range args {
		if !libs.IsPathFile(arg) {
			fmt.Printf("Invalid file path: %s\n", arg)
			continue
		}
		// Downscale the image with the specified scale factor
		downscaleImage(arg, *scaleFactor, overwrite)
	}
}
