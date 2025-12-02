package internal

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/charmbracelet/huh"
)

type PreviewGenerator struct{}

func NewPreviewGenerator() *PreviewGenerator {
	return &PreviewGenerator{}
}

func (pg *PreviewGenerator) CreatePreview(img image.Image) error {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	previewImg := image.NewRGBA(bounds)

	rgbaImg, ok := img.(*image.RGBA)
	if !ok {
		return fmt.Errorf("expected RGBA image for preview, but got %T", img)
	}

	for y := range height {
		for x := range width {
			c := rgbaImg.RGBAAt(x, y)
			previewImg.SetRGBA(x, y, color.RGBA{
				R: c.R * 17,
				G: c.G * 17,
				B: c.B * 17,
				A: c.A * 17,
			})
		}
	}

	filename, err := pg.getPreviewFilename()
	if err != nil {
		return fmt.Errorf("failed to get preview filename: %w", err)
	}

	return pg.savePreviewImage(previewImg, filename)
}

func (pg *PreviewGenerator) getPreviewFilename() (string, error) {
	var filename string

	err := huh.NewInput().Title("Enter name of preview image relative to current directory: ").
		Prompt("?").
		Value(&filename).
		DescriptionFunc(func() string {
			if _, err := os.Stat(filename); os.IsNotExist(err) {
				return "This preview file is just so you have an idea of what it should look like."
			}
			return fmt.Sprintf("WARNING!: file %s already exists", filename)
		}, &filename).Run()
	if err != nil {
		return filename, err
	}
	return filename, nil
}

func (pg *PreviewGenerator) savePreviewImage(img image.Image, filename string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create preview file: %w", err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		return fmt.Errorf("failed to encode preview image: %w", err)
	}

	fmt.Printf("Preview image saved as: %s\n", filename)
	return nil
}
