package greeting

import (
	"fmt"
	"os"
	"path/filepath"
)

func SayHelloToWorld() {
	fmt.Println("Saying hello to the World")
}

func FindMin (array []int) int {

		min := array[0]

		for i:= 1; i < len(array); i++ {
			if min < array[i] {
				min = i
			}
		}


	return min
}

func SearchFiles(rootPath string, files *[]string) {

	entries, err := os.ReadDir(rootPath)
	if err != nil {
		return 
	}

	for _, entry := range entries {

		fullPath := filepath.Join(rootPath, entry.Name())

		if entry.IsDir() {
			SearchFiles(fullPath, files)
		} else {
			*files = append(*files, fullPath)
		}
	}
}