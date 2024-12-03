package libs

import (
	"os"
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
