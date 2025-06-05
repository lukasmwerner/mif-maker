package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/lukasmwerner/mif-maker/internal"
)

func main() {
	var filename string
	err := huh.NewInput().Title("Enter name of image relative to current directory: ").
		Prompt("?").
		Validate(func(s string) error {
			if _, err := os.Stat(s); os.IsNotExist(err) {
				return fmt.Errorf("file %s does not exist", s)
			}
			return nil
		}).Value(&filename).Run()
	if err != nil {
		fmt.Printf("Error getting filename: %v\n", err)
		os.Exit(1)
	}

	img, err := internal.LoadImage(filename)
	if err != nil {
		fmt.Printf("Error loading image: %v\n", err)
		os.Exit(1)
	}

	uint4Converter := internal.NewUint4Converter()
	previewGenerator := internal.NewPreviewGenerator()
	mifGenerator := internal.NewMifGenerator()

	resizedImg := internal.Resize(img)

	uint4Img, err := uint4Converter.CreateU4(resizedImg)
	if err != nil {
		fmt.Printf("Error creating uint4 data: %v\n", err)
		os.Exit(1)
	}

	internal.CompareImages(internal.Image{
		Img:   img,
		Title: "Original Data",
	}, internal.Image{
		Img:   uint4Img,
		Title: "MIF Data (uint4)",
	})

	err = previewGenerator.CreatePreview(uint4Img)
	if err != nil {
		fmt.Printf("Error creating preview: %v\n", err)
	}

	err = mifGenerator.WriteMif(uint4Img, "mifData.mif")
	if err != nil {
		fmt.Printf("Error writing MIF file: %v\n", err)
		os.Exit(1)
	}
}
