//go:build windows
// +build windows

package libs

import (
	"fmt"
	"os"
)

// ParsePathArgs parses command line arguments and combines them into paths if they exist.
func ParsePathArgs() []string {
	args := []string{}
	currentArg := ""
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
		if len(currentArg) > 0 {
			currentArg += " " + arg
		} else {
			currentArg = arg
		}
		if _, err := os.Stat(currentArg); err == nil {
			args = append(args, currentArg)
			fmt.Println("Added", currentArg)
			currentArg = ""
		}
	}
	if len(currentArg) > 0 {
		args = append(args, currentArg)
	}
	return args
}
