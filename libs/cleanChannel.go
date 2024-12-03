package libs

import (
	"image"
	"image/color"
)

// Function to remove the red channel
func RemoveRedChannel(img image.Image) image.Image {
	return processImage(img, func(r, g, b, a uint32) color.RGBA {
		return color.RGBA{
			R: 0,
			G: uint8(g >> 8),
			B: uint8(b >> 8),
			A: uint8(a >> 8),
		}
	})
}

// Function to remove the green channel
func RemoveGreenChannel(img image.Image) image.Image {
	return processImage(img, func(r, g, b, a uint32) color.RGBA {
		return color.RGBA{
			R: uint8(r >> 8),
			G: 0,
			B: uint8(b >> 8),
			A: uint8(a >> 8),
		}
	})
}

// Function to remove the blue channel
func RemoveBlueChannel(img image.Image) image.Image {
	return processImage(img, func(r, g, b, a uint32) color.RGBA {
		return color.RGBA{
			R: uint8(r >> 8),
			G: uint8(g >> 8),
			B: 0,
			A: uint8(a >> 8),
		}
	})
}

// Generic image processing function
func processImage(img image.Image, modifyColor func(r, g, b, a uint32) color.RGBA) image.Image {
	// Create a new RGBA image to store the modified pixels
	bounds := img.Bounds()
	outputImg := image.NewRGBA(bounds)

	// Iterate over each pixel and apply the modification
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			r, g, b, a := originalColor.RGBA()
			newColor := modifyColor(r, g, b, a)
			outputImg.Set(x, y, newColor)
		}
	}

	return outputImg
}
