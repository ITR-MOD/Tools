package libs

import (
	"os"
	"path/filepath"
	"strings"
)

func IsPathFile(path string) bool {
	// Check if the argument is a file
	notfailed := true
	fileInfo, err := os.Stat(path)
	if err != nil {
		notfailed = false
	}
	if fileInfo.IsDir() {
		notfailed = false
	}
	return notfailed
}

// isImageFile checks if a file is an image based on its extension.
func IsFileImage(filename string) bool {
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".bmp", ".gif", ".tiff", ".webp"}
	ext := strings.ToLower(filepath.Ext(filename))
	for _, imgExt := range imageExtensions {
		if ext == imgExt {
			return true
		}
	}
	return false
}
