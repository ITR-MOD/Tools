package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"strings"

	"github.com/ITR-MOD/Tools/libs"
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

		// Open the ORM texture file using OpenImage
		img, err := libs.ReadImage(filePath)
		if err != nil {
			fmt.Printf("Failed to open image %s: %s\n", filePath, err)
			continue
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

		// Save each channel as a separate image using libs.WriteImage
		if err := libs.WriteImage(occlusion, filepath.Join(dir, baseName+".occlu.png")); err != nil {
			fmt.Printf("Failed to save Occlusion image: %s\n", err)
		}
		if err := libs.WriteImage(roughness, filepath.Join(dir, baseName+".rough.png")); err != nil {
			fmt.Printf("Failed to save Roughness image: %s\n", err)
		}
		if err := libs.WriteImage(metallic, filepath.Join(dir, baseName+".metal.png")); err != nil {
			fmt.Printf("Failed to save Metallic image: %s\n", err)
		}

		fmt.Println("Textures successfully unpacked:")
		fmt.Printf("- Occlusion: %s.occlu.png\n", baseName)
		fmt.Printf("- Roughness: %s.rough.png\n", baseName)
		fmt.Printf("- Metallic: %s.metal.png\n", baseName)
	}
}
