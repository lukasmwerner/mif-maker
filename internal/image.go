package internal

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/webp"

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
	if height != 256 || width != 256 {
		fmt.Println(warningStyle.Render("ATTENTION: The image dimensions are not 256x256 and will be cropped/adjusted to 256x256"))
		maxRect := image.Rect(0, 0, 256, 256)
		croppedImg = image.NewRGBA(maxRect)
		for y := range 256 {
			for x := range 256 {
				var c color.Color
				if x < width && y < height {
					c = img.At(x, y)
				} else {
					c = color.Transparent
				}

				croppedImg.(*image.RGBA).Set(x, y, c)
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
	draw.Draw(rgbaImg, rgbaImg.Bounds(), img, img.Bounds().Min, draw.Src)
	return rgbaImg
}
