package internal

import (
	"fmt"
	"image"
	"image/color"
)

type Pixel4 struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type Uint4Converter struct{}

func NewUint4Converter() *Uint4Converter {
	return &Uint4Converter{}
}

func (u *Uint4Converter) CreateU4(img image.Image) (image.Image, error) {
	rgbaImg, ok := img.(*image.RGBA)
	if !ok {
		return nil, fmt.Errorf("expected RGBA image, but got %T", img)
	}

	bounds := rgbaImg.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	uint4Img := image.NewRGBA(bounds)

	for y := range height {
		for x := range width {
			c := rgbaImg.RGBAAt(x, y)
			uint4Img.SetRGBA(x, y, color.RGBA{
				R: c.R / 17,
				G: c.G / 17,
				B: c.B / 17,
				A: c.A / 17,
			})
		}
	}

	return uint4Img, nil
}

func (u *Uint4Converter) FormatPixel(p Pixel4) string {
	return fmt.Sprintf("%x%x%x%x", p.R, p.G, p.B, p.A)
}

func (u *Uint4Converter) FormatAddress(width, row, col int) string {
	address := row*width + col
	return fmt.Sprintf("%x", address)
}

func (u *Uint4Converter) ExtractU4Data(img image.Image) ([]string, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	rgbaImg, ok := img.(*image.RGBA)
	if !ok {
		return nil, fmt.Errorf("expected RGBA image, but got %T", img)
	}

	var data []string
	for y := range height {
		for x := range width {
			c := rgbaImg.RGBAAt(x, y)
			pixel := Pixel4{R: c.R, G: c.G, B: c.B, A: c.A}
			dataLine := fmt.Sprintf("%s: %s;", u.FormatAddress(width, y, x), u.FormatPixel(pixel))
			data = append(data, dataLine)
		}
	}
	return data, nil
}