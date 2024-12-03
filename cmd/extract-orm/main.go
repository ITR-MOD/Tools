package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/ITR-MOD/tools/libs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run orm-extract.go ./path-to-orm.png")
		return
	}

	args := libs.ParsePathArgs()

	for _, filePath := range args {
		if !libs.IsPathFile(filePath) {
			continue
		}
		// Open the ORM texture file
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Failed to open file: %s\n", err)
			return
		}
		defer file.Close()

		// Decode the image
		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Printf("Failed to decode image: %s\n", err)
			return
		}

		// Get base name and directory
		dir := filepath.Dir(filePath)
		baseName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))

		// Extract channels
		width := img.Bounds().Dx()
		height := img.Bounds().Dy()

		occlusion := image.NewGray(image.Rect(0, 0, width, height))
		roughness := image.NewGray(image.Rect(0, 0, width, height))
		metallic := image.NewGray(image.Rect(0, 0, width, height))

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				occlusion.SetGray(x, y, color.Gray{uint8(r >> 8)})
				roughness.SetGray(x, y, color.Gray{uint8(g >> 8)})
				metallic.SetGray(x, y, color.Gray{uint8(b >> 8)})
			}
		}

		// Save each channel as a separate image
		saveImage(filepath.Join(dir, baseName+".occlu.png"), occlusion)
		saveImage(filepath.Join(dir, baseName+".rough.png"), roughness)
		saveImage(filepath.Join(dir, baseName+".metal.png"), metallic)

		fmt.Println("Textures successfully unpacked:")
		fmt.Printf("- Occlusion: %s.occlu.png\n", baseName)
		fmt.Printf("- Roughness: %s.rough.png\n", baseName)
		fmt.Printf("- Metallic: %s.metal.png\n", baseName)
	}
}

func saveImage(filePath string, img image.Image) {
	outFile, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Failed to save image %s: %s\n", filePath, err)
		return
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		fmt.Printf("Failed to encode image %s: %s\n", filePath, err)
	}
}
