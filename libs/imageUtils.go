package libs

import (
	"image"
	"image/png"
	"os"
)

// WriteImageToPNG writes an image.Image to the specified output path as a PNG file.
func WriteImage(img image.Image, outputPath string) error {
	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Encode the image as PNG and write it to the file
	err = png.Encode(outFile, img)
	if err != nil {
		return err
	}

	return nil
}

func ReadImage(inputPath string) (image.Image, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}
