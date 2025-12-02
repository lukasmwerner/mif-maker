package internal

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

type InputHandler struct {
	reader *bufio.Reader
}

func NewInputHandler() *InputHandler {
	return &InputHandler{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (ih *InputHandler) GetImageFilename() (string, error) {
	fmt.Print("Enter name of image relative to current directory: ")
	filename, err := ih.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(filename), nil
}

func (ih *InputHandler) LoadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening image file: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("error decoding image: %w", err)
	}

	return img, nil
}

// sourced from : https://github.com/speakeasy-api/speakeasy/blob/main/internal/charm/utils.go#L87
func FileBasedSuggestions(relativeDir string, fileExtensions []string) []string {
	var validFiles []string
	workingDir, err := os.Getwd()
	if err != nil {
		return validFiles
	}

	targetDir := relativeDir
	if !filepath.IsAbs(targetDir) {
		targetDir = filepath.Join(workingDir, targetDir)
	}

	if targetDir == "" {
		targetDir = workingDir
	}

	files, err := os.ReadDir(targetDir)
	if err != nil {
		return validFiles
	}

	validDirectories := []string{}
	for _, file := range files {
		if !file.Type().IsDir() {
			for _, ext := range fileExtensions {
				if strings.HasSuffix(file.Name(), ext) {
					fileSuggestion := filepath.Join(relativeDir, file.Name())
					// Allows us to support current directory relative paths
					if relativeDir == "./" {
						fileSuggestion = relativeDir + file.Name()
					}
					validFiles = append(validFiles, fileSuggestion)
				}
			}
		} else {
			fileSuggestion := filepath.Join(relativeDir, file.Name())
			validDirectories = append(validDirectories, fileSuggestion)
		}
	}

	// HACK: this gets the folders at the bottom of the suggestions
	for _, dir := range validDirectories {
		validFiles = append(validFiles, dir)

	}

	return validFiles
}
