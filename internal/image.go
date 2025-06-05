package internal

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

var (
	warningStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#DFA000"))
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#F85552"))
	ruleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#8DA102"))
	titleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#34A77C"))
)

func LoadImage(filename string) (image.Image, error) {
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

type ImageProcessor struct{}

type Image struct {
	Img   image.Image
	Title string
}

func CompareImages(left Image, right Image) error {
	termWidth, _, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		return err
	}

	document := lipgloss.NewStyle().MaxWidth(termWidth)
	leftContent := ""
	rightContent := ""
	{
		img := left.Img
		bounds := left.Img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()
		leftContent = lipgloss.JoinVertical(
			lipgloss.Left,
			titleStyle.Render(left.Title),

			fmt.Sprintf("%s %T", ruleStyle.Render("Data type:"), img),
			fmt.Sprintf("%s %s", ruleStyle.Render("Image bounds:"), bounds),
			fmt.Sprintf("%s %d", ruleStyle.Render("Number of words:"), width*height),
			fmt.Sprintf("%s %d", ruleStyle.Render("Height (vertical/row) of image:"), height),
			fmt.Sprintf("%s %d", ruleStyle.Render("Width (horizontal/column) of image:"), width),
		)
	}
	{
		img := right.Img
		bounds := right.Img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()
		rightContent = lipgloss.JoinVertical(
			lipgloss.Left,
			titleStyle.Render(right.Title),

			fmt.Sprintf("%s %T", ruleStyle.Render("Data type:"), img),
			fmt.Sprintf("%s %s", ruleStyle.Render("Image bounds:"), bounds),
			fmt.Sprintf("%s %d", ruleStyle.Render("Number of words:"), width*height),
			fmt.Sprintf("%s %d", ruleStyle.Render("Height (vertical/row) of image:"), height),
			fmt.Sprintf("%s %d", ruleStyle.Render("Width (horizontal/column) of image:"), width),
		)
	}

	comparisonStyle := lipgloss.NewStyle().Padding(1)

	fmt.Println(document.Render(lipgloss.JoinHorizontal(lipgloss.Top, comparisonStyle.BorderRight(true).Render(leftContent), comparisonStyle.Render(rightContent))))
	return nil
}

func Resize(img image.Image) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	croppedImg := img
	if height > 256 || width > 256 {
		fmt.Println(warningStyle.Render("ATTENTION: The image dimensions are too large and will be cropped to 256x256"))
		maxRect := image.Rect(0, 0, 256, 256)
		croppedImg = image.NewRGBA(maxRect)
		for y := range min(height, 256) {
			for x := range min(width, 256) {
				croppedImg.(*image.RGBA).Set(x, y, img.At(x, y))
			}
		}
	}

	if _, ok := croppedImg.(*image.RGBA); ok {
		return croppedImg
	}

	return convertToRGBA(croppedImg)
}

func convertToRGBA(img image.Image) *image.RGBA {
	rgbaImg := image.NewRGBA(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			originalColor := img.At(x, y)
			r, g, b, a := originalColor.RGBA()
			rgbaImg.Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		}
	}

	for y := rgbaImg.Bounds().Min.Y; y < rgbaImg.Bounds().Max.Y; y++ {
		for x := rgbaImg.Bounds().Min.X; x < rgbaImg.Bounds().Max.X; x++ {
			r, g, b, _ := rgbaImg.At(x, y).RGBA()
			rgbaImg.Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255})
		}
	}

	return rgbaImg
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
