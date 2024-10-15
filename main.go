package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var categories = map[string]string{
	"js":   "Document",
	"py":   "Document",
	"png":  "Image",
	"gif":  "Image",
	"exe":  "Executable",
	"zip":  "zipFolder",
	"html": "web",
	"go":   "Document",
	"mp4":  "video",
}

func main() {
	if len(os.Args) < 2 {
		errors.New("Not entered a directory! Use: ")
		fmt.Println("go run main.go <file path>")
		return
	}

	directoryPath := os.Args[1]

	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		fmt.Println("Error: Directory path does not exist")
		return
	}

	files, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Println("Error reading files or files are corrupted:", err)
		return
	}

	makedirectory(directoryPath)

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Skipping directory:", file.Name())
			continue
		}

		extension := strings.Split(file.Name(), ".")
		ext := extension[1]

		if ext != "" {
			if category, ok := categories[ext]; ok {
				srcPath := filepath.Join(directoryPath, file.Name())
				destDir := filepath.Join(directoryPath, category)
				destPath := filepath.Join(destDir, file.Name())

				err := os.Rename(srcPath, destPath)
				if err != nil {
					fmt.Println("Error moving file:", err)
				} else {
					fmt.Println("Moved", file.Name(), "to", category, "directory")
				}
			} else {
				fmt.Println("Unknown file type for:", file.Name())
			}
		} else {
			fmt.Println("Skipping file without extension:", file.Name())
		}
	}
}

func makedirectory(directoryPath string) {
	for _, cat := range categories {
		categoryPath := filepath.Join(directoryPath, cat)

		if _, err := os.Stat(categoryPath); os.IsNotExist(err) {
			err := os.Mkdir(categoryPath, 0755)
			if err != nil {
				fmt.Println("Error creating directory:", cat)
			} else {
				fmt.Println("Created directory:", cat)
			}
		// } else {
		// 	fmt.Println(cat, "directory already exists!")
		// }
		}
	}
}