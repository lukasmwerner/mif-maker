package internal

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
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